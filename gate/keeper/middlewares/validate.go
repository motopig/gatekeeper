package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		if "123" != "123" {
			c.JSON(http.StatusForbidden, gin.H{"否逼得": "不能访问思密达"})
			c.Abort()
			return
		}
	}
}
