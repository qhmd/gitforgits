package user

import (
	"context"

	"github.com/qhmd/gitforgits/internal/domain/auth"
	"github.com/qhmd/gitforgits/internal/dto"
)

type UserRepository interface {
	GetUser(ctx context.Context, id int) (*auth.Auth, error)
	ListUser(ctx context.Context) ([]*auth.Auth, error)
	FindByEmail(ctx context.Context, email string) (*auth.Auth, error)
	DeleteUser(ctx context.Context, id int) error
	UpdateUser(ctx context.Context, users *dto.UserResponse, id int) (*dto.UserResponse, error)
}
