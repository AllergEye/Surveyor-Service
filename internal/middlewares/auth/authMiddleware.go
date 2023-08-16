package auth_middleware

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrUnexpectedSigningMethod = errors.New("unexpected jwt signing method")
)

type accessTokenClaims struct {
	jwt.RegisteredClaims
	ExpiresAt time.Time `json:"exp"`
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")

		if authorizationHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "no authorization header was provided"})
			c.Abort()
			return
		}

		bearerPrefix := "Bearer "
		if len(authorizationHeader) < len(bearerPrefix) || authorizationHeader[:len(bearerPrefix)] != bearerPrefix {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization format"})
			c.Abort()
			return
		}

		token := authorizationHeader[len(bearerPrefix):]
		if !isValidToken(token) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func isValidToken(tokenString string) bool {
	token, err := jwt.ParseWithClaims(tokenString, &accessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrUnexpectedSigningMethod
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return false
	}

	if _, ok := token.Claims.(*accessTokenClaims); !ok || !token.Valid {
		return false
	}

	return true
}
