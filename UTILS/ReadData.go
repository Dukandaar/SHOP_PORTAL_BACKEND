package utils

import (
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	"encoding/json"

	"github.com/kataras/iris/v12"
)

func ReadHeader(ctx iris.Context) map[string]interface{} {

	headers := make(map[string]interface{})
	headers[CONTENT_TYPE] = ctx.Request().Header.Get(CONTENT_TYPE)
	headers[ACCEPT] = ctx.Request().Header.Get(ACCEPT)
	headers[ACCEPT_ENCODING] = ctx.Request().Header.Get(ACCEPT_ENCODING)
	headers[TOKEN] = ctx.Request().Header.Get(TOKEN)
	headers[SKIP_TOKEN] = ctx.Request().Header.Get(SKIP_TOKEN)
	headers[CATCH_CONTROL] = ctx.Request().Header.Get(CATCH_CONTROL)

	return headers
}

func ReadQParams(ctx iris.Context) map[string]interface{} {

	qparams := make(map[string]interface{})

	for key, value := range ctx.URLParams() {
		qparams[key] = value
	}

	return qparams
}

func ReadGenerateTokenReqBody(ctx iris.Context, logPrefix string, CreateErrorResponse func(string, string, string) (structs.ErrorResponse1, int)) (structs.GenerateToken, interface{}, int) {
	body := structs.GenerateToken{}
	var response interface{}
	var rspCode = StatusOK
	err := json.NewDecoder(ctx.Request().Body).Decode(&body)
	if err != nil {
		response, rspCode = CreateErrorResponse("400008", "Error in decoding request body", logPrefix)
	}
	return body, response, rspCode
}

func ReadShopOwnerReqBody(ctx iris.Context, logPrefix string, CreateErrorResponse func(string, string, string) (structs.ErrorResponse1, int)) (structs.ShopOwner, interface{}, int) {
	body := structs.ShopOwner{}
	var response interface{}
	var rspCode = StatusOK
	err := json.NewDecoder(ctx.Request().Body).Decode(&body)
	if err != nil {
		response, rspCode = CreateErrorResponse("400008", "Error in decoding request body", logPrefix)
	}
	return body, response, rspCode
}

func ReadAllShopOwnerBody(ctx iris.Context, logPrefix string, CreateErrorResponse func(string, string, string) (structs.ErrorResponse1, int)) (structs.AllShopOwner, interface{}, int) {
	body := structs.AllShopOwner{}
	var response interface{}
	var rspCode = StatusOK
	err := json.NewDecoder(ctx.Request().Body).Decode(&body)
	if err != nil {
		response, rspCode = CreateErrorResponse("400008", "Error in decoding request body", logPrefix)
	}
	return body, response, rspCode
}

func ReadCustomerReqBody(ctx iris.Context, logPrefix string, CreateErrorResponse func(string, string, string) (structs.ErrorResponse1, int)) (structs.Customer, interface{}, int) {
	body := structs.Customer{}
	var response interface{}
	var rspCode = StatusOK
	err := json.NewDecoder(ctx.Request().Body).Decode(&body)
	if err != nil {
		response, rspCode = CreateErrorResponse("400008", "Error in decoding request body", logPrefix)
	}
	return body, response, rspCode
}

func ReadFilteredCustomerReqBody(ctx iris.Context, logPrefix string, CreateErrorResponse func(string, string, string) (structs.ErrorResponse1, int)) (structs.FilteredCustomer, interface{}, int) {
	body := structs.FilteredCustomer{}
	var response interface{}
	var rspCode = StatusOK
	err := json.NewDecoder(ctx.Request().Body).Decode(&body)
	if err != nil {
		response, rspCode = CreateErrorResponse("400008", "Error in decoding request body", logPrefix)
	}
	return body, response, rspCode
}

func ReadPostStockReqBody(ctx iris.Context, logPrefix string, CreateErrorResponse func(string, string, string) (structs.ErrorResponse1, int)) (structs.PostStock, interface{}, int) {
	body := structs.PostStock{}
	var response interface{}
	var rspCode = StatusOK
	err := json.NewDecoder(ctx.Request().Body).Decode(&body)
	if err != nil {
		response, rspCode = CreateErrorResponse("400008", "Error in decoding request body", logPrefix)
	}
	return body, response, rspCode
}

func ReadPutStockReqBody(ctx iris.Context, logPrefix string, CreateErrorResponse func(string, string, string) (structs.ErrorResponse1, int)) (structs.PutStock, interface{}, int) {
	body := structs.PutStock{}
	var response interface{}
	var rspCode = StatusOK
	err := json.NewDecoder(ctx.Request().Body).Decode(&body)
	if err != nil {
		response, rspCode = CreateErrorResponse("400008", "Error in decoding request body", logPrefix)
	}
	return body, response, rspCode
}

func ReadCustomerBillReqBody(ctx iris.Context, logPrefix string, CreateErrorResponse func(string, string, string) (structs.ErrorResponse1, int)) (structs.CustomerBill, interface{}, int) {
	body := structs.CustomerBill{}
	var response interface{}
	var rspCode = StatusOK
	err := json.NewDecoder(ctx.Request().Body).Decode(&body)
	if err != nil {
		response, rspCode = CreateErrorResponse("400008", "Error in decoding request body", logPrefix)
	}
	return body, response, rspCode
}

func ReadGetAllStockReqBody(ctx iris.Context, logPrefix string, CreateErrorResponse func(string, string, string) (structs.ErrorResponse1, int)) (structs.AllStock, interface{}, int) {
	body := structs.AllStock{}
	var response interface{}
	var rspCode = StatusOK
	err := json.NewDecoder(ctx.Request().Body).Decode(&body)
	if err != nil {
		response, rspCode = CreateErrorResponse("400008", "Error in decoding request body", logPrefix)
	}
	return body, response, rspCode
}
