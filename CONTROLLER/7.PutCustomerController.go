package controller

import (
	helper "SHOP_PORTAL_BACKEND/HELPER"
	service "SHOP_PORTAL_BACKEND/SERVICE"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	validator "SHOP_PORTAL_BACKEND/VALIDATOR"

	"github.com/kataras/iris/v12"
)

func PutCustomer(ctx iris.Context) {

	var response interface{}
	var errCodeStr string
	rspCode := utils.StatusOK

	logPrefix := ctx.Values().Get("logPrefix").(string)

	headers := utils.ReadHeader(ctx)
	qparams := utils.ReadQParams(ctx)
	reqBody, bodyError := utils.ReadCustomerReqBody(ctx)
	utils.Logger.Info(logPrefix, headers, qparams, reqBody)

	headerError, errCodeStr := validator.ValidateHeader(utils.PutCustomerHeaders, headers, ctx, logPrefix)
	if errCodeStr != utils.SUCCESS { // header error
		response, rspCode = helper.CreateErrorResponse(errCodeStr, headerError)
		utils.Logger.Error(logPrefix, headerError)
	} else {
		QparamsError, errCodeStr := validator.ValidateQParams(utils.PutCustomerQParams, qparams, logPrefix)
		if errCodeStr != utils.SUCCESS { // qparams error
			response, rspCode = helper.CreateErrorResponse(errCodeStr, QparamsError)
			utils.Logger.Error(logPrefix, QparamsError)
		} else {
			reqBodyError, errCodeStr := validator.ValidateCustomerReqBody(&reqBody, bodyError)
			if errCodeStr != utils.SUCCESS { // body error
				response, rspCode = helper.CreateErrorResponse(errCodeStr, reqBodyError)
				utils.Logger.Error(logPrefix, reqBodyError)
			} else {
				response, rspCode = service.PutCustomer(reqBody, ctx.URLParam(utils.OWNER_REG_ID), ctx.URLParam(utils.CUSTOMER_REG_ID), logPrefix)
			}
		}
	}

	utils.Logger.Info(logPrefix, response)

	ctx.ResponseWriter().WriteHeader(rspCode)
	ctx.JSON(response)
	utils.Logger.Info(logPrefix + "Request Completed.")
}
