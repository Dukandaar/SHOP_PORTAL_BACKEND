package service

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	helper "SHOP_PORTAL_BACKEND/HELPER"
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"database/sql"
	"time"
)

func PutCustomer(reqBody structs.Customer, OwnerRegId string, CustomerRegId string, logPrefix string) (interface{}, int) {
	var response interface{}
	rspCode := utils.StatusOK

	DB := database.ConnectDB()
	defer DB.Close()

	tx, err := DB.Begin() // Start transaction
	if err != nil {
		return helper.Set500ErrorResponse("Error starting transaction", "Error starting transaction:"+err.Error(), logPrefix)
	}

	defer func() {
		if r := recover(); r != nil || err != nil {
			utils.Logger.Error("Panic occurred during transaction:", r, err)
			tx.Rollback()
		}
		if rspCode != utils.StatusOK {
			tx.Rollback()
		}
	}()

	ServiceQuery := database.GetOwnerRowId() // Get Owner's row ID
	var ownerRowId string
	err = tx.QueryRow(ServiceQuery, OwnerRegId).Scan(&ownerRowId)
	if err != nil {
		return helper.Set500ErrorResponse("Error getting owner row ID", "Error getting owner row ID:"+err.Error(), logPrefix)
	}

	// Check if customer exists with same name, shopname and phone no
	ServiceQuery = database.CheckCustomerPresent()
	var rowId int
	var isActive string
	var customer_reg_id string

	err = tx.QueryRow(ServiceQuery, reqBody.Name, reqBody.ShopName, reqBody.PhoneNo).Scan(&rowId, &isActive, &customer_reg_id) // Use tx
	if err != nil {
		if err != sql.ErrNoRows {
			return helper.Set500ErrorResponse("Error checking customer", "Error checking customer:"+err.Error(), logPrefix)
		}
	}

	if isActive != utils.NULL_STRING {
		// Customer exists.

		if customer_reg_id == CustomerRegId {
			utils.Logger.Info("Row with reg_id ", CustomerRegId, " exists") // update row
		} else {
			utils.Logger.Info("Same data with reg_id ", customer_reg_id, " exists")
			response, rspCode = helper.CreateErrorResponse("400009", "Same data with reg_id "+customer_reg_id+" exists")
			return response, rspCode
		}
	}

	// Update customer details in DB (only if customer exists)
	ServiceQuery = database.UpdateCustomerData()

	_, err = tx.Exec(ServiceQuery, reqBody.Name, reqBody.ShopName, reqBody.PhoneNo, reqBody.RegDate, reqBody.Address, time.Now(), rowId) // Use rowId and tx
	if err != nil {
		return helper.Set500ErrorResponse("Error updating customer data", "Error updating customer data:"+err.Error(), logPrefix)
	} else {
		// update active and remark in owner_customer table
		ServiceQuery = database.UpdateOwnerCustomerData()

		_, err = tx.Exec(ServiceQuery, utils.ACTIVE_YES, reqBody.Remarks, ownerRowId, rowId)
		if err != nil {
			return helper.Set500ErrorResponse("Error updating owner_customer data", "Error updating owner_customer data:"+err.Error(), logPrefix)
		}
		response, rspCode = helper.CreateSuccessResponse("Updated customer details")
	}

	if rspCode == utils.StatusOK {
		err = tx.Commit()
		if err != nil {
			return helper.Set500ErrorResponse("Error committing transaction", "Error committing transaction:"+err.Error(), logPrefix)
		}
	}

	return response, rspCode
}
