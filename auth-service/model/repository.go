package model

import (
	"context"

	"github.com/qhmd/gitforgits/shared/models"
)

type AuthRepository interface {
	RegisterUser(ctx context.Context, auth *models.Auth) error
	UpdateMe(ctx context.Context, auth *models.Auth) (*models.Auth, error)
	GetUserByID(ctx context.Context, id uint) (*models.Auth, error)
	FindByEmail(ctx context.Context, email string) (*models.Auth, error)
	DeleteUser(ctx context.Context, id uint) error
	LogoutUser(ctx context.Context, token string) error
}
