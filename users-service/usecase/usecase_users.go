package usecase

import (
	"context"
	"fmt"

	"github.com/qhmd/gitforgits/shared/models"
	"github.com/qhmd/gitforgits/shared/utils"

	"github.com/qhmd/gitforgits/users-service/client"
	"github.com/qhmd/gitforgits/users-service/config"
	"github.com/qhmd/gitforgits/users-service/model"
)

type UsersUseCase struct {
	repo        model.UserRepository
	UsersClient *client.AuthServiceClient
}

func NewUsersUseCase(repo model.UserRepository, userClient *client.AuthServiceClient) *UsersUseCase {
	return &UsersUseCase{repo: repo, UsersClient: userClient}
}

func (uc *UsersUseCase) ListUser(ctx context.Context) ([]*models.Auth, error) {
	return uc.repo.ListUser(ctx)
}

func (uc *UsersUseCase) GetUserByID(ctx context.Context, id int) (*models.Auth, error) {
	return uc.repo.GetUser(ctx, id)
}

func (uc *UsersUseCase) UpdateUser(ctx context.Context, user *models.Auth, id int32) (*models.Auth, error) {
	pw, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	data := &models.Auth{
		Name:     user.Name,
		Email:    user.Email,
		Password: pw,
		Role:     user.Role,
	}

	return uc.repo.UpdateUser(ctx, data, id)
}

func (uc *UsersUseCase) DeleteUser(ctx context.Context, id int) error {
	return uc.repo.DeleteUser(ctx, id)
}

func (u *UsersUseCase) RegisterUser(ctx context.Context, us *models.Auth) error {
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
