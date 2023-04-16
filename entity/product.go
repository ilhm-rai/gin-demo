package entity

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string
	UserID      uint `gorm:"not null"`
	User        *User
}
