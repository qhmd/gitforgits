package middleware

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/qhmd/gitforgits/internal/dto"
	"github.com/qhmd/gitforgits/utils"
)

func ValidateAuthLogin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req dto.LoginRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
		}
		if err := utils.Validate.Struct(req); err != nil {
			validationError := err.(validator.ValidationErrors)
			errors := make(map[string]string)
			for _, e := range validationError {
				errors[e.Field()] = utils.MsgForTag(e)
			}
			return c.Status(400).JSON(errors)
		}
		c.Locals("validateAuth", req)
		return c.Next()
	}
}

func ValidateAuthReg() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req dto.RegisterRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
		}

		if err := utils.Validate.Struct(req); err != nil {
			validationError := err.(validator.ValidationErrors)
			errors := make(map[string]string)
			for _, e := range validationError {
				errors[e.Field()] = utils.MsgForTag(e)
			}
			return c.Status(400).JSON(errors)
		}
		c.Locals("validateAuth", req)
		return c.Next()
	}
}
