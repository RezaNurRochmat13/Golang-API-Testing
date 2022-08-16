package controller

import (
	"golang-api-testing/config"
	"golang-api-testing/model"
	"net/http"

	uuid "github.com/satori/go.uuid"

	"github.com/gin-gonic/gin"
)

// Global database config
var initializeDatabaseConnection = config.DatabaseConn()

// GetAllBooks func
func GetAllBooks(c *gin.Context) {

	var (
		books      []model.Books
		countBooks int64
	)

	initializeDatabaseConnection.Table("books").
		Select("books.uuid_books, books.book_name, books.book_writer, books.book_publisher," +
			"books.book_description, book_category.book_category_name").
		Joins("INNER JOIN book_category ON book_category.uuid_book_category=books.uuid_book_category").
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

	var books model.Books

	if err := initializeDatabaseConnection.First(&books, UUIDBooks).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Not found"})
	} else {
		initializeDatabaseConnection.Table("book").
			Select("books.uuid_books, books.book_name, books.book_writer, books.book_publisher," +
				"books.book_description, book_category.book_category_name").
			Joins("INNER JOIN book_category ON book_category.uuid_book_category=books.uuid_book_category").
			Find(&books)

		c.JSON(http.StatusOK, gin.H{"data": books})
	}
}

// CreateNewProducts func does create new record product
func CreateNewBooks(c *gin.Context) {

	UUIDBooks := uuid.NewV4()
	BookName := c.Param("BookName")
	BookWriter := c.Param("BookWriter")
	BookPublisher := c.Param("BookPublisher")
	BookDescription := c.Param("BookDescription")
	UUIDBookCategory := c.Param("UUIDBookCategory")

	createBookPayload := model.Book{
		UUIDBooks:        UUIDBooks,
		BookName:         BookName,
		BookWriter:       BookWriter,
		BookPublisher:    BookPublisher,
		BookDescription:  BookDescription,
		UUIDBookCategory: UUIDBookCategory,
	}

	c.BindJSON(&createBookPayload)

	initializeDatabaseConnection.Create(createBookPayload)

	c.JSON(http.StatusOK, gin.H{"message": "Inserted sucessfully"})
}

// UpdateProducts func does update record product
func UpdateBooks(c *gin.Context) {

	var booksModel model.Book

	UUIDBooks := c.Param("UUIDBooks")
	BookName := c.Param("BookName")
	BookWriter := c.Param("BookWriter")
	BookPublisher := c.Param("BookPublisher")
	BookDescription := c.Param("BookDescription")
	UUIDBookCategory := c.Param("UUIDBookCategory")

	updateBookPayload := model.Book{
		BookName:         BookName,
		BookWriter:       BookWriter,
		BookPublisher:    BookPublisher,
		BookDescription:  BookDescription,
		UUIDBookCategory: UUIDBookCategory,
	}

	c.BindJSON(&updateBookPayload)

	if err := initializeDatabaseConnection.First(&booksModel, UUIDBooks).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Data not found"})
	} else {
		initializeDatabaseConnection.Model(&booksModel).Updates(updateBookPayload)

		c.JSON(http.StatusOK, gin.H{"message": "Updated successfully"})
	}
}

// DeleteProducts func does delete record product
func DeleteBooks(c *gin.Context) {

	var booksModel model.Book

	UUIDBooks := c.Param("UUIDBooks")

	if err := initializeDatabaseConnection.First(&booksModel, UUIDBooks).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Data not found"})
	} else {
		initializeDatabaseConnection.Where("books.uuid_books = ?", UUIDBooks).
			Find(&booksModel).Delete(&booksModel)

		c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
	}
}

func ExampleTesting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"total": 1,
		"data":  "Example",
		"count": 1})
}

func GetPeople(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "People"})
}
