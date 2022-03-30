package api

import (
	"database/sql"

	"github.com/AbdulkarimOgaji/kkmoney/db"
	"github.com/gin-gonic/gin"
)

type DB struct {
	driver *sql.DB
}

func StartServer() {
	driver := db.CreateConnection()

	db.InitializeTables(driver)
	dbService := DB{
		driver: driver,
	}
	r := gin.Default()
	// require_auth := r.Group("/api/v1").Use()
	// {
	// 	require_auth.GET("")
	// }

	r.GET("api/v1/getUsers", dbService.getUsers)
	r.GET("api/v1/getUsers/:user-id", dbService.getUserById)
	r.GET("api/v1/getAccts", dbService.getAccounts)
	r.GET("api/v1/getAccts/:acct-id", dbService.getAcctById)
	r.GET("api/v1/getTxns", dbService.getTxns)
	r.GET("api/v1/getTxns/:txn-id", dbService.getTxnsById)
	r.GET("api/v1/getUserAccts/:user-id", dbService.getUserAccounts)

	r.POST("api/v1/createUser", dbService.createUser)
	r.POST("api/v1/createTxn", dbService.createTxn)
	r.POST("api/v1/createAcct", dbService.createAcct)

	r.PUT("api/v1/editUser", dbService.updateUser)
	r.PUT("api/v1/editAcct", dbService.updateAcct)

	// connect to mysql
	//initialize gin
	// set up routes
}
