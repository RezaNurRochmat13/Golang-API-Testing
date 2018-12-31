package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ProductController unit test
func TestGetAllProducts(t *testing.T) {
	// Declare testing routes
	initializeRouteTest := GinEngine()

	// Declare HTTP Initialization
	recorderTesting := httptest.NewRecorder()
	requestTesting, _ := http.NewRequest("GET", "/api/v1/books", nil)
	initializeRouteTest.ServeHTTP(recorderTesting, requestTesting)

	// Test suites
	assert.Equal(t, http.StatusOK, recorderTesting.Code)

}

// TestGetDetailProducts unit test
func TestGetDetailProducts(t *testing.T) {
	// Declare testing routes
	initializeRouteTest := GinEngine()

	// Declare HTTP Initialization
	recorderTesting := httptest.NewRecorder()
	requestTesting, _ := http.NewRequest("GET", "/api/v1/books/:UUIDProducts", nil)
	initializeRouteTest.ServeHTTP(recorderTesting, requestTesting)

	// Test suites
	assert.Equal(t, http.StatusOK, recorderTesting.Code)
}

// TestCreateProducts unit test
func TestCreateProducts(t *testing.T) {
	// Declare testing routes
	initializeRouteTest := GinEngine()

	// Declare HTTP Initialization
	recorderTesting := httptest.NewRecorder()
	requestTesting, _ := http.NewRequest("POST", "/api/v1/books", nil)
	initializeRouteTest.ServeHTTP(recorderTesting, requestTesting)

	// Check error when request trigerred
	assert.Equal(t, http.StatusOK, recorderTesting.Code)
}
