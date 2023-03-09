package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"publickey-repogitory/controller/auth"
	"publickey-repogitory/controller/common"
	"publickey-repogitory/controller/user"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("template/*")
	r.Static("/static", "static")

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("session", store))

	publicGroup := r.Group("/")
	{
		publicGroup.GET("/ping", common.Ping)
	}

	loggedInGroup := r.Group("/")
	{
		loggedInGroup.GET("/user", user.UserIndex)
		loggedInGroup.POST("/user", user.UserCreate)
		loggedInGroup.GET("/logout", auth.Logout)
	}

	notLoggedInGroup := r.Group("/")
	{
		notLoggedInGroup.GET("/login", auth.LoginPage)
		notLoggedInGroup.POST("/login", auth.Login)
	}

	return r
}
