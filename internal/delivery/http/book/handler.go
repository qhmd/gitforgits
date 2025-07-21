package book

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	bukuStruct "github.com/qhmd/gitforgits/internal/domain/book"
	"github.com/qhmd/gitforgits/internal/dto"
	"github.com/qhmd/gitforgits/internal/middleware"

	"github.com/qhmd/gitforgits/internal/usecase/book"
)

type BookHandler struct {
	Usecase *book.BookUseCase
}

func NewBookHandler(app *fiber.App, uc *book.BookUseCase) {
	h := &BookHandler{Usecase: uc}
	app.Get("/books", h.ListBook)
	app.Get("/books/:id<^[0-9]+$>", h.GetBookByID)
	app.Post("/books", middleware.ValidateBook(), h.Create)
	app.Delete("/books/:id<^[0-9]+$>", h.Delete)
	app.Put("/books/:id<^[0-9]+$>", middleware.ValidateBook(), h.Update)
}

func (h *BookHandler) ListBook(c *fiber.Ctx) error {
	books, err := h.Usecase.List(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(books)
}

func (h *BookHandler) GetBookByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	book, err := h.Usecase.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "not found"})
	}
	return c.JSON(book)
}

func (h *BookHandler) Create(c *fiber.Ctx) error {
	req := c.Locals("validateBook").(dto.BookRequest)
	book := &bukuStruct.Book{
		Title:  req.Title,
		Author: req.Author,
		Page:   req.Page,
	}
	if err := h.Usecase.Create(c.Context(), book); err != nil {
		if err == bukuStruct.ErrBookTitleExists {
			return c.Status(409).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(book)
}

func (h *BookHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}
	req := c.Locals("validateBook").(dto.BookRequest)
	// Ambil data dari DB dulu
	existing, err := h.Usecase.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "book not found"})
	}
	// Update datanya
	existing.Title = req.Title
	existing.Author = req.Author
	existing.Page = req.Page

	if err := h.Usecase.Update(c.Context(), existing); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "update successfully",
		"data":    existing,
	})

}

func (h *BookHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}
	err = h.Usecase.Delete(c.Context(), id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "book not found"})
	}
	return c.Status(200).JSON(fiber.Map{"message": "delete successfully"})
}
