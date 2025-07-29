package dto

type BookRequest struct {
	Title  string `json:"title" validate:"required" example:"How To Become Backend Engineer"`
	Author string `json:"author" validate:"required,min=4,max=50,alphaSpace" example:"John Smith"`
	Page   int    `json:"page" validate:"required,gt=0" example:"205"`
}

type RegisterRequest struct {
	Name     string `json:"name" validate:"required,min=4,max=50,alphaSpace,alphaMin4"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
