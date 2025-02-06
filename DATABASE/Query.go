package database

func GetUniquePartner() string {
	query := `
			SELECT
				p.full_name, p.company_name, p.ph_no, p.address, p.reg_date
			FROM
				shop.partner as p
				where p.id = $1;
			`
	return query
}

func GetAllPartner() string {
	query := `
			SELECT
				p.full_name, p.company_name
			FROM
				shop.shop_partner AS sp
			JOIN
				shop.partner AS p ON sp.partner_id = p.id
			WHERE
				sp.shop_id = $1;
			`
	return query
}

func GetPartnerBalance() string {
	query := `
			SELECT
				p.full_name,
				p.company_name,
				b.gold_due,
				b.silver_due,
				b.cash_due
			FROM
				shop.partner AS p
			LEFT JOIN
				shop.balance AS b ON p.id = b.partner_id
			where p.id = $1;
			`
	return query
}
