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
		rsp := utils.CodeMap["500001"]
		rsp.Description = "Error in Inserting Shop Owner Data"
		response = helper.CreateErrorResponse(rsp)
		utils.Logger.Error(rsp.Description, err.Error())
	} else {
		response = helper.CreateSuccessResponse("Shop Owner Added Successfully")
		utils.Logger.Info(response)
	}

	return response
}
