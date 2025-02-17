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
	Gold      float32 `json:"gold"`
	Silver    float32 `json:"silver"`
	Cash      float32 `json:"cash"`
	IsActive  string  `json:"is_active"`
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
	Name     string  `json:"name"`
	ShopName string  `json:"shop_name"`
	GstIN    string  `json:"gst_in"`
	RegId    string  `json:"reg_id"`
	PhoneNo  string  `json:"phone_no"`
	RegDate  string  `json:"reg_date"`
	Address  string  `json:"address"`
	Remarks  string  `json:"remarks"`
	Gold     float32 `json:"gold"`
	Silver   float32 `json:"silver"`
	Cash     float32 `json:"cash"`
	IsActive string  `json:"isActive"`
}
