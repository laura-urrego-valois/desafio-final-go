package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {

		verb := c.Request.Method
		time1 := time.Now()
		path := c.Request.URL.Path

		c.Next()

		var size int
		if c.Writer != nil {
			size = c.Writer.Size()
		}

		time2 := time.Since(time1)

		fmt.Printf("\nverb: %v\ntime: %v\npath: %v\nsize: %v\ntime tx: %v\n", verb, time1, path, size, time2)
	}
}
