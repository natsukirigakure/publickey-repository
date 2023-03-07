package controller

import (
	"github.com/gin-gonic/gin"
	"publickey-repogitory/controller/user"
)

func Router() *gin.Engine {
	r := gin.Default()
	userController := user.Controller{}
	r.GET("/ping", pingController)
	r.GET("/user", userController.Index)
	r.POST("/user", userController.Create)
	return r
}
