package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	m "github.com/AbdulkarimOgaji/kkmoney/api/middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type TokenResponse struct {
	Token   string
	Success bool
}

type LoginDetails struct {
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}

func (db *DB) loginHandler(c *gin.Context) {
	var payload LoginDetails
	err := c.BindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request body",
			"error":   err,
		})
	}
	var passwordDb string
	// check databse if username and password exists. if so return success and token

	err = db.driver.QueryRow("SELECT passwordHash FROM users WHERE email = ?", payload.Email).Scan(&passwordDb)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": fmt.Sprintf("User with email %v does not exist in the database", payload.Email),
			"success": false,
			"error":   err,
		})
		return
	}
	// validate user
	if err := bcrypt.CompareHashAndPassword([]byte(passwordDb), []byte(payload.Password)); err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "username and password does not match",
			"success": false,
			"error":   err,
		})
		return
	}
	// generate token
	claims := jwt.MapClaims{
		"email":     payload.Email,
		"ExpiresAt": 15000,
		"IssuedAt":  time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(m.SecretKey)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to signed new token",
			"error":   err,
		})
		return
	}
	resp := TokenResponse{
		Token:   tokenString,
		Success: true,
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "login successful",
		"success": true,
		"payload": gin.H{
			"tokenResponse": resp,
		},
	})
}
