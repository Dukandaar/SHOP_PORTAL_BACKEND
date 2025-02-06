package service

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	helper "SHOP_PORTAL_BACKEND/HELPER"
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
)

func PostShopOwner(reqBody structs.ShopOwner) interface{} {

	var response interface{}

	DB := database.ConnectDB()
	defer DB.Close()

	ServiceQuery := database.InsertShopOwnerData()
	_, err := DB.Exec(ServiceQuery, reqBody.ShopName, reqBody.OwnerName, reqBody.RegDate, reqBody.PhNo, reqBody.Address, utils.ACTIVE_YES, reqBody.Remarks)
	if err != nil {
		utils.Logger.Error("Error in Inserting Shop Owner Data")
		rsp := utils.CodeMap["500001"]
		response = helper.CreateErrorResponse(rsp, "Error in Inserting Shop Owner Data")
	} else {
		response = helper.CreateSuccessResponse("Shop Owner Added Successfully")
	}

	return response
}
