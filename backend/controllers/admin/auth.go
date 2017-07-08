package admin

import (
	"encoding/base64"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/motopig/gatekeeper/backend/common"
	"github.com/motopig/gatekeeper/backend/model"
	"github.com/motopig/gatekeeper/gate/keeper/util"
)

var session sessions.Session

func Login(c *gin.Context) {

	c.HTML(http.StatusOK, "login.tmpl", gin.H{
		"body": "admin website",
	})
}

func DoLogin(c *gin.Context) {

	name := c.PostForm("name")
	password := c.PostForm("password")

	// check user auth
	user := new(User)
	model.Gorm().Where(&User{Name: name, Password: util.Mdfive(password)}).First(user)
	if user.Name == "" {
		c.String(http.StatusForbidden, "Hello Sucker %s %s", "FFFuck", "OOOf")
		c.Abort()
		return
	}

	session = sessions.Default(c)
	session.Options(sessions.Options{
		Path:     "/",
		Domain:   common.Hconfig.String("session::domain"),
		MaxAge:   7200,
		Secure:   false,
		HttpOnly: true,
	})
	session.Set("uname", base64.StdEncoding.EncodeToString([]byte(name)))
	session.Save()
	c.Redirect(http.StatusMovedPermanently, "/admin/home")
}

func AuthRequire(c *gin.Context) {
	session = sessions.Default(c)
	u := session.Get("uname")

	if u == nil {
		c.Redirect(http.StatusMovedPermanently, "/admin/login")
	}
	c.Next()
}
