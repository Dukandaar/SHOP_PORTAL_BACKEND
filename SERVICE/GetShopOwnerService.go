package service

import (
	helper "SHOP_PORTAL_BACKEND/HELPER"
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"database/sql"

	database "SHOP_PORTAL_BACKEND/DATABASE"
)

func GetShopOwner(regId string) (interface{}, int) {

	var response interface{}
	rspCode := utils.StatusOK

	var shopName string
	var ownerName string
	var phNo string
	var regDate string
	var address string
	var remarks string
	var gold float32
	var silver float32
	var cash float32

	DB := database.ConnectDB()
	defer DB.Close()

	ServiceQuery := database.GetShopOwnerData()
	err := DB.QueryRow(ServiceQuery, regId).Scan(&shopName, &ownerName, &phNo, &regDate, &address, &remarks, &gold, &silver, &cash)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.Logger.Info("Data for reg_id", regId, "does not exists")
			response, rspCode = helper.CreateErrorResponse("404001", "Data for reg_id "+regId+" does not exists")
		} else {
			utils.Logger.Error(err.Error())
			response, rspCode = helper.CreateErrorResponse("500001", "Error in getting row")
		}
	}

	if rspCode == utils.StatusOK {
		response = structs.ShopOwnerDetailsSubResponse{
			ShopName:  shopName,
			OwnerName: ownerName,
			PhNo:      phNo,
			RegDate:   regDate,
			Address:   address,
			Remarks:   remarks,
			Gold:      gold,
			Silver:    silver,
			Cash:      cash,
		}
	}

	return response, rspCode
}
