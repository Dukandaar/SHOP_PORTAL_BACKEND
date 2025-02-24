package validator

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"fmt"
)

func ValidateQParams(reqApiQParams map[string]bool, apiQParams map[string]interface{}, logPrefix string) (string, string) {

	DB := database.DB

	// owner_reg_id
	if reqApiQParams[utils.OWNER_REG_ID] {

		if apiQParams[utils.OWNER_REG_ID] == utils.NULL_STRING {
			return "Missing owner_reg_id", "400003"
		}

		regId, _ := apiQParams[utils.OWNER_REG_ID].(string)

		if len(regId) != 10 {
			return "Invalid owner_reg_id length", "400004"
		}

		ServiceQuery := database.CheckValidOwnerRegId()
		var exists bool
		err := DB.QueryRow(ServiceQuery, regId).Scan(&exists)
		if err != nil {
			errMsg := fmt.Sprintf("Error in checking if row with reg_id %s exists", regId)
			utils.Logger.Error(logPrefix, err.Error())
			return errMsg, "500001"
		}

		if exists {
			utils.Logger.Info(logPrefix, "Row with reg_id : ", regId, " exists")
		} else {
			utils.Logger.Info(logPrefix, "Row with reg_id ", regId, " does not exist")
			return "Owner Registration ID does not exist", "400004"
		}

	}

	// customer_reg_id
	if reqApiQParams[utils.CUSTOMER_REG_ID] {

		if apiQParams[utils.CUSTOMER_REG_ID] == utils.NULL_STRING {
			return "Missing customer_reg_id", "400003"
		}

		regId, _ := apiQParams[utils.CUSTOMER_REG_ID].(string)

		if len(regId) != 12 {
			return "Invalid customer_reg_id length", "400004"
		}

		ServiceQuery := database.CheckValidCustomerRegId()
		var exists bool
		err := DB.QueryRow(ServiceQuery, regId).Scan(&exists)
		if err != nil {
			errMsg := fmt.Sprintf("Error in checking if row with reg_id %s exists", regId)
			utils.Logger.Error(logPrefix, err.Error())
			return errMsg, "500001"
		}

		if exists {
			utils.Logger.Info(logPrefix, "Row with reg_id : ", regId, " exists")
		} else {
			utils.Logger.Info(logPrefix, "Row with reg_id ", regId, " does not exist")
			return "Customer Registration ID does not exist", "404001"
		}
	}

	// stock_id
	if reqApiQParams[utils.STOCK_ID] {

		if apiQParams[utils.STOCK_ID] == utils.NULL_INT {
			return "Missing stock_id", "400003"
		}

		stockId, _ := apiQParams[utils.STOCK_ID].(string)

		ServiceQuery := database.CheckValidStockId()
		var exists bool
		err := DB.QueryRow(ServiceQuery, stockId).Scan(&exists)
		if err != nil {
			errMsg := fmt.Sprintf("Error in checking if row with stock_id %s exists", stockId)
			utils.Logger.Error(logPrefix, err.Error())
			return errMsg, "500001"
		}

		if exists {
			utils.Logger.Info(logPrefix, "Row with stock_id : ", stockId, " exists")
		} else {
			utils.Logger.Info(logPrefix, "Row with stock_id ", stockId, " does not exist")
			return "Stock ID does not exist", "404001"
		}
	}

	// bill_id
	if reqApiQParams[utils.BILL_ID] {

		if apiQParams[utils.BILL_ID] == utils.NULL_INT {
			return "Missing bill_id", "400003"
		}

		billId, _ := apiQParams[utils.BILL_ID].(string)

		ServiceQuery := database.CheckValidBillId()
		var exists bool
		err := DB.QueryRow(ServiceQuery, billId).Scan(&exists)
		if err != nil {
			errMsg := fmt.Sprintf("Error in checking if row with bill_id %s exists", billId)
			utils.Logger.Error(logPrefix, err.Error())
			return errMsg, "500001"
		}

		if exists {
			utils.Logger.Info(logPrefix, "Row with bill_id : ", billId, " exists")
		} else {
			utils.Logger.Info(logPrefix, "Row with bill_id ", billId, " does not exist")
			return "Bill ID does not exist", "404001"
		}
	}

	return utils.NULL_STRING, utils.SUCCESS
}
