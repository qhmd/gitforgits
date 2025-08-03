package database

import (
	"fmt"

	"github.com/qhmd/gitforgits/book-service/model"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	if err := db.AutoMigrate(&model.Book{}, &model.Category{}); err != nil {
		fmt.Println("Migration gagal:", err)
		return
	}
	fmt.Println("Migration berhasil")
}
