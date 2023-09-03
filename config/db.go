package config

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbDriver = "mysql"
	username = "root"
	password = "R@k1b+++"
	hostname = "127.0.0.1:3306"
	dbName   = "gotodo"
)

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func Connect() *sql.DB {
	db, err := sql.Open(dbDriver, dsn(""))

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Database connection Successful!")
	}

	err = db.Ping()

	if err != nil {
		fmt.Println("Database connection unsuccessful!")
	}

	_, err = db.Exec(`CREATE DATABASE IF NOT EXISTS gotodo;`)

	if err != nil {
		fmt.Println(err)
	}

	_, err = db.Exec(`USE gotodo`)

	if err != nil {
		fmt.Println(err)
	}
	return db
}

func CreateDB() error {
	db, err := sql.Open(dbDriver, dsn(""))

	if err != nil {
		fmt.Println(err)
	}

	// Connection pool

	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(1000)
	db.SetConnMaxLifetime(time.Hour)

	err = db.Ping()

	if err != nil {
		fmt.Println(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS gotodo.todos(
			id INT AUTO_INCREMENT,
			item TEXT NOT NULL,
			completed BOOLEAN DEFAULT FALSE,

			PRIMARY KEY(id)
		);
	`)

	if err != nil {
		fmt.Println(err, "Can't create table!")
	}

	return nil
}

func LogRequest(date time.Time) {
	db := Connect()

	_, err := db.Exec(`
		INSERT INTO `+"`request_logs`"+` (date, count)
		VALUES (?, 1)
		ON DUPLICATE KEY UPDATE count = count + 1;
	`, date.Format("2006-01-02"))

	if err != nil {
		fmt.Println("Error Logging Request!", err)
	} else {
		fmt.Println("Request Logged Successfully")
	}
}
