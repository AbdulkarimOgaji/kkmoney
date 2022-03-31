package api

import (
	"database/sql"

	"github.com/AbdulkarimOgaji/kkmoney/api/middleware"
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
	require_auth := r.Group("/api/v1").Use(middleware.AuthorizeClient())
	{
		require_auth.GET("api/v1/getUsers", dbService.getUsers)
		require_auth.GET("api/v1/getUsers/:user-id", dbService.getUserById)
		require_auth.GET("api/v1/getAccts", dbService.getAccounts)
		require_auth.GET("api/v1/getAccts/:acct-id", dbService.getAcctById)
		require_auth.GET("api/v1/getTxns", dbService.getTxns)
		require_auth.GET("api/v1/getTxns/:txn-id", dbService.getTxnsById)
		require_auth.GET("api/v1/getUserAccts/:user-id", dbService.getUserAccounts)
		require_auth.POST("api/v1/createTxn", dbService.createTxn)
		require_auth.POST("api/v1/createAcct", dbService.createAcct)

		require_auth.PUT("api/v1/editUser/:user-id", dbService.updateUser)
		require_auth.PUT("api/v1/editAcct/:acct-id", dbService.updateAcct)

		require_auth.DELETE("api/v1/deleteUser/:user-id", dbService.deleteUser)
		require_auth.DELETE("api/v1/deleteAcct/:acct-id", dbService.deleteAcct)

	}

	r.POST("api/v1/createUser", dbService.createUser)
	r.GET("api/c1/login", dbService.loginHandler)

	r.Run(":8000")
	// connect to mysql
	//initialize gin
	// set up routes
}
