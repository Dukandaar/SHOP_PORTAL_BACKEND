package structs

type ErrorResponse struct {
	Response ErrorSubResponse `json:"rsp"`
}

type ErrorSubResponse struct {
	Stat        string               `json:"stat"`
	Payload     ErrorPayloadResponse `json:"payload"`
	Description string               `json:"description"`
}

type ErrorPayloadResponse struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

type SuccessResponse struct {
	Response SuccessSubResponse `json:"rsp"`
}

type SuccessSubResponse struct {
	Stat        string                 `json:"stat"`
	Payload     SuccessPayloadResponse `json:"payload"`
	Description string                 `json:"description"`
}

type SuccessPayloadResponse struct {
	Message string `json:"msg"`
}

// Generate Token Response
type GenerateTokenResponse struct {
	Response GenerateTokenSubResponse `json:"rsp"`
}

type GenerateTokenSubResponse struct {
	Stat        string                       `json:"stat"`
	Payload     GenerateTokenPayloadResponse `json:"payload"`
	Description string                       `json:"description"`
}

type GenerateTokenPayloadResponse struct {
	Token string `json:"token"`
}

// Post Shop Owner Response
type PostShopOwnerResponse struct {
	Response PostShopOwnerSubResponse `json:"rsp"`
}

type PostShopOwnerSubResponse struct {
	Stat        string                       `json:"stat"`
	Payload     PostShopOwnerPayloadResponse `json:"payload"`
	Description string                       `json:"description"`
}

type PostShopOwnerPayloadResponse struct {
	RegId string `json:"reg_id"`
	Key   string `json:"key"`
}

// Get Shop Owner Response
type GetShopOwnerResponse struct {
	Response GetShopOwnerSubResponse `json:"rsp"`
}

type GetShopOwnerSubResponse struct {
	Stat        string                      `json:"stat"`
	Payload     GetShopOwnerPayloadResponse `json:"payload"`
	Description string                      `json:"description"`
}

