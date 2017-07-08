package keeper

import (
	"github.com/gin-gonic/gin"
	"github.com/motopig/gatekeeper/gate/keeper/api"
	"github.com/motopig/gatekeeper/gate/keeper/middlewares"
	"github.com/motopig/gatekeeper/gate/keeper/system"
)

func Routers() {

	gin.SetMode(system.GetConfig("runtime::mode"))
	r := gin.New()

	authGroup := r.Group("/api").Use(
		middlewares.AuthRequire(),
		middlewares.Blacklist(),
		middlewares.Validate())
	{
		authGroup.POST("/home", api.Home)
		authGroup.POST("/dc", api.Keep)
	}

	r.Run(":" + system.GetConfig("runtime::port"))
}
