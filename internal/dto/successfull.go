package dto

import "github.com/qhmd/gitforgits/internal/domain/book"

type SuccessfullUpdate struct {
	Data    book.Book `json:"data"`
	Message string    `json:"message" example:"update successfully"`
}

type DeleteSuccesfullu struct {
	Message string `json:"error" example:"delete successfully"`
}
