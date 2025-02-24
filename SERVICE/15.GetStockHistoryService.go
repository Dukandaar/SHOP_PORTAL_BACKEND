package service

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	helper "SHOP_PORTAL_BACKEND/HELPER"
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"database/sql"
	"fmt"
)

func GetStockHistory(ownerRegID string, stockId int, logPrefix string) (interface{}, int) {

	var response interface{}
	var rspCode = utils.StatusOK

	rsp := make([]structs.StockHistorySubResponse, 0)

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

	// Get Stock histroy
	ServiceQuery := database.GetDetailedStockHistory()
	var prevBalance float64
	var newBalance float64
	var reason string
	var remarks string
	var createdAt string
	var id int64
	var billId int64
	var itemName string
	var weight float64
	var less float64
	var netWeight float64
	var tunch float64
	var fine float64
	var discount float64
	var amount float64

	rows, err := tx.Query(ServiceQuery, stockId)

	if err != nil {
		if err == sql.ErrNoRows {
			utils.Logger.Info(logPrefix, "Data for stockId ", stockId, " and reg_id ", ownerRegID, " does not exist")
			msg := fmt.Sprintf("Data for stockId %d and reg_id %s does not exist", stockId, ownerRegID)
			response, rspCode = helper.CreateErrorResponse("404001", msg)
			return response, rspCode
		}
		return helper.Set500ErrorResponse("Error in getting stock history row", "Error in getting stock history row:"+err.Error(), logPrefix)
	}

	// Get transactoin details
	for rows.Next() {
		err = rows.Scan(&prevBalance, &newBalance, &reason, &remarks, &createdAt, &id, &billId, &itemName, &weight, &less, &netWeight, &tunch, &fine, &discount, &amount)
		if err != nil {
			return helper.Set500ErrorResponse("Error in getting stock row", "Error getting stock: "+err.Error(), logPrefix)
		}

		rsp = append(rsp, structs.StockHistorySubResponse{
			PrevBalance: prevBalance,
			NewBalance:  newBalance,
			Reason:      reason,
			Remarks:     remarks,
			CreatedAt:   createdAt,
			Transaction: structs.TransactionResponse{
				Id:        id,
				BillId:    billId,
				ItemName:  itemName,
				Weight:    weight,
				Less:      less,
				NetWeight: netWeight,
				Tunch:     tunch,
				Fine:      fine,
				Discount:  discount,
				Amount:    amount,
			},
		})
	}

	if rspCode == utils.StatusOK {
		err = tx.Commit()
		if err != nil {
			return helper.Set500ErrorResponse("Error in committing transaction", "Error committing transaction:"+err.Error(), logPrefix)
		}
		response = structs.StockHistoryResponse{
			Stat:                    "OK",
			Count:                   len(rsp),
			StockHistorySubResponse: rsp,
		}
	}

	return response, rspCode
}
