package entity

import (
	"gorm.io/gorm"
	"time"
)

type Abstract struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
