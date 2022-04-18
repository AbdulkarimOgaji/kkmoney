package api

import (
	"context"
	"net/http"
	"strconv"

	db "github.com/AbdulkarimOgaji/kkmoney3/db/sqlc"
	"github.com/gin-gonic/gin"
)

type ListAccountsRequest struct {
	PageSize int `form:"pageNum" binding:"required"`
	PageId   int `form:"pageId" binding:"required"`
}

func (s *Server) ListAccounts(c *gin.Context) {
	var queryParams ListAccountsRequest
	err := c.BindQuery(&queryParams)
	args := db.GetAccountsArgs{
		Limit:  queryParams.PageSize,
		Offset: (queryParams.PageId - 1) * queryParams.PageSize,
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	accounts, err := s.queries.GetAccounts(context.Background(), args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, accounts)
}

func (s *Server) CreateAccount(c *gin.Context) {
	var reqBody db.CreateAccountArgs
	err := c.BindJSON(&reqBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	account, err := s.queries.CreateAccount(context.Background(), reqBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, account)
}

func (s *Server) EditAccount(c *gin.Context) {
	var reqBody db.UpdateAccountArgs
	err := c.BindJSON(&reqBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	account, err := s.queries.UpdateAccount(context.Background(), reqBody)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, account)

}

func (s *Server) DeleteAccount(c *gin.Context) {
	stringId := c.Query("id")
	id, err := strconv.ParseInt(stringId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err = s.queries.DeleteAccount(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, "Success")
}
