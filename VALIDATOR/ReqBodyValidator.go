package validator

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	helper "SHOP_PORTAL_BACKEND/HELPER"
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"fmt"
	"time"
)

func ValidateGenerateTokenReqBody(body *structs.GenerateToken, logPrefix string) (interface{}, int) {

	if body.OwnerRegId == utils.NULL_STRING {
		return helper.CreateErrorResponse("400005", "Missing reg_id", logPrefix)
	}

	if len(body.OwnerRegId) > utils.OWNER_REG_ID_MAX_LEN {
		return helper.CreateErrorResponse("400007", "reg_id length greater than 10", logPrefix)
	}

	var exists bool
	ServiceQuery := database.CheckValidOwnerRegId()
	db := database.DB
	err := db.QueryRow(ServiceQuery, body.OwnerRegId).Scan(&exists)
	if err != nil {
		errMsg := fmt.Sprintf("[DB ERROR] Error in checking if row with reg_id %s exists : ", body.OwnerRegId)
		return helper.Set500ErrorResponse(errMsg, err.Error(), logPrefix)
	}

	if exists {
		utils.Logger.Info(logPrefix, "Row with owner reg_id : ", body.OwnerRegId, " exists")
	} else {
		utils.Logger.Info(logPrefix, "Row with owner reg_id ", body.OwnerRegId, " does not exist")
		return helper.CreateErrorResponse("404001", "Data for reg_id "+body.OwnerRegId+" does not exist", logPrefix)
	}

	if body.Key == utils.NULL_STRING {
		return helper.CreateErrorResponse("400005", "Missing key", logPrefix)
	} else {
		errMsg, errCodeStr := ValidateKey(body.Key, body.OwnerRegId, logPrefix)
		if errMsg != utils.NULL_STRING {
			return helper.CreateErrorResponse(errCodeStr, errMsg, logPrefix)
		}
	}

	return utils.NULL_STRING, utils.StatusOK
}

func ValidateShopOwnerReqBody(body *structs.ShopOwner, logPrefix string) (interface{}, int) {

	if body.ShopName == utils.NULL_STRING {
		return helper.CreateErrorResponse("400005", "Missing shop_name", logPrefix)
	}

	if len(body.ShopName) > utils.SHOP_NAME_MAX_LEN {
		return helper.CreateErrorResponse("400007", "shop_name length greater than 255", logPrefix)
	}

	if body.OwnerName == utils.NULL_STRING {
		return helper.CreateErrorResponse("400005", "Missing owner_name", logPrefix)
	}

	if len(body.OwnerName) > utils.OWNER_NAME_MAX_LEN {
		return helper.CreateErrorResponse("400007", "owner_name length greater than 255", logPrefix)
	}

	if body.RegDate == utils.NULL_STRING {
		return helper.CreateErrorResponse("400005", "Missing reg_date", logPrefix)
	} else {
		_, err := time.Parse("2006-01-02", body.RegDate) // YYYY-MM-DD
		if err != nil {
			return helper.CreateErrorResponse("400006", "Invalid date format", logPrefix)
		}
	}

	if body.GstIN == utils.NULL_STRING {
		body.GstIN = "NOT PROVIDED   "
	}

	if len(body.GstIN) != utils.GST_IN_MAX_LEN {
		return helper.CreateErrorResponse("400007", "gst_in length not equal to 15", logPrefix)
	}

	if body.PhoneNo == utils.NULL_STRING {
		return helper.CreateErrorResponse("400005", "Missing phone_no", logPrefix)
	}

	if len(body.PhoneNo) > utils.PHONE_NO_MAX_LEN {
		return helper.CreateErrorResponse("400007", "phone_no length greater than 10", logPrefix)
	}

	if body.Address == utils.NULL_STRING {
		return helper.CreateErrorResponse("400005", "Missing address", logPrefix)
	}

	if len(body.Address) > utils.ADDRESS_MAX_LEN {
		return helper.CreateErrorResponse("400007", "address length greater than 255", logPrefix)
	}

	return utils.NULL_STRING, utils.StatusOK
}

func ValidateAllShopOwnerBody(body *structs.AllShopOwner, logPrefix string) (interface{}, int) {

	if body.IsActive == utils.NULL_STRING {
		return helper.CreateErrorResponse("400005", "Missing is_active", logPrefix)
	} else {
		if body.IsActive != utils.ACTIVE_YES && body.IsActive != utils.ACTIVE_NO && body.IsActive != utils.ALL {
			return helper.CreateErrorResponse("400006", "Invalid is_active", logPrefix)
		}
	}

	return utils.NULL_STRING, utils.StatusOK
}

