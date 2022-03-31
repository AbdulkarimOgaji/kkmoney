package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (db *DB) deleteUser(c *gin.Context) {
	id := c.Param("user-id")
	stmt, err := db.driver.Prepare(deleteUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid request body",
			"payload": nil,
		})
		return
	}
	_, err = stmt.Exec(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to delete user",
			"payload": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "user deleted successfully",
		"payload": nil,
	})

}
