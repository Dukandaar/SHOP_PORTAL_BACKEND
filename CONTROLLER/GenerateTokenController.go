package controller

import (
	helper "SHOP_PORTAL_BACKEND/HELPER"
	service "SHOP_PORTAL_BACKEND/SERVICE"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	validator "SHOP_PORTAL_BACKEND/VALIDATOR"

	"github.com/kataras/iris/v12"
)

func GenerateToken(ctx iris.Context) {
	logPrefix := ctx.Values().Get("logPrefix").(string)

	headers := utils.ReadHeader(ctx, utils.GenerateTokenHeaders)
	reqBody, bodyError := utils.ReadGenerateTokenReqBody(ctx)
	utils.Logger.Info(headers, reqBody)

	headerError := validator.ValidateHeader(utils.GenerateTokenHeaders, headers)
	reqBodyError := validator.ValidateGenerateTokenReqBody(&reqBody, bodyError)

	errMsg := helper.CheckError(headerError, utils.NULL_STRING, reqBodyError)

	var response interface{}
	rspCode := utils.StatusOK

	if errMsg != utils.NULL_STRING {
		response, rspCode = helper.CreateErrorResponse("400001", errMsg)
	} else {
		response, rspCode = service.GenerateToken(reqBody)
	}

	utils.Logger.Info(logPrefix, response)

	ctx.ResponseWriter().WriteHeader(rspCode)
	ctx.JSON(response)
	utils.Logger.Info(logPrefix + " Request Completed.")
}
