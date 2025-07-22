package auth

import "gorm.io/gorm"

type Auth struct {
	gorm.Model
	Name     string
	Email    string
	Role     string
	Password string
}
