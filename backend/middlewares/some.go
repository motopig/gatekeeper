package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Dododo(s string) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(`DoDoDo`)
		return
	}
}
