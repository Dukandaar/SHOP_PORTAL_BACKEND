package database

import (
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"fmt"
	"strings"
	"time"
)

func InsertShopOwnerData() string {
	query := `
		INSERT INTO
			shop.owner (shop_name, owner_name, reg_id, gst_in, phone_no, is_active, reg_date, address, remarks, key, created_at, updated_at)
		VALUES
			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);
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

func CheckRegIdPresent() string {
	query := `
		SELECT EXISTS 
		(SELECT 1 FROM shop.owner 
		WHERE reg_id = $1)
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
        SELECT c.id, oc.is_active, c.reg_id
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
			UPDATE 
				shop.customer c
			SET
				name = $1,
				company_name = $2,
				ph_no = $3,
				reg_date = $4,
				address = $5,
				updated_at = $6
			WHERE c.id = $7;
	`
	return query
}

func UpdateCustomerOwnerData() string {
	query := `
			UPDATE 
				shop.owner_customer
			SET
				is_active = $1,
				remark = $2
			WHERE owner_id = $3 AND customer_id = $4;
	`
	return query
}

func GetCustomerData() string {
	query := `
        SELECT
            c.name,
            c.company_name,
            c.ph_no,
            c.reg_date::text,
            c.address,
            oc.remark,
            COALESCE(b_gold.balance, 0) as gold,
            COALESCE(b_silver.balance, 0) as silver,
            COALESCE(b_cash.balance, 0) as cash,
			oc.is_active
        FROM
            shop.customer c
        LEFT JOIN
            shop.owner_customer oc ON c.id = oc.customer_id
        LEFT JOIN LATERAL (
            SELECT balance
            FROM shop.balance
            WHERE customer_id = c.id AND type = 'Gold'
        ) AS b_gold ON TRUE
        LEFT JOIN LATERAL (
            SELECT balance
            FROM shop.balance
            WHERE customer_id = c.id AND type = 'Silver'
        ) AS b_silver ON TRUE
        LEFT JOIN LATERAL (
            SELECT balance
            FROM shop.balance
            WHERE customer_id = c.id AND type = 'Cash'
        ) AS b_cash ON TRUE
        WHERE
            c.reg_id = $1 AND oc.owner_id = $2;
    `
	return query
}

func GetAllCustomerData(isActiveStates string) string {
	query := `
	SELECT
		c.name,
		c.company_name,
		c.reg_id,
		c.ph_no,
		c.reg_date::text,
		c.address,
		oc.remark,
		COALESCE(b_gold.balance, 0) as gold,
		COALESCE(b_silver.balance, 0) as silver,
		COALESCE(b_cash.balance, 0) as cash,
		oc.is_active
	FROM
		shop.customer c
	LEFT JOIN
		shop.owner_customer oc ON c.id = oc.customer_id
	LEFT JOIN LATERAL (
		SELECT balance
		FROM shop.balance
		WHERE customer_id = c.id AND type = 'Gold'
	) AS b_gold ON TRUE
	LEFT JOIN LATERAL (
		SELECT balance
		FROM shop.balance
		WHERE customer_id = c.id AND type = 'Silver'
	) AS b_silver ON TRUE
	LEFT JOIN LATERAL (
		SELECT balance
		FROM shop.balance
		WHERE customer_id = c.id AND type = 'Cash'
	) AS b_cash ON TRUE
	WHERE c.id IN (SELECT customer_id FROM shop.owner_customer); -- Get all customers related to an owner

`

	if isActiveStates == utils.ACTIVE_YES {
		query += " AND oc.is_active = 'Y'" // Filter on owner_customer.is_active
	} else if isActiveStates == utils.ACTIVE_NO {
		query += " AND oc.is_active = 'N'" // Filter on owner_customer.is_active
	}
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

func GetCustomerRegId() string {
	return `SELECT reg_id FROM shop.customer WHERE reg_id = $1;`
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
func GetFilteredCustomerData(filter structs.FilteredCustomer) string {
	query := `
        SELECT
            c.name,
            c.company_name,
            c.reg_id,
            c.ph_no,
            c.reg_date::text,
            c.address,
            COALESCE(oc.remark, '') as remark,
            COALESCE(b_gold.balance, 0) as gold,
            COALESCE(b_silver.balance, 0) as silver,
            COALESCE(b_cash.balance, 0) as cash,
            oc.is_active
        FROM
            shop.customer c
        LEFT JOIN
            shop.owner_customer oc ON c.id = oc.customer_id
        LEFT JOIN LATERAL (
            SELECT balance
            FROM shop.balance
            WHERE customer_id = c.id AND type = 'Gold'
        ) AS b_gold ON TRUE
        LEFT JOIN LATERAL (
            SELECT balance
            FROM shop.balance
            WHERE customer_id = c.id AND type = 'Silver'
        ) AS b_silver ON TRUE
        LEFT JOIN LATERAL (
            SELECT balance
            FROM shop.balance
            WHERE customer_id = c.id AND type = 'Cash'
        ) AS b_cash ON TRUE
        WHERE oc.owner_id = $1`

	whereClauses := []string{}

	if filter.RegId != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("c.reg_id = '%s'", filter.RegId))
	}
	if filter.Name != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("c.name ILIKE '%%%s%%'", filter.Name))
	}
	if filter.CompanyName != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("c.company_name ILIKE '%%%s%%'", filter.CompanyName))
	}
	if filter.PhNo != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("c.ph_no ILIKE '%%%s%%'", filter.PhNo))
	}
	if filter.RegDate != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("c.reg_date = '%s'", filter.RegDate))
	}
	if filter.IsActive != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("oc.is_active = '%s'", filter.IsActive))
	}

	if filter.DateInterval.Type != "" {
		switch filter.DateInterval.Type {
		case "Today":
			whereClauses = append(whereClauses, "c.reg_date = CURRENT_DATE")
		case "Week":
			whereClauses = append(whereClauses, "c.reg_date BETWEEN CURRENT_DATE - INTERVAL '7 days' AND CURRENT_DATE")
		case "Month":
			whereClauses = append(whereClauses, "c.reg_date BETWEEN date_trunc('month', CURRENT_DATE) AND CURRENT_DATE")
		case "Year":
			whereClauses = append(whereClauses, "c.reg_date BETWEEN date_trunc('year', CURRENT_DATE) AND CURRENT_DATE")
		case "All":
			// No additional WHERE clause needed for "All"
		case "Custom":
			if filter.DateInterval.Start != "" && filter.DateInterval.End != "" {
				startDate, _ := time.Parse("2006-01-02", filter.DateInterval.Start)
				endDate, _ := time.Parse("2006-01-02", filter.DateInterval.End)
				whereClauses = append(whereClauses, fmt.Sprintf("c.reg_date BETWEEN '%s' AND '%s'", startDate.Format("2006-01-02"), endDate.Format("2006-01-02")))
			}
		default:
			utils.Logger.Warn("Invalid DateInterval Type:", filter.DateInterval.Type) // Log invalid type
		}
	}
	if filter.Id > 0 {
		whereClauses = append(whereClauses, fmt.Sprintf("c.id = %d", filter.Id))
	}

	if len(whereClauses) > 0 {
		query += " AND " + strings.Join(whereClauses, " AND ")
	}

	utils.Logger.Info(query)
	return query
}
