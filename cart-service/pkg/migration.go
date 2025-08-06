package pkg

import (
	"fmt"

	"github.com/qhmd/gitforgits/cart-service/model"

	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	if err := db.AutoMigrate(&model.Cart{}, &model.CartItem{}); err != nil {
		fmt.Println("Migration gagal:", err)
		return
	}
	fmt.Println("Migration berhasil")
}
