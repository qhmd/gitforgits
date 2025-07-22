package dto

type BookRequest struct {
	Title  string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required,min=4,max=50,alphaSpace"`
	Page   int    `json:"page" validate:"required,gt=0"`
}

type RegisterRequest struct {
	Name     string `json:"name" validate:"required,min=4,max=50,alphaSpace,alphaMin4"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
