package database

import (
	"fmt"

	"github.com/qhmd/gitforgits/internal/domain/book"
	"gorm.io/gorm"
)

func RunMigaration(db *gorm.DB) {
	db.AutoMigrate(&book.Book{})
	fmt.Println("berhasil")
}
