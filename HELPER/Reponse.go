package helper

import (
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
)

func CreateErrorResponse(code string, des string) (structs.ErrorResponse, int) {
	rsp := utils.CodeMap[code]
	rsp.Description = des

	return structs.ErrorResponse{
		Stat: "Fail",
		ErrorSubResponse: structs.ErrorSubResponse{
			ErrorCode:       rsp.StatusCode,
			ErrorMsg:        rsp.Message,
			ErrorDescrition: rsp.Description,
		},
	}, rsp.StatusCode
}

func CreateSuccessResponse(message string) (structs.SuccessResponse, int) {
	return structs.SuccessResponse{
		Stat: "OK",
		SuccessSubResponse: structs.SuccessSubResponse{
			SuccessMsg: message,
		},
	}, utils.StatusOK
}
