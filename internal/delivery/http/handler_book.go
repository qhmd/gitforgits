package http

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/qhmd/gitforgits/config"
	bukuStruct "github.com/qhmd/gitforgits/internal/domain/book"
	"github.com/qhmd/gitforgits/internal/dto"
	"github.com/qhmd/gitforgits/internal/middleware"
	"github.com/qhmd/gitforgits/internal/usecase"
)

type BookHandler struct {
	Usecase *usecase.BookUseCase
}

func NewBookHandler(app *fiber.App, uc *usecase.BookUseCase) {
	h := &BookHandler{Usecase: uc}
	app.Get("/books", h.ListBook)
	app.Get("/books/:id<^[0-9]+$>", h.GetBookByID)
	app.Post("/books", middleware.JWT(), middleware.ValidateBook(), h.Create)
	app.Delete("/books/:id<^[0-9]+$>", h.Delete)
	app.Put("/books/:id<^[0-9]+$>", middleware.ValidateBook(), h.Update)
}

// ListBook godoc
// @Summary Get all books
// @Description Retrieve all books from the database
// @Tags Books
// @Accept json
// @Produce json
// @Success 200 {array} dto.SuccessGetListBook
// @Failure 500 {object} dto.ErrorResponse
// @Router /books [get]
func (h *BookHandler) ListBook(c *fiber.Ctx) error {
	books, err := h.Usecase.List(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "successfully get list book",
		"data":    books,
	})
}

// GetBookByID godoc
// @Summary Get book by ID
// @Description Retrieve a single book by its ID
// @Tags Books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} dto.SuccessGetBook
// @Failure 404 {object} dto.BookNotFoundResponse
// @Router /books/{id} [get]
func (h *BookHandler) GetBookByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	book, err := h.Usecase.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "books not found"})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "successfully get the book",
		"data":    book,
	})
}

// Create godoc
// @Summary Create a new book
// @Description Add a new book to the database
// @Tags Books
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param book body dto.BookRequest true "Book data"
// @Success 201 {object} dto.SuccessfullCreate
// @Failure 409 {object} dto.TitleAlreadytaken
// @Failure 400 {object} dto.MissingAuthorization
// @Failure 500 {object} dto.ErrorResponse
// @Security ApiKeyAuth
// @Router /books [post]
func (h *BookHandler) Create(c *fiber.Ctx) error {
	req := c.Locals("validateBook").(dto.BookRequest)
	book := &bukuStruct.Book{
		Title:  req.Title,
		Author: req.Author,
		Page:   req.Page,
	}
	if err := h.Usecase.Create(c.Context(), book); err != nil {
		if err == config.ErrBookTitleExists {
			return c.Status(409).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(
		fiber.Map{
			"message": "successfully add the book",
			"data":    book,
		})
}

// Update godoc
// @Summary Update book by ID
// @Description Update book information by its ID, u have to login first to access your access token
// @Tags Books
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Book ID"
// @Param book body dto.BookRequest true "Updated book data"
// @Success 200 {object} dto.SuccessfullUpdate
// @Failure 400 {object} dto.InvalidId
// @Failure 404 {object} dto.BookNotFoundResponse
// @Failure 409 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security ApiKeyAuth
// @Router /books/{id} [put]
func (h *BookHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}
	req := c.Locals("validateBook").(dto.BookRequest)
	existing, err := h.Usecase.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "book not found"})
	}
	existing.Title = req.Title
	existing.Author = req.Author
	existing.Page = req.Page

	if err := h.Usecase.Update(c.Context(), existing); err != nil {
		if err == config.ErrBookTitleExists {
			return c.Status(409).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "update successfully",
		"data":    existing,
	})

}

// Delete godoc
// @Summary Delete book by ID
// @Description Remove a book from the database using its ID
// @Tags Books
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Book ID"
// @Success 200 {object} dto.DeleteSuccesfullu
// @Failure 400 {object} dto.InvalidId
// @Failure 404 {object} dto.BookNotFoundResponse
// @Router /books/{id} [delete]
func (h *BookHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}
	_, err = h.Usecase.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "book not found"})
	}
	err = h.Usecase.Delete(c.Context(), id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"message": "delete successfully"})
}
