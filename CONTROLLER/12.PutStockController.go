package controller

import (
	helper "SHOP_PORTAL_BACKEND/HELPER"
	service "SHOP_PORTAL_BACKEND/SERVICE"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	validator "SHOP_PORTAL_BACKEND/VALIDATOR"
	"strconv"

	"github.com/kataras/iris/v12"
)

func PutStock(ctx iris.Context) {

	var response interface{}
	var errCodeStr string
	rspCode := utils.StatusOK

	logPrefix := ctx.Values().Get("logPrefix").(string)

	headers := utils.ReadHeader(ctx)
	qparams := utils.ReadQParams(ctx)
	reqBody, _ := utils.ReadPutStockReqBody(ctx)
	utils.Logger.Info(logPrefix, headers, qparams, reqBody)

	headerError, errCodeStr := validator.ValidateHeader(utils.PutStockHeaders, headers, ctx, logPrefix)
	if errCodeStr != utils.SUCCESS { // header error
		response, rspCode = helper.CreateErrorResponse(errCodeStr, headerError)
		utils.Logger.Error(logPrefix, headerError)
	} else {
		QparamsError, errCodeStr := validator.ValidateQParams(utils.PutStockQParams, qparams, logPrefix)
		if errCodeStr != utils.SUCCESS { // qparams error
			response, rspCode = helper.CreateErrorResponse(errCodeStr, QparamsError)
			utils.Logger.Error(logPrefix, QparamsError)
		} else {
			sId, _ := strconv.Atoi(ctx.URLParam(utils.STOCK_ID))
			response, rspCode = service.PutStock(reqBody, ctx.URLParam(utils.OWNER_REG_ID), sId, logPrefix)
		}
	}

	utils.Logger.Info(logPrefix, response)

	ctx.ResponseWriter().WriteHeader(rspCode)
	ctx.JSON(response)
	utils.Logger.Info(logPrefix + "Request Completed.")
}
