package service

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	helper "SHOP_PORTAL_BACKEND/HELPER"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"database/sql"
)

func DeleteStock(ownerRegId string, stockId int, logPrefix string) (interface{}, int) {

	var response interface{}
	rspCode := utils.StatusOK

	DB := database.DB

	tx, err := DB.Begin()
	if err != nil {
		return helper.Create500ErrorResponse("[DB ERROR 0061] Error starting transaction", "Error starting transaction: "+err.Error(), logPrefix)
	}

	defer tx.Rollback()

	ownerRowId, err := helper.GetOwnerId(ownerRegId, tx)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.CreateErrorResponse("404001", "Owner not found", logPrefix)
		}
		return helper.Create500ErrorResponse("[DB ERROR 0140] Error getting owner row ID", "Error getting owner row ID: "+err.Error(), logPrefix)
	}

	ServiceQuery := database.DeleteStock()
	stockName := utils.NULL_STRING
	err = tx.QueryRow(ServiceQuery, stockId, ownerRowId).Scan(&stockName)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.CreateErrorResponse("404001", "Stock not found for this owner", logPrefix)
		}
		return helper.Create500ErrorResponse("[DB ERROR 0141] Error in deleting stock", "Error in deleting stock: "+err.Error(), logPrefix)
	}

	if rspCode == utils.StatusOK {
		err = tx.Commit()
		if err != nil {
			return helper.Create500ErrorResponse("[DB ERROR 0142] Error committing transaction", "Error committing transaction: "+err.Error(), logPrefix)
		}

		response, rspCode = helper.CreateSuccessResponse("Stock deleted successfully", stockName+" deleted successfully.", logPrefix)
	}

	return response, rspCode
}
