package model

import (
	"context"

	"github.com/qhmd/gitforgits/shared/models"
)

type UserRepository interface {
	GetUser(ctx context.Context, id int) (*models.Auth, error)
	ListUser(ctx context.Context) ([]*models.Auth, error)
	FindByEmail(ctx context.Context, email string) (*models.Auth, error)
	DeleteUser(ctx context.Context, id int) error
	UpdateUser(ctx context.Context, users *models.Auth, id int32) (*models.Auth, error)

	RegisterUser(ctx context.Context, auth *models.Auth) error
}
