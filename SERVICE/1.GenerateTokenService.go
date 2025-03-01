package service

import (
	helper "SHOP_PORTAL_BACKEND/HELPER"
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
)

func GenerateToken(reqBody structs.GenerateToken, logPrefix string) (interface{}, int) {

	var response interface{}
	rspCode := utils.StatusOK

	token, err := helper.GenerateJWT(reqBody.Key)
	if err != nil {
		utils.Logger.Error(logPrefix, err.Error())
		response, rspCode = helper.Create500ErrorResponse("Error in generating Token", err.Error(), logPrefix)
	} else {
		utils.Logger.Info(logPrefix, "Generated JWT: ", token)
		response, rspCode = helper.CreateGenerateTokenResponse(token, "Token created successfully", logPrefix)
	}

	return response, rspCode
}
