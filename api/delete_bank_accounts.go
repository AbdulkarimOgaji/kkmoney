package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (db *DB) deleteAcct(c *gin.Context) {
	id := c.Param("acct-id")
	stmt, err := db.driver.Prepare(deleteAcct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request",
			"payload": nil,
			"error":   err,
		})
		return
	}
	r, err := stmt.Exec(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to delete account",
			"payload": nil,
			"error":   err,
		})
		return
	}
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
		"message": "account deleted successfully",
		"payload": nil,
	})

}
