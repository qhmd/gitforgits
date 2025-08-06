package model

import (
	"context"
)

type CartRepository interface {
	AddItemToCart(ctx context.Context, userID uint, bookID uint) error
	GetCartByUserID(ctx context.Context, userID uint) (*Cart, error)
	RemoveItemFromCart(ctx context.Context, itemID uint) error
	ClearCartByUserID(ctx context.Context, userID uint) error
}
