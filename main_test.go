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
	requestTesting, error := http.NewRequest("GET", "/api/v1/product", nil)
	initializeRouteTest.ServeHTTP(recorderTesting, requestTesting)

	// Test suites

	// Check error when request trigerred
	if error == nil {
		assert.Equal(t, http.StatusOK, recorderTesting.Code)
	} else {
		assert.Equal(t, http.StatusInternalServerError, recorderTesting.Code)
	}

}

// TestGetDetailProducts unit test
func TestGetDetailProducts(t *testing.T) {
	// Declare testing routes
	initializeRouteTest := GinEngine()

	// Declare HTTP Initialization
	recorderTesting := httptest.NewRecorder()
	requestTesting, error := http.NewRequest("GET", "/api/v1/product/:UUIDProducts", nil)
	initializeRouteTest.ServeHTTP(recorderTesting, requestTesting)

	// Test suites

	// Check error when request trigerred
	if error == nil {
		assert.Equal(t, http.StatusOK, recorderTesting.Code)
	} else {
		assert.Equal(t, http.StatusInternalServerError, recorderTesting.Code)
	}
}

// TestCreateProducts unit test
func TestCreateProducts(t *testing.T) {
	// Declare testing routes
	initializeRouteTest := GinEngine()

	// Declare HTTP Initialization
	recorderTesting := httptest.NewRecorder()
	requestTesting, error := http.NewRequest("POST", "/api/v1/product", nil)
	initializeRouteTest.ServeHTTP(recorderTesting, requestTesting)

	// Check error when request trigerred
	if error == nil {
		assert.Equal(t, http.StatusOK, recorderTesting.Code)
	} else {
		assert.Equal(t, http.StatusInternalServerError, recorderTesting.Code)
	}
}
