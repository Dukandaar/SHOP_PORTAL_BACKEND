package controller

import (
	helper "SHOP_PORTAL_BACKEND/HELPER"
	service "SHOP_PORTAL_BACKEND/SERVICE"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	validator "SHOP_PORTAL_BACKEND/VALIDATOR"

	"github.com/kataras/iris/v12"
)

func PostShopOwner(ctx iris.Context) {

	var response interface{}
	rspCode := utils.StatusOK
	logPrefix := ctx.Values().Get("logPrefix").(string)

	headers := utils.ReadHeader(ctx)
	reqBody, response, rspCode := utils.ReadShopOwnerReqBody(ctx, logPrefix, helper.CreateErrorResponse)
	utils.LogRequest(logPrefix, ctx, reqBody)

	if rspCode == utils.StatusOK {
		response, rspCode = validator.ValidateHeader(utils.PostShopOwnerHeaders, headers, ctx, logPrefix)
		if rspCode == utils.StatusOK {
			response, rspCode = validator.ValidateShopOwnerReqBody(&reqBody, logPrefix)
			if rspCode == utils.StatusOK {
				response, rspCode = service.PostShopOwner(reqBody, logPrefix)
			}
		}
	}

	utils.LogResponse(logPrefix, response)
	ctx.ResponseWriter().WriteHeader(rspCode)
	ctx.JSON(response)
	utils.Logger.Info(logPrefix + "Request Completed.")
}
