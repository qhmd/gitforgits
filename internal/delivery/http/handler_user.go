package http

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/qhmd/gitforgits/config"
	authDto "github.com/qhmd/gitforgits/internal/dto/auth"
	"github.com/qhmd/gitforgits/internal/middleware"
	"github.com/qhmd/gitforgits/internal/usecase"
)

type UserHandler struct {
	uc *usecase.UsersUseCase
}

func NewHandlerUser(app *fiber.App, uc *usecase.UsersUseCase) {
	h := &UserHandler{uc: uc}
	app.Get("/admin/users/:id<^[0-9]+$>", h.GetUserByID)
	app.Get("/admin/users/", h.GetListUsers)
	app.Put("/admin/users/:id<^[0-9]+$>", middleware.ValidateUserUpdate(), h.UpdateUsers)
	app.Delete("admin/users/:id<^[0-9]+$>", h.DeleteUserByID)
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}
	data, err := h.uc.GetUserByID(c.Context(), id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "user not found"})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "successfully find the user",
		"data":    data,
	})
}

func (h *UserHandler) GetListUsers(c *fiber.Ctx) error {
	data, err := h.uc.ListUser(c.Context())
	if err != nil {
		c.Status(500).JSON(fiber.Map{"error": "something went wrong"})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "successfully get list user",
		"data":    data,
	})
}

func (h *UserHandler) UpdateUsers(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}
	user := c.Locals("validateUser").(authDto.UserResponse)

	_, err = h.uc.GetUserByID(c.Context(), id)
	fmt.Print(err)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "user not found"})
	}

	updateUser, err := h.uc.UpdateUser(c.Context(), &user, id)
	if err != nil {
		if errors.Is(err, config.ErrUserExists) {
			return c.Status(409).JSON(fiber.Map{
				"error": config.ErrUserExists,
			})
		}
		return c.Status(500).JSON(fiber.Map{"error": "something went wrong"})
	}
	fmt.Print("updated user : ", updateUser)
	return c.Status(200).JSON(fiber.Map{
		"message": "successfully update the user",
		"data":    updateUser,
	})
}

func (h *UserHandler) DeleteUserByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}
	_, err = h.uc.GetUserByID(c.Context(), id)
	fmt.Print(err)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "user not found"})
	}
	if err = h.uc.DeleteUser(c.Context(), id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "something went wrong"})
	}
	return c.Status(200).JSON(fiber.Map{"message": "successfully delete"})
}
