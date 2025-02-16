package database

import (
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() *sql.DB {
	logPrefix := ("[" + time.Now().Format("2006-01-02 15:04:05") + "] ")
	driverName := "postgres"
	dsn := "user=postgres password=Post321 host=127.0.0.1 port=5432 dbname=postgres sslmode=disable"
	DB, err := sql.Open(driverName, dsn)
	if err != nil {
		utils.Logger.Error(logPrefix + "Connection unsuccessful..!!!")
		return DB
	}
	pingErr := DB.Ping()
	if pingErr != nil {
		utils.Logger.Error(logPrefix + pingErr.Error())
		return DB
	}
	utils.Logger.Info(logPrefix + "Connection Successful..!!!")
	return DB
}
