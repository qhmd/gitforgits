package repo

import (
	"context"

	"github.com/qhmd/gitforgits/cart-service/model"
	"gorm.io/gorm"
)

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) model.CartRepository {
	return &cartRepository{db: db}
}

func (r *cartRepository) AddItemToCart(ctx context.Context, userID uint, bookID uint) error {
	var cart model.Cart
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).FirstOrCreate(&cart, model.Cart{UserID: userID}).Error; err != nil {
		return err
	}

	item := model.CartItem{
		CartID: cart.ID,
		BookID: bookID,
	}
	return r.db.WithContext(ctx).Create(&item).Error
}

func (r *cartRepository) GetCartByUserID(ctx context.Context, userID uint) (*model.Cart, error) {
	var cart model.Cart
	if err := r.db.WithContext(ctx).
		Preload("Items").
		Where("user_id = ?", userID).
		First(&cart).Error; err != nil {
		return nil, err
	}
	return &cart, nil
}

func (r *cartRepository) RemoveItemFromCart(ctx context.Context, itemID uint) error {
	return r.db.WithContext(ctx).Delete(&model.CartItem{}, itemID).Error
}

func (r *cartRepository) ClearCartByUserID(ctx context.Context, userID uint) error {
	var cart model.Cart
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(&cart).Error; err != nil {
		return err
	}
	return r.db.WithContext(ctx).Where("cart_id = ?", cart.ID).Delete(&model.CartItem{}).Error
}
