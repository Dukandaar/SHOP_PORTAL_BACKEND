package service

import (
	helper "SHOP_PORTAL_BACKEND/HELPER"
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"

	database "SHOP_PORTAL_BACKEND/DATABASE"
)

func GetFilteredCustomer(reqBody structs.FilteredCustomer, ownerRegID string) (interface{}, int) {
	var response interface{}
	rspCode := utils.StatusOK

	DB := database.ConnectDB()
	defer DB.Close()

	ServiceQuery := database.GetOwnerRowId() // Get Owner's row ID
	var ownerRowId string
	err := DB.QueryRow(ServiceQuery, ownerRegID).Scan(&ownerRowId)
	if err != nil {
		return helper.SetErrorResponse("Error getting owner row ID", "Error getting owner row ID:"+err.Error())
	}

	ServiceQuery = database.GetFilteredCustomerData(reqBody)
	rows, err := DB.Query(ServiceQuery, ownerRowId)
	if err != nil {
		utils.Logger.Error(err.Error())
		response, rspCode = helper.CreateErrorResponse("500001", "Error in getting filtered rows")
		return response, rspCode
	}

	rsp := make([]structs.CustomerDetailsSubResponse, 0)
	for rows.Next() {
		var name string
		var shopName string
		var regId string
		var phoneNo string
		var regDate string
		var address string
		var remarks string
		var gold float32
		var silver float32
		var cash float32
		var isActive string
		err = rows.Scan(&name, &shopName, &regId, &phoneNo, &regDate, &address, &remarks, &gold, &silver, &cash, &isActive)
		if err != nil {
			utils.Logger.Error(err.Error())
			response, rspCode = helper.CreateErrorResponse("500001", "Error in getting filtered rows")
			return response, rspCode
		}
		rsp = append(rsp, structs.CustomerDetailsSubResponse{
			Name:        name,
			CompanyName: shopName,
			RegId:       regId,
			PhNo:        phoneNo,
			RegDate:     regDate,
			Address:     address,
			Remarks:     remarks,
			Gold:        gold,
			Silver:      silver,
			Cash:        cash,
			IsActive:    isActive,
		})
	}

	response = structs.AllCustomerDetailsResponse{
		Stat:                       "OK",
		Count:                      len(rsp),
		CustomerDetailsSubResponse: rsp,
	}
	return response, rspCode
}
