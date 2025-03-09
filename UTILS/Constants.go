package utils

import "math"

const (
	// API NAME
	SERVER_UP                         = "SERVER_UP"
	GENERATE_TOKEN                    = "GENERATE_TOKEN"
	POST_SHOP_OWNER                   = "POST_SHOP_OWNER"
	PUT_SHOP_OWNER                    = "PUT_SHOP_OWNER"
	GET_SHOP_OWNER                    = "GET_SHOP_OWNER"
	GET_ALL_SHOP_OWNER                = "GET_ALL_SHOP_OWNER"
	POST_CUSTOMER                     = "POST_CUSTOMER"
	PUT_CUSTOMER                      = "PUT_CUSTOMER"
	GET_CUSTOMER                      = "GET_CUSTOMER"
	GET_ALL_CUSTOMER                  = "GET_ALL_CUSTOMER"
	GET_FILTERED_CUSTOMER             = "GET_FILTERED_CUSTOMER"
	POST_STOCK                        = "POST_STOCK"
	PUT_STOCK                         = "PUT_STOCK"
	GET_STOCK                         = "GET_STOCK"
	GET_ALL_STOCK                     = "GET_ALL_STOCK"
	GET_STOCK_HISTORY                 = "GET_STOCK_HISTORY"
	GET_PREVIOUS_BALANCE              = "GET_PREVIOUS_BALANCE"
	POST_CUSTOMER_BILL                = "POST_CUSTOMER_BILL"
	PUT_CUSTOMER_TRANSACTION          = "PUT_CUSTOMER_TRANSACTION"
	GET_CUSTOMER_BILL                 = "GET_CUSTOMER_BILL"
	GET_ALL_CUSTOMER_BILL             = "GET_ALL_CUSTOMER_BILL"
	GET_FILTERED_CUSTOMER_TRANSACTION = "GET_FILTERED_CUSTOMER_TRANS"
	GET_ALL_OWNER_BILL                = "GET_ALL_OWNER_BILL"
	GET_PREVIOUS_BILL_NO              = "GET_PREVIOUS_BILL_NO"
	DELETE_STOCK                      = "DELETE_STOCK"

	// DEFAULTS
	NULL_STRING = ""
	NULL_INT    = 0
	NULL_FLOAT  = 0.0
	MAX_INT     = math.MaxInt8
	MAX_FLOAT   = math.MaxFloat64
	OK          = "OK"
	ACTIVE_YES  = "Y"
	ACTIVE_NO   = "N"
	ALL         = "All"
	SUCCESS     = "200000"
	GOLD        = "Gold"
	SILVER      = "Silver"
	BOTH        = "Both"
	CASH        = "Cash"
	UPI         = "UPI"
	NEFT        = "NEFT"
	RTGS        = "RTGS"
	CHEQUE      = "Cheque"
	CARD        = "Card"
	OTHER       = "Other"
	TODAY       = "Today"
	WEEK        = "Week"
	MONTH       = "Month"
	YEAR        = "Year"
	CUSTOM      = "Custom"
	BUY         = "Buy"
	SELL        = "Sell"
	TRUE        = "True"
	FALSE       = "False"
	WHOLESALE   = "WholeSale"
	RETAIL      = "Retail"
	FINE        = "Fine"
	AMOUNT      = "Amount"

	// MAX FIELD LENGTH
	SHOP_NAME_MAX_LEN       = 255
	OWNER_NAME_MAX_LEN      = 255
	SHOP_REG_ID_MAX_LEN     = 10
	PHONE_NO_MAX_LEN        = 10
	GST_IN_MAX_LEN          = 15
	CUSTOMER_NAME_MAX_LEN   = 255
	shop_name_MAX_LEN       = 255
	OWNER_REG_ID_MAX_LEN    = 10
	CUSTOMER_REG_ID_MAX_LEN = 15
	ADDRESS_MAX_LEN         = 255
	ITEM_NAME_MAX_LEN       = 255

	// HEADER NAMES
	ACCEPT_ENCODING = "Accept-Encoding"
	CONTENT_TYPE    = "Content-Type"
	CACHE_CONTROL   = "Cache-Control"
	ACCEPT          = "Accept"
	TOKEN           = "Token"
	SKIP_TOKEN      = "Skip-Token"

	// QPARAMS NAMES
	OWNER_REG_ID    = "owner_reg_id"
	CUSTOMER_REG_ID = "customer_reg_id"
	STOCK_ID        = "stock_id"
	BILL_ID         = "bill_id"
)

