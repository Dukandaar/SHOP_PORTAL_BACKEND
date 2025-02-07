package helper

import utils "SHOP_PORTAL_BACKEND/UTILS"

func ErrorResponse(errorCode string, description string) utils.Codes {
	rsp := utils.CodeMap[errorCode]
	rsp.Description = description
	return rsp
}

func CheckError(headerError string, qparamsError string, bodyError string) (utils.Codes, int) {

	rspBody := ErrorResponse("200001", "SUCCESS")
	rspCode := utils.StatusOK

	if headerError != utils.NULL_STRING {
		utils.Logger.Error(headerError)
		rspBody = ErrorResponse("400001", headerError)
		rspCode = utils.StatusBadRequest
	} else if qparamsError != utils.NULL_STRING {
		utils.Logger.Error(qparamsError)
		rspBody = ErrorResponse("400002", qparamsError)
		rspCode = utils.StatusBadRequest
	} else if bodyError != utils.NULL_STRING {
		utils.Logger.Error(bodyError)
		rspBody = ErrorResponse("400003", bodyError)
		rspCode = utils.StatusBadRequest
	}

	return rspBody, rspCode
}
