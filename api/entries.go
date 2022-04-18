package api

import (
	"context"
	"net/http"
	"strconv"

	db "github.com/AbdulkarimOgaji/kkmoney3/db/sqlc"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetEntry(c *gin.Context) {
	stringId := c.Query("id")
	id, err := strconv.ParseInt(stringId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	entry, err := s.queries.GetEntry(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, entry)
}

func (s *Server) EditEntry(c *gin.Context) {
	var reqBody db.UpdateEntryArgs
	err := c.BindJSON(&reqBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	entry, err := s.queries.UpdateEntryAmount(context.Background(), reqBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, entry)
}

func (s *Server) CreateEntry(c *gin.Context) {
	var reqBody db.CreateEntryArgs
	err := c.BindJSON(&reqBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	entry, err := s.queries.CreateEntry(context.Background(), reqBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, entry)
}

func (s *Server) DeleteEntry(c *gin.Context) {
	stringId := c.Query("id")
	id, err := strconv.ParseInt(stringId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err = s.queries.DeleteEntry(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, "success")
}

func (s *Server) ListEntries(c *gin.Context) {
	var queryParams ListAccountsRequest
	err := c.BindQuery(&queryParams)
	args := db.GetEntriesArgs{
		Limit:  queryParams.PageSize,
		Offset: (queryParams.PageId - 1) * queryParams.PageSize,
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	entries, err := s.queries.GetEntries(context.Background(), args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, entries)
}
