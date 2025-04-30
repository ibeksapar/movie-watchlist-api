package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			return
		}

		client := resty.New()
		resp, err := client.R().
			SetHeader("Authorization", authHeader).
			Get("http://localhost:8081/validate")

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Error calling validate token", "details": err.Error()})
			return
		}

		if resp.StatusCode() != http.StatusOK {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token", "details": resp.Status()})
            return
        }

		c.Next()
	}
}
