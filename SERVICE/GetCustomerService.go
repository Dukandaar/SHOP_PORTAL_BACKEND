package service

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	helper "SHOP_PORTAL_BACKEND/HELPER"
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"database/sql"
)

func GetCustomer(owner_reg_id string, customer_reg_id string) (interface{}, int) {

	var response interface{}
	rspCode := utils.StatusOK

	var shopName string
	var Name string
	var phNo string
	var regDate string
	var address string
	var remarks string
	var gold float32
	var silver float32
	var cash float32
	var isActive string

	DB := database.ConnectDB()
	defer DB.Close()

	ServiceQuery := database.GetOwnerRowId() // Get Owner's row ID
	var ownerRowId string
	err := DB.QueryRow(ServiceQuery, owner_reg_id).Scan(&ownerRowId)
	if err != nil {
		return helper.SetErrorResponse("Error getting owner row ID", "Error getting owner row ID:"+err.Error())
	}

	ServiceQuery = database.GetCustomerData()
	err = DB.QueryRow(ServiceQuery, customer_reg_id, ownerRowId).Scan(&shopName, &Name, &phNo, &regDate, &address, &remarks, &gold, &silver, &cash, &isActive)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.Logger.Info("Data for reg_id ", customer_reg_id, " does not exists")
			response, rspCode = helper.CreateErrorResponse("404001", "Data for reg_id "+customer_reg_id+" does not exists")
		} else {
			utils.Logger.Error(err.Error())
			response, rspCode = helper.CreateErrorResponse("500001", "Error in getting row")
		}
	}

	if rspCode == utils.StatusOK {

		response = structs.CustomerDetailsResponse{
			Stat: "OK",
			CustomerDetailsSubResponse: structs.CustomerDetailsSubResponse{
				Name:        shopName,
				CompanyName: Name,
				RegId:       customer_reg_id,
				PhNo:        phNo,
				RegDate:     regDate,
				Address:     address,
				Remarks:     remarks,
				Gold:        gold,
				Silver:      silver,
				Cash:        cash,
				IsActive:    isActive,
			},
		}
	}

	return response, rspCode
}
