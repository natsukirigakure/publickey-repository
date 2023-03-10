package entity

import "golang.org/x/crypto/bcrypt"

type User struct {
	Abstract
	username string `form:"username" binding:"required" gorm:"unique; not null"`
	password string `form:"password" binding:"required"`
}

type UserId struct {
	Identifier Identifier
}

func GenerateUserId() UserId {
	return UserId{GenerateIdentifier()}
}

func NewUserId(id string) UserId {
	return UserId{NewIdentifier(id)}
}

func GeneratePasswordHash(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(string()), 16)
	return string(hashed)
}

func NewUser() {

}
