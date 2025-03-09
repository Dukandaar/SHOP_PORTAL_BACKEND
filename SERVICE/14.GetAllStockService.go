package service

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	helper "SHOP_PORTAL_BACKEND/HELPER"
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"database/sql"
)

func GetAllStock(metalType string, ownerRegID string, logPrefix string) (interface{}, int) {

	var response interface{}
	var rspCode = utils.StatusOK
	var stockPayloads = make([]structs.OwnerStockPayloadResponse, 0)

	DB := database.DB

	tx, err := DB.Begin()
	if err != nil {
		return helper.Create500ErrorResponse("[DB ERROR 0066] Error starting transaction", "Error starting transaction: "+err.Error(), logPrefix)
	}
	defer tx.Rollback()

	// Get owner row id
	ownerRowId, err := helper.GetOwnerId(ownerRegID, tx)
	if err != nil {
		return helper.Create500ErrorResponse("[DB ERROR 0067] Error getting owner row ID", "Error getting owner row ID: "+err.Error(), logPrefix)
	}

	ServiceQuery := utils.NULL_STRING
	var rows *sql.Rows

	if metalType != utils.ALL {
		ServiceQuery = database.GetAllStock()
		rows, err = tx.Query(ServiceQuery, ownerRowId, metalType)
	} else {
		ServiceQuery = database.GetAllStock1()
		rows, err = tx.Query(ServiceQuery, ownerRowId)
	}

	var id int
	var itemName string
	var tunch sql.NullFloat64
	var weight sql.NullFloat64
	var updatedAt string

	if err != nil {
		if err == sql.ErrNoRows {
			return helper.CreateErrorResponse("404002", "Stock Not Found", logPrefix)
		}
		return helper.Create500ErrorResponse("[DB ERROR 0068] Error getting stock", "Error getting stock: "+err.Error(), logPrefix)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id, &itemName, &tunch, &weight, &updatedAt)
		if err != nil {
			return helper.Create500ErrorResponse("[DB ERROR 0069] Error scanning row", "Error scanning row: "+err.Error(), logPrefix)
		}

		stockPayloads = append(stockPayloads, structs.OwnerStockPayloadResponse{
			Id:        id,
			ItemName:  itemName,
			Tunch:     tunch.Float64,
			Weight:    weight.Float64,
			UpdatedAt: updatedAt,
		})
	}

	if rspCode == utils.StatusOK {
		err = tx.Commit()
		if err != nil {
			return helper.Create500ErrorResponse("[DB ERROR 0070] Error committing transaction", "Error committing transaction: "+err.Error(), logPrefix)
		}
		utils.Logger.Info(logPrefix, "Transaction committed successfully")

		response = structs.OwnerAllStockResponse{
			Response: structs.OwnerAllStockSubResponse{
				Stat:        utils.OK,
				Count:       len(stockPayloads),
				Payload:     stockPayloads,
				Description: "All Stock retrieved successfully",
			},
		}
	}
	return response, rspCode
}
