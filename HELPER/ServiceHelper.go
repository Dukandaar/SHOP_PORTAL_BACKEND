package helper

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"database/sql"
	"encoding/json"
)

func GetOwnerId(ownerRegId string, tx *sql.Tx) (int, error) {
	ServiceQuery := database.GetOwnerRowId()
	var ownerRowId int
	err := tx.QueryRow(ServiceQuery, ownerRegId).Scan(&ownerRowId)
	return ownerRowId, err
}

func GetCustomerId(customerRegId string, ownerRowId int, tx *sql.Tx) (int, error) {
	ServiceQuery := database.GetCustomerId()
	var customerId int
	err := tx.QueryRow(ServiceQuery, customerRegId, ownerRowId).Scan(&customerId)
	return customerId, err
}

func CheckIfCustomerBelongsToOwner(customerId int, ownerRowId int, tx *sql.Tx) (bool, error) {
	ServiceQuery := database.CheckIfCustomerBelongsToOwner()
	var isActive string
	err := tx.QueryRow(ServiceQuery, customerId, ownerRowId).Scan(&isActive)
	return isActive == utils.TRUE, err
}

func OwnerAllBill(ownerRowId int, customerRowId int, tx *sql.Tx, logPrefix string) (structs.CustomerAllBillResponse, interface{}, int) {
	var response structs.CustomerAllBillResponse
	var errRsp interface{}
	errCode := utils.StatusOK

	AllBill := make([]structs.CustomerBillSubResponse, 0)

	ServiceQuery := database.GetAllBill()
	rows, err := tx.Query(ServiceQuery, ownerRowId, customerRowId)
	if err != nil {
		errRsp, errCode = Set500ErrorResponse("Error in getting owner bill row", "Error getting owner Bill row ID:"+err.Error(), logPrefix)
		return response, errRsp, errCode
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var billNo int
		var Type string
		var metal string
		var rate float64
		var date string
		var remarks string
		var customerDetailsJSON string
		var transactionDetailsJSON string
		var paymentDetailsJSON string
		var createdAt string
		var updatedAt string

		err := rows.Scan(&id, &billNo, &Type, &metal, &rate, &date, &remarks, &customerDetailsJSON, &transactionDetailsJSON, &paymentDetailsJSON, &createdAt, &updatedAt)
		if err != nil {
			errRsp, errCode = Set500ErrorResponse("Error in scanning row", "Error getting owner row ID:"+err.Error(), logPrefix)
			return response, errRsp, errCode // Return immediately on error
		}

		var customerDetails structs.Customer
		err = json.Unmarshal([]byte(customerDetailsJSON), &customerDetails)
		if err != nil {
			errRsp, errCode = Set500ErrorResponse("Error unmarshaling customer details", err.Error(), logPrefix)
			return response, errRsp, errCode // Return immediately on error
		}

		var transactionDetails []structs.Transaction
		err = json.Unmarshal([]byte(transactionDetailsJSON), &transactionDetails)
		if err != nil {
			errRsp, errCode = Set500ErrorResponse("Error unmarshaling transaction details", err.Error(), logPrefix)
			return response, errRsp, errCode // Return immediately on error
		}

		var paymentDetails structs.Payment
		err = json.Unmarshal([]byte(paymentDetailsJSON), &paymentDetails)
		if err != nil {
			errRsp, errCode = Set500ErrorResponse("Error unmarshaling payment details", err.Error(), logPrefix)
			return response, errRsp, errCode // Return immediately on error
		}

		OneBill := structs.CustomerBillSubResponse{
			Id:                 id,
			BillNo:             billNo,
			Type:               Type,
			Metal:              metal,
			Rate:               rate,
			Date:               date,
			Remarks:            remarks,
			CustomerDetails:    customerDetails,
			TransactionDetails: transactionDetails,
			PaymentDetails:     paymentDetails,
			CreatedAt:          createdAt,
			UpdatedAt:          updatedAt,
		}

		AllBill = append(AllBill, OneBill)
	}

	if errCode == utils.StatusOK {
		response = structs.CustomerAllBillResponse{
			Stat:                    "OK",
			Count:                   len(AllBill),
			CustomerBillSubResponse: AllBill,
		}
	}
	return response, errRsp, errCode
}
