package model

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID uint       `gorm:"not null"` // relasi ke user
	Items  []CartItem `gorm:"foreignKey:CartID"`
}

type CartItem struct {
	gorm.Model
	CartID uint `gorm:"not null"`
	BookID uint `gorm:"not null"`
}