func ValidateCustomerReqBody(body *structs.Customer, logPrefix string) (interface{}, int) {

	if body.Name == utils.NULL_STRING {
		return helper.CreateErrorResponse("400005", "Missing name", logPrefix)
	}

	if len(body.Name) > utils.CUSTOMER_NAME_MAX_LEN {
		return helper.CreateErrorResponse("400007", "name length greater than 255", logPrefix)
	}

	if body.ShopName == utils.NULL_STRING {
		return helper.CreateErrorResponse("400005", "Missing shop_name", logPrefix)
	}

	if len(body.ShopName) > utils.SHOP_NAME_MAX_LEN {
		return helper.CreateErrorResponse("400007", "shop_name length greater than 255", logPrefix)
	}

	if body.GstIN == utils.NULL_STRING {
		body.GstIN = "NOT PROVIDED   "
	}

	if len(body.GstIN) != utils.GST_IN_MAX_LEN {
		return helper.CreateErrorResponse("400007", "gst_in length not equal to 15", logPrefix)
	}

	if body.PhoneNo == utils.NULL_STRING {
		return helper.CreateErrorResponse("400005", "Missing phone_no", logPrefix)
	}

	if len(body.PhoneNo) != 10 {
		return helper.CreateErrorResponse("400006", "Invalid phone_no", logPrefix)
	}

	if body.RegDate == utils.NULL_STRING {
		return helper.CreateErrorResponse("400005", "Missing reg_date", logPrefix)
	} else {
		_, err := time.Parse("2006-01-02", body.RegDate) // YYYY-MM-DD
		if err != nil {
			return helper.CreateErrorResponse("400006", "Invalid date format", logPrefix)
		}
	}

	if body.Address == utils.NULL_STRING {
		return helper.CreateErrorResponse("400005", "Missing address", logPrefix)
	}

	if len(body.Address) > utils.ADDRESS_MAX_LEN {
		return helper.CreateErrorResponse("400007", "address length greater than 255", logPrefix)
	}

	if body.Remarks == utils.NULL_STRING {
		return helper.CreateErrorResponse("400005", "Missing remarks", logPrefix)
	}

	return utils.NULL_STRING, utils.StatusOK
}

func ValidateFilteredCustomerReqBody(body *structs.FilteredCustomer, logPrefix string) (interface{}, int) {

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
			return helper.CreateErrorResponse("400005", "Missing reqBody fields", logPrefix)
		}
	}

	if dateInterval.Type == utils.CUSTOM {
		_, err := time.Parse("2006-01-02", dateInterval.Start) // YYYY-MM-DD
		if err != nil {
			return helper.CreateErrorResponse("400006", "Invalid start date format", logPrefix)
		}

		_, err = time.Parse("2006-01-02", dateInterval.End) // YYYY-MM-DD
		if err != nil {
			return helper.CreateErrorResponse("400006", "Invalid end date format", logPrefix)
		}
	}

	return utils.NULL_STRING, utils.StatusOK

}

func ValidatePostStockReqBody(body *structs.PostStock, logPrefix string) (interface{}, int) {

	if body.ItemName == utils.NULL_STRING {
		return helper.CreateErrorResponse("400005", "Missing item_name", logPrefix)
	}

	if len(body.ItemName) > utils.ITEM_NAME_MAX_LEN {
		return helper.CreateErrorResponse("400007", "item_name length greater than 255", logPrefix)
	}

	if body.Type == utils.NULL_STRING {
		return helper.CreateErrorResponse("400005", "Missing type", logPrefix)
	}

	if body.Type != utils.GOLD && body.Type != utils.SILVER && body.Type != utils.CASH {
		return helper.CreateErrorResponse("400006", "Invalid type", logPrefix)
	}

	if body.Weight == utils.NULL_INT {
		return helper.CreateErrorResponse("400005", "Missing weight", logPrefix)
	}

	return utils.NULL_STRING, utils.StatusOK
}

func ValidatePutStockReqBody(body *structs.PutStock, logPrefix string) (interface{}, int) {

	if body.PrevWeight == utils.NULL_INT {
		return helper.CreateErrorResponse("400005", "Missing prev_weight", logPrefix)
	}

	if body.CurrentWeight == utils.NULL_INT {
		return helper.CreateErrorResponse("400005", "Missing current_weight", logPrefix)
	}

	return utils.NULL_STRING, utils.StatusOK
}

func ValidatePostCustomerBillReqBody(body *structs.CustomerBill, logPrefix string) (interface{}, int) {
	return utils.NULL_STRING, utils.StatusOK
}

func ValidateGetAllStockReqBody(body *structs.AllStock, logPrefix string) (interface{}, int) {

	if body.Type == utils.NULL_STRING {
		return helper.CreateErrorResponse("400005", "Missing type", logPrefix)
	}

	if body.Type != utils.GOLD && body.Type != utils.SILVER && body.Type != utils.ALL {
		return helper.CreateErrorResponse("400006", "Invalid type", logPrefix)
	}

	return utils.NULL_STRING, utils.StatusOK
}
