package service

import (
	helper "SHOP_PORTAL_BACKEND/HELPER"
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"

	database "SHOP_PORTAL_BACKEND/DATABASE"
)

func GetFilteredCustomer(reqBody structs.FilteredCustomer, ownerRegID string, logPrefix string) (interface{}, int) {
	var response interface{}
	rspCode := utils.StatusOK

	DB := database.DB

	// Start transaction
	tx, err := DB.Begin()
	if err != nil {
		return helper.Set500ErrorResponse("Error starting transaction", "Error starting transaction:"+err.Error(), logPrefix)
	}

	defer func() {
		if r := recover(); r != nil || err != nil {
			utils.Logger.Error(logPrefix, "Panic occurred during transaction:", r, err)
			tx.Rollback()
		}
	}()

	ServiceQuery := database.GetOwnerRowId() // Get Owner's row ID
	var ownerRowId string
	err = tx.QueryRow(ServiceQuery, ownerRegID).Scan(&ownerRowId)
	if err != nil {
		return helper.Set500ErrorResponse("Error getting owner row ID", "Error getting owner row ID:"+err.Error(), logPrefix)
	}

	ServiceQuery = database.GetFilteredCustomerData(reqBody)
	rows, err := tx.Query(ServiceQuery, ownerRowId)
	if err != nil {
		utils.Logger.Error(err.Error())
		response, rspCode = helper.CreateErrorResponse("500001", "Error in getting filtered rows")
		return response, rspCode
	}

	rsp := make([]structs.CustomerDetailsSubResponse, 0)
	for rows.Next() {
		var name string
		var shopName string
		var GstIN string
		var regId string
		var phoneNo string
		var regDate string
		var address string
		var remarks string
		var gold float64
		var silver float64
		var cash float64
		var isActive string
		err = rows.Scan(&shopName, &name, &GstIN, &regId, &regDate, &phoneNo, &isActive, &address, &remarks, &gold, &silver, &cash)
		if err != nil {
			utils.Logger.Error(err.Error())
			response, rspCode = helper.CreateErrorResponse("500001", "Error in getting filtered rows")
			return response, rspCode
		}
		rsp = append(rsp, structs.CustomerDetailsSubResponse{
			Name:     name,
			ShopName: shopName,
			GstIN:    GstIN,
			RegId:    regId,
			PhoneNo:  phoneNo,
			RegDate:  regDate,
			Address:  address,
			Remarks:  remarks,
			Gold:     gold,
			Silver:   silver,
			Cash:     cash,
			IsActive: isActive,
		})
	}

	response = structs.AllCustomerDetailsResponse{
		Stat:                       "OK",
		Count:                      len(rsp),
		CustomerDetailsSubResponse: rsp,
	}
	return response, rspCode
}
