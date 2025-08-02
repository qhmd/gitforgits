package database

import (
	"fmt"

	"github.com/qhmd/gitforgits/book-service/model"
	"github.com/qhmd/gitforgits/shared/models"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	if err := db.AutoMigrate(&model.Book{}, &models.Auth{}); err != nil {
		fmt.Println("Migration gagal:", err)
		return
	}
	fmt.Println("Migration berhasil")
}
