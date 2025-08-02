package repo

import (
	"context"
	"errors"

	"github.com/qhmd/gitforgits/book-service/model"
	"gorm.io/gorm"
)

type mysqlBookRepository struct {
	db *gorm.DB
}

// CreateBook implements book.BookRepository.
func (m *mysqlBookRepository) CreateBook(ctx context.Context, book *model.Book) error {
	return m.db.WithContext(ctx).Create(book).Error
}

// DeleteBookByID implements book.BookRepository.
func (m *mysqlBookRepository) DeleteBookByID(ctx context.Context, id int) error {
	result := m.db.WithContext(ctx).Delete(&model.Book{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetBookByID implements book.BookRepository.
func (m *mysqlBookRepository) GetBookByID(ctx context.Context, id int) (*model.Book, error) {
	var b model.Book
	err := m.db.WithContext(ctx).First(&b, id).Error
	return &b, err
}

// ListBook implements book.BookRepository.
func (m *mysqlBookRepository) ListBook(ctx context.Context) ([]*model.Book, error) {
	var books []*model.Book
	err := m.db.WithContext(ctx).Find(&books).Error
	return books, err
}

// UpdateBook implements book.BookRepository.
func (m *mysqlBookRepository) UpdateBook(ctx context.Context, book *model.Book) error {
	return m.db.WithContext(ctx).Where("id = ?", book.Model.ID).Updates(book).Error
}

func (m *mysqlBookRepository) GetBookByTitle(ctx context.Context, title string) (*model.Book, error) {
	var b model.Book
	if err := m.db.WithContext(ctx).Where("title = ?", title).First(&b).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &b, nil
}

func NewMySQLBookRepository(db *gorm.DB) model.BookRepository {
	return &mysqlBookRepository{db: db}
}
