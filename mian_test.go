package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nihankhan/go-todo/internal"
	"github.com/stretchr/testify/assert"
)

func TestMainServer(t *testing.T) {
	r := internal.Routers()

	// Create a test request and recorder
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	// Serve the request through the router
	r.ServeHTTP(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)
}
