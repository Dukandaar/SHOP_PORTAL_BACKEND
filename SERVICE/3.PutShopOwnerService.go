package service

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	helper "SHOP_PORTAL_BACKEND/HELPER"
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"database/sql"
	"time"
)

func PutShopOwner(reqBody structs.ShopOwner, OwnerRegId string, logPrefix string) (interface{}, int) {

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
	var reg_id string
	var isActive string

	err = tx.QueryRow(ServiceQuery, reqBody.OwnerName, reqBody.ShopName, reqBody.PhoneNo).Scan(&rowId, &reg_id, &isActive)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.Logger.Info(logPrefix, "Data for reg_id ", OwnerRegId, " does not exist")
			response, rspCode = helper.CreateErrorResponse("404001", "Data for reg_id "+OwnerRegId+" does not exist", logPrefix)
			return response, rspCode
		}
		response, rspCode = helper.Set500ErrorResponse("Error in getting row", "Error in getting row:"+err.Error(), logPrefix)
		return response, rspCode
	}

	if isActive != utils.NULL_STRING {
		if reg_id == OwnerRegId {
			utils.Logger.Info(logPrefix, "Data with reg_id ", OwnerRegId, " exists") // update row
		} else {
			utils.Logger.Info(logPrefix, "Same data with reg_id ", OwnerRegId, " exists")
			response, rspCode = helper.CreateErrorResponse("400009", "Same data with reg_id "+OwnerRegId+" exists", logPrefix)
			return response, rspCode
		}
	}

	// update details in DB
	ServiceQuery = database.UpdateShopOwnerData()
	_, err = tx.Exec(ServiceQuery, reqBody.ShopName, reqBody.OwnerName, reqBody.GstIN, reqBody.PhoneNo, utils.ACTIVE_YES, reqBody.RegDate, reqBody.Address, reqBody.Remarks, time.Now(), OwnerRegId)
	if err != nil {
		response, rspCode = helper.Set500ErrorResponse("Error in updating row", "Error in updating row:"+err.Error(), logPrefix)
	} else {
		response, rspCode = helper.CreateSuccessResponse("Updated owner with reg_id : " + OwnerRegId)
	}

	if rspCode == utils.StatusOK {
		err = tx.Commit()
		if err != nil {
			return helper.Set500ErrorResponse("Error committing transaction", "Error committing transaction:"+err.Error(), logPrefix)
		}
		utils.Logger.Info(logPrefix, "Transaction committed")
	}

	return response, rspCode
}
