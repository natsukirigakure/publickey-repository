package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/natsukirigakure/publickey-repogitory/service/user"
	"net/http"
)

func UserIndex(c *gin.Context) {
	var s user.Service
	p, err := s.GetAll()

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, p)
	}
}

func UserCreate(c *gin.Context) {
	var s user.Service
	p, err := s.CreateModel("", "")

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusCreated, p)
	}
}
