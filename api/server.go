package api

import (
	"database/sql"
	"log"
	"net/http"
	"time"

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
		require_auth.GET("/getUsers", dbService.getUsers)
		require_auth.GET("/getUsers/:user-id", dbService.getUserById)
		require_auth.GET("/getAccts", dbService.getAccounts)
		require_auth.GET("/getAccts/:acct-id", dbService.getAcctById)
		require_auth.GET("/getTxns", dbService.getTxns)
		require_auth.GET("/getTxns/:txn-id", dbService.getTxnsById)
		require_auth.GET("/getUserAccts/:user-id", dbService.getUserAccounts)
		require_auth.POST("/createTxn", dbService.createTxn)
		require_auth.POST("/createAcct", dbService.createAcct)

		require_auth.PUT("/editUser/:user-id", dbService.updateUser)
		require_auth.PUT("/editAcct/:acct-id", dbService.updateAcct)

		require_auth.DELETE("/deleteUser/:user-id", dbService.deleteUser)
		require_auth.DELETE("/deleteAcct/:acct-id", dbService.deleteAcct)

	}

	r.POST("api/v1/createUser", dbService.createUser)
	r.GET("api/v1/login", dbService.loginHandler)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

}
