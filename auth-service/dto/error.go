package dto

type ErrorResponseAuth struct {
	Success bool   `json:"success" example:"false"`
	Error   string `json:"error" example:"something went wrong"`
}

type ErrorResponseLogin struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"try another email"`
	Error   string `json:"error" example:"email already used"`
}

type ErrorUnauthorized struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"missing refresh token"`
}
