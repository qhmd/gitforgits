package middleware

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/qhmd/gitforgits/internal/dto"
)

func ValidateBook() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req dto.BookRequest

		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
		}

		validate := validator.New()
		if err := validate.Struct(req); err != nil {
			validationErrors := err.(validator.ValidationErrors)

			errors := make(map[string]string)
			for _, e := range validationErrors {
				errors[e.Field()] = msgForTag(e)
			}

			return c.Status(400).JSON(errors)
		}

		c.Locals("validateBook", req)
		return c.Next()
	}
}

func msgForTag(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return e.Field() + " is required"
	case "gt":
		return e.Field() + " must be greater than " + e.Param()
	default:
		return e.Field() + " is not valid"
	}
}
