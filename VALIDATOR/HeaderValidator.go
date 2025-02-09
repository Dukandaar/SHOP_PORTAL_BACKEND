package validator

import (
	utils "SHOP_PORTAL_BACKEND/UTILS"

	"github.com/kataras/iris/v12"
)

func ValidateHeader(reqApiHeader map[string]bool, apiHeader map[string]interface{}, ctx iris.Context) (string, string) {

	// Token validation
	if reqApiHeader[utils.TOKEN] {
		if apiHeader[utils.TOKEN] == utils.NULL_STRING {
			return "Missing Token header", "400001"
		}

		token, _ := apiHeader[utils.TOKEN].(string)
		regId := ctx.URLParam("reg_id")

		if len(regId) != 10 {
			return "Invalid reg_id length", "400004"
		}

		errMsg, rspCode := ValidateToken(token, regId)
		if rspCode != utils.SUCCESS {
			return errMsg, rspCode
		}

	}

	// Content-Type
	if reqApiHeader[utils.CONTENT_TYPE] {
		if apiHeader[utils.CONTENT_TYPE] == utils.NULL_STRING {
			return "Missing Content-Type header", "400001"
		}
		valid := false
		for _, validHeaderValue := range utils.ValidHeaders[utils.CONTENT_TYPE] {
			if apiHeader[utils.CONTENT_TYPE] == validHeaderValue {
				valid = true
				break
			}
		}
		if !valid {
			return "Invalid Content-Type header", "400002"
		}
	}

	// Accept
	if reqApiHeader[utils.ACCEPT] {
		if apiHeader[utils.ACCEPT] == utils.NULL_STRING {
			return "Missing Accept header", "400001"
		}
		valid := false
		for _, validHeaderValue := range utils.ValidHeaders[utils.ACCEPT] {
			if apiHeader[utils.ACCEPT] == validHeaderValue {
				valid = true
				break
			}
		}
		if !valid {
			return "Invalid Accept header", "400002"
		}
	}

	// Accept-Encoding
	if reqApiHeader[utils.ACCEPT_ENCODING] {
		if apiHeader[utils.ACCEPT_ENCODING] == utils.NULL_STRING {
			return "Missing Accept-Encoding header", "400001"
		}
		valid := false
		for _, validHeaderValue := range utils.ValidHeaders[utils.ACCEPT_ENCODING] {
			if apiHeader[utils.ACCEPT_ENCODING] == validHeaderValue {
				valid = true
				break
			}
		}
		if !valid {
			return "Invalid Accept-Encoding header", "400002"
		}
	}

	return utils.NULL_STRING, utils.SUCCESS
}
