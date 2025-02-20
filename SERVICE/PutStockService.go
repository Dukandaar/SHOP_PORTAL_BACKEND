package service

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	helper "SHOP_PORTAL_BACKEND/HELPER"
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"time"
)

func PutStock(reqBody structs.PutStock, ownerRegId string, stockId string, logPrefix string) (interface{}, int) {

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
		}
		if rspCode != utils.StatusOK {
			tx.Rollback()
		}
	}()

	ServiceQuery := database.GetOwnerRowId()
	var ownerRowId string
	err = tx.QueryRow(ServiceQuery, ownerRegId).Scan(&ownerRowId)
	if err != nil {
		return helper.Set500ErrorResponse("Error in getting row", "Error getting owner row ID:"+err.Error(), logPrefix)
	}

	// Stock id is validated, it is present, update values

	ServiceQuery = database.UpdateStockData()
	_, err = tx.Exec(ServiceQuery, reqBody.Tunch, reqBody.Weight, time.Now(), stockId, ownerRowId)
	if err != nil {
		utils.Logger.Error(err.Error())
		return helper.Set500ErrorResponse("500001", "Error in updating stock", logPrefix)
	} else {
		response, rspCode = helper.CreateSuccessResponse("Stock updated successfully")
	}

	if rspCode == utils.StatusOK {
		err = tx.Commit()
		if err != nil {
			return helper.Set500ErrorResponse("Error committing transaction", "Error committing transaction:"+err.Error(), logPrefix)
		}
	}

	return response, rspCode
}
