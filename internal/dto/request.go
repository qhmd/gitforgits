package dto

type BookRequest struct {
	Title  string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required"`
	Page   int    `json:"page" validate:"required,gt=0"`
}
