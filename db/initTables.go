package db

import (
	"database/sql"
	"log"
)

func InitializeTables(db *sql.DB) {
	// create users table
	stmt, err := db.Prepare(USERS_TABLE)
	if err != nil {
		log.Fatal("statement creation Failed: ", err)
	}
	_, err = stmt.Exec()

	if err != nil {
		log.Fatal("failed to create table Users", err)
	}
	// create accounts table
	stmt, err = db.Prepare(ACCOUNTS_TABLE)
	if err != nil {
		log.Fatal("statement creation Failed: ", err)
	}

	_, err = stmt.Exec()

	if err != nil {
		log.Fatal("failed to create table Accounts", err)
	}

	// create transactions table
	stmt, err = db.Prepare(TRANSACTIONS_TABLE)
	if err != nil {
		log.Fatal("statement creation Failed: ", err)
	}
	_, err = stmt.Exec()

	if err != nil {
		log.Fatal("failed to create table Transactions", err)
	}
	log.Println("All tables created successfully")
}
