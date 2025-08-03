package users

import "github.com/qhmd/gitforgits/shared/models"

type SuccessGetUser struct {
	Success bool         `json:"success" example:"true"`
	Message string       `json:"message" example:"successfully get the user"`
	Data    *models.Auth `json:"data"`
}

type SuccessGetList struct {
	Success bool           `json:"success" example:"true"`
	Message string         `json:"message" example:"successfully get list user"`
	Data    []*models.Auth `json:"data"`
}

type SuccessDeleteUser struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"success delete user"`
}
