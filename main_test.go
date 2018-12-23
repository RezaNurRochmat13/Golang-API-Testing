package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ProductController unit test
func TestProductController(t *testing.T) {
	// Declare testing routes
	initializeRouteTest := GinEngine()

	// Declare HTTP Initialization
	recorderTesting := httptest.NewRecorder()
	requestTesting, _ := http.NewRequest("GET", "/api/v1/product", nil)
	initializeRouteTest.ServeHTTP(recorderTesting, requestTesting)

	// Test suites
	assert.Equal(t, 200, recorderTesting.Code)

}
