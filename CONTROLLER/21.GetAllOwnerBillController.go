package controller

import (
	service "SHOP_PORTAL_BACKEND/SERVICE"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	validator "SHOP_PORTAL_BACKEND/VALIDATOR"

	"github.com/kataras/iris/v12"
)

func GetAllOwnerBill(ctx iris.Context) {

	var response interface{}
	rspCode := utils.StatusOK
	logPrefix := ctx.Values().Get("logPrefix").(string)

	headers := utils.ReadHeader(ctx)
	qparams := utils.ReadQParams(ctx)
	utils.LogRequest(logPrefix, ctx, response)

	if rspCode == utils.StatusOK {
		response, rspCode = validator.ValidateHeader(utils.GetAllOwnerBillHeaders, headers, ctx, logPrefix)
		if rspCode == utils.StatusOK {
			response, rspCode = validator.ValidateQParams(utils.GetAllOwnerBillQParams, qparams, logPrefix)
			if rspCode == utils.StatusOK {
				response, rspCode = service.GetAllOwnerBill(ctx.URLParam(utils.OWNER_REG_ID), logPrefix)
			}
		}
	}
	ctx.ResponseWriter().WriteHeader(rspCode)
	ctx.JSON(response)
	utils.Logger.Info(logPrefix + "Request Completed.")

}
