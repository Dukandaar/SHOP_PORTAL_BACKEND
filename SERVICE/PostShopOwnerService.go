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

// Helper function to set error response and log (NO ROLLBACK HERE)

func PostShopOwner(reqBody structs.ShopOwner) (interface{}, int) {
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
		if rspCode != utils.StatusOK {
			tx.Rollback()
		}
	}()

	ServiceQuery := database.CheckOwnerPresent()
	var rowId int
	var reg_id string
	var isActive string

	err = tx.QueryRow(ServiceQuery, reqBody.OwnerName, reqBody.ShopName, reqBody.PhNo).Scan(&rowId, &reg_id, &isActive)
	if err == sql.ErrNoRows { // Shop owner NOT found (proceed with insertion)
		ServiceQuery = database.InsertShopOwnerData()
		date, _ := time.Parse("2006-01-02", reqBody.RegDate)
		regId := maths.GenerateRegID()
		key, errMsg := maths.GenerateKey(regId) // Key generation *before* transaction

		if errMsg != utils.NULL_STRING {
			return helper.SetErrorResponse(errMsg, "Key generation error:"+errMsg)
		} else {
			_, err = tx.Exec(ServiceQuery, reqBody.ShopName, reqBody.OwnerName, regId, reqBody.PhNo, utils.ACTIVE_YES, date, reqBody.Address, reqBody.Remarks, key, time.Now(), time.Now())
			if err != nil {
				return helper.SetErrorResponse("Error inserting Shop Owner Data", "Error inserting Shop Owner Data:"+err.Error())
			} else {
				response, rspCode = helper.CreateSuccessResponse("Shop Owner Added Successfully with regId: " + regId)
			}
		}
	} else if err != nil { // Error checking if owner exists
		return helper.SetErrorResponse("Error checking Shop Owner Data", "Error checking Shop Owner Data:"+err.Error())
	} else { // Shop owner ALREADY exists (proceed with update if not active)
		if isActive == utils.ACTIVE_YES {
			response, rspCode = helper.CreateErrorResponse("400008", "Shop Owner with same details is already present")
			return response, rspCode
		} else {
			ServiceQuery = database.ToggleShopOwnerActiveStatus()
			_, err = tx.Exec(ServiceQuery, utils.ACTIVE_YES, time.Now(), rowId)
			if err != nil {
				return helper.SetErrorResponse("Error activating Shop Owner", "Error activating Shop Owner:"+err.Error())
			} else {
				response, rspCode = helper.CreateSuccessResponse("Shop Owner Activated Successfully with regId: " + reg_id)
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
