package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/motopig/gatekeeper/backend/common"
	"github.com/motopig/gatekeeper/backend/controllers/admin"
	"github.com/motopig/gatekeeper/backend/middlewares"
)

func Routers() {

	gin.SetMode("debug")
	r := gin.New()

	r.Static("/resources", "./backend/resources")
	r.StaticFile("/favicon.ico", "./favicon.ico")

	r.LoadHTMLGlob("backend/templates/*/*")

	sessionConn := common.Hconfig.String("session::domain") + ":" + common.Hconfig.String("session::port")
	sessionSecret := common.Hconfig.String("session::secret")
	sessionDatabase := common.Hconfig.String("session::database")
	store, _ := sessions.NewRedisStoreWithDB(10, "tcp", sessionConn, sessionSecret, sessionDatabase, []byte(common.Hconfig.String("session::keysecret")))

	r.Use(sessions.Sessions("hodor", store))

	loginGroup := r.Group("/admin")
	{
		loginGroup.GET("/login", admin.Login)
		loginGroup.POST("/dologin", admin.DoLogin)
	}

	authGroup := r.Group("/admin").Use(admin.AuthRequire, middlewares.Dododo("huake"))
	{
		authGroup.GET("/home", admin.Home)
		authGroup.GET("/service", admin.Services)
		authGroup.GET("/health/:service/:datacenter", admin.Health)
		authGroup.GET("/nodes", admin.Nodes)
		authGroup.GET("/detail/:node/:datacenter", admin.Dnode)
		authGroup.GET("/kvs", admin.Kvs)
		authGroup.GET("/keys/*key", admin.Keys)
		authGroup.PUT("/keys/*key", admin.Put)
		authGroup.DELETE("/keys/*key", admin.Del)
		authGroup.GET("/dc", admin.DataCenter)
	}

	r.Run(":9002")
}
