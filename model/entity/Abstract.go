package entity

import (
	"gorm.io/gorm"
	"time"
)

type Abstract struct {
	ID        Identifier `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
