package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func CreateConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:grandmaster002@(localhost:3306)/kkmoney")
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}
	return db
}
