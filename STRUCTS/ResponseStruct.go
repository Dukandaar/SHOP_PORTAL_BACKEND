package structs

type ErrorResponse struct {
	Stat             string           `json:"stat"`
	ErrorSubResponse ErrorSubResponse `json:"rsp"`
}

type ErrorSubResponse struct {
	ErrorCode       int    `json:"code"`
	ErrorMsg        string `json:"msg"`
	ErrorDescrition string `json:"description"`
}

type SuccessResponse struct {
	Stat               string             `json:"stat"`
	SuccessSubResponse SuccessSubResponse `json:"rsp"`
}

type SuccessSubResponse struct {
	SuccessMsg string `json:"msg"`
}

type SuccessRegIdResponse struct {
	Stat               string                  `json:"stat"`
	SuccessSubResponse SuccessRegIdSubResponse `json:"rsp"`
}

type SuccessIdResponse struct {
	Stat               string               `json:"stat"`
	SuccessSubResponse SuccessIdSubResponse `json:"rsp"`
}

type SuccessIdSubResponse struct {
	SuccessMsg string `json:"msg"`
	Id         int    `json:"id"`
}

type CreateOwnerSuccessResponseWithIdKey struct {
	Stat               string                                 `json:"stat"`
	SuccessSubResponse CreateOwnerSuccessSubResponseWithIdKey `json:"rsp"`
}

type CreateOwnerSuccessSubResponseWithIdKey struct {
	SuccessMsg string `json:"msg"`
	RegId      string `json:"reg_id"`
	Key        string `json:"key"`
}

type SuccessRegIdSubResponse struct {
	SuccessMsg string `json:"msg"`
	RegId      string `json:"reg_id"`
}

type TokenResponse struct {
	Stat             string           `json:"stat"`
	TokenSubResponse TokenSubResponse `json:"rsp"`
}

type TokenSubResponse struct {
	Token string `json:"token"`
}

type ShopOwnerDetailsResponse struct {
	Stat                        string                      `json:"stat"`
	ShopOwnerDetailsSubResponse ShopOwnerDetailsSubResponse `json:"rsp"`
}

type ShopOwnerDetailsSubResponse struct {
	ShopName  string  `json:"shop_name"`
	OwnerName string  `json:"owner_name"`
	GstIN     string  `json:"gst_in"`
	PhoneNo   string  `json:"phone_no"`
	RegDate   string  `json:"reg_date"`
	Address   string  `json:"address"`
	Remarks   string  `json:"remarks"`
	Gold      float64 `json:"gold"`
	Silver    float64 `json:"silver"`
	Cash      float64 `json:"cash"`
	IsActive  string  `json:"is_active"`
	BillCount int     `json:"bill_count"`
}

type AllShopOwnerDetailsResponse struct {
	Stat                           string                        `json:"stat"`
	Count                          int                           `json:"count"`
	AllShopOwnerDetailsSubResponse []ShopOwnerDetailsSubResponse `json:"rsp"`
}

type CustomerDetailsResponse struct {
	Stat                       string                     `json:"stat"`
	Count                      int                        `json:"count"`
	CustomerDetailsSubResponse CustomerDetailsSubResponse `json:"rsp"`
}

type AllCustomerDetailsResponse struct {
	Stat                       string                       `json:"stat"`
	Count                      int                          `json:"count"`
	CustomerDetailsSubResponse []CustomerDetailsSubResponse `json:"rsp"`
}

type CustomerDetailsSubResponse struct {
	Name      string  `json:"name"`
	ShopName  string  `json:"shop_name"`
	GstIN     string  `json:"gst_in"`
	RegId     string  `json:"reg_id"`
	PhoneNo   string  `json:"phone_no"`
	RegDate   string  `json:"reg_date"`
	Address   string  `json:"address"`
	Remarks   string  `json:"remarks"`
	Gold      float64 `json:"gold"`
	Silver    float64 `json:"silver"`
	Cash      float64 `json:"cash"`
	IsActive  string  `json:"isActive"`
	BillCount int     `json:"bill_count"`
}

type CustomerPreviousBalanceResponse struct {
	Stat                               string                               `json:"stat"`
	CustomerPreviousBalanceSubResponse []CustomerPreviousBalanceSubResponse `json:"rsp"`
}

type CustomerPreviousBalanceSubResponse struct {
	RowId  int     `json:"row_id"`
	Gold   float64 `json:"gold"`
	Silver float64 `json:"silver"`
	Cash   float64 `json:"cash"`
}

type OwnerStockResponse struct {
	Stat                  string                `json:"stat"`
	OwnerStockSubResponse OwnerStockSubResponse `json:"rsp"`
}

type OwnerStockSubResponse struct {
	ItemName  string  `json:"item_name"`
	Tunch     float64 `json:"tunch"`
	Weight    float64 `json:"weight"`
	UpdatedAt string  `json:"updated_at"`
}

type OwnerAllStockResponse struct {
	Stat                  string                  `json:"stat"`
	Count                 int                     `json:"count"`
	OwnerStockSubResponse []OwnerStockSubResponse `json:"rsp"`
}

type StockHistoryResponse struct {
	Stat                    string                    `json:"stat"`
	Count                   int                       `json:"count"`
	StockHistorySubResponse []StockHistorySubResponse `json:"rsp"`
}

type StockHistorySubResponse struct {
	PrevBalance float64             `json:"prev_balance"`
	NewBalance  float64             `json:"new_balance"`
	Reason      string              `json:"reason"`
	Remarks     string              `json:"remarks"`
	CreatedAt   string              `json:"created_at"`
	Transaction TransactionResponse `json:"transaction"`
}

type TransactionResponse struct {
	Id        interface{} `json:"id"`
	BillId    interface{} `json:"bill_id"`
	ItemName  interface{} `json:"item_name"`
	Weight    interface{} `json:"weight"`
	Less      interface{} `json:"less"`
	NetWeight interface{} `json:"net_weight"`
	Tunch     interface{} `json:"tunch"`
	Fine      interface{} `json:"fine"`
	Discount  interface{} `json:"discount"`
	Amount    interface{} `json:"amount"`
}

type CustomerBillResponse struct {
	Stat                    string                  `json:"stat"`
	CustomerBillSubResponse CustomerBillSubResponse `json:"rsp"`
}

type CustomerBillSubResponse struct {
	BillNo             int           `json:"bill_no"`
	Type               string        `json:"type"`
	Metal              string        `json:"metal"`
	Rate               float64       `json:"rate"`
	Date               string        `json:"date"`
	Remarks            string        `json:"remarks"`
	CustomerDetails    Customer      `json:"customer_details"`
	TransactionDetails []Transaction `json:"transaction_details"`
	PaymentDetails     Payment       `json:"payment_details"`
	CreatedAt          string        `json:"created_at"`
	UpdatedAt          string        `json:"updated_at"`
}

type CustomerAllBillResponse struct {
	Stat                    string                    `json:"stat"`
	Count                   int                       `json:"count"`
	CustomerBillSubResponse []CustomerBillSubResponse `json:"rsp"`
}
