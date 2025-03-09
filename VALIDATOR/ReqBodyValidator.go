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
		errMsg := fmt.Sprintf("[DB ERROR 0000] Error in checking if row with reg_id %s exists : ", body.OwnerRegId)
		return helper.Create500ErrorResponse(errMsg, err.Error(), logPrefix)
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
		body.ShopName = "NOT PROVIDED"
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
		body.Address = "NOT PROVIDED"
	}

	if len(body.Address) > utils.ADDRESS_MAX_LEN {
		return helper.CreateErrorResponse("400007", "address length greater than 255", logPrefix)
	}

	if body.Remarks == utils.NULL_STRING {
		body.Remarks = "NOT PROVIDED"
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

	if body.Weight == nil {
		return helper.CreateErrorResponse("400005", "Missing weight", logPrefix)
	}

	return utils.NULL_STRING, utils.StatusOK
}

func ValidatePutStockReqBody(body *structs.PutStock, logPrefix string) (interface{}, int) {

	if body.PrevWeight == nil {
		return helper.CreateErrorResponse("400005", "Missing prev_weight", logPrefix)
	}

	if body.CurrentWeight == nil {
		return helper.CreateErrorResponse("400005", "Missing current_weight", logPrefix)
	}

	return utils.NULL_STRING, utils.StatusOK
}

func ValidateTransactionReqBody(body *[]structs.Transaction, logPrefix string) (interface{}, int) {

	transLength := len(*body)

	if transLength == 0 {
		return helper.CreateErrorResponse("400005", "Missing transactions", logPrefix)
	}

	for i := 0; i < transLength; i++ {
		transaction := (*body)[i]

		if transaction.ItemName == utils.NULL_STRING {
			return helper.CreateErrorResponse("400005", "Missing item_name", logPrefix)
		}

		if len(transaction.ItemName) > utils.ITEM_NAME_MAX_LEN {
			return helper.CreateErrorResponse("400007", "item_name length greater than 255", logPrefix)
		}

		if transaction.Weight == nil {
			return helper.CreateErrorResponse("400005", "Missing weight", logPrefix)
		}

		if *transaction.Weight < 0 || *transaction.Weight > utils.MAX_FLOAT {
			return helper.CreateErrorResponse("400006", "Invalid weight", logPrefix)
		}

		if transaction.Less == nil {
			return helper.CreateErrorResponse("400005", "Missing less", logPrefix)
		}

		if *transaction.Less < 0 || *transaction.Less > utils.MAX_FLOAT {
			return helper.CreateErrorResponse("400006", "Invalid less", logPrefix)
		}

		if transaction.NetWeight == nil {
			return helper.CreateErrorResponse("400005", "Missing net_weight", logPrefix)
		}

		if *transaction.NetWeight < 0 || *transaction.NetWeight > utils.MAX_FLOAT {
			return helper.CreateErrorResponse("400006", "Invalid net_weight", logPrefix)
		}

		if transaction.Tunch == nil {
			return helper.CreateErrorResponse("400005", "Missing tunch", logPrefix)
		}

		if *transaction.Tunch < 0 || *transaction.Tunch > utils.MAX_FLOAT {
			return helper.CreateErrorResponse("400006", "Invalid tunch", logPrefix)
		}

		if transaction.Fine == nil {
			return helper.CreateErrorResponse("400005", "Missing fine", logPrefix)
		}

		if *transaction.Fine < 0 || *transaction.Fine > utils.MAX_FLOAT {
			return helper.CreateErrorResponse("400006", "Invalid fine", logPrefix)
		}

		if transaction.Discount == nil {
			return helper.CreateErrorResponse("400005", "Missing discount", logPrefix)
		}

		if *transaction.Discount < 0 || *transaction.Discount > utils.MAX_FLOAT {
			return helper.CreateErrorResponse("400006", "Invalid discount", logPrefix)
		}

		if transaction.Amount == nil {
			return helper.CreateErrorResponse("400005", "Missing amount", logPrefix)
		}

		if *transaction.Amount < 0 || *transaction.Amount > utils.MAX_FLOAT {
			return helper.CreateErrorResponse("400006", "Invalid amount", logPrefix)
		}

		if transaction.IsActive == utils.NULL_STRING {
			return helper.CreateErrorResponse("400005", "Missing is_active", logPrefix)
		}

		if transaction.IsActive != utils.ACTIVE_YES && transaction.IsActive != utils.ACTIVE_NO {
			return helper.CreateErrorResponse("400006", "Invalid is_active", logPrefix)
		}
	}

	return utils.NULL_STRING, utils.StatusOK
}

