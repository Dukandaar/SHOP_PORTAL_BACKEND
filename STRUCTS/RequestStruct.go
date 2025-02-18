package structs

type GenerateToken struct {
	RegId string `json:"reg_id"`
	Key   string `json:"key"`
}

type ShopOwner struct {
	ShopName  string `json:"shop_name"`
	OwnerName string `json:"owner_name"`
	GstIN     string `json:"gst_in"`
	PhoneNo   string `json:"phone_no"`
	RegDate   string `json:"reg_date"`
	Address   string `json:"address"`
	Remarks   string `json:"remarks"`
}

type AllShopOwner struct {
	IsActive string `json:"is_active"`
}

type Customer struct {
	Name     string `json:"name"`
	ShopName string `json:"shop_name"`
	GstIN    string `json:"gst_in"`
	PhoneNo  string `json:"phone_no"`
	RegDate  string `json:"reg_date"`
	Address  string `json:"address"`
	Remarks  string `json:"remarks"`
}

type FilteredCustomer struct {
	Id           int          `json:"id"`
	RegId        string       `json:"reg_id"`
	Name         string       `json:"name"`
	ShopName     string       `json:"shop_name"`
	PhoneNo      string       `json:"phone_no"`
	RegDate      string       `json:"reg_date"`
	IsActive     string       `json:"is_active"`
	DateInterval DateInterval `json:"date_interval"`
}

type DateInterval struct {
	Type  string `json:"type"`
	Start string `json:"start"`
	End   string `json:"end"`
}

type PostStock struct {
	Type     string  `json:"type"`
	ItemName string  `json:"item_name"`
	Tunch    float32 `json:"tunch"`
	Weight   float32 `json:"weight"`
}

type PutStock struct {
	Tunch   float32 `json:"tunch"`
	Weight  float32 `json:"weight"`
	Remarks string  `json:"remarks"`
}
