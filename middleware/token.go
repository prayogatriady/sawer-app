package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/prayogatriady/sawer-app/model"
)

var secret = os.Getenv("JWT_SECRET")

func GenerateToken(user *model.UserEntity) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"expired":  time.Now().Local().Add(1 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// Get user_id via token
func ExtractToken(c *gin.Context) (userId int, err error) {
	authHeader := c.Request.Header.Get("Authorization")
	tokenString := strings.Split(authHeader, " ")[1]
	token, err := ValidateToken(tokenString)
	if err != nil {
		return
	}
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		userId = int(claims["user_id"].(float64))
		return
	}
	return
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func AuthMiddleware(c *gin.Context) {
	// get Bearer from Authorization Header
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "401 - Unauthorized",
			"msg":    "Unauthorized - Missing JWT Token",
		})
		c.Abort()
		return
	}

	// get token from Bearer
	tokenString := strings.Split(authHeader, " ")[1]

	// validate token
	_, err := ValidateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "401 - Unauthorized",
			"msg":    err.Error(),
		})
		c.Abort()
		return
	}
	c.Next()
}
