package middleware

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/qhmd/gitforgits/internal/dto"
	"github.com/qhmd/gitforgits/utils"
)

func ValidateBook() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req dto.BookRequest

		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
		}

		if err := utils.Validate.Struct(req); err != nil {
			validationErrors := err.(validator.ValidationErrors)

			errors := make(map[string]string)
			for _, e := range validationErrors {
				errors[e.Field()] = utils.MsgForTag(e)
			}

			return c.Status(400).JSON(fiber.Map{"error": errors})
		}

		c.Locals("validateBook", req)
		return c.Next()
	}
}
