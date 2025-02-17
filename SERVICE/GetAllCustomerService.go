package service

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	helper "SHOP_PORTAL_BACKEND/HELPER"
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"database/sql"
)

func GetAllCustomer(reqBody structs.AllCustomer, owner_reg_id string) (interface{}, int) {

	var response interface{}
	rspCount := 0
	rspCode := utils.StatusOK

	var name string
	var shopName string
	var regId string
	var phoneNo string
	var regDate string
	var address string
	var remarks string
	var gold float32
	var silver float32
	var cash float32
	var isActive string

	rsp := make([]structs.CustomerDetailsSubResponse, 0)

	DB := database.ConnectDB()
	defer DB.Close()

	ServiceQuery := database.GetAllCustomerData(reqBody.IsActive)
	rows, err := DB.Query(ServiceQuery)
	if err == nil {
		for rows.Next() {

			err = rows.Scan(&name, &shopName, &regId, &phoneNo, &regDate, &address, &remarks, &gold, &silver, &cash, &isActive)
			if err != nil {
				utils.Logger.Error(err.Error())
				response, rspCode = helper.CreateErrorResponse("500001", "Error in getting rows")
				return response, rspCode
			} else {

				rsp = append(rsp, structs.CustomerDetailsSubResponse{
					Name:     name,
					ShopName: shopName,
					RegId:    regId,
					PhoneNo:  phoneNo,
					RegDate:  regDate,
					Address:  address,
					Remarks:  remarks,
					Gold:     gold,
					Silver:   silver,
					Cash:     cash,
					IsActive: isActive,
				})
				rspCount++
			}
		}
	} else {
		if err == sql.ErrNoRows {
			utils.Logger.Info("No rows found")
			response, rspCode = helper.CreateSuccessResponse("No any owner found")
			return response, rspCode
		} else {
			utils.Logger.Error(err.Error())
			response, rspCode = helper.CreateErrorResponse("500001", "Error in getting rows")
			return response, rspCode
		}
	}

	response = structs.AllCustomerDetailsResponse{
		Stat:                       "OK",
		Count:                      rspCount,
		CustomerDetailsSubResponse: rsp,
	}

	return response, rspCode
}
