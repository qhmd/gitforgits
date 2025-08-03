package users

type ErrorResponse struct {
	Success bool   `json:"success" example:"false"`
	Error   string `json:"error" example:"something went wrong"`
}

type InvalidId struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"Your id is invalid"`
	Errors  string `json:"error" example:"Invalid id"`
}

type UserNotFoundResponse struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"user not found"`
	Errors  string `json:"errors" example:"user with id {id} does not exist"`
}

type EmailAlreadyUsed struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"Email already exist, choose another Email"`
	Errors  string `json:"error" example:"this is email already used"`
}
