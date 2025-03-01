package controller

import (
	service "SHOP_PORTAL_BACKEND/SERVICE"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	validator "SHOP_PORTAL_BACKEND/VALIDATOR"
	"strconv"

	"github.com/kataras/iris/v12"
)

func GetStock(ctx iris.Context) {

	var response interface{}
	rspCode := utils.StatusOK
	logPrefix := ctx.Values().Get("logPrefix").(string)

	headers := utils.ReadHeader(ctx)
	qparams := utils.ReadQParams(ctx)
	utils.LogRequest(logPrefix, ctx, response)

	if rspCode == utils.StatusOK {
		response, rspCode = validator.ValidateHeader(utils.GetStockHeaders, headers, ctx, logPrefix)
		if rspCode == utils.StatusOK {
			response, rspCode = validator.ValidateQParams(utils.GetStockQParams, qparams, logPrefix)
			if rspCode == utils.StatusOK {
				stockId, _ := strconv.Atoi(ctx.URLParam(utils.STOCK_ID))
				response, rspCode = service.GetStock(ctx.URLParam(utils.OWNER_REG_ID), stockId, logPrefix)
			}
		}
	}

	utils.LogResponse(logPrefix, response)
	ctx.ResponseWriter().WriteHeader(rspCode)
	ctx.JSON(response)
	utils.Logger.Info(logPrefix + "Request Completed.")
}
