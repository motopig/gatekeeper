package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthRequire() gin.HandlerFunc {

	// check auth
	return func(c *gin.Context) {
		//queryStr, err := util.PosToQuery(c)
		//action := c.PostForm("action")
		//app_id := c.PostForm("app_id")
		//if err != nil {
		//	c.JSON(http.StatusAccepted, gin.H{"msg": err.Error(), "code": 1})
		//	c.Abort()
		//	return
		//}
		can := true // todo check auth
		if can == true {
			c.JSON(http.StatusForbidden, gin.H{"msg": "no auth to access", "code": 1})
			c.Abort()
			return
		}
	}
}
