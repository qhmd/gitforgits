package handler

import (
	"context"
	"fmt"

	"github.com/qhmd/gitforgits/shared/models"
	authproto "github.com/qhmd/gitforgits/shared/proto/auth-proto"
	"github.com/qhmd/gitforgits/users-service/config"
	"github.com/qhmd/gitforgits/users-service/usecase"
)

type AuthGrcpHandler struct {
	authproto.UnsafeAuthServiceServer
	uc *usecase.UsersUseCase
}

func NewAuthGrcpHandler(u *usecase.UsersUseCase) *AuthGrcpHandler {
	return &AuthGrcpHandler{uc: u}
}

func (h *AuthGrcpHandler) CreateAuth(c context.Context, req *authproto.CreateAuthRequest) (*authproto.CreateAuthResponse, error) {
	fmt.Println("isi dari req : ", req)
	user := &models.Auth{
		Name:     req.GetName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Role:     req.GetRole(),
	}

	// ketika user di update
	if exist, err := h.uc.GetUserByID(c, int(req.GetId())); err != nil {
		return nil, err
	} else if exist != nil {
		if exist.ID == uint(req.Id) {
			user.Role = exist.Role
			_, err := h.uc.UpdateUser(c, user, req.GetId())
			if err != nil {
				return nil, err
			}
			resp := &authproto.CreateAuthResponse{
				Success: true,
				Message: "User updated successfully",
			}
			return resp, nil
		}
	}

	// ketika user registrasi
	if err := h.uc.RegisterUser(c, user); err != nil {
		fmt.Print("kalau error : ", err)
		if err == config.ErrUserExists {
			return nil, config.ErrUserExists
		}
		return nil, err
	}
	return &authproto.CreateAuthResponse{Success: true, Message: "User created successfully"}, nil
}
