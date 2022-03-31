package api

import (
	"log"
	"net/http"
	"time"

	models "github.com/AbdulkarimOgaji/kkmoney/db"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (db *DB) createUser(c *gin.Context) {
	var user models.UserStruct
	// for now leave it as bindjson.. later you will have to do c.Reques.FormValue("firstName") and
	// so on to get all the values you need. After that you will use bycrypt to create passwordhash then
	// store that in the db
	err := c.BindJSON(&user)
	user.CreatedTime = time.Now().String()[:19]
	if err != nil {
		log.Println(err)
		log.Println(user)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid request body",
			"payload": nil,
		})
		return
	}
	// generate passwordHash
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to generate password Hash",
			"error":   err,
		})
		return
	}
	stmt, err := db.driver.Prepare(createUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "sql statement error",
			"payload": nil,
			"error":   err,
		})
	}
	result, err := stmt.Exec(
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
		user.CreatedTime,
		passwordHash,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Internal Server Error",
			"error":   err,
		})
		return
	}
	newId, _ := result.LastInsertId()
	user.UserId = int(newId)
	// remove passwordHash from response
	user.PasswordHash = ""
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "user created successfully",
		"payload": gin.H{
			"newUser": user,
		},
	})
}