// headers
var GenerateTokenHeaders map[string]bool
var PostShopOwnerHeaders map[string]bool
var PutShopOwnerHeaders map[string]bool
var GetShopOwnerHeaders map[string]bool
var GetAllShopOwnerHeaders map[string]bool
var PostCustomerHeaders map[string]bool
var PutCustomerHeaders map[string]bool
var GetCustomerHeaders map[string]bool
var GetAllCustomerHeaders map[string]bool
var GetFilteredCustomerHeaders map[string]bool
var PostStockHeaders map[string]bool
var PutStockHeaders map[string]bool
var GetStockHeaders map[string]bool
var GetAllStockHeaders map[string]bool
var GetStockHistoryHeaders map[string]bool
var GetPreviousBalanceHeaders map[string]bool
var PostCustomerBillHeaders map[string]bool
var PutCustomerTransactionHeaders map[string]bool
var GetCustomerBillHeaders map[string]bool
var GetAllCustomerBillHeaders map[string]bool
var GetFilteredCustomerTransactionHeaders map[string]bool
var GetAllOwnerBillHeaders map[string]bool
var GetPreviousBillNoHeaders map[string]bool
var DeleteStockHeaders map[string]bool

// Qparams
var PutShopOwnerQParams map[string]bool
var GetShopOwnerQParams map[string]bool
var PostCustomerQParams map[string]bool
var PutCustomerQParams map[string]bool
var GetCustomerQParams map[string]bool
var GetAllCustomerQParams map[string]bool
var GetFilteredCustomerQParams map[string]bool
var PostStockQParams map[string]bool
var PutStockQParams map[string]bool
var GetStockQParams map[string]bool
var GetAllStockQParams map[string]bool
var GetStockHistoryQParams map[string]bool
var GetPreviousBalanceQParams map[string]bool
var PostCustomerBillQParams map[string]bool
var PutCustomerTransactionQParams map[string]bool
var GetCustomerBillQParams map[string]bool
var GetAllCustomerBillQParams map[string]bool
var GetFilteredCustomerTransactionQParams map[string]bool
var GetAllOwnerBillQParams map[string]bool
var GetPreviousBillNoQParams map[string]bool
var DeleteStockQParams map[string]bool

func SetApiHeaders() {
	PostShopOwnerHeaders = map[string]bool{CONTENT_TYPE: true, ACCEPT: true, ACCEPT_ENCODING: true}
	GenerateTokenHeaders = map[string]bool{CONTENT_TYPE: true, ACCEPT: true, ACCEPT_ENCODING: true}
	PutShopOwnerHeaders = map[string]bool{CONTENT_TYPE: true, ACCEPT: true, ACCEPT_ENCODING: true}
	GetShopOwnerHeaders = map[string]bool{ACCEPT: true, ACCEPT_ENCODING: true, CACHE_CONTROL: true}
	GetAllShopOwnerHeaders = map[string]bool{ACCEPT: true, ACCEPT_ENCODING: true, CACHE_CONTROL: true}
	PostCustomerHeaders = map[string]bool{TOKEN: true, CONTENT_TYPE: true, ACCEPT: true, ACCEPT_ENCODING: true}
	PutCustomerHeaders = map[string]bool{TOKEN: true, CONTENT_TYPE: true, ACCEPT: true, ACCEPT_ENCODING: true}
	GetCustomerHeaders = map[string]bool{TOKEN: true, ACCEPT: true, ACCEPT_ENCODING: true, CACHE_CONTROL: true}
	GetAllCustomerHeaders = map[string]bool{TOKEN: true, ACCEPT: true, ACCEPT_ENCODING: true, CACHE_CONTROL: true}
	GetFilteredCustomerHeaders = map[string]bool{TOKEN: true, CONTENT_TYPE: true, ACCEPT: true, ACCEPT_ENCODING: true, CACHE_CONTROL: true}
	PostStockHeaders = map[string]bool{TOKEN: true, CONTENT_TYPE: true, ACCEPT: true, ACCEPT_ENCODING: true}
	PutStockHeaders = map[string]bool{TOKEN: true, CONTENT_TYPE: true, ACCEPT: true, ACCEPT_ENCODING: true}
	GetStockHeaders = map[string]bool{TOKEN: true, ACCEPT: true, ACCEPT_ENCODING: true, CACHE_CONTROL: true}
	GetAllStockHeaders = map[string]bool{TOKEN: true, CONTENT_TYPE: true, ACCEPT: true, ACCEPT_ENCODING: true, CACHE_CONTROL: true}
	GetStockHistoryHeaders = map[string]bool{TOKEN: true, ACCEPT: true, ACCEPT_ENCODING: true, CACHE_CONTROL: true}
	GetPreviousBalanceHeaders = map[string]bool{TOKEN: true, ACCEPT: true, ACCEPT_ENCODING: true, CACHE_CONTROL: true}
	PostCustomerBillHeaders = map[string]bool{TOKEN: true, CONTENT_TYPE: true, ACCEPT: true, ACCEPT_ENCODING: true}
	PutCustomerTransactionHeaders = map[string]bool{TOKEN: true, CONTENT_TYPE: true, ACCEPT: true, ACCEPT_ENCODING: true}
	GetCustomerBillHeaders = map[string]bool{TOKEN: true, ACCEPT: true, ACCEPT_ENCODING: true, CACHE_CONTROL: true}
	GetAllCustomerBillHeaders = map[string]bool{TOKEN: true, ACCEPT: true, ACCEPT_ENCODING: true, CACHE_CONTROL: true}
	GetFilteredCustomerTransactionHeaders = map[string]bool{TOKEN: true, CONTENT_TYPE: true, ACCEPT: true, ACCEPT_ENCODING: true, CACHE_CONTROL: true}
	GetAllOwnerBillHeaders = map[string]bool{TOKEN: true, ACCEPT: true, ACCEPT_ENCODING: true, CACHE_CONTROL: true}
	GetPreviousBillNoHeaders = map[string]bool{TOKEN: true, ACCEPT: true, ACCEPT_ENCODING: true, CACHE_CONTROL: true}
	DeleteStockHeaders = map[string]bool{TOKEN: true, ACCEPT: true, ACCEPT_ENCODING: true}
}

