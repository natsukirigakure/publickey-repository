package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"publickey-repogitory/service/user"
)

type Controller struct{}

func (pc Controller) Index(c *gin.Context) {
	var s user.Service
	p, err := s.GetAll()

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, p)
	}
}

func (pc Controller) Create(c *gin.Context) {
	var s user.Service
	p, err := s.CreateModel(c)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusCreated, p)
	}
}
