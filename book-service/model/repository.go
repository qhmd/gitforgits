package model

import "context"

type BookRepository interface {
	GetBookByID(ctx context.Context, id int) (*Book, error)
	ListBook(ctx context.Context) ([]*Book, error)
	GetBookByCategory(ctx context.Context, id int) ([]*Book, error)
	CreateBook(ctx context.Context, book *Book) error
	UpdateBook(ctx context.Context, book *Book) error
	DeleteBookByID(ctx context.Context, id int) error
	GetBookByTitle(ctx context.Context, title string) (*Book, error)

	CreateCategory(ctx context.Context, category *Category) error
	DeleteCategory(ctx context.Context, id int) error
	UpdateCategory(ctx context.Context, category *Category) error
	GetCategory(ctx context.Context, category string) error
	ListCategory(ctx context.Context) ([]*Category, error)
	GetCategoryByID(ctx context.Context, id int) error
}
