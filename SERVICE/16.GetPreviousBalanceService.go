package service

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	helper "SHOP_PORTAL_BACKEND/HELPER"
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"database/sql"
)

func GetPreviousBalance(ownerRegId string, customerRegId string, logPrefix string) (interface{}, int) {

	var response interface{}
	rspCode := utils.StatusOK

	Db := database.DB

	tx, err := Db.Begin()
	if err != nil {
		return helper.Create500ErrorResponse("[DB ERROR 0078] Error starting transaction", "Error starting transaction:"+err.Error(), logPrefix)
	}
	defer tx.Rollback()

	// Get owner row id
	ownerRowId, err := helper.GetOwnerId(ownerRegId, tx)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.CreateErrorResponse("404001", "Owner not found", logPrefix)
		}
		return helper.Create500ErrorResponse("[DB ERROR 0079] Error getting owner row ID", "Error getting owner row ID:"+err.Error(), logPrefix)
	}

	// Get customer_id from reg_id
	customerId, err := helper.GetCustomerId(customerRegId, ownerRowId, tx)
	if err != nil {
		return helper.Create500ErrorResponse("[DB ERROR 0080] Error getting customer row ID", "Error getting customer row ID:"+err.Error(), logPrefix)
	}

	var gold float64
	var silver float64
	var cash float64
	var id int

	ServiceQuery := database.GetCustomerPreviousBalance()
	err = tx.QueryRow(ServiceQuery, customerId).Scan(&id, &gold, &silver, &cash)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.CreateErrorResponse("404003", "Previous Balance Not Found", logPrefix)
		}
		return helper.Create500ErrorResponse("[DB ERROR 0081] Error getting previous balance", "Error getting previous balance: "+err.Error(), logPrefix)
	}

	if rspCode == utils.StatusOK {
		err = tx.Commit()
		if err != nil {
			return helper.Create500ErrorResponse("[DB ERROR 0082] Error committing transaction", "Error committing transaction:"+err.Error(), logPrefix)
		}

		utils.Logger.Info(logPrefix, "Transaction committed Successfully")

		response = structs.CustomerPreviousBalanceResponse{
			Response: structs.CustomerPreviousBalanceSubResponse{
				Stat: utils.OK,
				Payload: structs.CustomerPreviousBalancePayloadResponse{
					RowId:  id,
					Gold:   gold,
					Silver: silver,
					Cash:   cash,
				},
				Description: "Previous Balance Fetched Successfully",
			},
		}
	}

	return response, rspCode
}
