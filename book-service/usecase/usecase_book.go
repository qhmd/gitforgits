package usecase

import (
	"context"
	"fmt"

	"github.com/qhmd/gitforgits/book-service/model"
	"github.com/qhmd/gitforgits/config"
)

type BookUseCase struct {
	repo model.BookRepository
}

func NewBookUsecase(repo model.BookRepository) *BookUseCase {
	return &BookUseCase{repo: repo}
}

func (u *BookUseCase) GetByID(ctx context.Context, id int) (*model.Book, error) {
	return u.repo.GetBookByID(ctx, id)
}

func (u *BookUseCase) List(ctx context.Context) ([]*model.Book, error) {
	return u.repo.ListBook(ctx)
}

func (u *BookUseCase) Create(ctx context.Context, b *model.Book) error {
	existing, err := u.repo.GetBookByTitle(ctx, b.Title)

	if err != nil {
		return err
	}

	if existing != nil {
		return config.ErrBookTitleExists
	}

	return u.repo.CreateBook(ctx, b)
}

func (u *BookUseCase) Delete(ctx context.Context, id int) error {
	return u.repo.DeleteBookByID(ctx, id)
}

func (u *BookUseCase) Update(ctx context.Context, b *model.Book) error {
	fmt.Println("ini" + b.Title)
	existing, err := u.repo.GetBookByTitle(ctx, b.Title)

	if err != nil {
		return err
	}
	if existing != nil {
		if existing.Title == b.Title {
			return u.repo.UpdateBook(ctx, b)
		}
		return config.ErrBookTitleExists
	}
	return u.repo.UpdateBook(ctx, b)
}
