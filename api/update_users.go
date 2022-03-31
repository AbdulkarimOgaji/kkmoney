package api

import (
	"log"
	"net/http"

	models "github.com/AbdulkarimOgaji/kkmoney/db"
	"github.com/gin-gonic/gin"
)

func (db *DB) updateUser(c *gin.Context) {
	id := c.Param("user-id")
	var user models.UserStruct
	err := c.BindJSON(&user)
	if err != nil {
		log.Println("Bind json error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid request body",
			"payload": nil,
			"error":   err,
		})
		return

	}
	stmt, err := db.driver.Prepare(updateUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid request body",
			"payload": nil,
		})
		return
	}
	r, err := stmt.Exec(
		user.FirstName,
		user.LastName,
		user.OtherName,
		user.Email,
		user.PhoneNum,
		user.OtherNum,
		user.Gender,
		user.Address,
		user.KinName,
		user.KinNumber,
		user.KinRelationship,
		id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to update user",
			"payload": nil,
			"error":   err,
		})
		return
	}
	n, _ := r.RowsAffected()
	if n < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Request caused no alterations to the database",
			"payload": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "updated user successfully",
		"payload": gin.H{
			"user": user,
		},
	})
}
