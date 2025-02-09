package utils

const (
	// API NAME
	GENERATE_TOKEN  = "GENERATE_TOKEN"
	POST_SHOP_OWNER = "POST_SHOP_OWNER"
	PUT_SHOP_OWNER  = "PUT_SHOP_OWNER"
	GET_SHOP_OWNER  = "GET_SHOP_OWNER"

	// EMPTY
	NULL_STRING = ""
	NULL_INT    = 0
	OK          = "OK"
	ACTIVE_YES  = "Y"
	ACTIVE_NO   = "N"
	SUCCESS     = "200000"
	GOLD        = "Gold"
	SILVER      = "Silver"
	CASH        = "Cash"

	// HEADER NAMES
	ACCEPT_ENCODING = "Accept-Encoding"
	CONTENT_TYPE    = "Content-Type"
	ACCEPT          = "Accept"
	TOKEN           = "Token"

	// QPARAMS NAMES
	OWNER_REG_ID = "owner_reg_id"

	JwtSecret = "YourStrongJwtSecretKeyHere"

	PrivateKeyPEM = `-----BEGIN RSA PRIVATE KEY-----
MIICWwIBAAKBgGhM3rfXDjV0hTJIrq5bvt+e+EqPVF8S0EHUGSVJpRagyZyBMlNd
JW4mPEryxG4zP19MS3pqLpMaZADNNvS/jW1pHfLOJwWRFwpAXLgGuT9Q/+j32S/B
ftAJJLDSHo6BcyJwaT9pVOmSIGsQMCl/1tiyof/rFgDpvt6OhdJENf67AgMBAAEC
gYAnW2pngNUxwqhIUzjnPmOGSpxyticmQRko3fonIeUT5tRtJCtzXaC0MeVqerU4
yEnPkiChLtQoWjsGOmnUbTvWYdnWCOLiSgV6CMmz7oTCraxWo3JUdd7ZJUOdhjYP
tSAdapANdu/iCrnrtKiBfkZah8cSTvs8dYvsRXW6M9MnAQJBAMWMQVg8fTlUoTBp
IqW/PtQZgrXxZPxU+OUK69ATEV3Zwy6z7e484t0r9SeCBCp4vcvDXqt0463eP9XY
6St6v8ECQQCHKVkm/AkhajZl2nu5ihYo+lAl8g/gb+6ovaFg/vF7IsFtFnrgtsAk
hunShnSZdpwclDMViyyWfj0HCd+JQx17AkEAnk4E11ax6s1c1lSKBVS6XnGLA45M
JMFbKwCTdAyzsAefl79sfauhCSf+rwhLmlVjkvQe2zsycNRXR2EpiUQ2gQJAPJUE
99taQFb6IPcoI8bIHf/scsWn5iJlp86vielb1aSDbGD6HMTtJLIwFgPcOXkXihvH
Ne3Ww3G76u66+ixSBwJAEfpA8nM9jZnyhQDW7BdC3DPutJnW0knV+YgDOOWcRrbL
E8oVALnplwqqVFu3C3ouRVECASS12wugh7yqw6QHlA==
-----END RSA PRIVATE KEY-----`

	// Replace with your actual RSA public key
	PublicKeyPEM = `-----BEGIN PUBLIC KEY-----
MIGeMA0GCSqGSIb3DQEBAQUAA4GMADCBiAKBgGhM3rfXDjV0hTJIrq5bvt+e+EqP
VF8S0EHUGSVJpRagyZyBMlNdJW4mPEryxG4zP19MS3pqLpMaZADNNvS/jW1pHfLO
JwWRFwpAXLgGuT9Q/+j32S/BftAJJLDSHo6BcyJwaT9pVOmSIGsQMCl/1tiyof/r
FgDpvt6OhdJENf67AgMBAAE=
-----END PUBLIC KEY-----`
)

// headers
var GenerateTokenHeaders map[string]bool
var PostShopOwnerHeaders map[string]bool
var PutShopOwnerHeaders map[string]bool
var GetShopOwnerHeaders map[string]bool

// Qparams
var PutShopOwnerQParams map[string]bool
var GetShopOwnerQParams map[string]bool

func SetApiHeaders() {
	PostShopOwnerHeaders = map[string]bool{CONTENT_TYPE: true, ACCEPT: true, ACCEPT_ENCODING: true}
	GenerateTokenHeaders = map[string]bool{CONTENT_TYPE: true, ACCEPT: true, ACCEPT_ENCODING: true}
	PutShopOwnerHeaders = map[string]bool{TOKEN: true, CONTENT_TYPE: true, ACCEPT: true, ACCEPT_ENCODING: true}
	GetShopOwnerHeaders = map[string]bool{TOKEN: true, ACCEPT: true, ACCEPT_ENCODING: true}
}

func SetApiQParams() {
	PutShopOwnerQParams = map[string]bool{OWNER_REG_ID: true}
	GetShopOwnerQParams = map[string]bool{OWNER_REG_ID: true}
}

var ValidHeaders map[string][]interface{}

func SetValidHeaders() {
	ValidHeaders = make(map[string][]interface{})
	ValidHeaders[CONTENT_TYPE] = []interface{}{"application/json", "text/plain", "application.json; charset=utf-8"}
	ValidHeaders[ACCEPT] = []interface{}{"application/json, text/plain", "*/*"}
	ValidHeaders[ACCEPT_ENCODING] = []interface{}{"gzip, deflate, br"}
}
