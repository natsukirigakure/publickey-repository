package auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "login",
	})
}

func Login(c *gin.Context) {
	s := sessions.Default(c)
	s.Set("loginUser", c.PostForm("userId"))
	s.Save()
	c.String(http.StatusOK, "ログインしました")
}

func Logout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.String(http.StatusOK, "ログアウトしました")
}
