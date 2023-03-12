package entity

import (
	"github.com/natsukirigakure/publickey-repogitory/utility"
	"gorm.io/gorm"
)

type User struct {
	Abstract
	Username string `form:"username" binding:"required" gorm:"unique; not null"`
	Password string `form:"password" binding:"required"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = utility.NewUlidString()
	return nil
}
