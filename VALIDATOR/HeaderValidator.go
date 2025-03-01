package validator

import (
	helper "SHOP_PORTAL_BACKEND/HELPER"
	utils "SHOP_PORTAL_BACKEND/UTILS"

	"github.com/kataras/iris/v12"
)

func ValidateHeader(reqApiHeader map[string]bool, apiHeader map[string]interface{}, ctx iris.Context, logPrefix string) (interface{}, int) {

	// Token validation
	skipToken, _ := apiHeader[utils.SKIP_TOKEN].(string)
	if reqApiHeader[utils.TOKEN] && skipToken != utils.TRUE {
		if apiHeader[utils.TOKEN] == utils.NULL_STRING {
			return helper.CreateErrorResponse("400001", "Missing Token header", logPrefix)
		}

		token, _ := apiHeader[utils.TOKEN].(string)
		regId := ctx.URLParam(utils.OWNER_REG_ID)

		if len(regId) != 10 {
			return helper.CreateErrorResponse("400004", "Invalid owner_reg_id length", logPrefix)
		}

		errMsg, rspCode := ValidateToken(token, regId, logPrefix)
		if rspCode != utils.SUCCESS {
			return helper.CreateErrorResponse(rspCode, errMsg, logPrefix)
		}

	}

	// Content-Type
	if reqApiHeader[utils.CONTENT_TYPE] {
		if apiHeader[utils.CONTENT_TYPE] == utils.NULL_STRING {
			return helper.CreateErrorResponse("400001", "Missing Content-Type header", logPrefix)
		}
		valid := false
		for _, validHeaderValue := range utils.ValidHeaders[utils.CONTENT_TYPE] {
			if apiHeader[utils.CONTENT_TYPE] == validHeaderValue {
				valid = true
				break
			}
		}
		if !valid {
			return helper.CreateErrorResponse("400002", "Invalid Content-Type header", logPrefix)
		}
	}

	// Accept
	if reqApiHeader[utils.ACCEPT] {
		if apiHeader[utils.ACCEPT] == utils.NULL_STRING {
			return helper.CreateErrorResponse("400001", "Missing Accept header", logPrefix)
		}
		valid := false
		for _, validHeaderValue := range utils.ValidHeaders[utils.ACCEPT] {
			if apiHeader[utils.ACCEPT] == validHeaderValue {
				valid = true
				break
			}
		}
		if !valid {
			return helper.CreateErrorResponse("400002", "Invalid Accept header", logPrefix)
		}
	}

	// Accept-Encoding
	if reqApiHeader[utils.ACCEPT_ENCODING] {
		if apiHeader[utils.ACCEPT_ENCODING] == utils.NULL_STRING {
			return helper.CreateErrorResponse("400001", "Missing Accept-Encoding header", logPrefix)
		}
		valid := false
		for _, validHeaderValue := range utils.ValidHeaders[utils.ACCEPT_ENCODING] {
			if apiHeader[utils.ACCEPT_ENCODING] == validHeaderValue {
				valid = true
				break
			}
		}
		if !valid {
			return helper.CreateErrorResponse("400002", "Invalid Accept-Encoding header", logPrefix)
		}
	}

	// Catch-Control
	if reqApiHeader[utils.CATCH_CONTROL] {
		if apiHeader[utils.CATCH_CONTROL] == utils.NULL_STRING {
			return helper.CreateErrorResponse("400001", "Missing Catch-Control header", logPrefix)
		}
		valid := false
		for _, validHeaderValue := range utils.ValidHeaders[utils.CATCH_CONTROL] {
			if apiHeader[utils.CATCH_CONTROL] == validHeaderValue {
				valid = true
				break
			}
		}
		if !valid {
			return helper.CreateErrorResponse("400002", "Invalid Catch-Control header", logPrefix)
		}
	}

	return utils.NULL_STRING, utils.StatusOK
}
