package service

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	helper "SHOP_PORTAL_BACKEND/HELPER"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"database/sql"
)

func GetPreviousBillNo(ownerRegID string, logPrefix string) (interface{}, int) {

	var response interface{}
	rspCode := utils.StatusOK

	Db := database.DB

	tx, err := Db.Begin()
	if err != nil {
		return helper.Create500ErrorResponse("[DB ERROR 0131] Error starting transaction", "Error starting transaction:"+err.Error(), logPrefix)
	}

	defer tx.Rollback()

	ownerRowId, err := helper.GetOwnerId(ownerRegID, tx)
	if err != nil {
		return helper.Create500ErrorResponse("[DB ERROR 0132] Error getting owner row ID", "Error getting owner row ID:"+err.Error(), logPrefix)
	}

	var bill_cnt int
	ServiceQuery := database.GetPreviousBillNo()
	err = tx.QueryRow(ServiceQuery, ownerRowId).Scan(&bill_cnt)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.CreateErrorResponse("404001", "Bill Not Found", logPrefix)
		}
		return helper.Create500ErrorResponse("[DB ERROR 0133] Error getting customer row ID", "Error getting customer row ID:"+err.Error(), logPrefix)
	}

	response, rspCode = helper.CreateSuccessWithIdResponse("Bill No fetched successfully", bill_cnt, logPrefix)

	if rspCode == utils.StatusOK {
		err = tx.Commit()
		if err != nil {
			return helper.Create500ErrorResponse("[DB ERROR 0134] Error committing transaction", "Error committing transaction:"+err.Error(), logPrefix)
		}
	}

	return response, rspCode
}
