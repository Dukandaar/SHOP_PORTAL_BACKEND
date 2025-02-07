package controller

import (
	helper "SHOP_PORTAL_BACKEND/HELPER"
	service "SHOP_PORTAL_BACKEND/SERVICE"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	validator "SHOP_PORTAL_BACKEND/VALIDATOR"

	"github.com/kataras/iris/v12"
)

func PostShopOwner(ctx iris.Context) {
	logPrefix := ctx.Values().Get("logPrefix").(string)

	headers := utils.ReadHeader(ctx, utils.PostShopOwnerHeaders)
	reqBody, bodyError := utils.ReadShopOwnerReqBody(ctx)
	utils.Logger.Info(headers, reqBody)

	headerError := validator.ValidateHeader(utils.PostShopOwnerHeaders, headers)
	reqBodyError := validator.ValidateShopOwnerReqBody(&reqBody, bodyError)

	rspBody, rspCode := helper.CheckError(headerError, utils.NULL_STRING, reqBodyError)

	var response interface{}

	if rspCode != utils.StatusOK {
		response = helper.CreateErrorResponse(rspBody)
	} else {
		response, rspCode = service.PostShopOwner(reqBody)
	}

	utils.Logger.Info(logPrefix, response)

	ctx.ResponseWriter().WriteHeader(rspCode)
	ctx.JSON(response)
	utils.Logger.Info(logPrefix + " Request Completed.")
}
