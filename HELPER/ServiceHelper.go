package helper

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"database/sql"
)

func GetOwnerId(ownerRegId string, tx *sql.Tx) (int, error) {
	ServiceQuery := database.GetOwnerRowId()
	var ownerRowId int
	err := tx.QueryRow(ServiceQuery, ownerRegId).Scan(&ownerRowId)
	return ownerRowId, err
}

func GetCustomerId(customerRegId string, ownerRowId int, tx *sql.Tx) (int, error) {
	ServiceQuery := database.GetCustomerId()
	var customerId int
	err := tx.QueryRow(ServiceQuery, customerRegId, ownerRowId).Scan(&customerId)
	return customerId, err
}

func CheckIfCustomerBelongsToOwner(customerId int, ownerRowId int, tx *sql.Tx) (bool, error) {
	ServiceQuery := database.CheckIfCustomerBelongsToOwner()
	var isActive string
	err := tx.QueryRow(ServiceQuery, customerId, ownerRowId).Scan(&isActive)
	return isActive == utils.TRUE, err
}