func SetApiQParams() {
	PutShopOwnerQParams = map[string]bool{OWNER_REG_ID: true}
	GetShopOwnerQParams = map[string]bool{OWNER_REG_ID: true}
	PostCustomerQParams = map[string]bool{OWNER_REG_ID: true}
	PutCustomerQParams = map[string]bool{OWNER_REG_ID: true, CUSTOMER_REG_ID: true}
	GetCustomerQParams = map[string]bool{OWNER_REG_ID: true, CUSTOMER_REG_ID: true}
	GetAllCustomerQParams = map[string]bool{OWNER_REG_ID: true}
	GetFilteredCustomerQParams = map[string]bool{OWNER_REG_ID: true}
	PostStockQParams = map[string]bool{OWNER_REG_ID: true}
	PutStockQParams = map[string]bool{OWNER_REG_ID: true, STOCK_ID: true}
	GetStockQParams = map[string]bool{OWNER_REG_ID: true, STOCK_ID: true}
	GetAllStockQParams = map[string]bool{OWNER_REG_ID: true}
	GetStockHistoryQParams = map[string]bool{OWNER_REG_ID: true, STOCK_ID: true}
	GetPreviousBalanceQParams = map[string]bool{OWNER_REG_ID: true, CUSTOMER_REG_ID: true}
	PostCustomerBillQParams = map[string]bool{OWNER_REG_ID: true}
	PutCustomerTransactionQParams = map[string]bool{OWNER_REG_ID: true, CUSTOMER_REG_ID: true, BILL_ID: true}
	GetCustomerBillQParams = map[string]bool{OWNER_REG_ID: true, BILL_ID: true}
	GetAllCustomerBillQParams = map[string]bool{OWNER_REG_ID: true, CUSTOMER_REG_ID: true}
	GetAllOwnerBillQParams = map[string]bool{OWNER_REG_ID: true}
	GetPreviousBillNoQParams = map[string]bool{OWNER_REG_ID: true}
	DeleteStockQParams = map[string]bool{OWNER_REG_ID: true, STOCK_ID: true}
}

var ValidHeaders map[string][]interface{}

func SetValidHeaders() {
	ValidHeaders = make(map[string][]interface{})
	ValidHeaders[CONTENT_TYPE] = []interface{}{"application/json", "text/plain", "application.json; charset=utf-8"}
	ValidHeaders[ACCEPT] = []interface{}{"application/json, text/plain", "*/*"}
	ValidHeaders[ACCEPT_ENCODING] = []interface{}{"gzip, deflate, br", "gzip, deflate, br, zstd", "gzip, deflate"}
	ValidHeaders[CACHE_CONTROL] = []interface{}{"no-cache"}
}
