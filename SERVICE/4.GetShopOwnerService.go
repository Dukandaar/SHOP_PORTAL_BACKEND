package service

import (
	helper "SHOP_PORTAL_BACKEND/HELPER"
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"database/sql"

	database "SHOP_PORTAL_BACKEND/DATABASE"
)

func GetShopOwner(regId string, logPrefix string) (interface{}, int) {

	var response interface{}
	rspCode := utils.StatusOK

	var rowId int
	var shopName string
	var ownerName string
	var GstIN string
	var PhoneNo string
	var regDate string
	var address string
	var remarks string
	var gold float64
	var silver float64
	var cash float64
	var isActive string
	var billCount int

	DB := database.DB

	tx, err := DB.Begin()
	if err != nil {
		return helper.Set500ErrorResponse("Error starting transaction", "Error starting transaction:"+err.Error(), logPrefix)
	}

	defer func() {
		if r := recover(); r != nil || err != nil {
			utils.Logger.Error(logPrefix, "Panic occurred during transaction:", r, err)
			tx.Rollback()
		}
	}()

	ServiceQuery := database.GetShopOwnerData()
	err = tx.QueryRow(ServiceQuery, regId).Scan(&rowId, &shopName, &ownerName, &GstIN, &PhoneNo, &regDate, &address, &remarks, &isActive, &gold, &silver, &cash)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.Logger.Info("Data for reg_id", regId, "does not exist")
			response, rspCode = helper.CreateErrorResponse("404001", "Data for reg_id "+regId+" does not exist", logPrefix)
		} else {
			utils.Logger.Error(err.Error())
			response, rspCode = helper.CreateErrorResponse("500001", "Error in getting row", logPrefix)
		}
	}

	ServiceQuery = database.GetOwnerBillCount()
	err = tx.QueryRow(ServiceQuery, rowId).Scan(&billCount)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.Logger.Info("Bill count data for reg_id", regId, "does not exist")
			response, rspCode = helper.CreateErrorResponse("404001", "Bill count data for reg_id "+regId+" does not exist", logPrefix)
		} else {
			utils.Logger.Error(err.Error())
			response, rspCode = helper.CreateErrorResponse("500001", "Error in getting row", logPrefix)
		}
	}

	if rspCode == utils.StatusOK {

		err = tx.Commit()
		if err != nil {
			utils.Logger.Error(err.Error())
			return helper.CreateErrorResponse("500001", "Error in committing transaction", logPrefix)
		}
		response = structs.ShopOwnerDetailsSubResponse{
			ShopName:  shopName,
			OwnerName: ownerName,
			GstIN:     GstIN,
			PhoneNo:   PhoneNo,
			RegDate:   regDate,
			Address:   address,
			Remarks:   remarks,
			Gold:      gold,
			Silver:    silver,
			Cash:      cash,
			IsActive:  isActive,
			BillCount: billCount,
		}
	}

	return response, rspCode
}
