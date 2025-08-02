package usecase

import (
	"context"

	"github.com/qhmd/gitforgits/users-service/model"
	"github.com/qhmd/gitforgits/utils"
)

type UsersUseCase struct {
	repo model.UserRepository
}

func NewUsersUseCase(repo model.UserRepository) *UsersUseCase {
	return &UsersUseCase{repo: repo}
}

func (uc *UsersUseCase) ListUser(ctx context.Context) ([]*models.Auth, error) {
	return uc.repo.ListUser(ctx)
}

func (uc *UsersUseCase) GetUserByID(ctx context.Context, id int) (*models.Auth, error) {
	return uc.repo.GetUser(ctx, id)
}

func (uc *UsersUseCase) UpdateUser(ctx context.Context, user *dto.UserResponse, id int) (*dto.UserResponse, error) {
	pw, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	data := &dto.UserResponse{
		RegisterRequest: dto.RegisterRequest{
			Name:     user.Name,
			Email:    user.Email,
			Password: pw,
		},
		Role: user.Role,
	}

	return uc.repo.UpdateUser(ctx, data, id)
}

func (uc *UsersUseCase) DeleteUser(ctx context.Context, id int) error {
	return uc.repo.DeleteUser(ctx, id)
}
