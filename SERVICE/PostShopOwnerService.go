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

func PostShopOwner(reqBody structs.ShopOwner) (interface{}, int) {

	var response interface{}
	rspCode := utils.StatusOK

	DB := database.ConnectDB()
	defer DB.Close()

	// check if owner with same shop_name, owner_name, phone_no is present
	ServiceQuery := database.CheckOwnerPresent()
	var rowId int
	var reg_id string
	var isActive string

	err := DB.QueryRow(ServiceQuery, reqBody.OwnerName, reqBody.ShopName, reqBody.PhNo).Scan(&rowId, &reg_id, &isActive)
	if err == sql.ErrNoRows {

		// insert row
		ServiceQuery = database.InsertShopOwnerData()
		date, _ := time.Parse("2006-01-02", reqBody.RegDate)
		regId := maths.GenerateRegID() // todo in future it should be unique by checking in DB
		key, errMsg := maths.GenerateKey(regId)

		if errMsg != utils.NULL_STRING {
			response, rspCode = helper.CreateErrorResponse("500001", errMsg)
			utils.Logger.Error(errMsg)

		} else {
			_, err = DB.Exec(ServiceQuery, reqBody.ShopName, reqBody.OwnerName, regId, reqBody.PhNo, utils.ACTIVE_YES, date, reqBody.Address, reqBody.Remarks, key, time.Now(), time.Now())
			if err != nil {
				response, rspCode = helper.CreateErrorResponse("500001", "Error in inserting Shop Owner Data")
				utils.Logger.Error(err.Error())

			} else {
				response, rspCode = helper.CreateSuccessResponse("Shop Owner Added Successfully with regId: " + regId)
			}
		}

	} else if err != nil {
		response, rspCode = helper.CreateErrorResponse("500001", "Error in checking Shop Owner Data")
		utils.Logger.Error(err.Error())

	} else {
		if isActive == utils.ACTIVE_YES {
			response, rspCode = helper.CreateErrorResponse("400008", "Shop Owner with same details is already present")
			utils.Logger.Error("Shop Owner with same details is already present")
		} else {
			ServiceQuery = database.ToggleShopOwnerActiveStatus() // update shop owner details
			_, err = DB.Exec(ServiceQuery, utils.ACTIVE_YES, time.Now(), rowId)
			if err != nil {
				response, rspCode = helper.CreateErrorResponse("500001", "Error in activating Shop Owner")
				utils.Logger.Error(err.Error())
			} else {
				response, rspCode = helper.CreateSuccessResponse("Shop Owner Activated Successfully with regId: " + reg_id)
			}
		}
	}

	return response, rspCode
}
