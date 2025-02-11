package validator

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"fmt"
	"time"
)

func ValidateGenerateTokenReqBody(body *structs.GenerateToken, bodyErr string) (string, string) {

	if bodyErr != utils.NULL_STRING {
		return bodyErr, "400007"
	}

	if body.RegId == utils.NULL_STRING {
		return "Missing reg_id", "400005"
	} else {
		var exists bool
		ServiceQuery := database.CheckValidOwnerRegId()
		db := database.ConnectDB()
		err := db.QueryRow(ServiceQuery, body.RegId).Scan(&exists)
		if err != nil {
			errMsg := fmt.Sprintf("Error in checking if row with reg_id %s exists", body.RegId)
			utils.Logger.Error(err.Error())
			return errMsg, "500001"
		}

		if exists {
			utils.Logger.Info("Row with reg_id : ", body.RegId, " exists")
		} else {
			utils.Logger.Info("Row with reg_id ", body.RegId, " does notexists")
			return "Owner Registration ID does not exist", "400004"
		}
	}

	if body.Key == utils.NULL_STRING {
		return "Missing key", "400003"
	} else {
		errMsg, errCodeStr := ValidateKey(body.Key, body.RegId)
		if errMsg != utils.NULL_STRING {
			return errMsg, errCodeStr
		}
	}

	return utils.NULL_STRING, utils.SUCCESS
}

func ValidateShopOwnerReqBody(body *structs.ShopOwner, bodyErr string) (string, string) {

	if bodyErr != utils.NULL_STRING {
		return bodyErr, "400007"
	}

	if body.ShopName == utils.NULL_STRING {
		return "Missing shop_name", "400005"
	}

	if body.OwnerName == utils.NULL_STRING {
		return "Missing owner_name", "400005"
	}

	if body.RegDate == utils.NULL_STRING {
		return "Missing reg_date", "400005"
	} else {
		_, err := time.Parse("2006-01-02", body.RegDate) // YYYY-MM-DD
		if err != nil {
			return "Invalid date format", "400006"
		}
	}

	if body.PhNo == utils.NULL_STRING {
		return "Missing ph_no", "400005"
	} else if len(body.PhNo) != 10 {
		return "Invalid ph_no", "400006"
	}

	if body.Address == utils.NULL_STRING {
		return "Missing address", "400005"
	}

	return utils.NULL_STRING, utils.SUCCESS
}

func ValidateAllShowOwnerBody(body *structs.AllShowOwner, bodyErr string) (string, string) {

	if bodyErr != utils.NULL_STRING {
		return bodyErr, "400007"
	}

	if body.IsActive == utils.NULL_STRING {
		return "Missing is_active", "400005"
	} else {
		if body.IsActive != utils.ACTIVE_YES && body.IsActive != utils.ACTIVE_NO && body.IsActive != utils.ALL {
			return "Invalid is_active", "400006"
		}
	}

	return utils.NULL_STRING, utils.SUCCESS
}

func ValidateCustomerReqBody(body *structs.Customer, bodyErr string) (string, string) {

	if bodyErr != utils.NULL_STRING {
		return bodyErr, "400007"
	}

	if body.Name == utils.NULL_STRING {
		return "Missing name", "400005"
	}

	if body.CompanyName == utils.NULL_STRING {
		return "Missing company_name", "400005"
	}

	if body.PhNo == utils.NULL_STRING {
		return "Missing ph_no", "400005"
	} else if len(body.PhNo) != 10 {
		return "Invalid ph_no", "400006"
	}

	if body.RegDate == utils.NULL_STRING {
		return "Missing reg_date", "400005"
	} else {
		_, err := time.Parse("2006-01-02", body.RegDate) // YYYY-MM-DD
		if err != nil {
			return "Invalid date format", "400006"
		}
	}

	if body.Address == utils.NULL_STRING {
		return "Missing address", "400005"
	}

	if body.Remarks == utils.NULL_STRING {
		return "Missing remarks", "400005"
	}

	return utils.NULL_STRING, utils.SUCCESS
}
