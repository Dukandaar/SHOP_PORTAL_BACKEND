package helper

import (
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"strconv"
)

// 500 ERROR RESPONSE
func Create500ErrorResponse(errMessage string, logMessage string, logPrefix string) (interface{}, int) { // 500 error only
	utils.Logger.Error(logPrefix, logMessage)
	return CreateErrorResponse("500001", errMessage, logPrefix)
}

// ERROR RESPONSE
func CreateErrorResponse(code string, des string, logPrefix string) (structs.ErrorResponse, int) {

	utils.Logger.Error(logPrefix, des)

	rsp := utils.CodeMap[code]
	errCode := rsp.StatusCode
	rsp.StatusCode, _ = strconv.Atoi(code)
	rsp.Description = des

	return structs.ErrorResponse{
		Response: structs.ErrorSubResponse{
			Stat: "Fail",
			Payload: structs.ErrorPayloadResponse{
				Code:    rsp.StatusCode,
				Message: rsp.Message,
			},
			Description: rsp.Description,
		},
	}, errCode
}

// SUCCESS RESPONSE
func CreateSuccessResponse(message string, description string, logPrefix string) (structs.SuccessResponse, int) {

	utils.Logger.Info(logPrefix + description)

	return structs.SuccessResponse{
		Response: structs.SuccessSubResponse{
			Stat: "OK",
			Payload: structs.SuccessPayloadResponse{
				Message: message,
			},
			Description: description,
		},
	}, utils.StatusOK
}

// POST GENERATE TOKEN RESPONSE
func CreateGenerateTokenResponse(token string, description string, logPrefix string) (structs.GenerateTokenResponse, int) {
	utils.Logger.Info(logPrefix + description)
	return structs.GenerateTokenResponse{
		Response: structs.GenerateTokenSubResponse{
			Stat: utils.OK,
			Payload: structs.GenerateTokenPayloadResponse{
				Token: token,
			},
			Description: description,
		},
	}, utils.StatusOK
}

// POST SHOP OWNER RESPONSE
func CreatePostOwnerResponse(key string, regID string, description string, logPrefix string) (structs.PostShopOwnerResponse, int) {
	utils.Logger.Info(logPrefix + description)
	return structs.PostShopOwnerResponse{
		Response: structs.PostShopOwnerSubResponse{
			Stat: "OK",
			Payload: structs.PostShopOwnerPayloadResponse{
				RegId: regID,
				Key:   key,
			},
			Description: description,
		},
	}, utils.StatusOK
}

// Post Customer Response
func CreatePostCustomerResponse(regID string, description string, logPrefix string) (structs.PostCustomerResponse, int) {
	utils.Logger.Info(logPrefix + description)
	return structs.PostCustomerResponse{
		Response: structs.PostCustomerSubResponse{
			Stat: "OK",
			Payload: structs.PostCustomerPayloadResponse{
				RegId: regID,
			},
			Description: description,
		},
	}, utils.StatusOK
}

// GET CUSTOMER RESPONSE
func CreateGetCustomerResponse(payload structs.GetCustomerPayloadResponse, description string, logPrefix string) (structs.GetCustomerResponse, int) {
	utils.Logger.Info(logPrefix + description)
	return structs.GetCustomerResponse{
		Response: structs.GetCustomerSubResponse{
			Stat:        "OK",
			Payload:     payload,
			Description: description,
		},
	}, utils.StatusOK
}

// GET ALL CUSTOMER RESPONSE
func CreateGetAllCustomerResponse(payload []structs.GetCustomerPayloadResponse, description string, logPrefix string) (structs.GetAllCustomerResponse, int) {
	utils.Logger.Info(logPrefix + description)
	return structs.GetAllCustomerResponse{
		Response: structs.GetAllCustomerSubResponse{
			Stat:        "OK",
			Count:       len(payload),
			Payload:     payload,
			Description: description,
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
