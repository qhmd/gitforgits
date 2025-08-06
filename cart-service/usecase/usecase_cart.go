package usecase

import (
	"context"
	"fmt"

	"github.com/qhmd/gitforgits/cart-service/model"
)

type CartUsecase interface {
	AddItemToCart(ctx context.Context, userID uint, bookID uint) error
	GetCartByUserID(ctx context.Context, userID uint) (*model.Cart, error)
	RemoveItemFromCart(ctx context.Context, itemID uint) error
	ClearCartByUserID(ctx context.Context, userID uint) error
}

type cartUsecase struct {
	repo model.CartRepository
}

func NewCartUsecase(repo model.CartRepository) *cartUsecase {
	return &cartUsecase{repo: repo}
}

func (u *cartUsecase) AddItemToCart(ctx context.Context, userID uint, bookID uint) error {
	fmt.Println("isi userd id dan book ", userID, " book : ", bookID)
	return u.repo.AddItemToCart(ctx, userID, bookID)
}

func (u *cartUsecase) GetCartByUserID(ctx context.Context, userID uint) (*model.Cart, error) {
	return u.repo.GetCartByUserID(ctx, userID)
}

func (u *cartUsecase) RemoveItemFromCart(ctx context.Context, itemID uint) error {
	fmt.Println("isi item id dan book ", itemID)
	return u.repo.RemoveItemFromCart(ctx, itemID)
}

func (u *cartUsecase) ClearCartByUserID(ctx context.Context, userID uint) error {
	return u.repo.ClearCartByUserID(ctx, userID)
}
