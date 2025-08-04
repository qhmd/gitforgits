package config

import "errors"

var (
	ErrBookTitleExists   = errors.New("book with this title already exists")
	ErrUserExists        = errors.New("email already used")
	ErrCategoryStillUsed = errors.New("category is still used by some book")
)
