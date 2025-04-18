package service

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	helper "SHOP_PORTAL_BACKEND/HELPER"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"database/sql"
)

func GetCustomerBill(ownerRegId string, billId int, logPrefix string) (interface{}, int) {

	var response interface{}
	var rspCode = utils.StatusOK

	DB := database.DB

	tx, err := DB.Begin()
	if err != nil {
		return helper.Create500ErrorResponse("[DB ERROR 0113] Error starting transaction", "Error starting transaction:"+err.Error(), logPrefix)
	}

	defer tx.Rollback()

	// Get customer Id from Bill
	ServiceQuery := database.GetCustomerIdFromBill()
	var customerId int
	err = tx.QueryRow(ServiceQuery, billId).Scan(&customerId)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.CreateErrorResponse("404001", "Bill Not Found", logPrefix)
		}
		return helper.Create500ErrorResponse("[DB ERROR 0114] Error getting customer row ID", "Error getting customer row ID:"+err.Error(), logPrefix)
	}

	// Get Owner's row ID
	ownerRowId, err := helper.GetOwnerId(ownerRegId, tx)
	if err != nil {
		return helper.Create500ErrorResponse("[DB ERROR 0115] Error getting owner row ID", "Error getting owner row ID:"+err.Error(), logPrefix)
	}

	// check if customer is for this owner only
	ServiceQuery = database.CheckCustomerOwnerPresent()
	var isActive string
	err = tx.QueryRow(ServiceQuery, ownerRowId, customerId).Scan(&isActive)
	if err != nil && err != sql.ErrNoRows {
		return helper.Create500ErrorResponse("[DB ERROR 0116] Error getting customer row ID", "Error getting customer row ID:"+err.Error(), logPrefix)
	}

	if isActive == utils.NULL_STRING {
		return helper.CreateErrorResponse("404001", "Bill Customer registered with another owner", logPrefix)
	}

	if isActive == utils.ACTIVE_NO {
		return helper.CreateErrorResponse("404001", "Customer is InActive", logPrefix)
	}

	// Get Bill details
	result, response, rspCode := helper.GetBill(billId, tx, logPrefix)
	if rspCode != utils.StatusOK {
		return response, rspCode
	}

	if rspCode == utils.StatusOK {
		err = tx.Commit()
		if err != nil {
			return helper.Create500ErrorResponse("[DB ERROR 0121] Error committing transaction", "Error committing transaction:"+err.Error(), logPrefix)
		}
		utils.Logger.Info(logPrefix, "Transaction committed")
		response = result
	}

	return response, rspCode
}
