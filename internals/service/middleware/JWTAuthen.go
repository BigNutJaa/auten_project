package middleware

import (
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// replace with token check
		//t := time.Now()

		//set example variable
		c.Set("example", "12345")

		// before request

		c.Next()
	}
}
