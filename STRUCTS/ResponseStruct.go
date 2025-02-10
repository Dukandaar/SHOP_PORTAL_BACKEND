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
	PhNo      string  `json:"ph_no"`
	RegDate   string  `json:"reg_date"`
	Address   string  `json:"address"`
	Remarks   string  `json:"remarks"`
	Gold      float32 `json:"gold"`
	Silver    float32 `json:"silver"`
	Cash      float32 `json:"cash"`
}

type AllShopOwnerDetailsResponse struct {
	Stat                           string                           `json:"stat"`
	Count                          int                              `json:"count"`
	AllShopOwnerDetailsSubResponse []AllShopOwnerDetailsSubResponse `json:"rsp"`
}
type AllShopOwnerDetailsSubResponse struct {
	ShopName  string  `json:"shopName"`
	OwnerName string  `json:"ownerName"`
	RegId     string  `json:"regId"`
	PhNo      string  `json:"phoneNo"`
	RegDate   string  `json:"regDate"`
	Address   string  `json:"address"`
	Remarks   string  `json:"remarks"`
	Gold      float32 `json:"gold"`
	Silver    float32 `json:"silver"`
	Cash      float32 `json:"cash"`
	IsActive  string  `json:"isActive"`
}

type CustomerDetailsResponse struct {
	Stat                       string                     `json:"stat"`
	CustomerDetailsSubResponse CustomerDetailsSubResponse `json:"rsp"`
}

type AllCustomerDetailsResponse struct {
	Stat                       string                       `json:"stat"`
	Count                      int                          `json:"count"`
	CustomerDetailsSubResponse []CustomerDetailsSubResponse `json:"rsp"`
}

type CustomerDetailsSubResponse struct {
	Name        string  `json:"name"`
	CompanyName string  `json:"company_name"`
	RegId       string  `json:"reg_id"`
	PhNo        string  `json:"ph_no"`
	RegDate     string  `json:"reg_date"`
	Address     string  `json:"address"`
	Remarks     string  `json:"remarks"`
	Gold        float32 `json:"gold"`
	Silver      float32 `json:"silver"`
	Cash        float32 `json:"cash"`
	IsActive    string  `json:"isActive"`
}
