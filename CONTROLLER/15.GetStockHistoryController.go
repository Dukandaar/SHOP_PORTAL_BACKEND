package controller

import (
	service "SHOP_PORTAL_BACKEND/SERVICE"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	validator "SHOP_PORTAL_BACKEND/VALIDATOR"
	"strconv"

	"github.com/kataras/iris/v12"
)

func GetStockHistory(ctx iris.Context) {

	var response interface{}
	rspCode := utils.StatusOK
	logPrefix := ctx.Values().Get("logPrefix").(string)

	headers := utils.ReadHeader(ctx)
	qparams := utils.ReadQParams(ctx)
	utils.LogRequest(logPrefix, ctx, response)

	if rspCode == utils.StatusOK {
		response, rspCode = validator.ValidateHeader(utils.GetStockHistoryHeaders, headers, ctx, logPrefix)
		if rspCode == utils.StatusOK {
			response, rspCode = validator.ValidateQParams(utils.GetStockHistoryQParams, qparams, logPrefix)
			if rspCode == utils.StatusOK {
				stockId, _ := strconv.Atoi(ctx.URLParam(utils.STOCK_ID))
				response, rspCode = service.GetStockHistory(ctx.URLParam(utils.OWNER_REG_ID), stockId, logPrefix)
			}
		}
	}

	utils.LogResponse(logPrefix, response)
	ctx.ResponseWriter().WriteHeader(rspCode)
	ctx.JSON(response)
	utils.Logger.Info(logPrefix, "Request Completed.")
}
