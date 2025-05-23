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
		return helper.Create500ErrorResponse("DB ERROR 0063] Error starting transaction", "Error starting transaction: "+err.Error(), logPrefix)
	}
	defer tx.Rollback()

	// Get owner row id
	ownerRowId, err := helper.GetOwnerId(ownerRegID, tx)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.CreateErrorResponse("404001", "Owner not found", logPrefix)
		}
		return helper.Create500ErrorResponse("[DB ERROR 0064] Error getting owner row ID", "Error getting owner row ID: "+err.Error(), logPrefix)
	}

	ServiceQuery := database.GetStock()
	var mtype string
	var itemName string
	var tunch float64
	var weight float64
	var updatedAt string

	err = tx.QueryRow(ServiceQuery, stockId, ownerRowId).Scan(&mtype, &itemName, &tunch, &weight, &updatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.CreateErrorResponse("404001", "Stock not found for this owner", logPrefix)
		}
		return helper.Create500ErrorResponse("[DB ERROR 0065] Error in getting stock", "Error getting stock: "+err.Error(), logPrefix)
	}

	if rspCode == utils.StatusOK {
		err = tx.Commit()
		if err != nil {
			return helper.Create500ErrorResponse("[DB ERROR 0066] Error committing transaction", "Error committing transaction: "+err.Error(), logPrefix)
		}
		utils.Logger.Info(logPrefix, "Transaction committed successfully")

		response = structs.OwnerStockResponse{
			Response: structs.OwnerStockSubResponse{
				Stat: utils.OK,
				Payload: structs.OwnerStockPayloadResponse{
					Id:        stockId,
					Type:      mtype,
					ItemName:  itemName,
					Tunch:     tunch,
					Weight:    weight,
					UpdatedAt: updatedAt,
				}, Description: "Successfully fetched stock",
			},
		}
	}

	return response, rspCode
}
