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
			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		RETURNING id;
	`
	return query
}

func InsertOwnerBillCount() string {
	query := `
		INSERT INTO
			shop.owner_bill_count (owner_id, bill_cnt, updated_at)
		VALUES
			($1, $2, $3);
	`
	return query
}

func GetOwnerBillCount() string {
	query := `
		SELECT
			bill_cnt
		FROM
			shop.owner_bill_count
		WHERE
			owner_id = $1;
	`
	return query
}

func UpdateOwnerBillCount() string {
	query := `
		UPDATE
			shop.owner_bill_count
		SET
			bill_cnt = bill_cnt + 1,
			updated_at = now()
		WHERE
			owner_id = $1;
	`
	return query
}

func InsertOwnerBalanceData() string {
	query := `
		INSERT INTO
			shop.balance (owner_id, gold, silver, cash, created_at, updated_at)
		VALUES
			($1, $2, $3, $4, $5, $6);
	`
	return query
}

func InsertCustomerBalanceData() string {
	query := `
		INSERT INTO
			shop.balance (customer_id, gold, silver, cash, created_at, updated_at)
		VALUES
			($1, $2, $3, $4, $5, $6);
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
			gst_in = $3,
			phone_no = $4,
			is_active = $5,
			reg_date = $6,
			address = $7,
			remarks = $8,
			updated_at = $9
		WHERE reg_id = $10;
	`
	return query
}

func GetShopOwnerData() string {
	query := `
		SELECT
			o.id,
			o.shop_name,
			o.owner_name,
			o.gst_in,
			o.phone_no,
			o.reg_date::text,
			o.address,
			o.remarks,
			o.is_active,
			b.gold,
			b.silver,
			b.cash
		FROM
			shop.owner o
		LEFT JOIN 
			shop.balance b 
		ON 
			o.id = b.owner_id
		WHERE
			o.reg_id = $1;`
	return query
}

func GetAllShopOwnerData(isActiveStates string) string {
	query := `
		SELECT
			o.shop_name,
			o.owner_name,
			o.gst_in,
			o.phone_no,
			o.reg_date::text,
			o.address,
			o.remarks,
			o.is_active,
			b.gold,
			b.silver,
			b.cash,
			obc.bill_cnt
		FROM
			shop.owner o
		LEFT JOIN 
			shop.balance b 
		ON 
			o.id = b.owner_id
		LEFT JOIN
			shop.owner_bill_count obc
		ON
			o.id = obc.owner_id
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
            shop.customer (is_active, owner_id, name, shop_name, reg_id, reg_date, phone_no, address, remarks, gst_in, created_at, updated_at)
        VALUES
            ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
        RETURNING id;
    `
	return query
}

func CheckCustomerPresent() string {
	query := `
        SELECT 
			id, is_active, reg_id
		FROM 
			shop.customer
		WHERE 
			owner_id = $1 AND name = $2 AND shop_name = $3 AND phone_no = $4;
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
				shop_name = $2,
				gst_in = $3,
				reg_date = $4,
				phone_no = $5,
				is_active = $6,
				address = $7,
				updated_at = $8
			WHERE 
				reg_id = $9;
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
			c.shop_name,
			c.name,
			c.gst_in,
			c.reg_date,
			c.phone_no,
			c.is_active,
			c.address,
			c.remarks,
			b.gold,
			b.silver,
			b.cash
		FROM
			shop.customer c
		LEFT JOIN
			shop.balance b 
		ON 
		    c.id = b.customer_id
		WHERE
			c.reg_id = $1 and c.owner_id = $2
    `
	return query
}

func GetAllCustomerData() string {
	query := `
        SELECT
			c.shop_name,
			c.name,
			c.gst_in,
			c.reg_id,
			c.phone_no,
			c.reg_date,
			c.is_active,
			c.address,
			c.remarks,
			b.gold,
			b.silver,
			b.cash
		FROM
			shop.customer c
		LEFT JOIN
			shop.balance b 
		ON 
		    c.id = b.customer_id
		WHERE
			c.owner_id = $1;
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
        UPDATE shop.customer
        SET
            is_active = $1,
            remarks = $2
        WHERE reg_id = $3;
    `
	return query
}

func GetOwnerCustomerData(ownerId int) string {
	query := `
        SELECT
            oc.customer_id,
            c.name as customer_name,
            c.shop_name,
            c.reg_id as customer_reg_id,
            c.reg_date::text as customer_reg_date,
            c.phone_no as customer_phone_no,
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
			c.shop_name,
			c.name,
			c.gst_in,
			c.reg_id,
			c.phone_no,
			c.reg_date,
			c.is_active,
			c.address,
			c.remarks,
			b.gold,
			b.silver,
			b.cash

		FROM
			shop.customer c
		LEFT JOIN
			shop.balance b 
		ON 
		    c.id = b.customer_id
		WHERE
			c.owner_id = $1`

	whereClauses := []string{}

	if filter.RegId != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("c.reg_id = '%s'", filter.RegId))
	}
	if filter.Name != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("c.name ILIKE '%%%s%%'", filter.Name))
	}
	if filter.ShopName != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("c.shop_name ILIKE '%%%s%%'", filter.ShopName))
	}
	if filter.PhoneNo != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("c.phone_no ILIKE '%%%s%%'", filter.PhoneNo))
	}
	if filter.RegDate != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("c.reg_date = '%s'", filter.RegDate))
	}
	if filter.IsActive != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("c.is_active = '%s'", filter.IsActive))
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
		}
	}
	if filter.Id > 0 {
		whereClauses = append(whereClauses, fmt.Sprintf("c.id = %d", filter.Id))
	}

	if len(whereClauses) > 0 {
		query += " AND " + strings.Join(whereClauses, " AND ")
	}
	return query
}

func CheckValidStockId() string {
	query := `
		SELECT EXISTS
		(SELECT 1 FROM shop.stock WHERE id = $1)
	`
	return query
}

func CheckStockPresent() string {
	query := `
		SELECT id
		FROM shop.stock
		WHERE owner_id = $1 AND type = $2 AND item_name = $3 AND tunch = $4;
	`
	return query
}

func InsertStockData() string {
	query := `
		INSERT INTO
			shop.stock (owner_id, type, item_name, tunch, weight, created_at, updated_at)
		VALUES
			($1, $2, $3, $4, $5, $6, $7)
		RETURNING id;
	`
	return query
}

func InsertStockHistoryData() string {
	query := `
		INSERT INTO
			shop.stock_history (stock_id, prev_balance, new_balance, reason, remarks, created_at)
		VALUES
			($1, $2, $3, $4, $5, $6);
	`
	return query
}

func UpdateStockData() string {
	query := `
		UPDATE
			shop.stock
		SET
			tunch = $1,
			weight = $2,
			updated_at = $3
		WHERE
			id = $4 and owner_id = $5;
	`
	return query
}

func GetStockData() string {
	query := `
		SELECT
			s.id,
			s.name,
			s.description,
			s.price,
			s.is_active,
			s.created_at,
			s.updated_at
		FROM
			shop.stock s
		WHERE
			s.id = $1;
	`
	return query
}

func GetCustomerPreviousBalance() string {
	query := `
		SELECT
			b.gold,
			b.silver,
			b.cash
		FROM
			shop.balance b
		WHERE
			b.customer_id = $1;
	`
	return query
}

func GetCustomerId() string {
	query := `
		SELECT
			c.id
		FROM
			shop.customer c
		WHERE
			c.reg_id = $1 and c.owner_id = $2;
	`
	return query
}

func CreateBill() string {
	query := `
		INSERT INTO
			shop.bill (bill_no, customer_id, type, metal, metal_rate, date, created_at, updated_at)
		VALUES
			($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id;
	`
	return query
}

func AddTransaction() string {
	query := `
		INSERT INTO
			shop.transaction (bill_id, item_name, weight, less, net_weight, tunch, fine, discount, amount)
		VALUES
			($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id;
	`
	return query
}

func GetStockId() string {
	query := `
		SELECT
			id,
			weight
		FROM
			shop.stock
		WHERE
			item_name = $1;
	`
	return query
}

func DecreaseStock() string {
	query := `
		UPDATE
			shop.stock
		SET
			weight =  $1
		WHERE
			id = $2;
	`
	return query
}

func AddStockHistory() string {
	query := `
		INSERT INTO
			shop.stock_history (stock_id, prev_balance, new_balance, reason, transaction_id, remarks, created_at)
		VALUES
			($1, $2, $3, $4, $5, $6, $7);
	`
	return query
}

func BillPayment() string {
	query := `
		INSERT INTO
			shop.payment (bill_id, customer_id, factor, new, prev, total, paid, rem, type, date, remarks, created_at, updated_at)
		VALUES
			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13);
	`
	return query
}

func UpdateCustomerBalance(Type string) string {
	query := `
		UPDATE
			shop.balance
		SET `

	if Type == utils.GOLD {
		query += `gold = $1,`
	} else if Type == utils.SILVER {
		query += `silver = $1,`
	} else if Type == utils.CASH {
		query += `cash = $1,`
	}

	query += `
			updated_at = $2
		WHERE
			customer_id = $3;`
	return query
}

func CheckValidBillId() string {
	query := `
		SELECT EXISTS
		(
			SELECT
				1
			FROM
				shop.bill
			WHERE
				id = $1
		)`
	return query
}

func UpdateBill() string {
	query := `
		UPDATE
			shop.bill
		SET
			type = $1,
			metal = $2,
			metal_rate = $3,
			date = $4,
			updated_at = $5
		WHERE
			id = $6;
	`
	return query
}

func GetBillTransactions() string {
	query := `
		SELECT
			id,
			item_name,
			weight,
			less,
			net_weight,
			tunch,
			fine,
			discount,
			amount
		FROM
			shop.transaction
		WHERE
			bill_id = $1;
	`
	return query
}
