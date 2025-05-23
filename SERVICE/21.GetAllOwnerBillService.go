package service

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	helper "SHOP_PORTAL_BACKEND/HELPER"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"database/sql"
)

func GetAllOwnerBill(ownerRegId string, logPrefix string) (interface{}, int) {

	var response interface{}
	rspCode := utils.StatusOK

	Db := database.DB

	tx, err := Db.Begin()
	if err != nil {
		return helper.Create500ErrorResponse("[DB ERROR 0133] Error starting transaction", "Error starting transaction:"+err.Error(), logPrefix)
	}

	defer tx.Rollback()

	ownerRowId, err := helper.GetOwnerId(ownerRegId, tx)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.CreateErrorResponse("404001", "Owner not found", logPrefix)
		}
		return helper.Create500ErrorResponse("[DB ERROR 0134] Error getting owner row ID", "Error getting owner row ID:"+err.Error(), logPrefix)
	}

	result, response, rspCode := helper.AllBill(ownerRowId, utils.NULL_INT, tx, logPrefix)
	if rspCode != utils.StatusOK {
		return response, rspCode
	}

	err = tx.Commit()
	if err != nil {
		return helper.Create500ErrorResponse("[DB ERROR 0135] Error committing transaction", "Error committing transaction:"+err.Error(), logPrefix)
	}

	return result, rspCode
}
