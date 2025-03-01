package service

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	helper "SHOP_PORTAL_BACKEND/HELPER"
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"database/sql"
)

func GetStock(ownerRegID string, stockId int, logPrefix string) (interface{}, int) {

	var response interface{}
	var rspCode = utils.StatusOK

	DB := database.DB

	tx, err := DB.Begin()
	if err != nil {
		return helper.Create500ErrorResponse("Error starting transaction", "Error starting transaction:"+err.Error(), logPrefix)
	}
	defer func() {
		if r := recover(); r != nil || rspCode != utils.StatusOK {
			utils.Logger.Error(logPrefix, "Panic occurred during transaction:", r)
			tx.Rollback()
		}
	}()

	// Get owner row id
	ownerRowId, err := helper.GetOwnerId(ownerRegID, tx)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.CreateErrorResponse("404001", "Owner Not Found", logPrefix)
		}
		return helper.Create500ErrorResponse("Error getting owner row ID", "Error getting owner row ID:"+err.Error(), logPrefix)
	}

	ServiceQuery := database.GetStock()
	var itemName string
	var tunch float64
	var weight float64
	var updatedAt string

	err = tx.QueryRow(ServiceQuery, stockId, ownerRowId).Scan(&itemName, &tunch, &weight, &updatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.CreateErrorResponse("404001", "Stock not found", logPrefix)
		}
		return helper.Create500ErrorResponse("Error in getting stock", "Error getting stock:"+err.Error(), logPrefix)
	}

	response = structs.OwnerStockResponse{
		Stat: "OK",
		OwnerStockSubResponse: structs.OwnerStockSubResponse{
			ItemName:  itemName,
			Tunch:     tunch,
			Weight:    weight,
			UpdatedAt: updatedAt,
		},
	}

	if rspCode == utils.StatusOK {
		err = tx.Commit()
		if err != nil {
			return helper.Create500ErrorResponse("Error committing transaction", "Error committing transaction:"+err.Error(), logPrefix)
		}
	}

	return response, rspCode
}
