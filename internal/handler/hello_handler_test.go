package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHelloHandler(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// Define routes
	router.GET("/hello", HelloHandler)

	// Create a mock HTTP request
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a recorder for the response
	recorder := httptest.NewRecorder()

	// Serve the HTTP request to the mock server
	router.ServeHTTP(recorder, req)

	// Assert the status code is 200 OK
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Assert the response body contains the expected message
	expectedResponse := `{"message":"Hello, World!"}`
	assert.Equal(t, expectedResponse, recorder.Body.String())
}
