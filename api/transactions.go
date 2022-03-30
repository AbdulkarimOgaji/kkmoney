package api

import (
	"net/http"
	"time"

	models "github.com/AbdulkarimOgaji/kkmoney/db"
	"github.com/gin-gonic/gin"
)

func (db *DB) getTxns(c *gin.Context) {
	var txns []models.TxnStruct

	sql := `
		SELECT * FROM transactions
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
	var tmp models.TxnStruct
	for rows.Next() {
		rows.Scan(&tmp.TxnId, &tmp.SenderId, &tmp.ReceiverId, &tmp.Amount, &tmp.TxnTime)
		txns = append(txns, tmp)
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "fetched transactions successfully",
		"payload": gin.H{
			"transactions": txns,
		},
	})
}

func (db *DB) getTxnsById(c *gin.Context) {
	id := c.Param("txn-id")
	sql := `
		SELECT * FROM transactions
		WHERE txnId = ?
	`
	var txn models.TxnStruct
	err := db.driver.QueryRow(sql, id).Scan(
		&txn.TxnId,
		&txn.SenderId,
		&txn.ReceiverId,
		&txn.Amount,
		&txn.TxnTime,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to fetch transaction",
			"payload": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "fetched transaction successfully",
		"payload": gin.H{
			"transaction": txn,
		},
	})

}

func (db *DB) createTxn(c *gin.Context) {
	var txn models.TxnStruct
	err := c.BindJSON(&txn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid request body",
			"payload": nil,
		})
		return
	}
	sql := `
	INSERT INTO transactions values(?, ?, ?, ?)
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

	// probably need to change this:
	txn.TxnTime = time.Now().String()
	result, err := stmt.Exec(txn.SenderId, txn.ReceiverId, txn.Amount, txn.TxnTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Internal Server Error",
			"error":   err,
		})
		return
	}
	newId, _ := result.LastInsertId()
	txn.TxnId = int(newId)

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "user created successfully",
		"payload": gin.H{
			"transaction": txn,
		},
	})
}
