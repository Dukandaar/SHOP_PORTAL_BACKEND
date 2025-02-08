package database

func InsertShopOwnerData() string {
	query := `
		INSERT INTO
			shop.owner (shop_name, owner_name, reg_id, phone_no, is_active, reg_date, address, remarks, created_at, updated_at)
		VALUES
			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);
	`
	return query
}
