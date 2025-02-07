package utils

const (
	// API NAME
	POST_SHOP_OWNER = "POST_SHOP_OWNER"

	// EMPTY
	NULL_STRING = ""
	NULL_INT    = 0
	OK          = "OK"
	ACTIVE_YES  = "Y"
	ACTIVE_NO   = "N"

	// HEADER NAMES
	ACCEPT_ENCODING = "Accept-Encoding"
	CONTENT_TYPE    = "Content-Type"
	ACCEPT          = "Accept"
)

var PostShopOwnerHeaders map[string]bool

func SetApiHeaders() {
	PostShopOwnerHeaders = map[string]bool{"Content-Type": true, "Accept": true, "Accept-Encoding": true}
}

var ValidHeaders map[string][]interface{}

func SetValidHeaders() {
	ValidHeaders = make(map[string][]interface{})
	ValidHeaders["Content-Type"] = []interface{}{"application/json", "text/plain", "application.json; charset=utf-8"}
	ValidHeaders["Accept"] = []interface{}{"application/json, text/plain", "*/*"}
	ValidHeaders["Accept-Encoding"] = []interface{}{"gzip, deflate, br"}
}
