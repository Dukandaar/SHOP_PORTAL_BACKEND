package service

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	helper "SHOP_PORTAL_BACKEND/HELPER"
	maths "SHOP_PORTAL_BACKEND/MATHS"
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"database/sql"
	"time"
)

func PostCustomer(reqBody structs.Customer, ownerRegID string, logPrefix string) (interface{}, int) {
	var response interface{}
	rspCode := utils.StatusOK

	DB := database.DB

	tx, err := DB.Begin()
	if err != nil {
		return helper.Create500ErrorResponse("[DB ERROR 0025] Error starting transaction", "Error starting transaction:"+err.Error(), logPrefix)
	}
	defer tx.Rollback()

	ownerRowID, err := helper.GetOwnerId(ownerRegID, tx)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.CreateErrorResponse("404001", "Owner not found", logPrefix)
		}
		return helper.Create500ErrorResponse("[DB ERROR 0026] Error getting owner row ID", "Error getting owner row ID:"+err.Error(), logPrefix)
	}

	ServiceQuery := database.CheckCustomerPresent()
	var rowID int
	var isActive string
	var customerRegID string

	err = tx.QueryRow(ServiceQuery, ownerRowID, reqBody.Name, reqBody.ShopName, reqBody.PhoneNo).Scan(&rowID, &isActive, &customerRegID)
	if err == sql.ErrNoRows { // Add New customer
		ServiceQuery = database.InsertCustomerData()
		date, _ := time.Parse("2006-01-02", reqBody.RegDate)
		regID := maths.GenerateCustomerRegID()

		if regID == utils.NULL_STRING {
			return helper.Create500ErrorResponse("Error in generating reg_id", "Error in generating reg_id", logPrefix)
		}

		err = tx.QueryRow(ServiceQuery, utils.ACTIVE_YES, ownerRowID, reqBody.Name, reqBody.ShopName, regID, date, reqBody.PhoneNo, reqBody.Address, reqBody.Remarks, reqBody.GstIN, time.Now(), time.Now()).Scan(&rowID)
		if err != nil {
			return helper.Create500ErrorResponse("[DB ERROR 0027] Error in inserting row in customer table", "Error inserting customer data: "+err.Error(), logPrefix)
		}

		utils.Logger.Info(logPrefix, "Inserted customer with reg_id:", regID)

		// insert new balance
		ServiceQuery = database.InsertCustomerBalanceData()
		_, err = tx.Exec(ServiceQuery, rowID, utils.NULL_FLOAT, utils.NULL_FLOAT, utils.NULL_FLOAT, time.Now(), time.Now())
		if err != nil {
			return helper.Create500ErrorResponse("[DB ERROR 0028] Error inserting Shop Owner Balance Data", "Error inserting Shop Owner Balance Data:"+err.Error(), logPrefix)
		}

		response, rspCode = helper.CreatePostCustomerResponse(regID, "Customer Added Successfully", logPrefix)

	} else if err != nil { // Database error checking for existing customer
		return helper.Create500ErrorResponse("[DB ERROR 0029] Error in checking row", "Error checking for existing customer: "+err.Error(), logPrefix)
	} else { // Existing customer
		if isActive == utils.ACTIVE_YES {
			utils.Logger.Info(logPrefix, "Same customer data exists")
			return helper.CreateErrorResponse("400009", "Same data exists with reg_id: "+customerRegID, logPrefix)
		}

		// Activate existing customer
		ServiceQuery = database.UpdateOwnerCustomerData()
		_, err = tx.Exec(ServiceQuery, utils.ACTIVE_YES, reqBody.Remarks, customerRegID)
		if err != nil {
			return helper.Create500ErrorResponse("[DB ERROR 0030] Error in updating active status", "Error updating customer status: "+err.Error(), logPrefix)
		}

		response, rspCode = helper.CreatePostCustomerResponse(customerRegID, "Customer Activated Successfully", logPrefix)
	}

	if rspCode == utils.StatusOK {
		err = tx.Commit()
		if err != nil {
			return helper.Create500ErrorResponse("[DB ERROR 0031] Error committing transaction", "Error committing transaction:"+err.Error(), logPrefix)
		}
		utils.Logger.Info(logPrefix, "Transaction committed")
	}

	return response, rspCode
}
