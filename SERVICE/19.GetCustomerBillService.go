package service

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	helper "SHOP_PORTAL_BACKEND/HELPER"
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"database/sql"
)

func GetBill(ownerRegId string, billId int, logPrefix string) (interface{}, int) {

	var response interface{}
	var rspCode = utils.StatusOK

	DB := database.DB

	tx, err := DB.Begin()
	if err != nil {
		return helper.Set500ErrorResponse("Error starting transaction", "Error starting transaction:"+err.Error(), logPrefix)
	}

	defer func() {
		if r := recover(); r != nil || rspCode != utils.StatusOK {
			utils.Logger.Error(logPrefix, "Panic occurred during transaction:", r)
			tx.Rollback()
			utils.Logger.Error(logPrefix, "Transaction rolled back")
		}
	}()

	// Get customer Id from Bill
	ServiceQuery := database.GetCustomerIdFromBill()
	var customerId int
	err = tx.QueryRow(ServiceQuery, billId).Scan(&customerId)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.CreateErrorResponse("404001", "Customer Not Found")
		}
		return helper.Set500ErrorResponse("Error getting customer row ID", "Error getting customer row ID:"+err.Error(), logPrefix)
	}

	// Get Owner's row ID
	ownerRowId, err := helper.GetOwnerId(ownerRegId, tx)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.CreateErrorResponse("404001", "Owner Not Found")
		}
		return helper.Set500ErrorResponse("Error getting owner row ID", "Error getting owner row ID:"+err.Error(), logPrefix)
	}

	// check if customer is for this owner only
	ServiceQuery = database.CheckCustomerOwnerPresent()
	var isActive string
	err = tx.QueryRow(ServiceQuery, ownerRowId, customerId).Scan(&isActive)
	if err != nil && err != sql.ErrNoRows {
		return helper.Set500ErrorResponse("Error getting customer row ID", "Error getting customer row ID:"+err.Error(), logPrefix)
	}

	if isActive == utils.NULL_STRING {
		return helper.CreateErrorResponse("404001", "Customer not found")
	}

	if isActive == utils.ACTIVE_NO {
		return helper.CreateErrorResponse("404001", "Customer is InActive")
	}

	// Generate Bill
	rsp := structs.CustomerBillSubResponse{}
	customerRsp := structs.Customer{}
	transactionRsp := structs.Transaction{}
	allTransactionsRsp := make([]structs.Transaction, 0)
	payment := structs.Payment{}
	var billNo int
	var Type string
	var metal string
	var rate float64
	var date string
	var created_at string
	var updated_at string
	var shopName string
	var name string
	var GstIN string
	var phoneNo string
	var regDate string
	var address string
	var id int
	var item_name string
	var weight float64
	var less float64
	var net_weight float64
	var tunch float64
	var fine float64
	var discount float64
	var amount float64
	var factor string
	var new float64
	var prev float64
	var total float64
	var paid float64
	var rem float64
	var payment_type string
	var remarks string

	// bill info

	ServiceQuery = database.GetBill()
	err = tx.QueryRow(ServiceQuery, billId).Scan(&billNo, &Type, &metal, &rate, &date, &created_at, &updated_at)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.CreateErrorResponse("404001", "Bill Not Found")
		}
		return helper.Set500ErrorResponse("Error getting bill", "Error getting bill:"+err.Error(), logPrefix)
	}

	rsp.BillNo = billNo
	rsp.Type = Type
	rsp.Metal = metal
	rsp.Rate = rate
	rsp.Date = date
	rsp.CreatedAt = created_at
	rsp.UpdatedAt = updated_at

	//customer info
	ServiceQuery = database.GetCustomerDataById()
	err = tx.QueryRow(ServiceQuery, customerId, ownerRowId).Scan(&shopName, &name, &GstIN, &regDate, &phoneNo, &isActive, &address, &remarks)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.CreateErrorResponse("404001", "Customer Not Found")
		}
		return helper.Set500ErrorResponse("Error getting customer", "Error getting customer:"+err.Error(), logPrefix)
	}

	customerRsp.ShopName = shopName
	customerRsp.Name = name
	customerRsp.GstIN = GstIN
	customerRsp.RegDate = regDate
	customerRsp.PhoneNo = phoneNo
	customerRsp.Address = address
	customerRsp.Remarks = remarks

	ServiceQuery = database.GetBillTransactions()
	row, err := tx.Query(ServiceQuery, billId)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.CreateErrorResponse("404001", "Bill Transaction  Not Found")
		}
		return helper.Set500ErrorResponse("Error getting bill transactions", "Error getting bill transactions:"+err.Error(), logPrefix)
	}
	for row.Next() {
		err = row.Scan(&id, &item_name, &weight, &less, &net_weight, &tunch, &fine, &discount, &amount)
		if err != nil {
			return helper.Set500ErrorResponse("Error getting bill transactions", "Error getting bill transactions:"+err.Error(), logPrefix)
		}

		transactionRsp.Id = id
		transactionRsp.ItemName = item_name
		transactionRsp.Weight = weight
		transactionRsp.Less = less
		transactionRsp.NetWeight = net_weight
		transactionRsp.Tunch = tunch
		transactionRsp.Fine = fine
		transactionRsp.Discount = discount
		transactionRsp.Amount = amount

		allTransactionsRsp = append(allTransactionsRsp, transactionRsp)
	}

	// get bill payement

	ServiceQuery = database.GetBillPayment()
	err = tx.QueryRow(ServiceQuery, billId).Scan(&factor, &new, &prev, &total, &paid, &rem, &payment_type, &remarks)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.CreateErrorResponse("404001", "Bill Payment Not Found")
		}
		return helper.Set500ErrorResponse("Error getting bill payment", "Error getting bill payment:"+err.Error(), logPrefix)
	}

	payment.Factor = factor
	payment.New = new
	payment.Prev = prev
	payment.Total = total
	payment.Paid = paid
	payment.Rem = rem
	payment.PaymentType = payment_type
	payment.Remarks = remarks

	rsp.CustomerDetails = customerRsp
	rsp.TransactionDetails = allTransactionsRsp
	rsp.PaymentDetails = payment

	response = structs.CustomerBillResponse{
		Stat:                    utils.OK,
		CustomerBillSubResponse: rsp,
	}

	if rspCode == utils.StatusOK {
		err = tx.Commit()
		if err != nil {
			return helper.Set500ErrorResponse("Error committing transaction", "Error committing transaction:"+err.Error(), logPrefix)
		}
		utils.Logger.Info(logPrefix, "Transaction committed")
	}

	return response, rspCode
}
