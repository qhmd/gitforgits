package dto

type BookRequest struct {
	Title       string  `json:"title" validate:"required,min=3,max=100" example:"How To Become Backend Engineer"`
	Author      string  `json:"author" validate:"required,min=4,max=50,alphaunicode" example:"John Smith"`
	Publisher   string  `json:"publisher" validate:"required,min=3,max=100" example:"TechBooks Publishing"`
	ReleaseDate string  `json:"release_date" validate:"required,datetime=2006-01-02" example:"2025-08-01"` // format YYYY-MM-DD
	Language    string  `json:"language" validate:"required,alphaunicode" example:"English"`
	Pages       int     `json:"pages" validate:"required,gt=0" example:"205"`
	Format      string  `json:"format" validate:"required,oneof=PDF EPUB MOBI" example:"PDF"`
	Description string  `json:"description" validate:"required,min=10,max=1000" example:"This book helps you become a backend engineer from scratch."`
	Price       float64 `json:"price" validate:"required,gt=0" example:"99.000"`
	FileURL     string  `json:"file_url" validate:"required,url" example:"https://example.com/files/book.pdf"`
	Thumbnail   string  `json:"thumbnail" validate:"required,url" example:"https://example.com/images/book-cover.jpg"`
	CategoryID  uint    `json:"category_id" validate:"required,gt=0" example:"1"`
}
