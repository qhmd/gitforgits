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
	UsersClient *client.UsersServiceClient
}

func NewUsersUseCase(repo model.UserRepository, userClient *client.UsersServiceClient) *UsersUseCase {
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
	fmt.Printf("isi exis : %v, dan isi err %v", data, err)

	uc.UsersClient.UpdateUsers(ctx, data, id)
	updateData, err := uc.repo.UpdateUser(ctx, data, id)
	if err != nil {
		return nil, err
	}

	return updateData, nil
}

func (uc *UsersUseCase) DeleteUser(ctx context.Context, id int) error {
	uc.UsersClient.DeleteUser(ctx, int32(id))
	results := uc.repo.DeleteUser(ctx, id)
	return results
}

func (u *UsersUseCase) RegisterUser(ctx context.Context, us *models.Auth) error {
	existing, err := u.repo.FindByEmail(ctx, us.Email)
	if err != nil {
		return err
	}

	if existing != nil {
		return config.ErrUserExists
	}

	return u.repo.RegisterUser(ctx, us)
}
