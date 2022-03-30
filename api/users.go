package api

import (
	"net/http"

	models "github.com/AbdulkarimOgaji/kkmoney/db"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (db *DB) getUsers(c *gin.Context) {
	var users []models.UserStruct

	sql := `
		SELECT (userId, firstName, lastName, otherName, email, phoneNum, otherNum, gender, address, kinName, kinNumber, kinRelationship) FROM users
	`
	rows, err := db.driver.Query(sql)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "failed to fetch users",
			"payload": nil,
		})
		return
	}
	var tmp models.UserStruct
	for rows.Next() {
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
			&tmp.KinRelationship)
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
	sql := `
		SELECT (userId, firstName, lastName, otherName, email, phoneNum, otherNum, gender, address, kinName, kinNumber, kinRelationship) FROM users
		WHERE userId = ?
	`
	var user models.UserStruct
	err := db.driver.QueryRow(sql, id).Scan(&user.UserId,
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
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to fetch user",
			"payload": nil,
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

func (db *DB) createUser(c *gin.Context) {
	var user models.UserStruct
	// for now leave it as bindjson.. later you will have to do c.Reques.FormValue("firstName") and
	// so on to get all the values you need. After that you will use bycrypt to create passwordhash then
	// store that in the db
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid request body",
			"payload": nil,
		})
		return
	}
	sql := `
	INSERT INTO users(firstName, lastName, otherName, email, phoneNum, otherNum, gender, address, kinName, kinNumber, kinRelationship, passwordHash) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	// generate passwordHash
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	stmt, _ := db.driver.Prepare(sql)
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

func (db *DB) updateUser(c *gin.Context) {
	id := c.Param("user-id")
	var user models.UserStruct
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid request body",
			"payload": nil,
		})
		return
	}
	sql := `
		UPDATE users
		SET firstName = ?
		SET lastName = ?
		SET otherName = ?
		SET email = ?
		SET phoneNum = ?
		SET otherNum = ?
		SET gender = ?
		SET address = ?
		SET kinName = ?
		SET kinNumber = ?
		SET kinRelationship = ?
		WHERE userId = ?
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
	r, err := stmt.Exec(
		sql,
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
			"message": "Failed to update user",
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

func (db *DB) deleteUser(c *gin.Context) {
	id := c.Param("user-id")
	sql := `
		DELETE users
		WHERE userId = ?
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
