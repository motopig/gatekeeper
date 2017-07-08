package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Blacklist() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(`黑名单`)
	}
}
