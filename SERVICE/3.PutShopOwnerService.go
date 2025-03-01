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
		return helper.Create500ErrorResponse("[DB ERROR 00015] Error starting transaction", "Error starting transaction:"+err.Error(), logPrefix)
	}

	defer tx.Rollback()

	ServiceQuery := database.CheckOwnerPresent()
	var rowId int
	var reg_id string
	var isActive string

	err = tx.QueryRow(ServiceQuery, reqBody.OwnerName, reqBody.ShopName, reqBody.PhoneNo).Scan(&rowId, &reg_id, &isActive)
	if err != nil && err != sql.ErrNoRows {
		return helper.Create500ErrorResponse("[DB ERROR 00016] Error in getting row", "Error in getting row:"+err.Error(), logPrefix)
	}

	if isActive != utils.NULL_STRING {
		if reg_id == OwnerRegId {
			utils.Logger.Info(logPrefix, "Data with reg_id ", OwnerRegId, " exists") // update row
		} else {
			return helper.CreateErrorResponse("400009", "Another owner with same name, shopname and phone_no. exists : reg_id : "+OwnerRegId, logPrefix)
		}
	}

	// update details in DB
	ServiceQuery = database.UpdateShopOwnerData()
	_, err = tx.Exec(ServiceQuery, reqBody.ShopName, reqBody.OwnerName, reqBody.GstIN, reqBody.PhoneNo, utils.ACTIVE_YES, reqBody.RegDate, reqBody.Address, reqBody.Remarks, time.Now(), OwnerRegId)
	if err != nil {
		return helper.Create500ErrorResponse("[DB ERROR 00017] Error in updating row", "Error in updating row:"+err.Error(), logPrefix)
	} else {
		response, rspCode = helper.CreateSuccessResponse("Updated Successfully.", "Updated owner with reg_id : "+OwnerRegId, logPrefix)
	}

	if rspCode == utils.StatusOK {
		err = tx.Commit()
		if err != nil {
			return helper.Create500ErrorResponse("[DB ERROR 00018] Error committing transaction", "Error committing transaction:"+err.Error(), logPrefix)
		}
		utils.Logger.Info(logPrefix, "Transaction committed")
	}

	return response, rspCode
}
