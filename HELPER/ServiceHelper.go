package helper

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	"database/sql"
)

func GetOwnerId(ownerRegId string, tx *sql.Tx) (int, error) {
	ServiceQuery := database.GetOwnerRowId()
	var ownerRowId int
	err := tx.QueryRow(ServiceQuery, ownerRegId).Scan(&ownerRowId)
	return ownerRowId, err
}

func GetCustomerId(customerRegId string, tx *sql.Tx) (int, error) {
	ServiceQuery := database.GetCustomerId()
	var customerId int
	err := tx.QueryRow(ServiceQuery, customerRegId).Scan(&customerId)
	return customerId, err
}
