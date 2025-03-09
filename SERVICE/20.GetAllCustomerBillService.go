package service

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	helper "SHOP_PORTAL_BACKEND/HELPER"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"database/sql"
)

func GetAllCustomerBill(ownerRegId string, customerRegId string, logPrefix string) (interface{}, int) {
	var response interface{}
	var rspCode = utils.StatusOK

	DB := database.DB

	tx, err := DB.Begin()
	if err != nil {
		return helper.Create500ErrorResponse("[DB ERROR 0122] Error starting transaction", "Error starting transaction:"+err.Error(), logPrefix)
	}

	defer tx.Rollback()

	// Get Owner's row ID
	ownerRowId, err := helper.GetOwnerId(ownerRegId, tx)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.CreateErrorResponse("404001", "Owner not found", logPrefix)
		}
		return helper.Create500ErrorResponse("[DB ERROR 0123] Error getting owner row ID", "Error getting owner row ID:"+err.Error(), logPrefix)
	}

	// Get customer Id
	customerId, err := helper.GetCustomerId(customerRegId, ownerRowId, tx)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.CreateErrorResponse("404001", "Customer registered for another owner", logPrefix)
		}
		return helper.Create500ErrorResponse("[DB ERROR 0124] Error getting customer row ID", "Error getting customer row ID:"+err.Error(), logPrefix)
	}

	// check if customer is for this owner only
	ServiceQuery := database.CheckCustomerOwnerPresent()
	var isActive string
	err = tx.QueryRow(ServiceQuery, ownerRowId, customerId).Scan(&isActive)
	if err != nil && err != sql.ErrNoRows {
		return helper.Create500ErrorResponse("[DB ERROR 0125] Error getting customer row ID", "Error getting customer row ID:"+err.Error(), logPrefix)
	}

	if isActive == utils.NULL_STRING {
		return helper.CreateErrorResponse("404001", "Customer registered with another owner", logPrefix)
	}

	if isActive == utils.ACTIVE_NO {
		return helper.CreateErrorResponse("404001", "Customer is InActive", logPrefix)
	}

	result, response, rspCode := helper.AllBill(ownerRowId, customerId, tx, logPrefix)
	if rspCode != utils.StatusOK {
		return response, rspCode
	}

	if rspCode == utils.StatusOK {
		err = tx.Commit()
		if err != nil {
			return helper.Create500ErrorResponse("[DB ERROR 0132] Error committing transaction", "Error committing transaction:"+err.Error(), logPrefix)
		}
		response = result
		utils.Logger.Info(logPrefix, "Transaction committed")
	}

	return response, rspCode

}
