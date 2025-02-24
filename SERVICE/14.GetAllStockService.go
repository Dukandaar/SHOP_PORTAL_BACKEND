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
	var rsp = make([]structs.OwnerStockSubResponse, 0)

	DB := database.DB

	tx, err := DB.Begin()
	if err != nil {
		return helper.Set500ErrorResponse("Error starting transaction", "Error starting transaction:"+err.Error(), logPrefix)
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
			return helper.CreateErrorResponse("404001", "Owner Not Found")
		}
		return helper.Set500ErrorResponse("Error getting owner row ID", "Error getting owner row ID:"+err.Error(), logPrefix)
	}

	ServiceQuery := database.GetAllStock()
	var itemName string
	var tunch float64
	var weight float64
	var updatedAt string

	rows, err := tx.Query(ServiceQuery, ownerRowId)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.CreateErrorResponse("404002", "Stock Not Found")
		}
		utils.Logger.Error(err.Error())
		return helper.Set500ErrorResponse("Error getting stock", "Error getting stock:"+err.Error(), logPrefix)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&itemName, &tunch, &weight, &updatedAt)
		if err != nil {
			utils.Logger.Error(err.Error())
			return helper.Set500ErrorResponse("Error scanning row", "Error scanning row:"+err.Error(), logPrefix)
		}

		rsp = append(rsp, structs.OwnerStockSubResponse{
			ItemName:  itemName,
			Tunch:     tunch,
			Weight:    weight,
			UpdatedAt: updatedAt,
		})
	}

	if rspCode == utils.StatusOK {
		err = tx.Commit()
		if err != nil {
			return helper.Set500ErrorResponse("Error committing transaction", "Error committing transaction:"+err.Error(), logPrefix)
		}

		response = structs.OwnerAllStockResponse{
			Stat:                  "OK",
			Count:                 len(rsp),
			OwnerStockSubResponse: rsp,
		}
	}

	return response, rspCode
}
