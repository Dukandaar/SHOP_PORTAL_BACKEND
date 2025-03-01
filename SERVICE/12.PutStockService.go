package service

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	helper "SHOP_PORTAL_BACKEND/HELPER"
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"time"
)

func PutStock(reqBody structs.PutStock, ownerRegId string, stockId int, logPrefix string) (interface{}, int) {

	var response interface{}
	rspCode := utils.StatusOK

	DB := database.DB

	tx, err := DB.Begin()
	if err != nil {
		return helper.Create500ErrorResponse("[DB ERROR 0057] Error starting transaction", "Error starting transaction: "+err.Error(), logPrefix)
	}

	defer tx.Rollback()

	ownerRowId, err := helper.GetOwnerId(ownerRegId, tx) // Get ownerRowId
	if err != nil {
		return helper.Create500ErrorResponse("[DB ERROR 0058] Error getting owner row ID", "Error getting owner row ID: "+err.Error(), logPrefix)
	}

	// Stock id is validated, it is present, update values
	ServiceQuery := database.UpdateStockData()
	_, err = tx.Exec(ServiceQuery, reqBody.Tunch, reqBody.CurrentWeight, time.Now(), stockId, ownerRowId)
	if err != nil {
		return helper.Create500ErrorResponse("[DB ERROR 0059] Error in updating stock", "Error in updating stock: "+err.Error(), logPrefix)
	}

	ServiceQuery = database.InsertStockHistoryData()
	_, err = tx.Exec(ServiceQuery, stockId, reqBody.PrevWeight, reqBody.CurrentWeight, utils.BUY, "Updated Stock", time.Now())
	if err != nil {
		return helper.Create500ErrorResponse("[DB ERROR 0060] Error in inserting row", "Error in inserting row: "+err.Error(), logPrefix)
	} else {
		response, rspCode = helper.CreateSuccessWithIdResponse("Stock updated successfully", stockId, logPrefix)
	}

	if rspCode == utils.StatusOK {
		err = tx.Commit()
		if err != nil {
			return helper.Create500ErrorResponse("[DB ERROR 0061] Error committing transaction", "Error committing transaction: "+err.Error(), logPrefix)
		}
		utils.Logger.Info(logPrefix, "Transaction committed successfully")
	}

	return response, rspCode
}
