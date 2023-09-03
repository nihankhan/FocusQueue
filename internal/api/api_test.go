package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	// Setup mock database
	db, mock, _ := sqlmock.New()
	defer db.Close()

	mock.ExpectQuery("SELECT \\* FROM gotodo.todos").WillReturnRows(sqlmock.NewRows([]string{"id", "item", "completed"}).AddRow(1, "Task 1", 0))

	// Create test request and recorder
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	// Call the Index function with the mock database
	Index(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Ensure the mock expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestAdd(t *testing.T) {
	// Setup mock database
	db, mock, _ := sqlmock.New()
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO todos").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Create test request and recorder
	reqBody := strings.NewReader("item=test")
	req := httptest.NewRequest("POST", "/add", reqBody)
	w := httptest.NewRecorder()

	// Set up a mock router with Gorilla Mux for the Add function
	r := mux.NewRouter()
	r.HandleFunc("/add", Add)

	// Serve the request through the mock router
	r.ServeHTTP(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusSeeOther, w.Code)

	// Ensure the mock expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// Similar tests for Delete and Complete functions
