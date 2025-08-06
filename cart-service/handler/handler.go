package handler

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/qhmd/gitforgits/cart-service/usecase"
)

type CartHandler struct {
	uc usecase.CartUsecase
}

func NewCartHandler(app *fiber.App, uc usecase.CartUsecase) {
	h := &CartHandler{uc: uc}

	app.Post("/cart", h.AddItemToCart)
	app.Get("/cart/:user_id", h.GetCartByUserID)
	app.Delete("/cart/item/:id", h.RemoveItemFromCart)
	app.Delete("/cart/user/:user_id", h.ClearCartByUserID)
}

type addItemRequest struct {
	UserID uint `json:"user_id"`
	BookID uint `json:"book_id"`
}

func (h *CartHandler) AddItemToCart(c *fiber.Ctx) error {
	var req addItemRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	if err := h.uc.AddItemToCart(c.Context(), req.UserID, req.BookID); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "failed to add item"})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "item added to cart"})
}

func (h *CartHandler) GetCartByUserID(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("user_id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid user id"})
	}

	cart, err := h.uc.GetCartByUserID(c.Context(), uint(userID))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "failed to fetch cart"})
	}

	return c.JSON(cart)
}

func (h *CartHandler) RemoveItemFromCart(c *fiber.Ctx) error {
	itemID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid item id"})
	}

	if err := h.uc.RemoveItemFromCart(c.Context(), uint(itemID)); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "failed to remove item"})
	}

	return c.JSON(fiber.Map{"message": "item removed"})
}

func (h *CartHandler) ClearCartByUserID(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("user_id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid user id"})
	}

	if err := h.uc.ClearCartByUserID(c.Context(), uint(userID)); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "failed to clear cart"})
	}

	return c.JSON(fiber.Map{"message": "cart cleared"})
}
