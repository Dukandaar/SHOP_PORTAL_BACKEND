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
		return helper.Create500ErrorResponse("Error starting transaction", "Error starting transaction:"+err.Error(), logPrefix)
	}

	defer func() {
		if r := recover(); r != nil || err != nil {
			utils.Logger.Error(logPrefix, "Panic occurred during transaction:", r, err)
			tx.Rollback()
		}
	}()

	ServiceQuery := database.GetOwnerRowId()
	var ownerRowId string
	err = tx.QueryRow(ServiceQuery, ownerRegId).Scan(&ownerRowId)
	if err != nil {
		return helper.Create500ErrorResponse("Error in getting row", "Error getting owner row ID:"+err.Error(), logPrefix)
	}

	// check stock with same item_name

	ServiceQuery = database.CheckStockPresent()
	var rowId int
	err = tx.QueryRow(ServiceQuery, ownerRowId, reqBody.Type, reqBody.ItemName).Scan(&rowId)
	if err != nil {
		if err == sql.ErrNoRows { // Stock NOT found
			utils.Logger.Info(logPrefix, "Stock NOT found")
		} else {
			return helper.Create500ErrorResponse("Error in getting row", "Error in getting row:"+err.Error(), logPrefix)
		}
	}

	if rowId > 0 { // Stock found
		utils.Logger.Info(logPrefix, "Stock found")
		return helper.CreateErrorResponse("400009", "Stock with same item name, type already present", logPrefix)
	}

	// insert stock
	var id int
	ServiceQuery = database.InsertStockData()
	err = tx.QueryRow(ServiceQuery, ownerRowId, reqBody.Type, reqBody.ItemName, reqBody.Tunch, reqBody.Weight, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		return helper.Create500ErrorResponse("Error in inserting row", "Error in inserting row:"+err.Error(), logPrefix)
	}

	// insert data in stock history
	ServiceQuery = database.InsertStockHistoryData()
	_, err = tx.Exec(ServiceQuery, id, utils.NULL_FLOAT, reqBody.Weight, utils.BUY, "Initial Stock", time.Now())
	if err != nil {
		return helper.Create500ErrorResponse("Error in inserting row", "Error in inserting row:"+err.Error(), logPrefix)
	} else {
		response, rspCode = helper.CreateSuccessResponseWithId("Stock added successfully", id)
	}

	// commit transaction
	if rspCode == utils.StatusOK {
		err = tx.Commit()
		if err != nil {
			return helper.Create500ErrorResponse("Error committing transaction", "Error committing transaction:"+err.Error(), logPrefix)
		}
	}

	return response, rspCode

}
