package entity

type User struct {
	ID       uint   `gorm:"primarykey"`
	username string `form:"username" binding:"required" gorm:"unique; not null"`
	password string `form:"password" binding:"required"`
}
