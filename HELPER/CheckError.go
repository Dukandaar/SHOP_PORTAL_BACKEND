package helper

import utils "SHOP_PORTAL_BACKEND/UTILS"

func ErrorResponse(errorCode string, description string) utils.Codes {
	rsp := utils.CodeMap[errorCode]
	rsp.Description = description
	return rsp
}

func CheckError(headerError string, qparamsError string, bodyError string) string {

	errMsg := utils.NULL_STRING

	if headerError != utils.NULL_STRING {
		utils.Logger.Error(headerError)
		errMsg = headerError
	} else if qparamsError != utils.NULL_STRING {
		utils.Logger.Error(qparamsError)
		errMsg = qparamsError
	} else if bodyError != utils.NULL_STRING {
		utils.Logger.Error(bodyError)
		errMsg = bodyError
	}

	return errMsg
}
