package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() *sql.DB {
	driverName := "postgres"
	dsn := "user=postgres password=Post321 host=127.0.0.1 port=5432 dbname=shop sslmode=disable"
	DB, err := sql.Open(driverName, dsn)
	if err != nil {
		log.Println("Connection unsucessfull!")
		return DB
	}
	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	log.Println("Connection Sucessfull!😍")
	return DB
}
