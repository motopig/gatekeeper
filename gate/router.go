package gate

import (
	"github.com/gin-gonic/gin"
	"github.com/motopig/gatekeeper/gate/keeper"
)

func Routers() {

	gin.SetMode(iniconf.String("runtime::mode"))
	r := gin.New()

	authGroup := r.Group("/api").Use(AuthRequire)
	{
		authGroup.POST("/home", keeper.Home)
		authGroup.GET("/dc", keeper.Keep)
	}

	r.Run(":" + iniconf.String("runtime::port"))
}
