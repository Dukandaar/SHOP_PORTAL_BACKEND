package controller

import (
	helper "SHOP_PORTAL_BACKEND/HELPER"
	service "SHOP_PORTAL_BACKEND/SERVICE"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	validator "SHOP_PORTAL_BACKEND/VALIDATOR"

	"github.com/kataras/iris/v12"
)

func GenerateToken(ctx iris.Context) {

	var response interface{}
	var errCodeStr string
	rspCode := utils.StatusOK

	logPrefix := ctx.Values().Get("logPrefix").(string)

	headers := utils.ReadHeader(ctx, utils.GenerateTokenHeaders)
	reqBody, bodyError := utils.ReadGenerateTokenReqBody(ctx)
	utils.Logger.Info(headers, reqBody)

	headerError, errCodeStr := validator.ValidateHeader(utils.GenerateTokenHeaders, headers)
	if errCodeStr != utils.SUCCESS {
		response, rspCode = helper.CreateErrorResponse(errCodeStr, headerError)
		utils.Logger.Error(headerError)
	} else {
		reqBodyError, errCodeStr := validator.ValidateGenerateTokenReqBody(&reqBody, bodyError)
		if errCodeStr != utils.SUCCESS {
			response, rspCode = helper.CreateErrorResponse(errCodeStr, reqBodyError)
			utils.Logger.Error(reqBodyError)
		} else {
			response, rspCode = service.GenerateToken(reqBody)
		}
	}

	utils.Logger.Info(logPrefix, response)

	ctx.ResponseWriter().WriteHeader(rspCode)
	ctx.JSON(response)
	utils.Logger.Info(logPrefix + " Request Completed.")
}
