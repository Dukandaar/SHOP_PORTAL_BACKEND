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

func PostCustomer(reqBody structs.Customer, OwnerRegId string, logPrefix string) (interface{}, int) {
	var response interface{}
	rspCode := utils.StatusOK

	DB := database.DB

	tx, err := DB.Begin()
	if err != nil {
		return helper.Create500ErrorResponse("Error starting transaction", "Error starting transaction:"+err.Error(), logPrefix)
	}

	defer func() {
		if r := recover(); r != nil || err != nil {
			utils.Logger.Error(logPrefix, "Panic occurred during transaction:", r, err)
			tx.Rollback()
		}
	}()

	ServiceQuery := database.GetOwnerRowId()
	var ownerRowId string
	err = tx.QueryRow(ServiceQuery, OwnerRegId).Scan(&ownerRowId)
	if err != nil {
		return helper.Create500ErrorResponse("Error in getting row", "Error getting owner row ID:"+err.Error(), logPrefix)
	}

	ServiceQuery = database.CheckCustomerPresent()
	var rowId int
	var isActive string
	var customerRegId string

	err = tx.QueryRow(ServiceQuery, ownerRowId, reqBody.Name, reqBody.ShopName, reqBody.PhoneNo).Scan(&rowId, &isActive, &customerRegId)
	if err == sql.ErrNoRows { // Add New customer
		ServiceQuery = database.InsertCustomerData()
		date, _ := time.Parse("2006-01-02", reqBody.RegDate)
		regId := maths.GenerateCustomerRegID()

		if regId == utils.NULL_STRING {
			return helper.Create500ErrorResponse("Error generating reg_id", "Error generating reg_id", logPrefix)
		}

		err = tx.QueryRow(ServiceQuery, utils.ACTIVE_YES, ownerRowId, reqBody.Name, reqBody.ShopName, regId, date, reqBody.PhoneNo, reqBody.Address, reqBody.Remarks, reqBody.GstIN, time.Now(), time.Now()).Scan(&rowId)
		if err != nil {
			helper.Create500ErrorResponse("Error in inserting row in customer table", "Error inserting customer data:"+err.Error(), logPrefix)
		} else {
			utils.Logger.Info(logPrefix, "Inserted customer with reg_id:", regId)
			// insert new balance
			ServiceQuery = database.InsertCustomerBalanceData()
			_, err = tx.Exec(ServiceQuery, rowId, utils.NULL_FLOAT, utils.NULL_FLOAT, utils.NULL_FLOAT, time.Now(), time.Now())
			if err != nil {
				return helper.Create500ErrorResponse("Error inserting Shop Owner Balance Data", "Error inserting Shop Owner Balance Data:"+err.Error(), logPrefix)
			}
			response, rspCode = helper.CreateSuccessResponse(regId, "Customer Added Successfully")
		}
	} else if err != nil { // Database error checking for existing customer
		return helper.Create500ErrorResponse("Error in checking row", "Error checking for existing customer:"+err.Error(), logPrefix)
	} else { // Existing customer
		if isActive == utils.ACTIVE_YES {
			utils.Logger.Info(logPrefix, "Same customer data exists")
			response, rspCode = helper.CreateErrorResponse("400009", "Same data exists with reg_id: "+customerRegId, logPrefix)
			return response, rspCode
		} else { // Activate existing customer
			ServiceQuery = database.UpdateOwnerCustomerData()
			_, err = tx.Exec(ServiceQuery, utils.ACTIVE_YES, reqBody.Remarks, customerRegId)
			if err != nil {
				return helper.Create500ErrorResponse("Error in updating active status", "Error updating customer status:"+err.Error(), logPrefix)
			} else {
				response, rspCode = helper.CreateSuccessResponse("Customer Activated Successfully", "Customer Activated Successfully")
			}
		}
	}

	if rspCode == utils.StatusOK {
		err = tx.Commit()
		if err != nil {
			return helper.Create500ErrorResponse("Error committing transaction", "Error committing transaction:"+err.Error(), logPrefix)
		}
	}

	return response, rspCode
}
