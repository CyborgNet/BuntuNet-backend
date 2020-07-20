package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
)

var router *gin.Engine

func main() {
	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()
	router.Use(static.Serve("/assets/css", static.LocalFile("./templates/assets/css", true)))
	router.Use(static.Serve("/assets/img", static.LocalFile("./templates/assets/img", true)))
	router.Use(static.Serve("/assets/js", static.LocalFile("./templates/assets/js", true)))
	router.Use(static.Serve("/panel/assets/css", static.LocalFile("./templates/assets/css", true)))
	router.Use(static.Serve("/panel/assets/img", static.LocalFile("./templates/assets/img", true)))
	router.Use(static.Serve("/panel/assets/js", static.LocalFile("./templates/assets/js", true)))
	router.LoadHTMLGlob("templates/*.html")
	store := cookie.NewStore([]byte("cmzjhobeielszohqnkethavecwxmyzuz"))
	router.Use(sessions.Sessions("session", store))
	initializeRoutes()
	router.Run("localhost:3716")

}

func render(c *gin.Context, data gin.H, templateName string) {
	session := sessions.Default(c)
	username := session.Get("username")
	password := session.Get("password")
	if !(username == nil && password == nil) {
		valid := isUserValid(username.(string), password.(string))
		vip := isUserVIP(username.(string), password.(string))
		data["is_logged_in"] = valid
		data["is_vip"] = vip
	} else {
		data["is_logged_in"] = false
		data["is_vip"] = false
	}
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}