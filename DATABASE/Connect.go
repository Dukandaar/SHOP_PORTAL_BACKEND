package database

import (
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	logPrefix := ("[" + time.Now().Format("2006-01-02 15:04:05") + "] ")
	driverName := "postgres"
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	utils.Logger.Info(logPrefix, user, password, host, port, dbname)
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", user, password, host, port, dbname)

	DB, err = sql.Open(driverName, dsn)
	if err != nil {
		utils.Logger.Error(logPrefix + "Connection unsuccessful..!!!")
		return DB
	}
	pingErr := DB.Ping()
	if pingErr != nil {
		utils.Logger.Error(logPrefix + pingErr.Error())
		return DB
	}

	DB.SetMaxOpenConns(10) // Maximum number of open connections
	DB.SetMaxIdleConns(5)  // Maximum number of connections in the idle connection pool

	utils.Logger.Info(logPrefix + "Connection Successful..!!!")
	return DB
}
