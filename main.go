package main

import (
	"database/sql"
	"log"

	"github.com/AbdulkarimOgaji/kkmoney3/api"
	db "github.com/AbdulkarimOgaji/kkmoney3/db/sqlc"
	"github.com/AbdulkarimOgaji/kkmoney3/util"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang/mock/mockgen/model"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Failed to load config files")
	}
	client, err := sql.Open("mysql", config.DbconnectionUri)
	if err != nil {
		log.Fatal("Failed to create client. Error: ", err)
	}
	defer client.Close()
	if err != nil {
		log.Fatal("Failed to connect")
	}

	// queries := db.NewStore(events, schedules, attendees)
	store := db.NewStore(client)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddr)
	if err != nil {
		log.Fatal("Failed to start Server")
	}
}
