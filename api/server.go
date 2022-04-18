package api

import (
	db "github.com/AbdulkarimOgaji/kkmoney3/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	queries db.Store
	router  gin.Engine
}

func NewServer(q db.Store) *Server {
	server := &Server{queries: q}
	router := gin.Default()

	router.GET("/accounts", server.ListAccounts)
	router.GET("/entries", server.ListEntries)
	router.GET("/entry", server.GetEntry)
	router.GET("/txns", server.GetTxn)

	router.POST("/createAccount", server.CreateAccount)
	router.POST("/createEntry", server.CreateEntry)
	router.POST("/createTxn", server.CreateTxn)

	router.PUT("/editAccount", server.EditAccount)
	router.PUT("/editEntry", server.EditEntry)
	router.PUT("/editTxn", server.EditTxn)

	router.DELETE("/deleteAccount", server.DeleteAccount)
	router.DELETE("/deleteEntry", server.DeleteEntry)
	router.DELETE("/deleteTxn", server.DeleteTxn)

	// Add routes Here
	server.router = *router
	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
