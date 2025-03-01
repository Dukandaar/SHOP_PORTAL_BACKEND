package service

import (
	helper "SHOP_PORTAL_BACKEND/HELPER"
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"database/sql"

	database "SHOP_PORTAL_BACKEND/DATABASE"
)

func GetFilteredCustomer(reqBody structs.FilteredCustomer, ownerRegID string, logPrefix string) (interface{}, int) {
	var response interface{}
	rspCode := utils.StatusOK

	DB := database.DB

	// Start transaction
	tx, err := DB.Begin()
	if err != nil {
		return helper.Create500ErrorResponse("[DB ERROR 0048] Error starting transaction", "Error starting transaction:"+err.Error(), logPrefix)
	}

	defer tx.Rollback()

	ownerRowID, err := helper.GetOwnerId(ownerRegID, tx)
	if err != nil {
		return helper.Create500ErrorResponse("[DB ERROR 0049] Error getting owner row ID", "Error getting owner row ID:"+err.Error(), logPrefix)
	}

	ServiceQuery := database.GetFilteredCustomerData(reqBody)
	rows, err := tx.Query(ServiceQuery, ownerRowID)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.CreateErrorResponse("404001", "Customer Not Found", logPrefix)
		}
		return helper.Create500ErrorResponse("[DB ERROR 0050] Error in getting filtered rows", "Error in getting filtered rows: "+err.Error(), logPrefix)
	}

	customerPayload := make([]structs.GetCustomerPayloadResponse, 0)
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
			response, rspCode = helper.CreateErrorResponse("500001", "Error in getting filtered rows", logPrefix)
			return response, rspCode
		}
		customerPayload = append(customerPayload, structs.GetCustomerPayloadResponse{
			ShopName: shopName,
			Name:     name,
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

	if rspCode == utils.StatusOK {
		err = tx.Commit()
		if err != nil {
			return helper.Create500ErrorResponse("[DB ERROR 0051] Error committing transaction", "Error committing transaction:"+err.Error(), logPrefix)
		}
		response = structs.GetAllCustomerResponse{
			Response: structs.GetAllCustomerSubResponse{
				Stat:        "OK",
				Count:       len(customerPayload),
				Payload:     customerPayload,
				Description: "All Customers Fetched Successfully",
			},
		}
	}

	return response, rspCode
}
