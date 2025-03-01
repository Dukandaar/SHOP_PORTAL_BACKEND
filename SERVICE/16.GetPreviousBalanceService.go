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
		return helper.Set500ErrorResponse("Error starting transaction", "Error starting transaction:"+err.Error(), logPrefix)
	}
	defer func() {
		if r := recover(); r != nil || err != nil {
			utils.Logger.Error(logPrefix, "Panic occurred during transaction:", r, err)
			tx.Rollback()
		}
	}()

	ServiceQuery := database.GetOwnerRowId() // Get Owner's row ID
	var ownerRowId int
	err = tx.QueryRow(ServiceQuery, ownerRegId).Scan(&ownerRowId)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.CreateErrorResponse("404001", "Owner Not Found", logPrefix)
		}
		return helper.Set500ErrorResponse("Error getting owner row ID", "Error getting owner row ID:"+err.Error(), logPrefix)
	}

	// Get customer_id from reg_id
	ServiceQuery = database.GetCustomerId()
	var customerId int
	err = tx.QueryRow(ServiceQuery, customerRegId, ownerRowId).Scan(&customerId)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.CreateErrorResponse("404002", "Customer Not Found", logPrefix)
		}
		return helper.Set500ErrorResponse("Error getting owner row ID", "Error getting owner row ID:"+err.Error(), logPrefix)
	}

	var gold float64
	var silver float64
	var cash float64

	ServiceQuery = database.GetCustomerPreviousBalance()
	err = tx.QueryRow(ServiceQuery, customerId).Scan(&gold, &silver, &cash)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.CreateErrorResponse("404003", "Previous Balance Not Found", logPrefix)
		}
		return helper.Set500ErrorResponse("Error getting previous balance", "Error getting previous balance:"+err.Error(), logPrefix)
	}

	if rspCode == utils.StatusOK {
		response = structs.CustomerPreviousBalanceResponse{
			Stat: "OK",
			CustomerPreviousBalanceSubResponse: []structs.CustomerPreviousBalanceSubResponse{
				{
					RowId:  customerId,
					Gold:   gold,
					Silver: silver,
					Cash:   cash,
				},
			},
		}

		err = tx.Commit()
		if err != nil {
			return helper.Set500ErrorResponse("Error committing transaction", "Error committing transaction:"+err.Error(), logPrefix)
		}
	}

	return response, rspCode
}
