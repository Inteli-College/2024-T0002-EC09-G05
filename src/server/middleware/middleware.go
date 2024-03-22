package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var SecretKey = []byte("teste")

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString, _ := c.Cookie("token")

		if tokenString == "" {

			tokenString = c.GetHeader("Authorization")

			if tokenString == "" {

				c.JSON(http.StatusUnauthorized, gin.H{"error": "No auth token found in request"})
				c.Abort()
				return
			}
		}

		bearerToken := strings.TrimPrefix(tokenString, "Bearer")

		token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return SecretKey, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := claims["user_id"]
			c.Set("user_id", userID)
			role := claims["role"]
			c.Set("role", role)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		c.Next()
	}
}
