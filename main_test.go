package main

import (
	"Golang-API-Testing/model"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ProductController unit test
func TestGetAllBooks(t *testing.T) {
	gin.SetMode(gin.TestMode)
	// Declare testing routes
	initializeRouteTest := GinEngine()

	// Declare HTTP Initialization
	recorderTesting := httptest.NewRecorder()
	requestTesting, err := http.NewRequest("GET", "/api/v1/books", nil)
	initializeRouteTest.ServeHTTP(recorderTesting, requestTesting)

	// Test suites
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, recorderTesting.Code)

}

// TestGetDetailProducts unit test
func TestGetDetailBooks(t *testing.T) {
	gin.SetMode(gin.TestMode)
	// Declare testing routes
	initializeRouteTest := GinEngine()

	// Declare HTTP Initialization
	recorderTesting := httptest.NewRecorder()
	requestTesting, err := http.NewRequest("GET", "/api/v1/books/:UUIDBooks", nil)
	initializeRouteTest.ServeHTTP(recorderTesting, requestTesting)

	// Test suites
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, recorderTesting.Code)
}

// TestCreateProducts unit test
func TestCreateBooks(t *testing.T) {
	gin.SetMode(gin.TestMode)
	// Declare testing routes
	initializeRouteTest := GinEngine()

	// Mocking Request
	booksMockStruct := model.BooksModel{
		UUIDBooks:       uuid.UUID{},
		BookName:        "Pendekar kesepian",
		BookWriter:      "Yusron F",
		BookPublisher:   "Reja Pub",
		BookDescription: "Reja itu ganteng",
	}

	requestBytes, _ := json.Marshal(booksMockStruct)
	requestReader := bytes.NewReader(requestBytes)

	// Declare HTTP Initialization
	recorderTesting := httptest.NewRecorder()
	requestTesting, err := http.NewRequest("POST", "/api/v1/books", requestReader)
	initializeRouteTest.ServeHTTP(recorderTesting, requestTesting)

	// when body is not found
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, recorderTesting.Code)

}
