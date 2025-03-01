package service

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	helper "SHOP_PORTAL_BACKEND/HELPER"
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"time"
)

func PutCustomerBill(reqBody structs.CustomerBill, ownerRegId string, customerRegId string, billId int, logPrefix string) (interface{}, int) {

	var response interface{}
	rspCode := utils.StatusOK

	DB := database.DB

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

	// Get owner row id
	ownerRowId, err := helper.GetOwnerId(ownerRegId, tx)
	if err != nil {
		return helper.Set500ErrorResponse("Error getting owner row ID", "Error getting owner row ID:"+err.Error(), logPrefix)
	}

	// Get customer row id
	customerRowId, err := helper.GetCustomerId(customerRegId, ownerRowId, tx)
	if err != nil {
		return helper.Set500ErrorResponse("Error getting customer row ID", "Error getting customer row ID:"+err.Error(), logPrefix)
	}

	// Update bill using bill id
	ServiceQuery := database.UpdateBill()
	_, err = tx.Exec(ServiceQuery, reqBody.Type, reqBody.Metal, reqBody.Rate, reqBody.Date, time.Now(), billId)
	if err != nil {
		return helper.Set500ErrorResponse("Error in updating bill", "Error in updating bill:"+err.Error(), logPrefix)
	}
	utils.Logger.Info(logPrefix, "Bill updated successfully")

	// Get transactoins for that bill
	ServiceQuery = database.GetBillTransactions()
	rows, err := tx.Query(ServiceQuery, billId)
	if err != nil {
		return helper.Set500ErrorResponse("Error in getting transactions", "Error in getting transactions:"+err.Error(), logPrefix)
	}
	defer rows.Close()

	IdNameMap := make(map[int]string)
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			return helper.Set500ErrorResponse("Error in scanning row", "Error in scanning row:"+err.Error(), logPrefix)
		}
		IdNameMap[id] = name
	}
	utils.Logger.Info(logPrefix, ownerRowId, customerRowId, IdNameMap)

	return response, rspCode
}
