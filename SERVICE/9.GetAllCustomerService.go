package service

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	helper "SHOP_PORTAL_BACKEND/HELPER"
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"database/sql"
)

func GetAllCustomer(owner_reg_id string, logPrefix string) (interface{}, int) {

	var response interface{}
	rspCode := utils.StatusOK

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

	customerPayload := make([]structs.GetCustomerPayloadResponse, 0)

	DB := database.DB

	tx, err := DB.Begin()
	if err != nil {
		return helper.Create500ErrorResponse("[DB ERROR 0043] Error starting transaction", "Error starting transaction:"+err.Error(), logPrefix)
	}
	defer tx.Rollback()

	ownerRowId, err := helper.GetOwnerId(owner_reg_id, tx)
	if err != nil {
		return helper.Create500ErrorResponse("[DB ERROR 0044] Error getting owner row ID", "Error getting owner row ID:"+err.Error(), logPrefix)
	}

	ServiceQuery := database.GetAllCustomerData()
	rows, err := tx.Query(ServiceQuery, ownerRowId)
	if err == nil {
		for rows.Next() {

			err = rows.Scan(&shopName, &name, &GstIN, &regId, &phoneNo, &regDate, &isActive, &address, &remarks, &gold, &silver, &cash)
			if err != nil {
				utils.Logger.Error(err.Error())
				return helper.Create500ErrorResponse("[DB ERROR 0045] Error in getting rows", "Error in getting rows: "+err.Error(), logPrefix)
			} else {

				customerPayload = append(customerPayload, structs.GetCustomerPayloadResponse{
					Name:     name,
					ShopName: shopName,
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
		}
	} else {
		if err == sql.ErrNoRows {
			return helper.CreateSuccessResponse("No any customer found", "No customer found for owner reg id : "+owner_reg_id, logPrefix)
		} else {
			return helper.Create500ErrorResponse("[DB ERROR 0046] Error in getting rows", "Error in getting rows:"+err.Error(), logPrefix)
		}
	}

	if rspCode == utils.StatusOK {
		err = tx.Commit()
		if err != nil {
			return helper.Create500ErrorResponse("[DB ERROR 0047] Error in committing transaction", "Error in committing transaction: "+err.Error(), logPrefix)
		}
		utils.Logger.Info(logPrefix + "Transaction committed successfully")
		response, rspCode = helper.CreateGetAllCustomerResponse(customerPayload, "All Customer found successfully for owner with regId : "+owner_reg_id, logPrefix)
	}

	return response, rspCode
}
