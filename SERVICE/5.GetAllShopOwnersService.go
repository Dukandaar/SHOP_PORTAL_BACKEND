package service

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	helper "SHOP_PORTAL_BACKEND/HELPER"
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"database/sql"
)

func GetAllShopOwner(IsActive string, logPrefix string) (interface{}, int) {

	var response interface{}
	rspCode := utils.StatusOK

	var shopName string
	var ownerName string
	var gst_in string
	var phoneNo string
	var regDate string
	var address string
	var remarks string
	var gold float64
	var silver float64
	var cash float64
	var isActive string
	var billCount int

	owners := make([]structs.GetShopOwnerPayloadResponse, 0)

	DB := database.DB

	tx, err := DB.Begin()
	if err != nil {
		return helper.Create500ErrorResponse("[DB ERROR 0021] Error starting transaction", "Error starting transaction:"+err.Error(), logPrefix)
	}

	defer tx.Rollback()

	ServiceQuery := database.GetAllShopOwnerData(IsActive)
	rows, err := tx.Query(ServiceQuery)
	if err == nil {
		for rows.Next() {

			err = rows.Scan(&shopName, &ownerName, &gst_in, &phoneNo, &regDate, &address, &remarks, &isActive, &gold, &silver, &cash, &billCount)
			if err != nil {
				return helper.Create500ErrorResponse("[DB ERROR 0022] Error in getting rows", "Error in getting rows:"+err.Error(), logPrefix)
			} else {
				owners = append(owners, structs.GetShopOwnerPayloadResponse{
					ShopName:  shopName,
					OwnerName: ownerName,
					GstIN:     gst_in,
					PhoneNo:   phoneNo,
					RegDate:   regDate,
					Address:   address,
					Remarks:   remarks,
					Gold:      gold,
					Silver:    silver,
					Cash:      cash,
					IsActive:  isActive,
					BillCount: billCount,
				})
			}
		}
	} else {
		if err == sql.ErrNoRows {
			return helper.CreateSuccessResponse("No any owner found", "No any owner found", logPrefix)
		} else {
			return helper.Create500ErrorResponse("[DB ERROR 0023] Error in getting rows", "Error in getting rows:"+err.Error(), logPrefix)
		}
	}

	if rspCode == utils.StatusOK {
		err = tx.Commit()
		if err != nil {
			return helper.Create500ErrorResponse("[DB ERROR 0024] Error in commiting transaction", "Error in commiting transaction:"+err.Error(), logPrefix)
		}
		utils.Logger.Info(logPrefix, "Transaction committed successfully")

		response = structs.GetAllShopOwnerResponse{
			Response: structs.GetAllShopOwnerSubResponse{
				Stat:        utils.OK,
				Count:       len(owners),
				Payload:     owners,
				Description: "All registered owner details.",
			},
		}
	}

	return response, rspCode
}
