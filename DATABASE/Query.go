package database

import utils "SHOP_PORTAL_BACKEND/UTILS"

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

func CheckValidOwnerRegId() string {
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

func GetOwnerRowId() string {
	query := `
        SELECT id
        FROM shop.owner
        WHERE reg_id = $1;
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

func GetShopOwnerData() string {
	query := `
		SELECT
			o.shop_name,
			o.owner_name,
			o.phone_no,
			o.reg_date::text,
			o.address,
			o.remarks,
			COALESCE(b_gold.balance, 0) as gold,
			COALESCE(b_silver.balance, 0) as silver,
			COALESCE(b_cash.balance, 0) as cash
		FROM
			shop.owner o
		LEFT JOIN LATERAL (
			SELECT balance
			FROM shop.balance
			WHERE owner_id = o.id AND type = 'Gold'
		) AS b_gold ON TRUE
		LEFT JOIN LATERAL (
			SELECT balance
			FROM shop.balance
			WHERE owner_id = o.id AND type = 'Silver'
		) AS b_silver ON TRUE
		LEFT JOIN LATERAL (
			SELECT balance
			FROM shop.balance
			WHERE owner_id = o.id AND type = 'Cash'
		) AS b_cash ON TRUE
		WHERE
			o.reg_id = $1;`
	return query
}

func GetAllShopOwnerData(isActiveStates string) string {
	query := `
		SELECT
			o.shop_name,
			o.owner_name,
			o.reg_id,
			o.phone_no,
			o.reg_date::text,
			o.address,
			o.remarks,
			COALESCE(b_gold.balance, 0) as gold,
			COALESCE(b_silver.balance, 0) as silver,
			COALESCE(b_cash.balance, 0) as cash,
			o.is_active
		FROM
			shop.owner o
		LEFT JOIN LATERAL (
			SELECT balance
			FROM shop.balance
			WHERE owner_id = o.id AND type = 'Gold'
		) AS b_gold ON TRUE
		LEFT JOIN LATERAL (
			SELECT balance
			FROM shop.balance
			WHERE owner_id = o.id AND type = 'Silver'
		) AS b_silver ON TRUE
		LEFT JOIN LATERAL (
			SELECT balance
			FROM shop.balance
			WHERE owner_id = o.id AND type = 'Cash'
		) AS b_cash ON TRUE
	`

	if isActiveStates == utils.ACTIVE_YES {
		query += " WHERE o.is_active = 'Y'"
	} else if isActiveStates == utils.ACTIVE_NO {
		query += " WHERE o.is_active = 'N'"
	} else {
		query += " WHERE o.is_active = 'Y' OR o.is_active = 'N'"
	}
	return query
}

func InsertCustomerData() string {
	query := `
        INSERT INTO
            shop.customer (name, company_name, reg_id, reg_date, ph_no, address, created_at, updated_at)
        VALUES
            ($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING id;
    `
	return query
}

func CheckCustomerPresent() string {
	query := `
        SELECT c.id, oc.is_active
        FROM shop.customer c
		JOIN shop.owner_customer oc ON oc.customer_id = c.id
        WHERE name = $1 and company_name = $2 and ph_no = $3;
    `
	return query
}

func CheckValidCustomerRegId() string {
	query := `
            SELECT EXISTS 
            (SELECT 1 FROM shop.customer 
            WHERE reg_id = $1)
    `
	return query
}

func UpdateCustomerData() string {
	query := `
        UPDATE shop.customer
        SET
            name = $1,
            company_name = $2,
            reg_date = $3,
            ph_no = $4,
            address = $5,
            updated_at = $6
        WHERE reg_id = $7;
    `
	return query
}

func GetCustomerData() string {
	query := `
        SELECT
            c.name,
            c.company_name,
            c.reg_id,
            c.reg_date::text,
            c.ph_no,
            c.address
        FROM
            shop.customer c
        WHERE
            c.reg_id = $1;
    `
	return query
}

func GetAllCustomerData() string {
	query := `
        SELECT
            c.name,
            c.company_name,
            c.reg_id,
            c.reg_date::text,
            c.ph_no,
            c.address
        FROM
            shop.customer c;
    `
	return query
}

// Owner Customer Function
func InsertOwnerCustomerData() string {
	query := `
        INSERT INTO
            shop.owner_customer (owner_id, customer_id, is_active, remark)
        VALUES
            ($1, $2, $3, $4);
    `
	return query
}

func CheckOwnerCustomerPresent() string {
	query := `
        SELECT id
        FROM shop.owner_customer
        WHERE owner_id = $1 AND customer_id = $2;
    `
	return query
}

func UpdateOwnerCustomerData() string {
	query := `
        UPDATE shop.owner_customer
        SET
            is_active = $1,
            remark = $2
        WHERE owner_id = $3 AND customer_id = $4;
    `
	return query
}

func GetOwnerCustomerData(ownerId int) string {
	query := `
        SELECT
            oc.customer_id,
            c.name as customer_name,
            c.company_name,
            c.reg_id as customer_reg_id,
            c.reg_date::text as customer_reg_date,
            c.ph_no as customer_ph_no,
            c.address as customer_address,
            oc.is_active,
            oc.remark
        FROM
            shop.owner_customer oc
        JOIN
            shop.customer c ON oc.customer_id = c.id
        WHERE
            oc.owner_id = $1;
    `
	return query
}
