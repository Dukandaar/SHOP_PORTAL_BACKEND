package structs

type ShopOwner struct {
	ShopName  string `json:"shop_name"`
	OwnerName string `json:"owner_name"`
	PhNo      string `json:"ph_no"`
	RegDate   string `json:"reg_date"`
	Address   string `json:"address"`
	Remarks   string `json:"remarks"`
}
