package validator

import (
	utils "SHOP_PORTAL_BACKEND/UTILS"

	"github.com/kataras/iris/v12"
)

func ValidateHeader(reqApiHeader map[string]bool, apiHeader map[string]interface{}, ctx iris.Context, logPrefix string) (string, string) {

	// Token validation
	skipToken, _ := apiHeader[utils.SKIP_TOKEN].(string)
	if reqApiHeader[utils.TOKEN] && skipToken != utils.TRUE {
		if apiHeader[utils.TOKEN] == utils.NULL_STRING {
			return "Missing Token header", "400001"
		}

		token, _ := apiHeader[utils.TOKEN].(string)
		regId := ctx.URLParam(utils.OWNER_REG_ID)

		if len(regId) != 10 {
			return "Invalid owner_reg_id length", "400004"
		}

		errMsg, rspCode := ValidateToken(token, regId, logPrefix)
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

	// Catch-Control
	if reqApiHeader[utils.CATCH_CONTROL] {
		if apiHeader[utils.CATCH_CONTROL] == utils.NULL_STRING {
			return "Missing Catch-Control header", "400001"
		}
		valid := false
		for _, validHeaderValue := range utils.ValidHeaders[utils.CATCH_CONTROL] {
			if apiHeader[utils.CATCH_CONTROL] == validHeaderValue {
				valid = true
				break
			}
		}
		if !valid {
			return "Invalid Catch-Control header", "400002"
		}
	}

	return utils.NULL_STRING, utils.SUCCESS
}
