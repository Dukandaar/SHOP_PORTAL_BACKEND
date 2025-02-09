package service

import (
	helper "SHOP_PORTAL_BACKEND/HELPER"
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
)

func GenerateToken(reqBody structs.GenerateToken) (interface{}, int) {

	var response interface{}
	rspCode := utils.StatusOK

	token, err := helper.GenerateJWT(reqBody.Key)
	if err != nil {
		utils.Logger.Error(err.Error())
		response, rspCode = helper.CreateErrorResponse("500001", "Error in generating Token")
	} else {
		utils.Logger.Info("Generated JWT:", token)
		response, rspCode = helper.CreateSuccessResponse(token)
	}

	return response, rspCode
}
