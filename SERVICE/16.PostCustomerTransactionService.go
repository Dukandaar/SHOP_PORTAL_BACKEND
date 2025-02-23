package service

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	helper "SHOP_PORTAL_BACKEND/HELPER"
	maths "SHOP_PORTAL_BACKEND/MATHS"
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"database/sql"
	"time"
)

func PostCustomerTransactionService(reqBody structs.CustomerBill, ownerRegId string, customerRegId string, logPrefix string) (interface{}, int) {

	var response interface{}
	rspCode := utils.StatusOK

	DB := database.DB

	tx, err := DB.Begin()
	if err != nil {
		return helper.Set500ErrorResponse("Error starting transaction", "Error starting transaction:"+err.Error(), logPrefix)
	}

	defer func() {
		if r := recover(); r != nil || err != nil {
			utils.Logger.Error(logPrefix, "Panic occurred during transaction:", r, err)
			tx.Rollback()
			utils.Logger.Error(logPrefix, "Transaction rolled back")
		}
	}()

	// Get Owner's row ID
	ServiceQuery := database.GetOwnerRowId()
	var ownerRowId int
	err = tx.QueryRow(ServiceQuery, ownerRegId).Scan(&ownerRowId)
	if err != nil {
		return helper.Set500ErrorResponse("Error in getting row", "Error getting owner row ID:"+err.Error(), logPrefix)
	}

	// check customer id
	var regId string
	var isActive string
	var customerRowId int
	if customerRegId == utils.NULL_STRING {

		// Check if customer exists
		ServiceQuery = database.CheckCustomerPresent()
		err = tx.QueryRow(ServiceQuery, ownerRowId, reqBody.CustomerDetails.Name, reqBody.CustomerDetails.ShopName, reqBody.CustomerDetails.PhoneNo).Scan(&customerRowId, &isActive, &regId)
		if err != nil {
			if err == sql.ErrNoRows {
				utils.Logger.Info(logPrefix, "Customer does not exist")
			} else {
				return helper.Set500ErrorResponse("Error in getting row", "Error getting customer row ID:"+err.Error(), logPrefix)
			}
		} else {
			utils.Logger.Info(logPrefix, "Customer exists with id:", customerRowId)
			return helper.CreateErrorResponse("400009", "Customer already exists with reg_id:"+regId)
		}
		// Add Customer Details and return customer id
		ServiceQuery = database.InsertCustomerData()
		regId := maths.GenerateCustomerRegID()
		if regId == utils.NULL_STRING {
			return helper.Set500ErrorResponse("Error generating reg_id", "Error generating reg_id", logPrefix)
		}

		err = tx.QueryRow(ServiceQuery, utils.ACTIVE_YES, ownerRowId, reqBody.CustomerDetails.Name, reqBody.CustomerDetails.ShopName, regId, reqBody.Date, reqBody.CustomerDetails.PhoneNo, reqBody.CustomerDetails.Address, reqBody.Remarks, reqBody.CustomerDetails.GstIN, time.Now(), time.Now()).Scan(&customerRowId)
		if err != nil {
			return helper.Set500ErrorResponse("Error in creating customer", "Error in creating customer:"+err.Error(), logPrefix)
		}
		utils.Logger.Info(logPrefix, "Customer created with id:", customerRowId)

	} else {
		ServiceQuery = database.GetCustomerId()
		err = tx.QueryRow(ServiceQuery, customerRegId, ownerRowId).Scan(&customerRowId)
		if err != nil {
			return helper.Set500ErrorResponse("Error in getting row", "Error getting customer row ID:"+err.Error(), logPrefix)
		}
	}

	// Create bill and return bill id
	ServiceQuery = database.CreateBill()
	var billId int
	err = tx.QueryRow(ServiceQuery, reqBody.BillNo, customerRowId, reqBody.Type, reqBody.Metal, reqBody.Rate, reqBody.Date, time.Now(), time.Now()).Scan(&billId)
	if err != nil {
		return helper.Set500ErrorResponse("Error in creating bill", "Error in creating bill:"+err.Error(), logPrefix)
	} else {
		utils.Logger.Info(logPrefix, "Bill created with id:", billId)
	}

	// Add transactions
	itemCount := len(reqBody.TransactionDetails)

	for i := 0; i < itemCount; i++ {

		// add item
		var transcId int
		ServiceQuery = database.AddTransaction()
		err = tx.QueryRow(ServiceQuery, billId, reqBody.TransactionDetails[i].ItemName, reqBody.TransactionDetails[i].Weight, reqBody.TransactionDetails[i].Less, reqBody.TransactionDetails[i].NetWeight, reqBody.TransactionDetails[i].Tunch, reqBody.TransactionDetails[i].Fine, reqBody.TransactionDetails[i].Discount, reqBody.TransactionDetails[i].Amount).Scan(&transcId)
		if err != nil {
			return helper.Set500ErrorResponse("Error in adding transaction", "Error in adding transaction:"+err.Error(), logPrefix)
		}
		utils.Logger.Info(logPrefix, "Transaction added with id:", transcId)

		// get Item stock id
		var stockId int
		var prev_balance float64
		ServiceQuery = database.GetStockId()
		err = tx.QueryRow(ServiceQuery, reqBody.TransactionDetails[i].ItemName).Scan(&stockId, &prev_balance)
		if err != nil {
			return helper.Set500ErrorResponse("Error in getting stock id", "Error getting stock id:"+err.Error(), logPrefix)
		}
		utils.Logger.Info(logPrefix, "Stock id:", stockId)

		// decrease amount from stock
		ServiceQuery = database.DecreaseStock()
		_, err = tx.Exec(ServiceQuery, (prev_balance - reqBody.TransactionDetails[i].NetWeight), stockId)
		if err != nil {
			return helper.Set500ErrorResponse("Error in decreasing stock", "Error in decreasing stock:"+err.Error(), logPrefix)
		}
		utils.Logger.Info(logPrefix, "Stock decreased by:", reqBody.TransactionDetails[i].NetWeight)

		// update stock histroy table
		ServiceQuery = database.AddStockHistory()
		_, err = tx.Exec(ServiceQuery, stockId, prev_balance, (prev_balance - reqBody.TransactionDetails[i].NetWeight), utils.SELL, transcId, time.Now(), time.Now())
		if err != nil {
			return helper.Set500ErrorResponse("Error in adding stock history", "Error in adding stock history:"+err.Error(), logPrefix)
		}
		utils.Logger.Info(logPrefix, "Stock history added for stock id:", stockId)

	}

	// Add payment
	ServiceQuery = database.BillPayment()
	_, err = tx.Exec(ServiceQuery, billId, customerRowId, reqBody.PaymentDetails.Factor, reqBody.PaymentDetails.New, reqBody.PaymentDetails.Prev, reqBody.PaymentDetails.Total, reqBody.PaymentDetails.Paid, reqBody.PaymentDetails.Rem, reqBody.PaymentDetails.PaymentType, reqBody.Date, reqBody.PaymentDetails.Remarks, time.Now(), time.Now())
	if err != nil {
		return helper.Set500ErrorResponse("Error in adding payment", "Error in adding payment:"+err.Error(), logPrefix)
	}
	utils.Logger.Info(logPrefix, "Payment added for bill id:", billId)

	// update customer balance
	ServiceQuery = database.UpdateCustomerBalance(reqBody.Metal)
	_, err = tx.Exec(ServiceQuery, reqBody.PaymentDetails.Rem, time.Now(), customerRowId)
	if err != nil {
		return helper.Set500ErrorResponse("Error in updating customer balance", "Error in updating customer balance:"+err.Error(), logPrefix)
	}
	utils.Logger.Info(logPrefix, "Customer balance updated for customer id:", customerRowId)

	// update owner bill count
	ServiceQuery = database.UpdateOwnerBillCount()
	_, err = tx.Exec(ServiceQuery, ownerRowId)
	if err != nil {
		return helper.Set500ErrorResponse("Error in updating owner bill count", "Error in updating owner bill count:"+err.Error(), logPrefix)
	}
	utils.Logger.Info(logPrefix, "Owner bill count updated for owner id:", ownerRowId)

	if rspCode == utils.StatusOK {
		err = tx.Commit()
		if err != nil {
			return helper.Set500ErrorResponse("Error in committing transaction", "Error committing transaction:"+err.Error(), logPrefix)
		}
		response, rspCode = helper.CreateSuccessResponse("Bill created successfully")
		utils.Logger.Info(logPrefix, "Transaction committed")
	}

	return response, rspCode
}
