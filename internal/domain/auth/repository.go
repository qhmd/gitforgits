package auth

import "context"

type AuthRepository interface {
	RegisterUser(ctx context.Context, auth *Auth) error
	FindByEmail(ctx context.Context, email string) (*Auth, error)
	DeleteUser(ctx context.Context, id uint) error
	GetUserByID(ctx context.Context, id uint) (*Auth, error)
	ListUser(ctx context.Context) ([]*Auth, error)
	LogoutUser(ctx context.Context, token string) error
}
