package entity

type User struct {
	Username string `form:"username" binding:"required" gorm:"unique; not null"`
	Password string `form:"password" binding:"required"`
}
