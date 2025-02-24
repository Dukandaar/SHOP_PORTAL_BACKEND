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
	Tunch    float64 `json:"tunch"`
	Weight   float64 `json:"weight"`
}

type PutStock struct {
	Tunch   float64 `json:"tunch"`
	Weight  float64 `json:"weight"`
	Remarks string  `json:"remarks"`
}

type Transaction struct {
	ItemName  string  `json:"item_name"`
	Weight    float64 `json:"weight"`
	Less      float64 `json:"less"`
	NetWeight float64 `json:"net_weight"`
	Tunch     float64 `json:"tunch"`
	Fine      float64 `json:"fine"`
	Discount  float64 `json:"discount"`
	Amount    float64 `json:"amount"`
}

type Payment struct {
	Factor      string  `json:"factor"`
	New         float64 `json:"new"`
	Prev        float64 `json:"prev"`
	Total       float64 `json:"total"`
	Paid        float64 `json:"paid"`
	Rem         float64 `json:"rem"`
	PaymentType string  `json:"payment_type"`
	Remarks     string  `json:"remarks"`
}

type CustomerBill struct {
	BillNo             int           `json:"bill_no"`
	Type               string        `json:"type"`
	Metal              string        `json:"metal"`
	Rate               float64       `json:"rate"`
	Date               string        `json:"date"`
	Remarks            string        `json:"remarks"`
	CustomerDetails    Customer      `json:"customer_details"`
	TransactionDetails []Transaction `json:"transaction_details"`
	PaymentDetails     Payment       `json:"payment_details"`
}

type AllStock struct {
	Type string `json:"type"`
}
