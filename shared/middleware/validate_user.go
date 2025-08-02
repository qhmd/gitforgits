package middleware

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/qhmd/gitforgits/shared/dto"
	"github.com/qhmd/gitforgits/shared/utils"
)

func ValidateUserUpdate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req dto.UserResponse
		if err := c.BodyParser(&req); err != nil {
			return utils.ErrorResponse(c, fiber.StatusBadRequest, "invalid request body", err)
		}
		if err := utils.Validate.Struct(req); err != nil {
			ValidationError := err.(validator.ValidationErrors)
			errors := make(map[string]string)
			for _, e := range ValidationError {
				errors[e.Field()] = utils.MsgForTag(e)
			}
			return utils.ErrorResponse(c, fiber.StatusBadRequest, "validation error", errors)
		}
		c.Locals("validateUser", req)
		return c.Next()
	}
}
