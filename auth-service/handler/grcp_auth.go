package handler

import (
	"context"
	"fmt"

	"github.com/qhmd/gitforgits/auth-service/usecase"
	"github.com/qhmd/gitforgits/shared/models"
	userproto "github.com/qhmd/gitforgits/shared/proto/users-proto"
	"gorm.io/gorm"
)

type UsersGrcpHandler struct {
	userproto.UnsafeUsersServiceServer
	uc *usecase.AuthUseCase
}

func NewUsersGrcpHandler(uc *usecase.AuthUseCase) *UsersGrcpHandler {
	return &UsersGrcpHandler{uc: uc}
}

// DeleteUser implements users_proto.UsersServiceServer.
func (h *UsersGrcpHandler) DeleteUser(c context.Context, req *userproto.DeleteUserRequest) (*userproto.Response, error) {
	err := h.uc.DeleteUserByID(c, uint(req.Id))
	if err != nil {
		res := &userproto.Response{Success: false, Message: "failed to deleted"}
		return res, err
	}
	res := &userproto.Response{Success: true, Message: "successfully deleted"}
	return res, nil
}

func (h *UsersGrcpHandler) UpdateUser(c context.Context, req *userproto.UserRequestData) (*userproto.Response, error) {
	fmt.Println("isi dari req : ", req)

	user := &models.Auth{
		Model:    gorm.Model{ID: uint(req.Id)},
		Name:     req.GetName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Role:     req.GetRole(),
	}
	updateUser, err := h.uc.UpdateMe(c, user)
	if err != nil {
		return nil, err
	}

	fmt.Println("isi updateUser : ", updateUser)
	fmt.Println("isi req : ", req)
	resp := &userproto.Response{
		Success: true,
		Message: "User update successfully",
	}
	return resp, nil
}
