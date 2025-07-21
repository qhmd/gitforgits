package book

import "errors"

var (
	ErrBookTitleExists = errors.New("book with this title already exists")
)
