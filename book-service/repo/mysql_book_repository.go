package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/qhmd/gitforgits/book-service/config"
	"github.com/qhmd/gitforgits/book-service/model"
	"gorm.io/gorm"
)

type mysqlBookRepository struct {
	db *gorm.DB
}

// GetBookByCategory implements model.BookRepository.
func (m *mysqlBookRepository) GetBookByCategory(ctx context.Context, id int) ([]*model.Book, error) {
	var book []*model.Book
	err := m.db.WithContext(ctx).Model(book).Where("category_id = ?", id).Find(&book).Error
	return book, err
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

// CreateCategory implements model.BookRepository.
func (m *mysqlBookRepository) CreateCategory(ctx context.Context, category *model.Category) error {
	return m.db.WithContext(ctx).Create(category).Error
}

// DeleteCategory implements model.BookRepository.
func (m *mysqlBookRepository) DeleteCategory(ctx context.Context, id int) error {
	var count int64
	m.db.WithContext(ctx).Model(&model.Book{}).Where("category_id = ?", id).Count(&count)
	fmt.Print("jumalhnya", count)
	if count > 0 {
		return config.ErrCategoryStillUsed
	}

	result := m.db.WithContext(ctx).Delete(&model.Category{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// ListCategory implements model.BookRepository.
func (m *mysqlBookRepository) ListCategory(ctx context.Context) ([]*model.Category, error) {
	var c []*model.Category
	err := m.db.WithContext(ctx).Find(&c).Error
	return c, err
}

// UpdateCategory implements model.BookRepository.
func (m *mysqlBookRepository) UpdateCategory(ctx context.Context, category *model.Category) error {
	return m.db.WithContext(ctx).Where("id = ?", category.Model.ID).Updates(category).Error
}

func (m *mysqlBookRepository) GetCategoryByID(ctx context.Context, id int) error {
	var c *model.Category
	if err := m.db.WithContext(ctx).Where("id = ?", id).First(&c).Error; err != nil {
		return err
	}
	return nil
}

// GetCategory implements model.BookRepository.
func (m *mysqlBookRepository) GetCategory(ctx context.Context, category string) error {
	var c *model.Category
	if err := m.db.WithContext(ctx).Where("name = ?", category).First(&c).Error; err != nil {
		return err
	}
	return nil
}

func NewMySQLBookRepository(db *gorm.DB) model.BookRepository {
	return &mysqlBookRepository{db: db}
}
