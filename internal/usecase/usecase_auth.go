package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/qhmd/gitforgits/config"
	"github.com/qhmd/gitforgits/internal/domain/auth"
	"github.com/qhmd/gitforgits/internal/dto"
	"github.com/qhmd/gitforgits/utils"
)

type AuthUseCase struct {
	repo auth.AuthRepository
}

func NewAuthUsecase(repo auth.AuthRepository) *AuthUseCase {
	return &AuthUseCase{repo: repo}
}

func (u *AuthUseCase) RegisterUser(ctx context.Context, us *auth.Auth) error {
	existing, err := u.repo.FindByEmail(ctx, us.Email)
	fmt.Printf("isi exis : %v, dan isi err %v", existing, err)
	if err != nil {
		return err
	}

	if existing != nil {
		return config.ErrUserExists
	}
	return u.repo.RegisterUser(ctx, us)
}

func (u *AuthUseCase) LoginUser(ctx context.Context, email, password string) (*auth.Auth, error) {
	user, err := u.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid email or password")
	}
	return user, nil
}

func (u *AuthUseCase) Me(ctx context.Context, email string) (*dto.UserResponse, error) {
	user, err := u.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	userReponse := &dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}
	return userReponse, nil
}

func (u *AuthUseCase) UpdateMe(ctx context.Context, email string) (*dto.RegisterRequest, error) {
	user, err := u.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	userUpdateReponse := &dto.RegisterRequest{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	return userUpdateReponse, nil
}

func (u *AuthUseCase) GetUserByID(ctx context.Context, id uint) (*auth.Auth, error) {
	return u.repo.GetUserByID(ctx, id)
}
