package database

import (
	config "SHOP_PORTAL_BACKEND/CONFIG"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"database/sql"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	driverName := config.DRIVER_NAME
	user := config.DB_USER
	password := config.DB_PASSWORD
	host := config.DB_HOST
	port := config.DB_PORT
	dbname := config.DB_NAME

	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", user, password, host, port, dbname)

	DB, err = sql.Open(driverName, dsn)
	if err != nil {
		utils.Logger.Error("Connection unsuccessful..!!!")
		return DB
	}
	pingErr := DB.Ping()
	if pingErr != nil {
		utils.Logger.Error(pingErr.Error())
		return DB
	}

	DB.SetMaxOpenConns(10) // Maximum number of open connections
	DB.SetMaxIdleConns(5)  // Maximum number of connections in the idle connection pool

	utils.Logger.Info("Connection Successful..!!!")
	return DB
}
