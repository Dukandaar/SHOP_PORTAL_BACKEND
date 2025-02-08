package validator

import (
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"time"
)

func ValidateHeader(reqApiHeader map[string]bool, apiHeader map[string]interface{}) string {

	// Content-Type
	if reqApiHeader[utils.CONTENT_TYPE] {
		if apiHeader[utils.CONTENT_TYPE] == utils.NULL_STRING {
			return "Missing Content-Type header"
		}
		valid := false
		for _, validHeaderValue := range utils.ValidHeaders[utils.CONTENT_TYPE] {
			if apiHeader[utils.CONTENT_TYPE] == validHeaderValue {
				valid = true
				break
			}
		}
		if !valid {
			return "Invalid Content-Type header"
		}
	}

	// Accept
	if reqApiHeader[utils.ACCEPT] {
		if apiHeader[utils.ACCEPT] == utils.NULL_STRING {
			return "Missing Accept header"
		}
		valid := false
		for _, validHeaderValue := range utils.ValidHeaders[utils.ACCEPT] {
			if apiHeader[utils.ACCEPT] == validHeaderValue {
				valid = true
				break
			}
		}
		if !valid {
			return "Invalid Accept header"
		}
	}

	// Accept-Encoding
	if reqApiHeader[utils.ACCEPT_ENCODING] {
		if apiHeader[utils.ACCEPT_ENCODING] == utils.NULL_STRING {
			return "Missing Accept-Encoding header"
		}
		valid := false
		for _, validHeaderValue := range utils.ValidHeaders[utils.ACCEPT_ENCODING] {
			if apiHeader[utils.ACCEPT_ENCODING] == validHeaderValue {
				valid = true
				break
			}
		}
		if !valid {
			return "Invalid Accept-Encoding header"
		}
	}

	return utils.NULL_STRING
}

func ValidateQParams(qparams map[string]interface{}) error {
	return nil
}

func ValidateGenerateTokenReqBody(body *structs.GenerateToken, bodyErr string) string {

	if bodyErr != utils.NULL_STRING {
		return bodyErr
	}

	if body.RegId == utils.NULL_STRING {
		return "Missing reg_id"
	}

	if body.Key == utils.NULL_STRING {
		return "Missing key"
	}

	return utils.NULL_STRING
}

func ValidateShopOwnerReqBody(body *structs.ShopOwner, bodyErr string) string {

	if bodyErr != utils.NULL_STRING {
		return bodyErr
	}

	if body.ShopName == utils.NULL_STRING {
		return "Missing shop_name"
	}

	if body.OwnerName == utils.NULL_STRING {
		return "Missing owner_name"
	}

	if body.RegDate == utils.NULL_STRING {
		return "Missing reg_date"
	} else {
		_, err := time.Parse("2006-01-02", body.RegDate) // YYYY-MM-DD
		if err != nil {
			return "Invalid date format"
		}
	}

	if body.PhNo == utils.NULL_STRING {
		return "Missing ph_no"
	} else if len(body.PhNo) != 10 {
		return "Invalid ph_no"
	}

	if body.Address == utils.NULL_STRING {
		return "Missing address"
	}

	return utils.NULL_STRING
}
