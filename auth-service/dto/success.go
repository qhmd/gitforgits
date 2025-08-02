package dto

import (
	"github.com/qhmd/gitforgits/shared/models"
)

type SuccessRegis struct {
	Success bool         `json:"success" example:"true"`
	Data    *models.Auth `json:"data"`
	Message string       `json:"message" example:"successfull created"`
}

type SuccessLogin struct {
	Success bool         `json:"success" example:"true"`
	Data    *models.Auth `json:"data"`
	Message string       `json:"message" example:"success to login"`
}

type SuccessUpdate struct {
	Success bool         `json:"success" example:"true"`
	Data    *models.Auth `json:"data"`
	Message string       `json:"message" example:"updated successfully"`
}

type SuccessLogout struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"success to log out"`
}

type SuccessAccessToken struct {
	Success bool   `json:"success" example:"true"`
	Data    string `json:"data" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im1hZGRla2U0N..."`
	Message string `json:"message" example:"success to access token"`
}
