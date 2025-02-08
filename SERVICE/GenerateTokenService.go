package service

import (
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
)

func GenerateToken(reqBody structs.GenerateToken) (interface{}, int) {

	var response interface{}
	rspCode := utils.StatusOK

	return response, rspCode
}
