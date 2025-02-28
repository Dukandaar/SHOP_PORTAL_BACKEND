package validator

import (
	config "SHOP_PORTAL_BACKEND/CONFIG"
	helper "SHOP_PORTAL_BACKEND/HELPER"
	utils "SHOP_PORTAL_BACKEND/UTILS"
)

func DecodeKey(encryptedRegID string) (string, string) {
	privKey, err := helper.ParsePrivateKey(config.PRIVATE_KEY) // Parse private key
	if err != nil {
		return utils.NULL_STRING, "Error in parsing private key"
	}

	decodedEncryptedID, err := helper.Base64Decode(encryptedRegID)
	if err != nil {
		return utils.NULL_STRING, "Error in base64 decoding encrypted ID"
	}

	decryptedID, err := helper.Decrypt(decodedEncryptedID, privKey)
	if err != nil {
		return utils.NULL_STRING, "Error in RSA decryption"
	}

	return decryptedID, utils.NULL_STRING
}

func ValidateKey(key string, regId string, logPrefix string) (string, string) {

	decreptedRegId, errMsg := DecodeKey(key)
	utils.Logger.Info(logPrefix, "Decrypted reg_id: ", decreptedRegId)

	if errMsg != utils.NULL_STRING {
		return errMsg, "400006"
	} else if decreptedRegId != regId {
		return "Invalid key for reg_id", "400006"
	}
	utils.Logger.Info(logPrefix, "Valid key for reg_id: ", regId)
	return utils.NULL_STRING, utils.SUCCESS
}

func ValidateToken(token string, regId string, logPrefix string) (string, string) {

	decryptedID, err := helper.ParseAndDecryptJWT(token)
	if err != nil {
		utils.Logger.Error(logPrefix, err.Error())
		return "Invalid token", "400002"
	}

	if decryptedID != regId {
		return "Invalid token for reg_id", "400002"
	}
	utils.Logger.Info(logPrefix, "RegId in token is: ", decryptedID)

	return utils.NULL_STRING, utils.SUCCESS
}
