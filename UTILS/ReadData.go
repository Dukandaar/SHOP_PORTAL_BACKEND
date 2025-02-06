package utils

import (
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	"encoding/json"

	"github.com/kataras/iris/v12"
)

// todo
func ReadHeader(ctx iris.Context) string {
	h1 := NULL_STRING
	h2 := NULL_STRING
	h3 := NULL_STRING

	h1 = ctx.GetHeader("Content-Type")
	h2 = ctx.GetHeader("Accept")
	h3 = ctx.GetHeader("Authorization")

	if h1 == NULL_STRING {
		return "MISSING_CONTENT_TYPE"
	}

	if h2 == NULL_STRING {
		return "MISSING_ACCEPT"
	}

	if h3 == NULL_STRING {
		return "MISSING_AUTHORIZATION"
	}

	return NULL_STRING
}

func ReadQParams(ctx iris.Context) (int, string) {

	shop_id := NULL_INT
	shop_id, _ = ctx.URLParamInt("shop_id")

	if shop_id == NULL_INT {
		return NULL_INT, "MISSING_SHOP_ID"
	}

	return shop_id, NULL_STRING
}

func ReadShopOwnerReqBody(ctx iris.Context) (structs.ShopOwner, string) {
	body := structs.ShopOwner{}
	err := json.NewDecoder(ctx.Request().Body).Decode(&body)
	if err != nil {
		return body, "Error in decoding request body"
	}
	return body, NULL_STRING
}
