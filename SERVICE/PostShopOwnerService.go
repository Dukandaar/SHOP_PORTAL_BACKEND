package service

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	helper "SHOP_PORTAL_BACKEND/HELPER"
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"time"
)

func PostShopOwner(reqBody structs.ShopOwner) (interface{}, int) {

	var response interface{}
	rspCode := utils.StatusOK

	DB := database.ConnectDB()
	defer DB.Close()

	ServiceQuery := database.InsertShopOwnerData()
	date, _ := time.Parse("2006-01-02", reqBody.RegDate)
	_, err := DB.Exec(ServiceQuery, reqBody.ShopName, reqBody.OwnerName, date, reqBody.PhNo, reqBody.Address, utils.ACTIVE_YES, reqBody.Remarks)
	if err != nil {
		rsp := utils.CodeMap["500001"]
		rsp.Description = "Error in inserting Shop Owner Data"
		response = helper.CreateErrorResponse(rsp)
		rspCode = utils.StatusInternalServerError
		utils.Logger.Error(rsp.Description, err.Error())
	} else {
		response = helper.CreateSuccessResponse("Shop Owner Added Successfully")
		rspCode = utils.StatusOK
	}

	return response, rspCode
}
