package maths

import (
	config "SHOP_PORTAL_BACKEND/CONFIG"
	database "SHOP_PORTAL_BACKEND/DATABASE"
	helper "SHOP_PORTAL_BACKEND/HELPER"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"database/sql"
	"math/rand"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1s, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of int63 values to get 64 random letters
)

func RandStringBytesMaskImpr(n int) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func GenerateShopRegID(tx *sql.Tx) string {
	regId := RandStringBytesMaskImpr(10)

	ServiceQuery := database.CheckRegIdPresent()
	var exists bool
	err := tx.QueryRow(ServiceQuery, regId).Scan(&exists)
	if err != nil {
		utils.Logger.Error(err.Error())
		return utils.NULL_STRING
	}

	if exists {
		return GenerateShopRegID(tx)
	} else {
		return regId
	}
}

func GenerateCustomerRegID() string {
	regId := RandStringBytesMaskImpr(12)

	ServiceQuery := database.CheckRegIdPresent()
	var exists bool
	err := database.DB.QueryRow(ServiceQuery, regId).Scan(&exists)
	if err != nil {
		utils.Logger.Error(err.Error())
		return utils.NULL_STRING
	}

	if exists {
		return GenerateCustomerRegID()
	} else {
		return regId
	}
}

func GenerateKey(regId string) (string, string) {
	pubKey, err := helper.ParsePublicKey(config.PUBLIC_KEY) // Parse public key
	if err != nil {
		return utils.NULL_STRING, "Error in parsing public key"
	}

	encryptedID, err := helper.Encrypt(regId, pubKey) // Encrypt ID with public key
	if err != nil {
		return utils.NULL_STRING, "Error in encrypting ID"
	}

	return helper.Base64Encode(encryptedID), utils.NULL_STRING
}
