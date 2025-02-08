package utils

import (
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	"encoding/json"

	"github.com/kataras/iris/v12"
)

// todo
func ReadHeader(ctx iris.Context, reqApiHeader map[string]bool) map[string]interface{} {

	headers := make(map[string]interface{})

	headers[CONTENT_TYPE] = ctx.GetHeader(CONTENT_TYPE)
	headers[ACCEPT] = ctx.GetHeader(ACCEPT)
	headers[ACCEPT_ENCODING] = ctx.GetHeader(ACCEPT_ENCODING)

	return headers
}

// func ReadQParams(ctx iris.Context) map[string]interface{} {

// 	shop_id := NULL_INT
// 	shop_id, _ = ctx.URLParamInt("shop_id")

// 	return shop_id, NULL_STRING
// }

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
