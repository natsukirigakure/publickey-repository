package user

import (
	"github.com/gin-gonic/gin"
	"github.com/natsukirigakure/publickey-repogitory/db"
	"github.com/natsukirigakure/publickey-repogitory/model/entity"
	"github.com/natsukirigakure/publickey-repogitory/utility"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
}

type User entity.User

func (s Service) CreateModel(username string, password string) (User, error) {
	db := db.GetDB()

	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	u := User{
		Abstract: entity.Abstract{ID: utility.NewUlidString()},
		Username: username,
		Password: string(hashed),
	}

	if err := db.Create(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

func (s Service) GetAll() ([]User, error) {
	db := db.GetDB()
	var u []User

	if err := db.Find(&u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func (s Service) UpdateById(id string, c *gin.Context) (User, error) {
	db := db.GetDB()
	var u User

	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	db.Save(&u)
	return u, nil
}

func (s Service) DeleteById(id string) error {
	db := db.GetDB()
	var u User

	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return err
	}

	return nil
}
