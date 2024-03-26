package middleware

import (
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			c.JSON(http.StatusUnauthorized, errors.New("token not found"))
			c.Abort()
			return
		}
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, errors.New("invalid token"))
			c.Abort()
			return
		}
		c.Next()
	}
}
