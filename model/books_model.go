package model

import "github.com/satori/go.uuid"

// BooksModel struct
type Books struct {
	UUIDBooks        uuid.UUID
	BookName         string
	BookWriter       string
	BookPublisher    string
	BookDescription  string
	BookCategoryName string
}

type Book struct {
	UUIDBooks        uuid.UUID
	BookName         string
	BookWriter       string
	BookPublisher    string
	BookDescription  string
	UUIDBookCategory string
}
