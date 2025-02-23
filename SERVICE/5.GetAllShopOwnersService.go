package service

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	helper "SHOP_PORTAL_BACKEND/HELPER"
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"database/sql"
)

func GetAllShopOwner(reqBody structs.AllShopOwner, logPrefix string) (interface{}, int) {

	var response interface{}
	rspCount := 0
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

	rsp := make([]structs.ShopOwnerDetailsSubResponse, 0)

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

	ServiceQuery := database.GetAllShopOwnerData(reqBody.IsActive)
	rows, err := tx.Query(ServiceQuery)
	if err == nil {
		for rows.Next() {

			err = rows.Scan(&shopName, &ownerName, &gst_in, &phoneNo, &regDate, &address, &remarks, &isActive, &gold, &silver, &cash, &billCount)
			if err != nil {
				if err == sql.ErrNoRows {
					utils.Logger.Info(logPrefix, "No rows found")
					response, rspCode = helper.CreateSuccessResponse("No any owner found")
					return response, rspCode
				}
				response, rspCode = helper.Set500ErrorResponse("Error in getting rows", "Error in getting rows:"+err.Error(), logPrefix)
				return response, rspCode
			} else {
				rsp = append(rsp, structs.ShopOwnerDetailsSubResponse{
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
				rspCount++
			}
		}
	} else {
		if err == sql.ErrNoRows {
			utils.Logger.Info(logPrefix, "No rows found")
			response, rspCode = helper.CreateSuccessResponse("No any owner found")
			return response, rspCode
		} else {
			response, rspCode = helper.Set500ErrorResponse("Error in getting rows", "Error in getting rows:"+err.Error(), logPrefix)
			return response, rspCode
		}
	}

	if rspCode != utils.StatusOK {
		err = tx.Commit()
		if err != nil {
			response, rspCode = helper.Set500ErrorResponse("Error in commiting transaction", "Error in commiting transaction:"+err.Error(), logPrefix)
			return response, rspCode
		}
	} else {
		response = structs.AllShopOwnerDetailsResponse{
			Stat:                           "OK",
			Count:                          rspCount,
			AllShopOwnerDetailsSubResponse: rsp,
		}
	}

	return response, rspCode
}
