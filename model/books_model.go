package model

import "github.com/satori/go.uuid"

// BooksModel struct
type BooksModel struct {
	UUIDBooks        uuid.UUID
	BookName         string
	BookWriter       string
	BookPublisher    string
	BookDescription  string
	BookCategoryName string
}
