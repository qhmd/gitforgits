package middleware

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/qhmd/gitforgits/internal/dto"
	"github.com/qhmd/gitforgits/utils"
)

func ValidateUserUpdate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req dto.UserResponse
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "invalid body request"})
		}
		if err := utils.Validate.Struct(req); err != nil {
			ValidationError := err.(validator.ValidationErrors)
			errors := make(map[string]string)
			for _, e := range ValidationError {
				errors[e.Field()] = utils.MsgForTag(e)
			}
			return c.Status(400).JSON(errors)
		}
		c.Locals("validateUser", req)
		return c.Next()
	}
}
