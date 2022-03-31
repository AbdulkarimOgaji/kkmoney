package api

import (
	"log"
	"net/http"
	"time"

	models "github.com/AbdulkarimOgaji/kkmoney/db"
	"github.com/gin-gonic/gin"
)

func (db *DB) createAcct(c *gin.Context) {
	var acct models.AcctStruct
	err := c.BindJSON(&acct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid request body",
			"payload": nil,
			"error":   err,
		})
		return
	}
	acct.CreatedTime = time.Now().String()[:19]
	acct.AcctNum = generateAcctNumber(acct.UserId)

	stmt, err := db.driver.Prepare(createAcct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid request body",
			"payload": nil,
			"error":   err,
		})
		return
	}
	result, err := stmt.Exec(acct.UserId, acct.AcctType, acct.AcctNum, acct.CreatedTime)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Internal Server Error",
			"error":   err,
			"payload": nil,
		})
		return
	}
	newId, _ := result.LastInsertId()
	acct.AcctId = int(newId)

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "bank account created successfully",
		"payload": gin.H{
			"newAcct": acct,
		},
	})
}

// change this later so that the account number is set according to the account type
func generateAcctNumber(userId int) int {
	return 2000000000 + userId
}
