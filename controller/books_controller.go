package controller

import (
	"Golang-API-Testing/config"
	"Golang-API-Testing/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllBooks func
func GetAllBooks(c *gin.Context) {
	initializeDatabaseConnection := config.DatabaseConn()

	var (
		books      []model.BooksModel
		countBooks int
	)

	initializeDatabaseConnection.Table("book").
		Select("book.uuid_books, book.book_name, book.book_writer, book.book_publisher," +
			"book.book_description, book_category.book_category_name").
		Joins("INNER JOIN book_category ON book_category.uuid_book_category=book.uuid_book_category").
		Find(&books).Count(&countBooks)

	c.JSON(http.StatusOK, gin.H{
		"total": countBooks,
		"data":  books,
		"count": countBooks})
}

// GetDetailProducts func does get detail products
func GetDetailBooks(c *gin.Context) {
	initializeDatabaseConnection := config.DatabaseConn()

	UUIDBooks := c.Param("UUIDBooks")

	var books model.BooksModel

	if initializeDatabaseConnection.Table("book").Where("book.uuid_books = ?", UUIDBooks).Find(&books).RecordNotFound() {
		c.JSON(http.StatusNoContent, gin.H{"message": "Not found"})
	} else {
		initializeDatabaseConnection.Table("book").
			Select("book.uuid_books, book.book_name, book.book_writer, book.book_publisher," +
				"book.book_description, book_category.book_category_name").
			Joins("INNER JOIN book_category ON book_category.uuid_book_category=book.uuid_book_category").
			Find(&books)

		c.JSON(http.StatusOK, gin.H{"data": books})
	}
}

// CreateNewProducts func does create new record product
func CreateNewBooks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Inserted sucessfully"})
}

// UpdateProducts func does update record product
func UpdateBooks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Updated successfully"})
}

// DeleteProducts func does delete record product
func DeleteBooks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}
