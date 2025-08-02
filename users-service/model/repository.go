package model

import (
	"context"

	authDto "github.com/qhmd/gitforgits/auth-service/dto"
	"github.com/qhmd/gitforgits/shared/models"
)

type UserRepository interface {
	GetUser(ctx context.Context, id int) (*models.Auth, error)
	ListUser(ctx context.Context) ([]*models.Auth, error)
	FindByEmail(ctx context.Context, email string) (*models.Auth, error)
	DeleteUser(ctx context.Context, id int) error
	UpdateUser(ctx context.Context, users *authDto.UserResponse, id int) (*authDto.UserResponse, error)
}
