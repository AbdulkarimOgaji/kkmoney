package api

import (
	"log"
	"net/http"

	models "github.com/AbdulkarimOgaji/kkmoney/db"
	"github.com/gin-gonic/gin"
)

func (db *DB) getAccounts(c *gin.Context) {
	var accts []models.AcctStruct
	rows, err := db.driver.Query(getAccts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "failed to fetch accounts",
			"payload": nil,
			"error":   err,
		})
		return
	}
	var tmp models.AcctStruct
	for rows.Next() {
		rows.Scan(&tmp.AcctId, &tmp.UserId, &tmp.CurrentBal, &tmp.AcctType, &tmp.AcctNum, &tmp.CreatedTime)
		accts = append(accts, tmp)
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "fetched accounts successfully",
		"payload": gin.H{
			"accounts": accts,
		},
	})
}

func (db *DB) getUserAccounts(c *gin.Context) {
	var accts []models.AcctStruct
	id := c.Param("user-id")
	rows, err := db.driver.Query(getUserAccts, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "failed to fetch accounts",
			"payload": nil,
			"error":   err,
		})
		return
	}
	var tmp models.AcctStruct
	for rows.Next() {
		rows.Scan(&tmp.AcctId, &tmp.UserId, &tmp.CurrentBal, &tmp.AcctType, &tmp.AcctNum, &tmp.CreatedTime)
		accts = append(accts, tmp)
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "fetched user's accounts successfully",
		"payload": gin.H{
			"accounts": accts,
		},
	})
}

func (db *DB) getAcctById(c *gin.Context) {
	id := c.Param("acct-id")
	var acct models.AcctStruct
	err := db.driver.QueryRow(getAcctById, id).Scan(
		&acct.AcctId,
		&acct.UserId,
		&acct.CurrentBal,
		&acct.AcctType,
		&acct.AcctNum,
		&acct.CreatedTime,
	)
	if err != nil {
		log.Println("queryrow error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to fetch account",
			"payload": nil,
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "fetched account successfully",
		"payload": gin.H{
			"account": acct,
		},
	})

}
