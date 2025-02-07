package database

func InsertShopOwnerData() string {
	query := `
		INSERT INTO
			shop.shop_owner (shop_name, owner_name, reg_date, ph_no, address, is_active, remarks)
		VALUES
			($1, $2, $3, $4, $5, $6, $7);
	`
	return query
}
