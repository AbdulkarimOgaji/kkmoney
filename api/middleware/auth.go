package middleware

import (
	"fmt"
	"log"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const SecretKey = "MyNameIsAbdulkarim"

func AuthorizeClient() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("This is the middleware")
		tokenString := c.GetHeader("access_token")
		if tokenString == "" {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "access token not provided in request header",
				"success": false,
				"error":   "unauthorized",
			})
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
			}
			return secretKey, nil
		})
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "invalid access token",
				"success": false,
				"error":   err,
			})
			return
		}
		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Next()
		} else {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "invalid access token",
				"success": false,
				"error":   err,
			})
		}
	}

}
