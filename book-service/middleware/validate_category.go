package middleware

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/qhmd/gitforgits/book-service/dto"
	"github.com/qhmd/gitforgits/shared/utils"
)

func ValidateCategory() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req dto.CategoryRequest

		if err := c.BodyParser(&req); err != nil {
			return utils.ErrorResponse(c, fiber.StatusBadRequest, "invalid request body", err)
		}

		if err := utils.Validate.Struct(req); err != nil {
			validationErrors := err.(validator.ValidationErrors)

			errors := make(map[string]string)
			for _, e := range validationErrors {
				errors[e.Field()] = utils.MsgForTag(e)
			}
			return utils.ErrorResponse(c, fiber.StatusBadRequest, "validation error", errors)

		}

		c.Locals("validateCategory", req)
		return c.Next()
	}
}
