package controller

import (
	utils "SHOP_PORTAL_BACKEND/UTILS"

	"github.com/kataras/iris/v12"
)

func PutCustomerTransaction(ctx iris.Context) {

	var response interface{}
	// var errCodeStr string
	rspCode := utils.StatusOK

	logPrefix := ctx.Values().Get("logPrefix").(string)

	// headers := utils.ReadHeader(ctx)
	// qparams := utils.ReadQParams(ctx)
	// reqBody, bodyError := utils.ReadCustomerTransactionReqBody(ctx)
	// utils.Logger.Info(logPrefix, headers, qparams, reqBody)

	// headerError, errCodeStr := validator.ValidateHeader(utils.PutCustomerTransactionHeaders, headers, ctx)
	// if errCodeStr != utils.SUCCESS { // header error
	// 	response, rspCode = helper.CreateErrorResponse(errCodeStr, headerError)
	// 	utils.Logger.Error(logPrefix, headerError)
	// } else {
	// 	QparamsError, errCodeStr := validator.ValidateQParams(utils.PutCustomerTransactionQParams, qparams)
	// 	if errCodeStr != utils.SUCCESS { // qparams error
	// 		response, rspCode = helper.CreateErrorResponse(errCodeStr, QparamsError)
	// 		utils.Logger.Error(logPrefix, QparamsError)
	// 	} else {
	// 		reqBodyError, errCodeStr := validator.ValidateCustomerTransactionReqBody(&reqBody, bodyError)
	// 		if errCodeStr != utils.SUCCESS { // body error
	// 			response, rspCode = helper.CreateErrorResponse(errCodeStr, reqBodyError)
	// 			utils.Logger.Error(logPrefix, reqBodyError)
	// 		} else {
	// 			response, rspCode = service.PutCustomerTransaction(reqBody, ctx.URLParam(utils.OWNER_REG_ID), ctx.URLParam(utils.CUSTOMER_REG_ID))
	// 		}
	// 	}
	// }

	utils.Logger.Info(logPrefix, response)

	ctx.ResponseWriter().WriteHeader(rspCode)
	ctx.JSON(response)
	utils.Logger.Info(logPrefix + "Response Completed.")
}
