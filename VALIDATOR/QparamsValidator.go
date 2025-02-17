package validator

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"fmt"
)

func ValidateQParams(reqApiQParams map[string]bool, apiQParams map[string]interface{}) (string, string) {

	DB := database.ConnectDB()
	defer DB.Close()

	// owner_reg_id
	if reqApiQParams[utils.OWNER_REG_ID] {

		if apiQParams[utils.OWNER_REG_ID] == utils.NULL_STRING {
			return "Missing owner_reg_id", "400003"
		}

		regId, _ := apiQParams[utils.OWNER_REG_ID].(string)

		if len(regId) != 10 {
			return "Invalid reg_id length", "400004"
		}

		ServiceQuery := database.CheckValidOwnerRegId()
		var exists bool
		err := DB.QueryRow(ServiceQuery, regId).Scan(&exists)
		if err != nil {
			errMsg := fmt.Sprintf("Error in checking if row with reg_id %s exists", regId)
			utils.Logger.Error(err.Error())
			return errMsg, "500001"
		}

		if exists {
			utils.Logger.Info("Row with reg_id : ", regId, " exists")
		} else {
			utils.Logger.Info("Row with reg_id ", regId, " does not exist")
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
			return "Invalid reg_id length", "400004"
		}

		ServiceQuery := database.CheckValidCustomerRegId()
		var exists bool
		err := DB.QueryRow(ServiceQuery, regId).Scan(&exists)
		if err != nil {
			errMsg := fmt.Sprintf("Error in checking if row with reg_id %s exists", regId)
			utils.Logger.Error(err.Error())
			return errMsg, "500001"
		}

		if exists {
			utils.Logger.Info("Row with reg_id : ", regId, " exists")
		} else {
			utils.Logger.Info("Row with reg_id ", regId, " does not exist")
			return "Customer Registration ID does not exist", "400004"
		}
	}

	return utils.NULL_STRING, utils.SUCCESS
}
