package helper

import (
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"strconv"
)

func SetErrorResponse(errMessage string, logMessage string) (interface{}, int) {
	utils.Logger.Error(logMessage)
	return CreateErrorResponse("500001", errMessage) // Or appropriate error code
}

func CreateErrorResponse(code string, des string) (structs.ErrorResponse, int) {
	rsp := utils.CodeMap[code]
	errCode := rsp.StatusCode
	rsp.StatusCode, _ = strconv.Atoi(code)
	rsp.Description = des

	return structs.ErrorResponse{
		Stat: "Fail",
		ErrorSubResponse: structs.ErrorSubResponse{
			ErrorCode:       rsp.StatusCode,
			ErrorMsg:        rsp.Message,
			ErrorDescrition: rsp.Description,
		},
	}, errCode
}

func CreateSuccessResponse(message string) (structs.SuccessResponse, int) {
	return structs.SuccessResponse{
		Stat: "OK",
		SuccessSubResponse: structs.SuccessSubResponse{
			SuccessMsg: message,
		},
	}, utils.StatusOK
}
