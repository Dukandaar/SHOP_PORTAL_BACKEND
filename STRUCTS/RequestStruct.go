package structs

type ShopOwner struct {
	ShopName  string `json:"shop_name"`
	OwnerName string `json:"owner_name"`
	RegDate   string `json:"reg_date"`
	PhNo      string `json:"ph_no"`
	Address   string `json:"address"`
	Remarks   string `json:"remarks"`
}
