package api

import (
	"net/http"

	"github.com/motopig/gatekeeper/gate/keeper/system"

	"strconv"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"miss": "you"})
}

func Keep(c *gin.Context) {

	service := c.PostForm("service")
	action := c.PostForm("action")

	se, err := system.ParseService(service)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"hi": "huake you!"})
	}
	fullUrl := se.Address + ":" + strconv.Itoa(se.Port) + "/" + action

	// 默认 POST todo GET
	ret, _ := system.Post(fullUrl, c.Request.PostForm)
	c.JSON(http.StatusOK, gin.H(ret))
}
