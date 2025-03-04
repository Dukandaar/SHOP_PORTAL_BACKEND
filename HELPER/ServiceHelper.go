package helper

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
)

func GetDateTime(timestampStr string) (string, string) {
	timestamp, err := time.Parse(time.RFC3339Nano, timestampStr) // Use RFC3339Nano
	if err != nil {
		fmt.Println("Error parsing timestamp:", err)
		return "", ""
	}

	// Extract date and time
	date := timestamp.Format("2006-01-02")
	timeStr := timestamp.Format("03:04 PM") // 12-hour format with AM/PM
	return date, timeStr
}

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

func GetBill(billID int, tx *sql.Tx, logPrefix string) (structs.CustomerBillResponse, interface{}, int) {
	var response structs.CustomerBillResponse
	var errRsp interface{}
	errCode := utils.StatusOK

	ServiceQuery := database.GetBill()
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

	err := tx.QueryRow(ServiceQuery, billID).Scan(&id, &billNo, &Type, &metal, &rate, &date, &remarks, &customerDetailsJSON, &transactionDetailsJSON, &paymentDetailsJSON, &createdAt, &updatedAt)
	if err != nil {
		errRsp, errCode = Create500ErrorResponse("[DB ERROR 0117] Error in getting owner bill row", "Error getting owner Bill row ID:"+err.Error(), logPrefix)
		return response, errRsp, errCode
	}

	var customerDetails structs.Customer
	err = json.Unmarshal([]byte(customerDetailsJSON), &customerDetails)
	if err != nil {
		errRsp, errCode = Create500ErrorResponse("[DB ERROR 0118] Error unmarshaling customer details", err.Error(), logPrefix)
		return response, errRsp, errCode
	}

	var transactionDetails []structs.Transaction
	err = json.Unmarshal([]byte(transactionDetailsJSON), &transactionDetails)
	if err != nil {
		errRsp, errCode = Create500ErrorResponse("[DB ERROR 0119] Error unmarshaling transaction details", err.Error(), logPrefix)
		return response, errRsp, errCode
	}

	var paymentDetails structs.Payment
	err = json.Unmarshal([]byte(paymentDetailsJSON), &paymentDetails)
	if err != nil {
		errRsp, errCode = Create500ErrorResponse("[DB ERROR 0120] Error unmarshaling payment details", err.Error(), logPrefix)
		return response, errRsp, errCode
	}

	response = structs.CustomerBillResponse{
		Response: structs.CustomerBillSubResponse{
			Stat: "Success",
			Payload: structs.BillPayloadResponse{
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
			},
			Description: "Customer bill fetched successfully",
		},
	}

	return response, errRsp, errCode
}

func AllBill(ownerRowID int, customerRowID int, tx *sql.Tx, logPrefix string) (structs.AllBillResponse, interface{}, int) {
	var response structs.AllBillResponse
	var errRsp interface{}
	errCode := utils.StatusOK

	allBill := make([]structs.BillPayloadResponse, 0)

	ServiceQuery := database.GetAllBill()
	rows, err := tx.Query(ServiceQuery, ownerRowID, customerRowID)
	if err != nil {
		errRsp, errCode = Create500ErrorResponse("[DB ERROR 0126] Error in getting owner bill row", "Error getting owner Bill row ID:"+err.Error(), logPrefix)
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
		var timeStr string // Corrected variable name
		var remarks string
		var customerDetailsJSON string
		var transactionDetailsJSON string
		var paymentDetailsJSON string
		var createdAt string
		var updatedAt string

		err := rows.Scan(&id, &billNo, &Type, &metal, &rate, &date, &timeStr, &remarks, &customerDetailsJSON, &transactionDetailsJSON, &paymentDetailsJSON, &createdAt, &updatedAt)
		if err != nil {
			errRsp, errCode = Create500ErrorResponse("[DB ERROR 0127] Error in scanning row", "Error getting owner row ID:"+err.Error(), logPrefix)
			return response, errRsp, errCode // Return immediately on error
		}

		var customerDetails structs.Customer
		err = json.Unmarshal([]byte(customerDetailsJSON), &customerDetails)
		if err != nil {
			errRsp, errCode = Create500ErrorResponse("[DB ERROR 0128] Error unmarshaling customer details", err.Error(), logPrefix)
			return response, errRsp, errCode // Return immediately on error
		}

		var transactionDetails []structs.Transaction
		err = json.Unmarshal([]byte(transactionDetailsJSON), &transactionDetails)
		if err != nil {
			errRsp, errCode = Create500ErrorResponse("[DB ERROR 0129] Error unmarshaling transaction details", err.Error(), logPrefix)
			return response, errRsp, errCode // Return immediately on error
		}

		var paymentDetails structs.Payment
		err = json.Unmarshal([]byte(paymentDetailsJSON), &paymentDetails)
		if err != nil {
			errRsp, errCode = Create500ErrorResponse("[DB ERROR 0130] Error unmarshaling payment details", err.Error(), logPrefix)
			return response, errRsp, errCode // Return immediately on error
		}

		date, timeStr = GetDateTime(timeStr) // Corrected variable name
		if date == utils.NULL_STRING || timeStr == utils.NULL_STRING {
			errRsp, errCode = Create500ErrorResponse("[DB ERROR 0131] Error in getting date and time", "Error in getting date and time", logPrefix)
			return response, errRsp, errCode
		}

		oneBill := structs.BillPayloadResponse{
			Id:                 id,
			BillNo:             billNo,
			Type:               Type,
			Metal:              metal,
			Rate:               rate,
			Date:               date,
			Time:               timeStr,
			Remarks:            remarks,
			CustomerDetails:    customerDetails,
			TransactionDetails: transactionDetails,
			PaymentDetails:     paymentDetails,
			CreatedAt:          createdAt,
			UpdatedAt:          updatedAt,
		}

		allBill = append(allBill, oneBill)
	}

	description := utils.NULL_STRING
	if customerRowID == utils.NULL_INT {
		description = "All bills of owner is fetched successfully"
	} else {
		description = "All bills of customer is fetched successfully"
	}

	if errCode == utils.StatusOK {

		response = structs.AllBillResponse{
			Response: structs.AllBillSubResponse{
				Stat:        utils.OK,
				Count:       len(allBill),
				Payload:     allBill,
				Description: description,
			},
		}
	}
	return response, errRsp, errCode
}
