package database

func InsertShopOwnerData() string {
	query := `
		INSERT INTO
			shop.owner (shop_name, owner_name, reg_id, phone_no, is_active, reg_date, address, remarks, key, created_at, updated_at)
		VALUES
			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);
	`
	return query
}

func CheckOwnerPresent() string {
	query := `
		SELECT id, reg_id, is_active
		FROM shop.owner
		WHERE owner_name = $1 AND shop_name = $2 AND phone_no = $3;
	`
	return query
}

func CheckValidRegId() string {
	query := `
			SELECT EXISTS 
			(SELECT 1 FROM shop.owner 
			WHERE reg_id = $1)
	`
	return query
}

func ToggleShopOwnerActiveStatus() string {
	query := `
		UPDATE shop.owner
		SET
			is_active = $1,
			updated_at = $2
		WHERE id = $3;
	`
	return query
}

func UpdateShopOwnerData() string {
	query := `
		UPDATE shop.owner
		SET
			shop_name = $1,
			owner_name = $2,
			phone_no = $3,
			is_active = $4,
			reg_date = $5,
			address = $6,
			remarks = $7,
			updated_at = $8
		WHERE reg_id = $9;
	`
	return query
}
