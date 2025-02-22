package helper

import (
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"strconv"
)

func Set500ErrorResponse(errMessage string, logMessage string, logPrefix string) (interface{}, int) { // 500 error only
	utils.Logger.Error(logPrefix, logMessage)
	return CreateErrorResponse("500001", errMessage)
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

func CreateSuccessResponseWithRegId(message string, regIg string) (structs.SuccessRegIdResponse, int) {
	return structs.SuccessRegIdResponse{
		Stat: "OK",
		SuccessSubResponse: structs.SuccessRegIdSubResponse{
			SuccessMsg: message,
			RegId:      regIg,
		},
	}, utils.StatusOK
}

func CreateOwnerSuccessResponseWithIdKey(message string, regId string, key string) (structs.CreateOwnerSuccessResponseWithIdKey, int) {
	return structs.CreateOwnerSuccessResponseWithIdKey{
		Stat: "OK",
		SuccessSubResponse: structs.CreateOwnerSuccessSubResponseWithIdKey{
			SuccessMsg: message,
			RegId:      regId,
			Key:        key,
		},
	}, utils.StatusOK
}

func CreateSuccessResponseWithId(message string, id int) (structs.SuccessIdResponse, int) {
	return structs.SuccessIdResponse{
		Stat: "OK",
		SuccessSubResponse: structs.SuccessIdSubResponse{
			SuccessMsg: message,
			Id:         id,
		},
	}, utils.StatusOK
}
