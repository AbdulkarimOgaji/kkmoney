package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/AbdulkarimOgaji/kkmoney3/util"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbDriver        = "mysql"
	dbConnectionUri = "root:grandmaster002@tcp(localhost:3307)/kkmoney?parseTime=true"
)

var testingQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	config, err := util.LoadConfig("../../")
	testDB, err = sql.Open("mysql", config.DbconnectionUri)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	testingQueries = New(testDB)
	os.Exit(m.Run())
}
