package config

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	// Setup mock database
	db, mock, _ := sqlmock.New()
	defer db.Close()

	mock.ExpectPing()

	// Call the Connect function with the mock database
	connDB := Connect()

	// Check if the returned DB matches the mock DB
	assert.NotNil(t, connDB)

	// Ensure the mock expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCreateDB(t *testing.T) {
	// Setup mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// Add expectations for executed SQL
	mock.ExpectExec("CREATE DATABASE gotodo").
		WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec("USE gotodo").
		WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec("CREATE TABLE IF NOT EXISTS").
		WillReturnResult(sqlmock.NewResult(0, 0))

	// Call the CreateDB function with the mock database
	err = CreateDB()
	if err != nil {
		t.Fatal(err)
	}

	// Ensure the mock expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
