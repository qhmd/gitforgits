package dto

type ErrorResponse struct {
	Error string `json:"error" example:"something went wrong"`
}

type InvalidId struct {
	Error string `json:"error" example:"Invalid id"`
}

type BookNotFoundResponse struct {
	Error string `json:"error" example:"book not found"`
}

type TitleAlreadytaken struct {
	Error string `json:"error" example:"Email already taken"`
}

type MissingAuthorization struct {
	Error string `json:"error" example:"missing authorization header"`
}
