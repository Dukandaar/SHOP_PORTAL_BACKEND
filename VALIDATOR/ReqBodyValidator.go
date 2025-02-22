package validator

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"fmt"
	"time"
)

func ValidateGenerateTokenReqBody(body *structs.GenerateToken, bodyErr string, logPrefix string) (string, string) {

	if bodyErr != utils.NULL_STRING {
		return bodyErr, "400008"
	}

	if body.RegId == utils.NULL_STRING {
		return "Missing reg_id", "400005"
	}

	if len(body.RegId) > utils.SHOP_REG_ID_MAX_LEN {
		return "reg_id length greater than 10", "400007"
	}

	var exists bool
	ServiceQuery := database.CheckValidOwnerRegId()
	db := database.DB
	err := db.QueryRow(ServiceQuery, body.RegId).Scan(&exists)
	if err != nil {
		errMsg := fmt.Sprintf("Error in checking if row with reg_id %s exists", body.RegId)
		utils.Logger.Error(logPrefix, err.Error())
		return errMsg, "500001"
	}

	if exists {
		utils.Logger.Info(logPrefix, "Row with reg_id : ", body.RegId, " exists")
	} else {
		utils.Logger.Info(logPrefix, "Row with reg_id ", body.RegId, " does not exist")
		return "Owner Registration ID does not exist", "400006"
	}

	if body.Key == utils.NULL_STRING {
		return "Missing key", "400005"
	} else {
		errMsg, errCodeStr := ValidateKey(body.Key, body.RegId, logPrefix)
		if errMsg != utils.NULL_STRING {
			return errMsg, errCodeStr
		}
	}

	return utils.NULL_STRING, utils.SUCCESS
}

func ValidateShopOwnerReqBody(body *structs.ShopOwner, bodyErr string) (string, string) {

	if bodyErr != utils.NULL_STRING {
		return bodyErr, "400008"
	}

	if body.ShopName == utils.NULL_STRING {
		return "Missing shop_name", "400005"
	}

	if len(body.ShopName) > utils.SHOP_NAME_MAX_LEN {
		return "shop_name length greater than 255", "400007"
	}

	if body.OwnerName == utils.NULL_STRING {
		return "Missing owner_name", "400005"
	}

	if len(body.OwnerName) > utils.OWNER_NAME_MAX_LEN {
		return "owner_name length greater than 255", "400007"
	}

	if body.RegDate == utils.NULL_STRING {
		return "Missing reg_date", "400005"
	} else {
		_, err := time.Parse("2006-01-02", body.RegDate) // YYYY-MM-DD
		if err != nil {
			return "Invalid date format", "400006"
		}
	}

	if body.GstIN == utils.NULL_STRING {
		body.GstIN = "NOT PROVIDED   "
	}

	if len(body.GstIN) != utils.GST_IN_MAX_LEN {
		return "gst_in length not equal to 15", "400007"
	}

	if body.PhoneNo == utils.NULL_STRING {
		return "Missing phone_no", "400005"
	}

	if len(body.PhoneNo) > utils.PHONE_NO_MAX_LEN {
		return "phone_no length greater than 10", "400007"
	}

	if body.Address == utils.NULL_STRING {
		return "Missing address", "400005"
	}

	return utils.NULL_STRING, utils.SUCCESS
}

func ValidateAllShopOwnerBody(body *structs.AllShopOwner, bodyErr string) (string, string) {

	if bodyErr != utils.NULL_STRING {
		return bodyErr, "400008"
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
		return bodyErr, "400008"
	}

	if body.Name == utils.NULL_STRING {
		return "Missing name", "400005"
	}

	if len(body.Name) > utils.CUSTOMER_NAME_MAX_LEN {
		return "name length greater than 255", "400007"
	}

	if body.ShopName == utils.NULL_STRING {
		return "Missing shop_name", "400005"
	}

	if len(body.ShopName) > utils.SHOP_NAME_MAX_LEN {
		return "shop_name length greater than 255", "400007"
	}

	if body.GstIN == utils.NULL_STRING {
		body.GstIN = "NOT PROVIDED   "
	}

	if len(body.GstIN) != utils.GST_IN_MAX_LEN {
		return "gst_in length not equal to 15", "400007"
	}

	if body.PhoneNo == utils.NULL_STRING {
		return "Missing phone_no", "400005"
	} else if len(body.PhoneNo) != 10 {
		return "Invalid phone_no", "400006"
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

	if len(body.Address) > utils.ADDRESS_MAX_LEN {
		return "address length greater than 255", "400007"
	}

	if body.Remarks == utils.NULL_STRING {
		return "Missing remarks", "400005"
	}

	return utils.NULL_STRING, utils.SUCCESS
}

func ValidateFilteredCustomerReqBody(body *structs.FilteredCustomer, bodyErr string) (string, string) {

	if bodyErr != utils.NULL_STRING {
		return bodyErr, "400008"
	}

	id := body.Id
	regId := body.RegId
	name := body.Name
	ShopName := body.ShopName
	PhoneNo := body.PhoneNo
	regDate := body.RegDate
	isActive := body.IsActive
	dateInterval := body.DateInterval

	if (id == utils.NULL_INT) && (regId == utils.NULL_STRING) && (name == utils.NULL_STRING) && (ShopName == utils.NULL_STRING) && (PhoneNo == utils.NULL_STRING) && (regDate == utils.NULL_STRING) && (isActive == utils.NULL_STRING) {
		if (dateInterval.Type == utils.NULL_STRING) || (dateInterval.Type == utils.CUSTOM && (dateInterval.Start == utils.NULL_STRING || dateInterval.End == utils.NULL_STRING)) {
			return "Missing reqBody fields", "400005"
		}
	}

	if dateInterval.Type == utils.CUSTOM {
		_, err := time.Parse("2006-01-02", dateInterval.Start) // YYYY-MM-DD
		if err != nil {
			return "Invalid start date format", "400006"
		}

		_, err = time.Parse("2006-01-02", dateInterval.End) // YYYY-MM-DD
		if err != nil {
			return "Invalid end date format", "400006"
		}
	}

	return utils.NULL_STRING, utils.SUCCESS

}

func ValidatePostStockReqBody(body *structs.PostStock, bodyErr string) (string, string) {

	if bodyErr != utils.NULL_STRING {
		return bodyErr, "400008"
	}

	if body.ItemName == utils.NULL_STRING {
		return "Missing item_name", "400005"
	}

	if len(body.ItemName) > utils.ITEM_NAME_MAX_LEN {
		return "item_name length greater than 255", "400007"
	}

	if body.Type == utils.NULL_STRING {
		return "Missing type", "400005"
	}

	if body.Type != utils.GOLD && body.Type != utils.SILVER && body.Type != utils.CASH {
		return "Invalid type", "400006"
	}

	if body.Weight == utils.NULL_INT {
		return "Missing quantity", "400005"
	}

	return utils.NULL_STRING, utils.SUCCESS
}

func ValidatePutCustomerTransactionReqBody(body *structs.CustomerBill, bodyErr string) (string, string) {

	if bodyErr != utils.NULL_STRING {
		return bodyErr, "400008"
	}

	return utils.NULL_STRING, utils.SUCCESS
}
