package service

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	helper "SHOP_PORTAL_BACKEND/HELPER"
	maths "SHOP_PORTAL_BACKEND/MATHS"
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"database/sql"
	"strconv"
	"time"
)

func PostCustomer(reqBody structs.Customer, OwnerRegId string) (interface{}, int) {
	var response interface{}
	rspCode := utils.StatusOK

	DB := database.ConnectDB()
	defer DB.Close()

	tx, err := DB.Begin()
	if err != nil {
		return helper.SetErrorResponse("Error starting transaction", "Error starting transaction:"+err.Error())
	}

	defer func() {
		if r := recover(); r != nil || err != nil {
			utils.Logger.Error("Panic occurred during transaction:", r, err)
		}
		if rspCode != utils.StatusOK { // Rollback only if there was an error
			tx.Rollback()
		}
	}()

	ServiceQuery := database.GetOwnerRowId()
	var ownerRowId string
	err = tx.QueryRow(ServiceQuery, OwnerRegId).Scan(&ownerRowId)
	if err != nil {
		return helper.SetErrorResponse("Error in getting row", "Error getting owner row ID:"+err.Error())
	}

	ServiceQuery = database.CheckCustomerPresent()
	var rowId int
	var isActive string
	var customerRegId string

	err = tx.QueryRow(ServiceQuery, reqBody.Name, reqBody.CompanyName, reqBody.PhNo).Scan(&rowId, &isActive, &customerRegId)
	if err == sql.ErrNoRows { // New customer
		ServiceQuery = database.InsertCustomerData()
		date, _ := time.Parse("2006-01-02", reqBody.RegDate)
		regId := maths.GenerateRegID()
		err = tx.QueryRow(ServiceQuery, reqBody.Name, reqBody.CompanyName, regId, date, reqBody.PhNo, reqBody.Address, time.Now(), time.Now()).Scan(&rowId)
		if err != nil {
			helper.SetErrorResponse("Error in inserting row in customer table", "Error inserting customer data:"+err.Error())
		} else {
			utils.Logger.Info("Inserted customer with reg_id:", regId)
			ServiceQuery = database.InsertOwnerCustomerData()
			_, err = tx.Exec(ServiceQuery, ownerRowId, rowId, utils.ACTIVE_YES, reqBody.Remarks)
			if err != nil {
				helper.SetErrorResponse("Error in inserting row in owner_customer table", "Error inserting owner_customer data:"+err.Error())
			} else {
				response, rspCode = helper.CreateSuccessResponse("Customer Added Successfully " + "with reg_id [" + regId + "] and row_id [" + strconv.Itoa(rowId) + "]") // Capture both values
			}
		}
	} else if err != nil { // Database error checking for existing customer
		return helper.SetErrorResponse("Error in checking row", "Error checking for existing customer:"+err.Error())
	} else { // Existing customer
		if isActive == utils.ACTIVE_YES {
			utils.Logger.Info("Same customer data exists")
			response, rspCode = helper.CreateErrorResponse("400008", "Same data exists") // Capture both values
			return response, rspCode
		} else { // Activate existing customer
			ServiceQuery = database.UpdateOwnerCustomerData()
			_, err = tx.Exec(ServiceQuery, utils.ACTIVE_YES, reqBody.Remarks, ownerRowId, rowId)
			if err != nil {
				return helper.SetErrorResponse("Error in updating active status", "Error updating customer status:"+err.Error())
			} else {
				response, rspCode = helper.CreateSuccessResponse("Customer Activated Successfully") // Capture both values
			}
		}
	}

	if rspCode == utils.StatusOK {
		err = tx.Commit()
		if err != nil {
			return helper.SetErrorResponse("Error committing transaction", "Error committing transaction:"+err.Error())
		}
	}

	return response, rspCode
}
