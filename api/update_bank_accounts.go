package api

import (
	"log"
	"net/http"

	models "github.com/AbdulkarimOgaji/kkmoney/db"
	"github.com/gin-gonic/gin"
)

func (db *DB) updateAcct(c *gin.Context) {
	id := c.Param("acct-id")
	var acct models.AcctStruct
	err := c.BindJSON(&acct)
	if err != nil {
		log.Println("binding error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid request body",
			"payload": nil,
			"error":   err,
		})
		return
	}
	stmt, err := db.driver.Prepare(updateAcct)
	if err != nil {
		log.Println("sql error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid request body",
			"payload": nil,
			"error":   err,
		})
		return
	}
	r, err := stmt.Exec(acct.AcctType, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to update account",
			"payload": nil,
			"error":   err,
		})
		return
	}
	// return error if no row affected
	if n, _ := r.RowsAffected(); n < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Request caused no alterations to the database",
			"payload": nil,
			"error":   "Request caused no alterations",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "updated account successfully",
		"payload": gin.H{
			"acct": acct,
		},
	})
}
