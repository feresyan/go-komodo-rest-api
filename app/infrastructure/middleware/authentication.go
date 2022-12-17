package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gokomodo-assesment/internal/config"
	"net/http"
	"time"
)

func AuthenticationSellerMiddleware(c *gin.Context) {
	fmt.Println("Authentication middleware")

	// Get the cookie of request
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// Decode and validate token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.GetConfig().EnvSystem.SecretKey.Value), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check the expired
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Check role
		if claims["role"] != "seller" {
			c.String(http.StatusUnauthorized, "Menu ini hanya dapat diakses oleh seller")
			c.Abort()
		}
		// Continue
		c.Next()
	}
	c.AbortWithStatus(http.StatusUnauthorized)
}

func AuthenticationBuyerMiddleware(c *gin.Context) {
	fmt.Println("Authentication middleware")

	// Get the cookie of request
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// Decode and validate token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.GetConfig().EnvSystem.SecretKey.Value), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check the expired
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Check role
		if claims["role"] != "buyer" {
			c.String(http.StatusUnauthorized, "Menu ini hanya dapat diakses oleh buyer")
			c.Abort()
		}
		// Continue
		c.Next()
	}
	c.AbortWithStatus(http.StatusUnauthorized)
}
