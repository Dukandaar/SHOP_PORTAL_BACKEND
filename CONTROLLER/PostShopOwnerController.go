package controller

import (
	helper "SHOP_PORTAL_BACKEND/HELPER"
	service "SHOP_PORTAL_BACKEND/SERVICE"
	utils "SHOP_PORTAL_BACKEND/UTILS"

	"github.com/kataras/iris/v12"
)

func PostShopOwner(ctx iris.Context) {
	logPrefix := ctx.Values().Get("logPrefix").(string)

	headerError := utils.ReadHeader(ctx)
	_, qParamError := utils.ReadQParams(ctx)
	reqBody, reqBodyError := utils.ReadShopOwnerReqBody(ctx)

	utils.Logger.Info(logPrefix, ctx.Request())

	rspBody, rspCode := helper.CheckError(headerError, qParamError, reqBodyError)

	var rsp interface{}

	if rspCode != utils.StatusOK {
		rsp = helper.CreateErrorResponse(rspBody)

	} else {
		rsp = service.PostShopOwner(reqBody)
	}
	ctx.JSON(rsp)
	utils.Logger.Info(logPrefix + " Request Completed.")
}
