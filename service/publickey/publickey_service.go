package publickey

import (
	"github.com/gin-gonic/gin"
	"github.com/natsukirigakure/publickey-repogitory/db"
	"github.com/natsukirigakure/publickey-repogitory/model/entity"
)

type Service struct {
}

type PublicKey entity.PublicKey

func (s Service) CreateModel(c *gin.Context) (PublicKey, error) {
	db := db.GetDB()
	var pk PublicKey

	if err := db.Create(&pk).Error; err != nil {
		return pk, err
	}

	return pk, nil
}

func (s Service) GetAll() ([]PublicKey, error) {
	db := db.GetDB()
	var pk []PublicKey

	if err := db.Find(&pk).Error; err != nil {
		return nil, err
	}

	return pk, nil
}

func (s Service) UpdateById(id string, c *gin.Context) (PublicKey, error) {
	db := db.GetDB()
	var pk PublicKey

	if err := db.Where("id = ?", id).First(&pk).Error; err != nil {
		return pk, err
	}

	if err := c.BindJSON(&pk); err != nil {
		return pk, err
	}

	db.Save(&pk)
	return pk, nil
}

func (s Service) DeleteById(id string) error {
	db := db.GetDB()
	var pk PublicKey

	if err := db.Where("id = ?", id).First(&pk).Error; err != nil {
		return err
	}

	return nil
}
