package handler

import (
	"github.com/qhmd/gitforgits/auth-service/usecase"
	userproto "github.com/qhmd/gitforgits/shared/proto/users-proto"
)

type UsersGrcpHandler struct {
	userproto.UnimplementedUsersServiceServer
	uc *usecase.AuthUseCase
}

func NewUsersGrcpHandler(uc *usecase.AuthUseCase) *UsersGrcpHandler {
	return &UsersGrcpHandler{uc: uc}
}
