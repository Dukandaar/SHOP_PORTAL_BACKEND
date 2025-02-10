package controller

import (
	helper "SHOP_PORTAL_BACKEND/HELPER"
	service "SHOP_PORTAL_BACKEND/SERVICE"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	validator "SHOP_PORTAL_BACKEND/VALIDATOR"

	"github.com/kataras/iris/v12"
)

func PostCustomer(ctx iris.Context) {

	var response interface{}
	var errCodeStr string
	rspCode := utils.StatusOK

	logPrefix := ctx.Values().Get("logPrefix").(string)

	headers := utils.ReadHeader(ctx)
	qparams := utils.ReadQParams(ctx)
	reqBody, bodyError := utils.ReadCustomerReqBody(ctx)

	utils.Logger.Info(headers, qparams, reqBody)

	headerError, errCodeStr := validator.ValidateHeader(utils.PostCustomerHeaders, headers, ctx)
	if errCodeStr != utils.SUCCESS { // header error
		response, rspCode = helper.CreateErrorResponse(errCodeStr, headerError)
		utils.Logger.Error(headerError)
	} else {
		QparamsError, errCodeStr := validator.ValidateQParams(utils.PostCustomerQParams, qparams)
		if errCodeStr != utils.SUCCESS { // qparams error
			response, rspCode = helper.CreateErrorResponse(errCodeStr, QparamsError)
			utils.Logger.Error(QparamsError)
		} else {
			reqBodyError, errCodeStr := validator.ValidateCustomerReqBody(&reqBody, bodyError)
			if errCodeStr != utils.SUCCESS { // body error
				response, rspCode = helper.CreateErrorResponse(errCodeStr, reqBodyError)
				utils.Logger.Error(reqBodyError)
			} else {
				response, rspCode = service.PostCustomer(reqBody, ctx.URLParam(utils.OWNER_REG_ID))
			}
		}
	}

	utils.Logger.Info(logPrefix, response)

	ctx.ResponseWriter().WriteHeader(rspCode)
	ctx.JSON(response)
	utils.Logger.Info("Request Completed.")
}
