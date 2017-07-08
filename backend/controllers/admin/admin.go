package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/motopig/gatekeeper/backend/model"
)

type User struct {
	model.BaseModel
	Name     string
	Password string
	Super    int
}

func Home(c *gin.Context) {
	// home sweet home like women`s bubby，but it`s locked up !!

	c.HTML(http.StatusOK, "master.tmpl", gin.H{
		"body": "admin home page",
	})
}