func ValidatePaymentReqBody(body *structs.Payment, logPrefix string) (interface{}, int) {

	if body.Factor == utils.NULL_STRING {
		return helper.CreateErrorResponse("400005", "Missing factor", logPrefix)
	}

	if body.Factor != utils.AMOUNT && body.Factor != utils.FINE {
		return helper.CreateErrorResponse("400006", "Invalid factor", logPrefix)
	}

	if body.New == nil {
		return helper.CreateErrorResponse("400005", "Missing new", logPrefix)
	}

	if *body.New < 0 || *body.New > utils.MAX_FLOAT {
		return helper.CreateErrorResponse("400006", "Invalid new", logPrefix)
	}

	if body.Prev == nil {
		return helper.CreateErrorResponse("400005", "Missing prev", logPrefix)
	}

	if *body.Prev < 0 || *body.Prev > utils.MAX_FLOAT {
		return helper.CreateErrorResponse("400006", "Invalid prev", logPrefix)
	}

	if body.Total == nil {
		return helper.CreateErrorResponse("400005", "Missing total", logPrefix)
	}

	if *body.Total < 0 || *body.Total > utils.MAX_FLOAT {
		return helper.CreateErrorResponse("400006", "Invalid total", logPrefix)
	}

	if body.Paid == nil {
		return helper.CreateErrorResponse("400005", "Missing paid", logPrefix)
	}

	if *body.Paid < 0 || *body.Paid > utils.MAX_FLOAT {
		return helper.CreateErrorResponse("400006", "Invalid paid", logPrefix)
	}

	if body.Rem == nil {
		return helper.CreateErrorResponse("400005", "Missing rem", logPrefix)
	}

	if *body.Rem < 0 || *body.Rem > utils.MAX_FLOAT {
		return helper.CreateErrorResponse("400006", "Invalid rem", logPrefix)
	}

	return utils.NULL_STRING, utils.StatusOK
}

func ValidatePostCustomerBillReqBody(body *structs.CustomerBill, logPrefix string) (interface{}, int) {

	if body.BillNo == utils.NULL_INT {
		return helper.CreateErrorResponse("400005", "Missing bill_no", logPrefix)
	}

	if body.BillNo < 0 || body.BillNo > utils.MAX_INT {
		return helper.CreateErrorResponse("400006", "Invalid bill_no", logPrefix)
	}

	if body.Type == utils.NULL_STRING {
		return helper.CreateErrorResponse("400005", "Missing type", logPrefix)
	}

	if body.Type != utils.WHOLESALE && body.Type != utils.RETAIL {
		return helper.CreateErrorResponse("400006", "Invalid type", logPrefix)
	}

	if body.Metal == utils.NULL_STRING {
		return helper.CreateErrorResponse("400005", "Missing metal_type", logPrefix)
	}

	if body.Metal != utils.GOLD && body.Metal != utils.SILVER && body.Metal != utils.BOTH {
		return helper.CreateErrorResponse("400006", "Invalid metal_type", logPrefix)
	}

	if body.Rate == nil {
		return helper.CreateErrorResponse("400005", "Missing rate", logPrefix)
	}

	if *body.Rate < 0 || *body.Rate > utils.MAX_FLOAT {
		return helper.CreateErrorResponse("400006", "Invalid rate", logPrefix)
	}

	if body.Date == utils.NULL_STRING {
		return helper.CreateErrorResponse("400005", "Missing date", logPrefix)
	}

	_, err := time.Parse("2006-01-02", body.Date) // YYYY-MM-DD
	if err != nil {
		return helper.CreateErrorResponse("400006", "Invalid date format", logPrefix)
	}

	if body.Remarks == utils.NULL_STRING {
		body.Remarks = "NOT PROVIDED"
	}

	rsp, code := ValidateCustomerReqBody(&body.CustomerDetails, logPrefix)
	if code != utils.StatusOK {
		return rsp, code
	}

	rsp, code = ValidateTransactionReqBody(&body.TransactionDetails, logPrefix)
	if code != utils.StatusOK {
		return rsp, code
	}

	rsp, code = ValidatePaymentReqBody(&body.PaymentDetails, logPrefix)
	if code != utils.StatusOK {
		return rsp, code
	}

	return utils.NULL_STRING, utils.StatusOK
}

func ValidateGetAllStockReqBody(body *structs.AllStock, logPrefix string) (interface{}, int) {

	if body.Metal != utils.GOLD && body.Metal != utils.SILVER && body.Metal != utils.ALL {
		return helper.CreateErrorResponse("400006", "Invalid metal", logPrefix)
	}

	return utils.NULL_STRING, utils.StatusOK
}
