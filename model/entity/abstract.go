package entity

import (
	"gorm.io/gorm"
	"time"
)

type Abstract struct {
	ID        string `gorm:"primarykey; size:26; not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
