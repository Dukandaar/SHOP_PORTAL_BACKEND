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

	return headers
}

func ReadQParams(ctx iris.Context) map[string]interface{} {

	qparams := make(map[string]interface{})

	for key, value := range ctx.URLParams() {
		qparams[key] = value
	}

	return qparams
}

func ReadGenerateTokenReqBody(ctx iris.Context) (structs.GenerateToken, string) {
	body := structs.GenerateToken{}
	err := json.NewDecoder(ctx.Request().Body).Decode(&body)
	if err != nil {
		return body, "Error in decoding request body"
	}
	return body, NULL_STRING
}

func ReadShopOwnerReqBody(ctx iris.Context) (structs.ShopOwner, string) {
	body := structs.ShopOwner{}
	err := json.NewDecoder(ctx.Request().Body).Decode(&body)
	if err != nil {
		return body, "Error in decoding request body"
	}
	return body, NULL_STRING
}

func ReadAllShowOwnerBody(ctx iris.Context) (structs.AllShowOwner, string) {
	body := structs.AllShowOwner{}
	err := json.NewDecoder(ctx.Request().Body).Decode(&body)
	if err != nil {
		return body, "Error in decoding request body"
	}
	return body, NULL_STRING
}

func ReadCustomerReqBody(ctx iris.Context) (structs.Customer, string) {
	body := structs.Customer{}
	err := json.NewDecoder(ctx.Request().Body).Decode(&body)
	if err != nil {
		return body, "Error in decoding request body"
	}
	return body, NULL_STRING
}
