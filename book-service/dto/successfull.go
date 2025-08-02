package dto

import "github.com/qhmd/gitforgits/book-service/model"

type SuccessfullUpdate struct {
	Success bool       `json:"success" example:"true"`
	Data    model.Book `json:"data"`
	Message string     `json:"message" example:"update successfully"`
}

type SuccessfullCreate struct {
	Success bool       `json:"success" example:"true"`
	Data    model.Book `json:"data"`
	Message string     `json:"message" example:"successfully add the book"`
}

type SuccessGetBook struct {
	Success bool       `json:"success" example:"true"`
	Data    model.Book `json:"data"`
	Message string     `json:"message" example:"successfully get the book"`
}

type SuccessGetListBook struct {
	Success bool       `json:"success" example:"true"`
	Data    model.Book `json:"data"`
	Message string     `json:"message" example:"successfully get list book"`
}

type DeleteSuccesfully struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"delete successfully"`
}
