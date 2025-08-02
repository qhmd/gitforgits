package middleware

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"github.com/qhmd/gitforgits/shared/dto"
	"github.com/qhmd/gitforgits/shared/utils"
)

func ValidateAuthLogin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req dto.LoginRequest
		if err := c.BodyParser(&req); err != nil {
			return utils.ErrorResponse(c, fiber.StatusBadRequest, "invalid request body", err)

		}
		if err := utils.Validate.Struct(req); err != nil {
			validationError := err.(validator.ValidationErrors)
			errors := make(map[string]string)
			for _, e := range validationError {
				errors[e.Field()] = utils.MsgForTag(e)
			}
			return utils.ErrorResponse(c, fiber.StatusBadRequest, "validation error", errors)
		}
		c.Locals("validateAuth", req)
		return c.Next()
	}
}

func ValidateAuthReg() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req dto.RegisterRequest
		if err := c.BodyParser(&req); err != nil {
			return utils.ErrorResponse(c, fiber.StatusBadRequest, "invalid request body", err)
		}

		if err := utils.Validate.Struct(req); err != nil {
			validationError := err.(validator.ValidationErrors)
			errors := make(map[string]string)
			for _, e := range validationError {
				errors[e.Field()] = utils.MsgForTag(e)
			}
			return utils.ErrorResponse(c, fiber.StatusBadRequest, "validation error", errors)
		}
		c.Locals("validateAuth", req)
		return c.Next()
	}
}
