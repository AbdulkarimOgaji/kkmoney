package api

import (
	"log"
	"net/http"

	models "github.com/AbdulkarimOgaji/kkmoney/db"
	"github.com/gin-gonic/gin"
)

func (db *DB) getUsers(c *gin.Context) {
	var users []models.UserStruct
	rows, err := db.driver.Query(getUsers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "failed to fetch users",
			"payload": nil,
			"error":   err,
		})
		return
	}

	for rows.Next() {
		var tmp models.UserStruct
		rows.Scan(&tmp.UserId,
			&tmp.FirstName,
			&tmp.LastName,
			&tmp.OtherName,
			&tmp.Email,
			&tmp.PhoneNum,
			&tmp.OtherNum,
			&tmp.Gender,
			&tmp.Address,
			&tmp.KinName,
			&tmp.KinNumber,
			&tmp.KinRelationship,
			&tmp.CreatedTime,
		)
		users = append(users, tmp)
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "fetched users successfully",
		"payload": gin.H{
			"users": users,
		},
	})
}

func (db *DB) getUserById(c *gin.Context) {
	id := c.Param("user-id")
	var user models.UserStruct
	err := db.driver.QueryRow(getUserById, id).Scan(&user.UserId,
		&user.FirstName,
		&user.LastName,
		&user.OtherName,
		&user.Email,
		&user.PhoneNum,
		&user.OtherNum,
		&user.Gender,
		&user.Address,
		&user.KinName,
		&user.KinNumber,
		&user.KinRelationship,
		&user.CreatedTime,
	)
	if err != nil {
		log.Println("The nasty error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to fetch user",
			"payload": nil,
			"error":   err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "fetched user successfully",
		"payload": gin.H{
			"user": user,
		},
	})

}
