package dto

import "github.com/qhmd/gitforgits/internal/domain/book"

type SuccessfullUpdate struct {
	Data    book.Book `json:"data"`
	Message string    `json:"message" example:"update successfully"`
}

type SuccessfullCreate struct {
	Data    book.Book `json:"data"`
	Message string    `json:"message" example:"successfully add the book"`
}

type SuccessGetBook struct {
	Data    book.Book `json:"data"`
	Message string    `json:"message" example:"successfully get the book"`
}

type SuccessGetListBook struct {
	Data    book.Book `json:"data"`
	Message string    `json:"message" example:"successfully get list book"`
}

type DeleteSuccesfullu struct {
	Message string `json:"error" example:"delete successfully"`
}
