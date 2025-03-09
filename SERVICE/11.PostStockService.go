package service

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	helper "SHOP_PORTAL_BACKEND/HELPER"
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"database/sql"
	"time"
)

func PostStock(reqBody structs.PostStock, ownerRegId string, logPrefix string) (interface{}, int) {
	var response interface{}
	rspCode := utils.StatusOK

	DB := database.DB

	tx, err := DB.Begin()
	if err != nil {
		return helper.Create500ErrorResponse("[DB ERROR 0051] Error starting transaction", "Error starting transaction: "+err.Error(), logPrefix)
	}

	defer tx.Rollback()

	ownerRowId, err := helper.GetOwnerId(ownerRegId, tx) // Get ownerRowId
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.CreateErrorResponse("404001", "Owner not found", logPrefix)
		}
		return helper.Create500ErrorResponse("[DB ERROR 0052] Error in getting row", "Error getting owner row ID: "+err.Error(), logPrefix)
	}

	// check stock with same item_name
	ServiceQuery := database.CheckStockPresent() // tunch validation not added now : todo
	var rowId int
	err = tx.QueryRow(ServiceQuery, ownerRowId, reqBody.Type, reqBody.ItemName).Scan(&rowId)
	if err != nil {
		if err == sql.ErrNoRows { // Stock NOT found
			utils.Logger.Info(logPrefix, "Stock Not found")
		} else {
			return helper.Create500ErrorResponse("[DB ERROR 0053] Error in getting row", "Error in getting row: "+err.Error(), logPrefix)
		}
	}

	if rowId != utils.NULL_INT { // Stock found
		return helper.CreateErrorResponse("400009", "Stock with same item name, type already present", logPrefix)
	}

	// insert stock
	var id int
	ServiceQuery = database.InsertStockData()
	err = tx.QueryRow(ServiceQuery, ownerRowId, reqBody.Type, reqBody.ItemName, reqBody.Tunch, reqBody.Weight, utils.ACTIVE_YES, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		return helper.Create500ErrorResponse("[DB ERROR 0054] Error in inserting row", "Error in inserting row:"+err.Error(), logPrefix)
	}
	utils.Logger.Info(logPrefix, "Stock added successfully")

	// insert initial data in stock history
	ServiceQuery = database.InsertStockHistoryData()
	_, err = tx.Exec(ServiceQuery, id, utils.NULL_FLOAT, reqBody.Weight, utils.BUY, "Initial Stock", time.Now())
	if err != nil {
		return helper.Create500ErrorResponse("[DB ERROR 0055] Error in inserting row", "Error in inserting row: "+err.Error(), logPrefix)
	} else {
		response, rspCode = helper.CreateSuccessWithIdResponse("Stock added successfully", id, logPrefix)
	}
	utils.Logger.Info(logPrefix, "Stock history added successfully")

	// commit transaction
	if rspCode == utils.StatusOK {
		err = tx.Commit()
		if err != nil {
			return helper.Create500ErrorResponse("[DB ERROR 0056] Error committing transaction", "Error committing transaction:"+err.Error(), logPrefix)
		}
		utils.Logger.Info(logPrefix, "Transaction committed successfully")
	}

	return response, rspCode

}
