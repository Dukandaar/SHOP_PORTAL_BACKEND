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

func PostShopOwner(reqBody structs.ShopOwner, logPrefix string) (interface{}, int) {
	var response interface{}
	rspCode := utils.StatusOK

	DB := database.DB

	tx, err := DB.Begin()
	if err != nil {
		return helper.Set500ErrorResponse("Error starting transaction", "Error starting transaction:"+err.Error(), logPrefix)
	}

	defer func() {
		if r := recover(); r != nil || err != nil {
			utils.Logger.Error(logPrefix, "Panic occurred during transaction:", r, err)
			tx.Rollback()
		}
	}()

	ServiceQuery := database.CheckOwnerPresent()
	var rowId int
	var key string
	var reg_id string
	var isActive string

	err = tx.QueryRow(ServiceQuery, reqBody.OwnerName, reqBody.ShopName, reqBody.PhoneNo).Scan(&rowId, &reg_id, &isActive)
	if err == sql.ErrNoRows { // Shop owner NOT found (proceed with insertion)
		ServiceQuery = database.InsertShopOwnerData()

		date, _ := time.Parse("2006-01-02", reqBody.RegDate)
		regId := maths.GenerateShopRegID(tx)
		if regId == utils.NULL_STRING {
			return helper.Set500ErrorResponse("Error generating Shop Owner Registration ID", "Error generating Shop Owner Registration ID", logPrefix)
		}
		key, errMsg := maths.GenerateKey(regId) // Key generation *before* transaction

		if errMsg != utils.NULL_STRING {
			return helper.Set500ErrorResponse(errMsg, "Key generation error:"+errMsg, logPrefix)
		} else {
			err = tx.QueryRow(ServiceQuery, reqBody.ShopName, reqBody.OwnerName, regId, reqBody.GstIN, reqBody.PhoneNo, utils.ACTIVE_YES, date, reqBody.Address, reqBody.Remarks, key, time.Now(), time.Now()).Scan(&rowId)
			if err != nil {
				return helper.Set500ErrorResponse("Error inserting Shop Owner Data", "Error inserting Shop Owner Data:"+err.Error(), logPrefix)
			} else {
				// insert data in balance table
				ServiceQuery = database.InsertOwnerBalanceData()
				_, err = tx.Exec(ServiceQuery, rowId, utils.NULL_FLOAT, utils.NULL_FLOAT, utils.NULL_FLOAT, time.Now(), time.Now())
				if err != nil {
					return helper.Set500ErrorResponse("Error inserting Shop Owner Balance Data", "Error inserting Shop Owner Balance Data:"+err.Error(), logPrefix)
				}
				// insert intial bill count to 0 for owner
				ServiceQuery = database.InsertOwnerBillCount()
				_, err = tx.Exec(ServiceQuery, rowId, 0, time.Now())
				if err != nil {
					return helper.Set500ErrorResponse("Error inserting Shop Owner Bill Count Data", "Error inserting Shop Owner Bill Count Data:"+err.Error(), logPrefix)
				}
				response, rspCode = helper.CreateOwnerSuccessResponseWithIdKey("Shop Owner Added Successfully", regId, key)
			}
		}
	} else if err != nil { // Error checking if owner exists
		return helper.Set500ErrorResponse("Error checking Shop Owner Data", "Error checking Shop Owner Data: "+err.Error(), logPrefix)
	} else {
		// Shop owner ALREADY exists (proceed with update if not active)
		if isActive == utils.ACTIVE_YES {
			response, rspCode = helper.CreateErrorResponse("400009", "Shop Owner with same details is already present")
			return response, rspCode
		} else {
			ServiceQuery = database.ToggleShopOwnerActiveStatus()
			_, err = tx.Exec(ServiceQuery, utils.ACTIVE_YES, time.Now(), rowId)
			if err != nil {
				return helper.Set500ErrorResponse("Error activating Shop Owner", "Error activating Shop Owner: "+err.Error(), logPrefix)
			} else {
				response, rspCode = helper.CreateOwnerSuccessResponseWithIdKey("Shop Owner Activated Successfully", reg_id, key)
			}
		}
	}

	if rspCode == utils.StatusOK {
		err = tx.Commit()
		if err != nil {
			return helper.Set500ErrorResponse("Error committing transaction", "Error committing transaction:"+err.Error(), logPrefix)
		}
		utils.Logger.Info(logPrefix, "Transaction committed successfully")
	}

	return response, rspCode
}
