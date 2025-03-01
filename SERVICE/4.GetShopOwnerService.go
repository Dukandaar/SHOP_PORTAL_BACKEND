package service

import (
	helper "SHOP_PORTAL_BACKEND/HELPER"
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"database/sql"

	database "SHOP_PORTAL_BACKEND/DATABASE"
)

func GetShopOwner(ownerRegID string, logPrefix string) (interface{}, int) {

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
		return helper.Create500ErrorResponse("[DB ERROR 0019] Error starting transaction", "Error starting transaction:"+err.Error(), logPrefix)
	}

	defer tx.Rollback()

	ServiceQuery := database.GetShopOwnerData()
	err = tx.QueryRow(ServiceQuery, ownerRegID).Scan(&rowId, &shopName, &ownerName, &GstIN, &PhoneNo, &regDate, &address, &remarks, &isActive, &gold, &silver, &cash)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.CreateErrorResponse("404001", "Data for owner reg_id "+ownerRegID+" does not exist", logPrefix)
		} else {
			return helper.Create500ErrorResponse("[DB ERROR 0020] Error in getting row", "Error in getting row: "+err.Error(), logPrefix)
		}
	}

	ServiceQuery = database.GetOwnerBillCount()
	err = tx.QueryRow(ServiceQuery, rowId).Scan(&billCount)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.CreateErrorResponse("404001", "Bill count data for owner reg_id "+ownerRegID+" does not exist", logPrefix)
		} else {
			return helper.Create500ErrorResponse("[DB ERROR 00021] Error in getting row", "Error in getting row: "+err.Error(), logPrefix)
		}
	}

	if rspCode == utils.StatusOK {

		err = tx.Commit()
		if err != nil {
			return helper.Create500ErrorResponse("[DB ERROR 0022] Error in committing transaction", "Error committing transaction:"+err.Error(), logPrefix)
		}
		response = structs.GetShopOwnerResponse{
			Response: structs.GetShopOwnerSubResponse{
				Stat: utils.OK,
				Payload: structs.GetShopOwnerPayloadResponse{
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
				},
				Description: "Owner details for reg_id " + ownerRegID,
			},
		}
	}

	return response, rspCode
}
