package controller

import (
	helper "SHOP_PORTAL_BACKEND/HELPER"
	service "SHOP_PORTAL_BACKEND/SERVICE"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	validator "SHOP_PORTAL_BACKEND/VALIDATOR"

	"github.com/kataras/iris/v12"
)

func PostStock(ctx iris.Context) {

	var response interface{}
	rspCode := utils.StatusOK
	logPrefix := ctx.Values().Get("logPrefix").(string)

	headers := utils.ReadHeader(ctx)
	qparams := utils.ReadQParams(ctx)
	reqBody, response, rspCode := utils.ReadPostStockReqBody(ctx, logPrefix, helper.CreateErrorResponse)
	utils.LogRequest(logPrefix, ctx, reqBody)

	if rspCode == utils.StatusOK {
		response, rspCode = validator.ValidateHeader(utils.PostStockHeaders, headers, ctx, logPrefix)
		if rspCode == utils.StatusOK {
			response, rspCode = validator.ValidateQParams(utils.PostStockQParams, qparams, logPrefix)
			if rspCode == utils.StatusOK {
				response, rspCode = validator.ValidatePostStockReqBody(&reqBody, logPrefix)
				if rspCode == utils.StatusOK {
					response, rspCode = service.PostStock(reqBody, ctx.URLParam(utils.OWNER_REG_ID), logPrefix)
				}
			}
		}
	}

	utils.LogResponse(logPrefix, response)
	ctx.ResponseWriter().WriteHeader(rspCode)
	ctx.JSON(response)
	utils.Logger.Info(logPrefix + "Request Completed.")
}
