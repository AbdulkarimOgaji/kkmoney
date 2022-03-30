package api

import (
	"net/http"

	models "github.com/AbdulkarimOgaji/kkmoney/db"
	"github.com/gin-gonic/gin"
)

func (db *DB) getAccounts(c *gin.Context) {
	var accts []models.AcctStruct

	sql := `
		SELECT * FROM accounts
	`
	rows, err := db.driver.Query(sql)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "failed to fetch accounts",
			"payload": nil,
		})
		return
	}
	var tmp models.AcctStruct
	for rows.Next() {
		rows.Scan(&tmp.AcctId, &tmp.UserId, &tmp.CurrentBal, &tmp.AcctType, &tmp.AcctNum)
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
	sql := `
		SELECT * FROM accounts
		WHERE userId = ?
	`
	rows, err := db.driver.Query(sql, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "failed to fetch accounts",
			"payload": nil,
		})
		return
	}
	var tmp models.AcctStruct
	for rows.Next() {
		rows.Scan(&tmp.AcctId, &tmp.UserId, &tmp.CurrentBal, &tmp.AcctType, &tmp.AcctNum)
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
	sql := `
		SELECT * FROM users
		WHERE acctId = ?
	`
	var acct models.AcctStruct
	err := db.driver.QueryRow(sql, id).Scan(
		&acct.AcctId,
		&acct.UserId,
		&acct.CurrentBal,
		&acct.AcctType,
		&acct.AcctNum,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to fetch account",
			"payload": nil,
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

func (db *DB) createAcct(c *gin.Context) {
	var acct models.AcctStruct
	err := c.BindJSON(&acct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid request body",
			"payload": nil,
		})
		return
	}
	acct.AcctNum = generateAcctNumber(acct.UserId)
	sql := `
	INSERT INTO accounts(userId, acctType, acctNum) values(?, ?, ?)
	`

	stmt, err := db.driver.Prepare(sql)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid request body",
			"payload": nil,
			"error":   err,
		})
		return
	}
	result, err := stmt.Exec(acct.UserId, acct.AcctType, acct.AcctNum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Internal Server Error",
			"error":   err,
		})
		return
	}
	newId, _ := result.LastInsertId()
	acct.AcctId = int(newId)

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "user created successfully",
		"payload": gin.H{
			"newAcct": acct,
		},
	})
}

func (db *DB) updateAcct(c *gin.Context) {
	id := c.Param("acct-id")
	var acct models.AcctStruct
	err := c.BindJSON(&acct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid request body",
			"payload": nil,
		})
		return
	}
	sql := `
		UPDATE accounts
		SET acctType = ?
		WHERE acctId = ?
	`
	stmt, err := db.driver.Prepare(sql)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid request body",
			"payload": nil,
		})
		return
	}
	r, err := stmt.Exec(acct.AcctType, id)

	// return error if no row affected
	if n, _ := r.RowsAffected(); n < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Request caused no alterations to the database",
			"payload": nil,
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to update account",
			"payload": nil,
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

func (db *DB) deleteAcct(c *gin.Context) {
	id := c.Param("acct-id")
	sql := `
		DELETE accounts
		WHERE acctId = ?
	`
	stmt, err := db.driver.Prepare(sql)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request",
			"payload": nil,
		})
		return
	}
	r, err := stmt.Exec(id)
	if n, _ := r.RowsAffected(); n < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Request caused no alterations to the database",
			"payload": nil,
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to delete account",
			"payload": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "account deleted successfully",
		"payload": nil,
	})

}

// change this later so that the account number is set according to the account type
func generateAcctNumber(userId int) int {
	return 2000000000 + userId
}
