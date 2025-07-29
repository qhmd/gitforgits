package auth

import "gorm.io/gorm"

type Auth struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Role     string
	Password string
}
