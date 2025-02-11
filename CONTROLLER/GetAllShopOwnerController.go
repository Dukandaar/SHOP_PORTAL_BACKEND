package controller

import (
	helper "SHOP_PORTAL_BACKEND/HELPER"
	service "SHOP_PORTAL_BACKEND/SERVICE"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	validator "SHOP_PORTAL_BACKEND/VALIDATOR"

	"github.com/kataras/iris/v12"
)

func GetAllShopOwner(ctx iris.Context) {

	var response interface{}
	var errCodeStr string
	rspCode := utils.StatusOK

	logPrefix := ctx.Values().Get("logPrefix").(string)

	headers := utils.ReadHeader(ctx)
	reqBody, bodyError := utils.ReadAllShopOwnerBody(ctx)
	utils.Logger.Info(headers)

	headerError, errCodeStr := validator.ValidateHeader(utils.GetAllShopOwnerHeaders, headers, ctx)
	if errCodeStr != utils.SUCCESS {
		response, rspCode = helper.CreateErrorResponse(errCodeStr, headerError)
		utils.Logger.Error(headerError)
	} else {

		reqBodyError, errCodeStr := validator.ValidateAllShopOwnerBody(&reqBody, bodyError)
		if errCodeStr != utils.SUCCESS {
			response, rspCode = helper.CreateErrorResponse(errCodeStr, reqBodyError)
			utils.Logger.Error(reqBodyError)
		} else {
			response, rspCode = service.GetAllShopOwner(reqBody)
			if rspCode != utils.StatusOK {
				utils.Logger.Error("Error in getting all rows")
				response, rspCode = helper.CreateErrorResponse("500001", "Error in getting all rows")
			} else {
				utils.Logger.Info("Got all rows")
			}
		}
	}

	utils.Logger.Info(logPrefix, response)

	ctx.ResponseWriter().WriteHeader(rspCode)
	ctx.JSON(response)
	utils.Logger.Info(logPrefix + "Response Completed.")
}
