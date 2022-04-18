package api

import (
	"context"
	"net/http"
	"strconv"

	db "github.com/AbdulkarimOgaji/kkmoney3/db/sqlc"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetTxn(c *gin.Context) {
	stringId := c.Query("id")
	id, err := strconv.ParseInt(stringId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	txn, err := s.queries.GetTxn(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, txn)
}

func (s *Server) EditTxn(c *gin.Context) {
	var reqBody db.UpdateTxnArgs
	err := c.BindJSON(&reqBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	txn, err := s.queries.UpdateTxnAmount(context.Background(), reqBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, txn)
}

func (s *Server) CreateTxn(c *gin.Context) {
	var reqBody db.CreateTxnArgs
	err := c.BindJSON(&reqBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	txn, err := s.queries.CreateTxn(context.Background(), reqBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, txn)
}

func (s *Server) DeleteTxn(c *gin.Context) {
	stringId := c.Query("id")
	id, err := strconv.ParseInt(stringId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err = s.queries.DeleteTxn(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, "success")
}
