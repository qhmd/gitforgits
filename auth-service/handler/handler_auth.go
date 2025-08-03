package handler

import (
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/qhmd/gitforgits/auth-service/config"
	"github.com/qhmd/gitforgits/auth-service/usecase"
	"github.com/qhmd/gitforgits/shared/dto"
	"github.com/qhmd/gitforgits/shared/middleware"
	"github.com/qhmd/gitforgits/shared/models"
	"github.com/qhmd/gitforgits/shared/utils"
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

	hashedPassword, _ := utils.HashPassword(req.Password)

	user := &models.Auth{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Role:     "user",
	}
	if err := h.Usecase.RegisterUser(c.Context(), user); err != nil {
		if err == config.ErrUserExists {
			return utils.ErrorResponse(c, fiber.StatusConflict, err.Error(), nil)
		}
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "failed to register user: ", err.Error())
	}
	return utils.SuccessResponse(c, fiber.StatusCreated, "successfull created", user)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	req := c.Locals("validateAuth").(dto.LoginRequest)

	user, err := h.Usecase.LoginUser(c.Context(), req.Email, req.Password)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusConflict, err.Error(), nil)

	}
	refreshToken, _ := utils.GenerateRefreshToken(user.ID, req.Email, user.Name, user.Role)
	accessToken, _ := utils.GenerateAccessToken(user.ID, req.Email, user.Name, user.Role)

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(7 * 24 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	})
	data := &utils.RegisUserResponse{
		AccessToken: accessToken,
		User: map[string]any{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		},
	}
	return utils.SuccessResponse(c, fiber.StatusCreated, "success to login", data)
}

func (h *AuthHandler) Me(c *fiber.Ctx) error {
	emailUser := c.Locals("userEmail").(string)
	user, err := h.Usecase.Me(c.Context(), emailUser)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "failed to fetch data: ", err.Error())
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "success to fetch me", user)

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
	return utils.SuccessResponse(c, fiber.StatusOK, "success to log out", nil)
}

func (h *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refresh_token")
	if refreshToken == "" {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, "missing refresh token", nil)
	}

	token, err := utils.VerifyRefreshToken(refreshToken)
	if err != nil || !token.Valid {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, "invalid refresh token", nil)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, "invalid claims type", nil)
	}

	accessToken, _ := utils.GenerateAccessToken(
		uint(claims["id"].(float64)),
		claims["email"].(string),
		claims["name"].(string),
		claims["role"].(string),
	)
	return utils.SuccessResponse(c, fiber.StatusOK, "success to access token", accessToken)
}

func (h *AuthHandler) UpdateMe(c *fiber.Ctx) error {
	req := c.Locals("validateAuth").(dto.RegisterRequest)
	id := uint(c.Locals("userID").(float64))
	userData := &models.Auth{
		Model:    gorm.Model{ID: id},
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	fmt.Print("ID ", userData.ID)

	updatedReq, err := h.Usecase.UpdateMe(c.Context(), userData)
	if err != nil {
		fmt.Print("isi error ", err)
		if errors.Is(err, config.ErrUserExists) {
			return utils.ErrorResponse(c, fiber.StatusConflict, "try another email", err.Error())
		}
		return utils.SuccessResponse(c, fiber.StatusInternalServerError, "ssomething went wrong", err.Error())
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "update successfully", updatedReq)
}
