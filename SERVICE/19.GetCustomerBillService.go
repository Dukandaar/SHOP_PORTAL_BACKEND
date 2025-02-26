package service

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	helper "SHOP_PORTAL_BACKEND/HELPER"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"database/sql"
)

func GetBill(ownerRegId string, billId int, logPrefix string) (interface{}, int) {

	var response interface{}
	var rspCode = utils.StatusOK

	DB := database.DB

	tx, err := DB.Begin()
	if err != nil {
		return helper.Set500ErrorResponse("Error starting transaction", "Error starting transaction:"+err.Error(), logPrefix)
	}

	defer func() {
		if r := recover(); r != nil || rspCode != utils.StatusOK {
			utils.Logger.Error(logPrefix, "Panic occurred during transaction:", r)
			tx.Rollback()
			utils.Logger.Error(logPrefix, "Transaction rolled back")
		}
	}()

	// Get customer Id from Bill
	ServiceQuery := database.GetCustomerIdFromBill()
	var customerId int
	err = tx.QueryRow(ServiceQuery, billId).Scan(&customerId)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.CreateErrorResponse("404001", "Customer Not Found")
		}
		return helper.Set500ErrorResponse("Error getting customer row ID", "Error getting customer row ID:"+err.Error(), logPrefix)
	}

	// Get Owner's row ID
	ownerRowId, err := helper.GetOwnerId(ownerRegId, tx)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.CreateErrorResponse("404001", "Owner Not Found")
		}
		return helper.Set500ErrorResponse("Error getting owner row ID", "Error getting owner row ID:"+err.Error(), logPrefix)
	}

	// check if customer is for this owner only
	ServiceQuery = database.CheckCustomerOwnerPresent()
	var isActive string
	err = tx.QueryRow(ServiceQuery, ownerRowId, customerId).Scan(&isActive)
	if err != nil && err != sql.ErrNoRows {
		return helper.Set500ErrorResponse("Error getting customer row ID", "Error getting customer row ID:"+err.Error(), logPrefix)
	}

	if isActive == utils.NULL_STRING {
		return helper.CreateErrorResponse("404001", "Customer registered with another owner")
	}

	if isActive == utils.ACTIVE_NO {
		return helper.CreateErrorResponse("404001", "Customer is InActive")
	}

	result, response, rspCode := helper.AllBill(ownerRowId, customerId, tx, logPrefix)
	if rspCode != utils.StatusOK {
		return response, rspCode
	}

	if rspCode == utils.StatusOK {
		err = tx.Commit()
		if err != nil {
			return helper.Set500ErrorResponse("Error committing transaction", "Error committing transaction:"+err.Error(), logPrefix)
		}
		response = result
		utils.Logger.Info(logPrefix, "Transaction committed")
	}

	return response, rspCode
}
