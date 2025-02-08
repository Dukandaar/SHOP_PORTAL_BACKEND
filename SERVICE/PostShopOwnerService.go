package service

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	helper "SHOP_PORTAL_BACKEND/HELPER"
	maths "SHOP_PORTAL_BACKEND/MATHS"
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
	regId := maths.GenerateRegID() // todo in future it should be unique by checking in DB

	_, err := DB.Exec(ServiceQuery, reqBody.ShopName, reqBody.OwnerName, regId, reqBody.PhNo, utils.ACTIVE_YES, date, reqBody.Address, reqBody.Remarks, time.Now(), time.Now())
	if err != nil {
		rsp := utils.CodeMap["500001"]
		rsp.Description = "Error in inserting Shop Owner Data"
		response = helper.CreateErrorResponse(rsp)
		rspCode = utils.StatusInternalServerError
		utils.Logger.Error(rsp.Description, err.Error())
	} else {
		response = helper.CreateSuccessResponse("Shop Owner Added Successfully with regId: " + regId)
		rspCode = utils.StatusOK
	}

	return response, rspCode
}
