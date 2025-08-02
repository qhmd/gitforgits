package database

import (
	"fmt"

	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	if err := db.AutoMigrate(&models.Auth{}); err != nil {
		fmt.Println("Migration gagal:", err)
		return
	}
	fmt.Println("Migration berhasil")
}
