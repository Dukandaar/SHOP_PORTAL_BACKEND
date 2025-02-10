package service

import (
	helper "SHOP_PORTAL_BACKEND/HELPER"
	maths "SHOP_PORTAL_BACKEND/MATHS"
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"database/sql"
	"strconv"
	"time"

	database "SHOP_PORTAL_BACKEND/DATABASE"
)

func PostCustomer(reqBody structs.Customer, OwnerRegId string) (interface{}, int) {

	var response interface{}
	rspCode := utils.StatusOK

	DB := database.ConnectDB()
	defer DB.Close()

	// fetch row_id of owner with same reg_id is present
	ServiceQuery := database.GetOwnerRowId()
	var ownerRowId string
	err := DB.QueryRow(ServiceQuery, OwnerRegId).Scan(&ownerRowId)
	if err != nil {
		utils.Logger.Error(err.Error())
		response, rspCode = helper.CreateErrorResponse("500001", "Error in getting row")
		return response, rspCode
	}

	// check if customer with same company_name, name, phone_no is present
	ServiceQuery = database.CheckCustomerPresent()
	var rowId int
	var isActive string

	err = DB.QueryRow(ServiceQuery, reqBody.Name, reqBody.CompanyName, reqBody.PhNo).Scan(&rowId, &isActive)
	if err == sql.ErrNoRows {

		// insert row
		ServiceQuery = database.InsertCustomerData()
		date, _ := time.Parse("2006-01-02", reqBody.RegDate)
		regId := maths.GenerateRegID() // todo in future it should be unique by checking in DB
		err = DB.QueryRow(ServiceQuery, reqBody.Name, reqBody.CompanyName, regId, date, reqBody.PhNo, reqBody.Address, time.Now(), time.Now()).Scan(&rowId)
		if err != nil {
			utils.Logger.Error(err.Error())
			response, rspCode = helper.CreateErrorResponse("500001", "Error in inserting row in customer table")
		} else {
			utils.Logger.Info("Inserted row with reg_id", regId)
			ServiceQuery = database.InsertOwnerCustomerData()
			_, err = DB.Exec(ServiceQuery, rowId, ownerRowId, utils.ACTIVE_YES, reqBody.Remarks)
			if err != nil {
				utils.Logger.Error(err.Error())
				response, rspCode = helper.CreateErrorResponse("500001", "Error in inserting row in owner_customer table")
			} else {
				response, rspCode = helper.CreateSuccessResponse("Customer Added Successfully" + "with reg_id [" + regId + "] and row_id [" + strconv.Itoa(rowId) + "]")
			}
		}
	} else if err != nil {
		utils.Logger.Error(err.Error())
		response, rspCode = helper.CreateErrorResponse("500001", "Error in checking row")
	} else {

		if isActive == utils.ACTIVE_YES {
			utils.Logger.Info("Same data exists")
			response, rspCode = helper.CreateErrorResponse("400008", "Same data exists")
		} else {
			// update row
			ServiceQuery = database.UpdateOwnerCustomerData()
			_, err = DB.Exec(ServiceQuery, utils.ACTIVE_YES, reqBody.Remarks)
			if err != nil {
				utils.Logger.Error(err.Error())
				response, rspCode = helper.CreateErrorResponse("500001", "Error in updating active status")
			} else {
				utils.Logger.Info("Updated row")
				response, rspCode = helper.CreateSuccessResponse("Customer Activated Successfully")
			}
		}
	}

	return response, rspCode
}
