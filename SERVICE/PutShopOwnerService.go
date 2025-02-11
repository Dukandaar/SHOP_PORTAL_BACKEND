package service

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	helper "SHOP_PORTAL_BACKEND/HELPER"
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"database/sql"
	"time"
)

func PutShopOwner(reqBody structs.ShopOwner, OwnerRegId string) (interface{}, int) {

	var response interface{}
	rspCode := utils.StatusOK

	DB := database.ConnectDB()
	defer DB.Close()

	ServiceQuery := database.CheckOwnerPresent()
	var rowId int
	var reg_id string
	var isActive string

	err := DB.QueryRow(ServiceQuery, reqBody.OwnerName, reqBody.ShopName, reqBody.PhNo).Scan(&rowId, &reg_id, &isActive)
	if err != nil {
		if err != sql.ErrNoRows {
			utils.Logger.Error(err.Error())
			response, rspCode = helper.CreateErrorResponse("500001", "Error in checking row")
			return response, rspCode
		}
	}

	if isActive != utils.NULL_STRING {
		if reg_id == OwnerRegId {
			utils.Logger.Info("Row with reg_id ", OwnerRegId, " exists") // update row
		} else {
			utils.Logger.Info("Same data with reg_id", OwnerRegId, "exists")
			response, rspCode = helper.CreateErrorResponse("400008", "Same data with reg_id "+OwnerRegId+" exists")
			return response, rspCode
		}
	}

	// update details in DB
	ServiceQuery = database.UpdateShopOwnerData()
	_, err = DB.Exec(ServiceQuery, reqBody.ShopName, reqBody.OwnerName, reqBody.PhNo, utils.ACTIVE_YES, reqBody.RegDate, reqBody.Address, reqBody.Remarks, time.Now(), OwnerRegId)
	if err != nil {
		utils.Logger.Error(err.Error())
		response, rspCode = helper.CreateErrorResponse("500001", "Error in updating row")
	} else {
		response, rspCode = helper.CreateSuccessResponse("Updated row with reg_id : " + OwnerRegId)
	}

	return response, rspCode
}
