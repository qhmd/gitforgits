package http

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/qhmd/gitforgits/config"
	"github.com/qhmd/gitforgits/internal/domain/auth"
	"github.com/qhmd/gitforgits/internal/dto"
	"github.com/qhmd/gitforgits/internal/middleware"
	"github.com/qhmd/gitforgits/internal/usecase"
	"github.com/qhmd/gitforgits/utils"
	"gorm.io/gorm"
)

type AuthHandler struct {
	Usecase *usecase.AuthUseCase
}

func NewAuthHandler(app *fiber.App, uc *usecase.AuthUseCase) {
	h := &AuthHandler{Usecase: uc}
	app.Post("/auth/register", middleware.ValidateAuthReg(), h.Register)
	app.Post("/auth/login", middleware.ValidateAuthLogin(), h.Login)
	app.Post("/auth/logout", h.Logout)
	app.Get("/auth/me", middleware.JWT(), h.Me)
	app.Put("/auth/me/update", middleware.JWT(), middleware.ValidateAuthReg(), h.UpdateMe)
	app.Post("/auth/refresh", h.RefreshToken)
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	req := c.Locals("validateAuth").(dto.RegisterRequest)

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to hash password"})
	}

	user := &auth.Auth{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Role:     "user",
	}
	if err := h.Usecase.RegisterUser(c.Context(), user); err != nil {
		if err == config.ErrUserExists {
			return c.Status(409).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(500).JSON(fiber.Map{"error": "failed to register user: " + err.Error()})
	}
	return c.Status(201).JSON(user)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	req := c.Locals("validateAuth").(dto.LoginRequest)

	user, err := h.Usecase.LoginUser(c.Context(), req.Email, req.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}
	refreshToken, _ := utils.GenerateRefreshToken(user.ID, req.Email, user.Name)
	accessToken, _ := utils.GenerateAccessToken(user.ID, req.Email, user.Name)

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(7 * 24 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	})

	return c.Status(200).JSON(fiber.Map{
		"access_token": accessToken,
		"user": fiber.Map{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		},
	})
}

func (h *AuthHandler) Me(c *fiber.Ctx) error {
	emailUser := c.Locals("userEmail").(string)
	user, err := h.Usecase.Me(c.Context(), emailUser)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(user)
}

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	})
	return c.SendStatus(204)
}

func (h *AuthHandler) GetUserByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	user, err := h.Usecase.GetUserByID(c.Context(), uint(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "user not found"})
	}
	return c.JSON(user)

}

func (h *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refresh_token")
	if refreshToken == "" {
		return c.Status(401).JSON(fiber.Map{"error": "missing refresh token"})
	}

	token, err := utils.VerifyRefreshToken(refreshToken)
	if err != nil || !token.Valid {
		return c.Status(401).JSON(fiber.Map{"error": "invalid refresh token"})
	}

	claims := token.Claims.(jwt.MapClaims)

	accessToken, _ := utils.GenerateAccessToken(
		uint(claims["id"].(float64)),
		claims["email"].(string),
		claims["name"].(string),
	)

	return c.JSON(fiber.Map{
		"access_token": accessToken,
	})
}

func (h *AuthHandler) UpdateMe(c *fiber.Ctx) error {
	req := c.Locals("validateAuth").(dto.RegisterRequest)
	id := uint(c.Locals("userID").(float64))
	userData := &auth.Auth{
		Model:    gorm.Model{ID: id},
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	fmt.Print("ID ", userData.ID)

	updatedReq, err := h.Usecase.UpdateMe(c.Context(), userData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "something went wrong"})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "update successfully",
		"data":    updatedReq,
	})
}
