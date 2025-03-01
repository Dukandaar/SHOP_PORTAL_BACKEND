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

func PostShopOwner(reqBody structs.ShopOwner, logPrefix string) (interface{}, int) {
	var response interface{}
	rspCode := utils.StatusOK

	DB := database.DB

	tx, err := DB.Begin()
	if err != nil {
		return helper.Create500ErrorResponse("Error in starting transaction", "Error starting transaction:"+err.Error(), logPrefix)
	}
	defer tx.Rollback()

	serviceQuery := database.CheckOwnerPresent()
	var rowID int
	var key string
	var ownerRedID string
	var isActive string

	err = tx.QueryRow(serviceQuery, reqBody.OwnerName, reqBody.ShopName, reqBody.PhoneNo).Scan(&rowID, &ownerRedID, &isActive)
	if err == sql.ErrNoRows { // Shop owner NOT found (proceed with insertion)
		serviceQuery = database.InsertShopOwnerData()

		date, _ := time.Parse("2006-01-02", reqBody.RegDate)
		ownerRedID := maths.GenerateShopRegID(tx)
		if ownerRedID == utils.NULL_STRING {
			return helper.Create500ErrorResponse("Error in generating Shop Owner Registration ID", "Error generating Shop Owner Registration ID.", logPrefix)
		}
		key, errMsg := maths.GenerateKey(ownerRedID) // Key generation *before* transaction

		if errMsg != utils.NULL_STRING {
			return helper.Create500ErrorResponse("Error in generating key.", "Key generation error : "+errMsg, logPrefix)
		}

		err = tx.QueryRow(serviceQuery, reqBody.ShopName, reqBody.OwnerName, ownerRedID, reqBody.GstIN, reqBody.PhoneNo, utils.ACTIVE_YES, date, reqBody.Address, reqBody.Remarks, key, time.Now(), time.Now()).Scan(&rowID)
		if err != nil {
			return helper.Create500ErrorResponse("Error in inserting Shop Owner Data", "Error inserting Shop Owner Data: "+err.Error(), logPrefix)
		}
		utils.Logger.Info(logPrefix, "Shop Owner Data inserted successfully with row ID ", rowID)

		// insert intial data in balance table
		serviceQuery = database.InsertOwnerBalanceData()
		_, err = tx.Exec(serviceQuery, rowID, utils.NULL_FLOAT, utils.NULL_FLOAT, utils.NULL_FLOAT, time.Now(), time.Now())
		if err != nil {
			return helper.Create500ErrorResponse("Error in inserting Shop Owner Balance Data", "Error inserting Shop Owner Balance Data: "+err.Error(), logPrefix)
		}
		utils.Logger.Info(logPrefix, "Shop Owner Balance Data inserted successfully.")

		// insert initial bill count to 0 for owner
		serviceQuery = database.InsertOwnerBillCount()
		_, err = tx.Exec(serviceQuery, rowID, 0, time.Now())
		if err != nil {
			return helper.Create500ErrorResponse("Error in inserting Shop Owner Bill Count Data", "Error inserting Shop Owner Bill Count Data: "+err.Error(), logPrefix)
		}
		utils.Logger.Info(logPrefix, "Shop Owner Bill Count Data inserted successfully.")
		response, rspCode = helper.CreatePostOwnerResponse(key, ownerRedID, "Owner added successfully", logPrefix)

	} else if err != nil { // Error checking if owner exists
		return helper.Create500ErrorResponse("Error checking Shop Owner Data", "Error checking Shop Owner Data: "+err.Error(), logPrefix)
	} else {
		// Shop owner ALREADY exists (proceed with update if not active)
		if isActive == utils.ACTIVE_YES {
			response, rspCode = helper.CreateErrorResponse("400009", "Shop Owner with same details is already present", logPrefix)
			return response, rspCode
		}

		serviceQuery = database.UpdateShopOwnerData()
		_, err = tx.Exec(serviceQuery, reqBody.ShopName, reqBody.OwnerName, reqBody.GstIN, reqBody.PhoneNo, utils.ACTIVE_YES, reqBody.RegDate, reqBody.Address, reqBody.Remarks, time.Now(), ownerRedID)
		if err != nil {
			return helper.Create500ErrorResponse("Error in updating Shop Owner data", "rror in updating Shop Owner data: "+err.Error(), logPrefix)
		}
		response, rspCode = helper.CreatePostOwnerResponse(key, ownerRedID, "Owner updated successfully", logPrefix)
	}

	if rspCode == utils.StatusOK {
		err = tx.Commit()
		if err != nil {
			return helper.Create500ErrorResponse("Error in committing transaction", "Error in committing transaction: "+err.Error(), logPrefix)
		}
		utils.Logger.Info(logPrefix, "Transaction committed successfully")
	}

	return response, rspCode
}
