package service

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	helper "SHOP_PORTAL_BACKEND/HELPER"
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"database/sql"
)

func GetCustomer(owner_reg_id string, customer_reg_id string, logPrefix string) (interface{}, int) {

	var response interface{}
	rspCode := utils.StatusOK

	var shopName string
	var name string
	var GstIN string
	var phoneNo string
	var regDate string
	var address string
	var remarks string
	var gold float32
	var silver float32
	var cash float32
	var isActive string

	DB := database.DB

	ServiceQuery := database.GetOwnerRowId() // Get Owner's row ID
	var ownerRowId int
	err := DB.QueryRow(ServiceQuery, owner_reg_id).Scan(&ownerRowId)
	if err != nil {
		return helper.Set500ErrorResponse("Error getting owner row ID", "Error getting owner row ID:"+err.Error(), logPrefix)
	}

	ServiceQuery = database.GetCustomerData()
	err = DB.QueryRow(ServiceQuery, customer_reg_id, ownerRowId).Scan(&shopName, &name, &GstIN, &regDate, &phoneNo, &isActive, &address, &remarks, &gold, &silver, &cash)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.Logger.Info(logPrefix, "Data for reg_id ", customer_reg_id, " does not exist")
			response, rspCode = helper.CreateErrorResponse("404001", "Data for reg_id "+customer_reg_id+" does not exist")
		} else {
			utils.Logger.Error(err.Error())
			response, rspCode = helper.CreateErrorResponse("500001", "Error in getting row")
		}
	}

	if rspCode == utils.StatusOK {

		response = structs.CustomerDetailsResponse{
			Stat: "OK",
			CustomerDetailsSubResponse: structs.CustomerDetailsSubResponse{
				Name:     shopName,
				ShopName: name,
				GstIN:    GstIN,
				PhoneNo:  phoneNo,
				RegDate:  regDate,
				Address:  address,
				Remarks:  remarks,
				Gold:     gold,
				Silver:   silver,
				Cash:     cash,
				IsActive: isActive,
			},
		}
	}

	return response, rspCode
}
