package structs

type GenerateToken struct {
	RegId string `json:"reg_id"`
	Key   string `json:"key"`
}

type ShopOwner struct {
	ShopName  string `json:"shop_name"`
	OwnerName string `json:"owner_name"`
	PhNo      string `json:"ph_no"`
	RegDate   string `json:"reg_date"`
	Address   string `json:"address"`
	Remarks   string `json:"remarks"`
}
