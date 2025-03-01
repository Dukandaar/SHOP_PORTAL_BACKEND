package service

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	helper "SHOP_PORTAL_BACKEND/HELPER"
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"database/sql"
)

func GetCustomer(ownerRegID string, customerRegID string, logPrefix string) (interface{}, int) {
	var response interface{}
	rspCode := utils.StatusOK

	var shopName string
	var name string
	var gstIN string
	var phoneNo string
	var regDate string
	var address string
	var remarks string
	var gold float64
	var silver float64
	var cash float64
	var isActive string

	DB := database.DB

	tx, err := DB.Begin()
	if err != nil {
		return helper.Create500ErrorResponse("[DB ERROR 0039] Error starting transaction", "Error starting transaction: "+err.Error(), logPrefix)
	}
	defer tx.Rollback()

	ownerRowID, err := helper.GetOwnerId(ownerRegID, tx)
	if err != nil {
		return helper.Create500ErrorResponse("[DB ERROR 0040] Error getting owner row ID", "Error getting owner row ID:"+err.Error(), logPrefix)
	}

	serviceQuery := database.GetCustomerData()
	err = tx.QueryRow(serviceQuery, customerRegID, ownerRowID).Scan(&shopName, &name, &gstIN, &regDate, &phoneNo, &isActive, &address, &remarks, &gold, &silver, &cash)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.CreateErrorResponse("404001", "Data for reg_id "+customerRegID+" does not exist", logPrefix)
		} else {
			return helper.Create500ErrorResponse("[DB ERROR 0041] Error in getting row", "Error in getting row: "+err.Error(), logPrefix)
		}
	}

	if rspCode == utils.StatusOK {
		err = tx.Commit()
		if err != nil {
			return helper.Create500ErrorResponse("[DB ERROR 0042] Error committing transaction", "Error committing transaction: "+err.Error(), logPrefix)
		}

		customerPayload := structs.GetCustomerPayloadResponse{
			ShopName: shopName,
			Name:     name,
			GstIN:    gstIN,
			RegId:    customerRegID,
			PhoneNo:  phoneNo,
			RegDate:  regDate,
			Address:  address,
			Remarks:  remarks,
			Gold:     gold,
			Silver:   silver,
			Cash:     cash,
			IsActive: isActive,
		}

		response, rspCode = helper.CreateGetCustomerResponse(customerPayload, "Data for reg_id "+customerRegID+" fetched successfully", logPrefix)
	}

	return response, rspCode
}
