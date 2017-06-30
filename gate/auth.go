package gate

import (
	"github.com/gin-gonic/gin"
)

func AuthRequire(c *gin.Context) {

	// check auth
	//queryStr, err := PosToQuery(c)
	//action := c.PostForm("action")
	//app_id := c.PostForm("app_id")
	//if err != nil {
	//	c.JSON(http.StatusAccepted, gin.H{"msg": err.Error(), "code": 1})
	//}
	//can := true
	//if can == false {
	//	c.JSON(http.StatusForbidden, gin.H{"msg": "no auth to access", "code": 1})
	//	c.Abort()
	//}

	c.Next()
}