type GetShopOwnerPayloadResponse struct {
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

// Get All Shop Owner Response
type GetAllShopOwnerResponse struct {
	Response GetAllShopOwnerSubResponse `json:"rsp"`
}

type GetAllShopOwnerSubResponse struct {
	Stat        string                        `json:"stat"`
	Count       int                           `json:"count"`
	Payload     []GetShopOwnerPayloadResponse `json:"payload"`
	Description string                        `json:"description"`
}

// Post Customer Response
type PostCustomerResponse struct {
	Response PostCustomerSubResponse `json:"rsp"`
}

type PostCustomerSubResponse struct {
	Stat        string                      `json:"stat"`
	Payload     PostCustomerPayloadResponse `json:"payload"`
	Description string                      `json:"description"`
}

type PostCustomerPayloadResponse struct {
	RegId string `json:"reg_id"`
}

// Get Customer Response
type GetCustomerResponse struct {
	Response GetCustomerSubResponse `json:"rsp"`
}

type GetCustomerSubResponse struct {
	Stat        string                     `json:"stat"`
	Payload     GetCustomerPayloadResponse `json:"payload"`
	Description string                     `json:"description"`
}

type GetCustomerPayloadResponse struct {
	ShopName string  `json:"shop_name"`
	Name     string  `json:"owner_name"`
	RegId    string  `json:"reg_id"`
	GstIN    string  `json:"gst_in"`
	PhoneNo  string  `json:"phone_no"`
	RegDate  string  `json:"reg_date"`
	Address  string  `json:"address"`
	Remarks  string  `json:"remarks"`
	Gold     float64 `json:"gold"`
	Silver   float64 `json:"silver"`
	Cash     float64 `json:"cash"`
	IsActive string  `json:"is_active"`
}

// Get All Customer Response
type GetAllCustomerResponse struct {
	Response GetAllCustomerSubResponse `json:"rsp"`
}

type GetAllCustomerSubResponse struct {
	Stat        string                       `json:"stat"`
	Count       int                          `json:"count"`
	Payload     []GetCustomerPayloadResponse `json:"payload"`
	Description string                       `json:"description"`
}

// Stock Response
type PostStockResponse struct {
	Response PostStockSubResponse `json:"rsp"`
}

type PostStockSubResponse struct {
	Stat        string                   `json:"stat"`
	Payload     PostStockPayloadResponse `json:"payload"`
	Description string                   `json:"description"`
}

type PostStockPayloadResponse struct {
	Id int `json:"id"`
}

// Get Stock Response
type OwnerStockResponse struct {
	Response OwnerStockSubResponse `json:"rsp"`
}

type OwnerStockSubResponse struct {
	Stat        string                    `json:"stat"`
	Payload     OwnerStockPayloadResponse `json:"payload"`
	Description string                    `json:"description"`
}

type OwnerStockPayloadResponse struct {
	Id        int     `json:"id"`
	Type      string  `json:"type"`
	ItemName  string  `json:"item_name"`
	Tunch     float64 `json:"tunch"`
	Weight    float64 `json:"weight"`
	UpdatedAt string  `json:"updated_at"`
}

// Get All Stock Response
type OwnerAllStockResponse struct {
	Response OwnerAllStockSubResponse `json:"rsp"`
}

type OwnerAllStockSubResponse struct {
	Stat        string                      `json:"stat"`
	Count       int                         `json:"count"`
	Payload     []OwnerStockPayloadResponse `json:"payload"`
	Description string                      `json:"description"`
}

// Get Stock History Response
type StockHistoryResponse struct {
	Response StockHistorySubResponse `json:"rsp"`
}

type StockHistorySubResponse struct {
	Stat        string                        `json:"stat"`
	Count       int                           `json:"count"`
	Payload     []StockHistoryPayloadResponse `json:"payload"`
	Description string                        `json:"description"`
}

type StockHistoryPayloadResponse struct {
	PrevBalance float64             `json:"prev_balance"`
	NewBalance  float64             `json:"new_balance"`
	Reason      string              `json:"reason"`
	Remarks     string              `json:"remarks"`
	CreatedAt   string              `json:"created_at"`
	Transaction TransactionResponse `json:"transaction"`
}

// Get Previous Customer Balance
type CustomerPreviousBalanceResponse struct {
	Response CustomerPreviousBalanceSubResponse `json:"rsp"`
}

type CustomerPreviousBalanceSubResponse struct {
	Stat        string                                 `json:"stat"`
	Payload     CustomerPreviousBalancePayloadResponse `json:"payload"`
	Description string                                 `json:"description"`
}

type CustomerPreviousBalancePayloadResponse struct {
	RowId  int     `json:"row_id"`
	Gold   float64 `json:"gold"`
	Silver float64 `json:"silver"`
	Cash   float64 `json:"cash"`
}

type SuccessIdResponse struct {
	Stat               string               `json:"stat"`
	SuccessSubResponse SuccessIdSubResponse `json:"rsp"`
}

type SuccessIdSubResponse struct {
	SuccessMsg string `json:"msg"`
	Id         int    `json:"id"`
}

type SuccessRegIdSubResponse struct {
	SuccessMsg string `json:"msg"`
	RegId      string `json:"reg_id"`
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

type CustomerBillResponse struct {
	Response CustomerBillSubResponse `json:"rsp"`
}

type CustomerBillSubResponse struct {
	Stat        string              `json:"stat"`
	Payload     BillPayloadResponse `json:"payload"`
	Description string              `json:"description"`
}

type BillPayloadResponse struct {
	Id                 int           `json:"id"`
	BillNo             int           `json:"bill_no"`
	Type               string        `json:"type"`
	Metal              string        `json:"metal"`
	Rate               float64       `json:"rate"`
	Date               string        `json:"date"`
	Time               string        `json:"time"`
	Remarks            string        `json:"remarks"`
	CustomerDetails    Customer      `json:"customer_details"`
	TransactionDetails []Transaction `json:"transaction_details"`
	PaymentDetails     Payment       `json:"payment_details"`
	CreatedAt          string        `json:"created_at"`
	UpdatedAt          string        `json:"updated_at"`
}

type AllBillResponse struct {
	Response AllBillSubResponse `json:"rsp"`
}

type AllBillSubResponse struct {
	Stat        string                `json:"stat"`
	Count       int                   `json:"count"`
	Payload     []BillPayloadResponse `json:"payload"`
	Description string                `json:"description"`
}
