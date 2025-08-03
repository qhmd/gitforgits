package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string   `json:"title" example:"How To Become Backend Engineer"`
	Author      string   `json:"author" example:"John Smith"`
	Publisher   string   `json:"publisher" example:"TechBooks Publishing"`
	ReleaseDate string   `json:"release_date" example:"2025-08-01"` // YYYY-MM-DD
	Language    string   `json:"language" example:"English"`
	Pages       int      `json:"pages" example:"205"`
	Format      string   `json:"format" example:"PDF"`
	Description string   `json:"description" example:"This book helps you become a backend engineer from scratch."`
	Price       float64  `json:"price" example:"12.99"`
	FileURL     string   `json:"file_url" example:"https://example.com/files/book.pdf"`
	Thumbnail   string   `json:"thumbnail" example:"https://example.com/images/book-cover.jpg"`
	CategoryID  uint     `json:"category_id" example:"1"`
	Category    Category `gorm:"foreignKey:CategoryID" json:"category,omitzero"`
}

type Category struct {
	gorm.Model
	Name  string `json:"name" example:"technology"`
	Slug  string `json:"slug" example:"technology"`
	Books []Book `gorm:"foreignKey:CategoryID" json:"books,omitempty"`
}
