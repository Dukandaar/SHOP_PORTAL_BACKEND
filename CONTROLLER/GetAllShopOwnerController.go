package controller

import (
	helper "SHOP_PORTAL_BACKEND/HELPER"
	service "SHOP_PORTAL_BACKEND/SERVICE"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	validator "SHOP_PORTAL_BACKEND/VALIDATOR"

	"github.com/kataras/iris/v12"
)

func GetAllShopOwner(ctx iris.Context) {

	var response interface{}
	var errCodeStr string
	rspCode := utils.StatusOK

	logPrefix := ctx.Values().Get("logPrefix").(string)

	headers := utils.ReadHeader(ctx)
	reqBody, bodyError := utils.ReadAllShopOwnerBody(ctx)
	utils.Logger.Info(logPrefix, headers, reqBody)

	headerError, errCodeStr := validator.ValidateHeader(utils.GetAllShopOwnerHeaders, headers, ctx)
	if errCodeStr != utils.SUCCESS {
		response, rspCode = helper.CreateErrorResponse(errCodeStr, headerError)
		utils.Logger.Error(logPrefix, headerError)
	} else {

		reqBodyError, errCodeStr := validator.ValidateAllShopOwnerBody(&reqBody, bodyError)
		if errCodeStr != utils.SUCCESS {
			response, rspCode = helper.CreateErrorResponse(errCodeStr, reqBodyError)
			utils.Logger.Error(logPrefix, reqBodyError)
		} else {
			response, rspCode = service.GetAllShopOwner(reqBody, logPrefix)
		}
	}

	utils.Logger.Info(logPrefix, response)

	ctx.ResponseWriter().WriteHeader(rspCode)
	ctx.JSON(response)
	utils.Logger.Info(logPrefix + "Request Completed.")
}
