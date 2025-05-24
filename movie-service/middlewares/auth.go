package middlewares

import (
	"net/http"
	"os"

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
        authURL := os.Getenv("AUTH_SERVICE_URL")

        if authURL == "" {
            authURL = "http://localhost:8081"
        }

        resp, err := client.R().
            SetHeader("Authorization", authHeader).
            Get(authURL + "/validate")

        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "error":   "Error calling validate token",
                "details": err.Error(),
            })
            return
        }

        if resp.StatusCode() != http.StatusOK {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "error":   "Invalid token",
                "details": resp.Status(),
            })
            return
        }

        c.Next()
    }
}
