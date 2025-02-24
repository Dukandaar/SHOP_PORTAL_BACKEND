package service

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	helper "SHOP_PORTAL_BACKEND/HELPER"
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"database/sql"
)

func GetAllBill(ownerRegId string, customerRegId string, logPrefix string) (interface{}, int) {
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

	// Get Owner's row ID
	ownerRowId, err := helper.GetOwnerId(ownerRegId, tx)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.CreateErrorResponse("404001", "Owner Not Found")
		}
		return helper.Set500ErrorResponse("Error getting owner row ID", "Error getting owner row ID:"+err.Error(), logPrefix)
	}

	// Get customer Id
	customerId, err := helper.GetCustomerId(customerRegId, ownerRowId, tx)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.CreateErrorResponse("404001", "Customer Not Found")
		}
		return helper.Set500ErrorResponse("Error getting customer row ID", "Error getting customer row ID:"+err.Error(), logPrefix)
	}

	// Get Bills
	var billId int
	ServiceQuery := database.GetAllCustomerBillId()
	rows, err := tx.Query(ServiceQuery, customerId)
	if err != nil {
		return helper.Set500ErrorResponse("Error getting bills", "Error getting bills:"+err.Error(), logPrefix)
	}

	// bill Id array
	billIdArray := make([]int, 0)

	for rows.Next() {
		err := rows.Scan(&billId)
		if err != nil {
			return helper.Set500ErrorResponse("Error scanning row", "Error scanning row:"+err.Error(), logPrefix)
		}
		billIdArray = append(billIdArray, billId)
	}

	rsp1 := make([]structs.CustomerBillSubResponse, 0)

	for _, billId := range billIdArray {
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
		var isActive string
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

		ServiceQuery := database.GetBill()
		err := tx.QueryRow(ServiceQuery, billId).Scan(&billNo, &Type, &metal, &rate, &date, &created_at, &updated_at)
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

		rsp1 = append(rsp1, rsp)
	}

	response = structs.CustomerAllBillResponse{
		Stat:                    utils.OK,
		Count:                   len(billIdArray),
		CustomerBillSubResponse: rsp1,
	}

	return response, rspCode

}
