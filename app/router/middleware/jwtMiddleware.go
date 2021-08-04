package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)
// custom claims
type Claims struct {
	Id string      `json:"id"`
	jwt.StandardClaims
}

// jwt secret key
var jwtSecret = []byte("secret")

// validate JWT
func JwtMiddleware(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "token錯誤",
		})
		c.Abort()
		return
	}
	token := strings.Split(auth, "Bearer ")[1]
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return jwtSecret, nil
	})

	if err != nil {
		var message string
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors & jwt.ValidationErrorExpired != 0 {
				message = "登入已過期, 請重新登入"
			} else {
				message = "Token錯誤"
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": message,
		})
		c.Abort()
		return
	}
	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		c.Set("Id", claims.Id)
		c.Next()
	} else {
		c.Abort()
		return
	}
}