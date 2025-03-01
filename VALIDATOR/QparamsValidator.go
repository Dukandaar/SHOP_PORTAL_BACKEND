package validator

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	helper "SHOP_PORTAL_BACKEND/HELPER"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"fmt"
)

func ValidateQParams(reqApiQParams map[string]bool, apiQParams map[string]interface{}, logPrefix string) (interface{}, int) {

	DB := database.DB

	// owner_reg_id
	if reqApiQParams[utils.OWNER_REG_ID] {

		if apiQParams[utils.OWNER_REG_ID] == utils.NULL_STRING {
			return helper.CreateErrorResponse("400003", "Missing owner_reg_id", logPrefix)
		}

		regId, _ := apiQParams[utils.OWNER_REG_ID].(string)

		if len(regId) != 10 {
			return helper.CreateErrorResponse("400004", "Invalid owner_reg_id length", logPrefix)
		}

		ServiceQuery := database.CheckValidOwnerRegId()
		var exists bool
		err := DB.QueryRow(ServiceQuery, regId).Scan(&exists)
		if err != nil {
			errMsg := fmt.Sprintf("[DB ERROR 0001] Error in checking if owner with reg_id %s exists", regId)
			return helper.Create500ErrorResponse(errMsg, err.Error(), logPrefix)
		}

		if exists {
			utils.Logger.Info(logPrefix, "Owner with reg_id : ", regId, " exists")
		} else {
			return helper.CreateErrorResponse("404001", "Owner for reg_id "+regId+" does not exist", logPrefix)
		}

	}

	// customer_reg_id
	if reqApiQParams[utils.CUSTOMER_REG_ID] {

		if apiQParams[utils.CUSTOMER_REG_ID] == utils.NULL_STRING {
			return helper.CreateErrorResponse("400003", "Missing customer_reg_id", logPrefix)
		}

		regId, _ := apiQParams[utils.CUSTOMER_REG_ID].(string)

		if len(regId) != 12 {
			return helper.CreateErrorResponse("400004", "Invalid customer_reg_id length", logPrefix)
		}

		ServiceQuery := database.CheckValidCustomerRegId()
		var exists bool
		err := DB.QueryRow(ServiceQuery, regId).Scan(&exists)
		if err != nil {
			errMsg := fmt.Sprintf("[DB ERROR 0002] Error in checking if customer with reg_id %s exists", regId)
			return helper.Create500ErrorResponse(errMsg, err.Error(), logPrefix)
		}

		if exists {
			utils.Logger.Info(logPrefix, "Customer with reg_id : ", regId, " exists")
		} else {
			return helper.CreateErrorResponse("404001", "Customer for reg_id "+regId+" does not exist", logPrefix)
		}
	}

	// stock_id
	if reqApiQParams[utils.STOCK_ID] {

		if apiQParams[utils.STOCK_ID] == utils.NULL_INT {
			return helper.CreateErrorResponse("400003", "Missing stock_id", logPrefix)
		}

		stockId, _ := apiQParams[utils.STOCK_ID].(string)

		ServiceQuery := database.CheckValidStockId()
		var exists bool
		err := DB.QueryRow(ServiceQuery, stockId).Scan(&exists)
		if err != nil {
			errMsg := fmt.Sprintf("[DB ERROR 0003] Error in checking if row with stock_id %s exists", stockId)
			return helper.Create500ErrorResponse(errMsg, err.Error(), logPrefix)
		}

		if exists {
			utils.Logger.Info(logPrefix, "Row with stock_id : ", stockId, " exists")
		} else {
			return helper.CreateErrorResponse("404001", "Stock ID does not exist", logPrefix)
		}
	}

	// bill_id
	if reqApiQParams[utils.BILL_ID] {

		if apiQParams[utils.BILL_ID] == utils.NULL_INT {
			return helper.CreateErrorResponse("400003", "Missing bill_id", logPrefix)
		}

		billId, _ := apiQParams[utils.BILL_ID].(string)

		ServiceQuery := database.CheckValidBillId()
		var exists bool
		err := DB.QueryRow(ServiceQuery, billId).Scan(&exists)
		if err != nil {
			errMsg := fmt.Sprintf("[DB ERROR 0004] Error in checking if row with bill_id %s exists", billId)
			return helper.Create500ErrorResponse(errMsg, err.Error(), logPrefix)
		}

		if exists {
			utils.Logger.Info(logPrefix, "Row with bill_id : ", billId, " exists")
		} else {
			return helper.CreateErrorResponse("404001", "Bill ID does not exist", logPrefix)
		}
	}

	return utils.NULL_STRING, utils.StatusOK
}
