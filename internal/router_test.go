package internal

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouters(t *testing.T) {
	r := Routers()

	// Create test request and recorder
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	// Serve the request through the router
	r.ServeHTTP(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)
}
