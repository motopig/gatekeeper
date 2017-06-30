package keeper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"miss": "you"})
}

func Keep(c *gin.Context) {

}
