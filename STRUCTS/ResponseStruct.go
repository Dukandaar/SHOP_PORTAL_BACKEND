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
	Id        int64   `json:"id"`
	BillId    int64   `json:"bill_id"`
	ItemName  string  `json:"item_name"`
	Weight    float64 `json:"weight"`
	Less      float64 `json:"less"`
	NetWeight float64 `json:"net_weight"`
	Tunch     float64 `json:"tunch"`
	Fine      float64 `json:"fine"`
	Discount  float64 `json:"discount"`
	Amount    float64 `json:"amount"`
}
