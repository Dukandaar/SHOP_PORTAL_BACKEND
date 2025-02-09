package maths

import (
	helper "SHOP_PORTAL_BACKEND/HELPER"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"math/rand"
	"os"
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

func GenerateRegID() string {
	return RandStringBytesMaskImpr(10)
}

func GenerateKey(regId string) (string, string) {
	pubKey, err := helper.ParsePublicKey(os.Getenv("PublicKeyPEM")) // Parse public key
	if err != nil {
		return utils.NULL_STRING, "Error in parsing public key"
	}

	encryptedID, err := helper.Encrypt(regId, pubKey) // Encrypt ID with public key
	if err != nil {
		return utils.NULL_STRING, "Error in encrypting ID"
	}

	return helper.Base64Encode(encryptedID), utils.NULL_STRING
}
