package helper

import (
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
)

func CreateErrorResponse(rsp utils.Codes, description string) structs.ErrorResponse {
	return structs.ErrorResponse{
		Stat: "Fail",
		ErrorSubResponse: structs.ErrorSubResponse{
			ErrorCode:       rsp.StatusCode,
			ErrorMsg:        rsp.Message,
			ErrorDescrition: description,
		},
	}
}

func CreateSuccessResponse(message string) structs.SuccessResponse {
	return structs.SuccessResponse{
		Stat: "OK",
		SuccessSubResponse: structs.SuccessSubResponse{
			SuccessMsg: message,
		},
	}
}
