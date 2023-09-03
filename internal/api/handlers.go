package api

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/nihankhan/go-todo/config"
	"github.com/nihankhan/go-todo/models"
)

var (
	id        int
	item      string
	completed int

	view = template.Must(template.ParseFiles("./templates/index.html"))
	db   = config.Connect()
	_    = config.CreateDB()
)

func Index(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Query(`SELECT * FROM gotodo.todos`)

	if err != nil {
		fmt.Println(err)
	}

	var todos []models.Todo

	for stmt.Next() {
		err = stmt.Scan(&id, &item, &completed)

		if err != nil {
			fmt.Println(err)
		}

		todo := models.Todo{
			Id:        id,
			Item:      item,
			Completed: completed,
		}

		todos = append(todos, todo)
	}

	requestCount := GetRequestCount()

	data := models.View{
		Todos:        todos,
		RequestCount: requestCount,
	}

	_ = view.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	item := r.FormValue("item")

	tx, err := db.Begin()

	if err != nil {
		log.Fatal(err)
	}

	_, err = tx.Exec(`INSERT INTO todos (item) VALUE(?)`, item)

	if err != nil {
		log.Fatal(err)
		tx.Rollback()
	}

	err = tx.Commit()

	if err != nil {
		log.Fatal(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	tx, err := db.Begin()

	if err != nil {
		log.Fatal(err)
	}

	_, err = tx.Exec(`DELETE FROM todos WHERE id = ?`, id)

	if err != nil {
		log.Fatal(err)
		tx.Rollback()
	}

	err = tx.Commit()

	if err != nil {
		log.Fatal(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Complete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	tx, err := db.Begin()

	if err != nil {
		log.Fatal(err)
	}

	_, err = tx.Exec(`UPDATE todos SET completed = 1 WHERE id = ?`, id)

	if err != nil {
		log.Fatal(err)

		tx.Rollback()
	}

	err = tx.Commit()

	if err != nil {
		log.Fatal(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func GetRequestCount() int {
	var count int

	//q := `SELECT COALESCE(SUM(count), 0) FROM ` + "`request_logs`"

	q := `SELECT COALESCE(SUM(count), 0) FROM gotodo.request_logs`

	//log.Println(q)

	err := db.QueryRow(q).Scan(&count)

	if err != nil {
		log.Println("Error Fetching request count!", err)
	} else {
		log.Println("Successfully Fetching request count.")
	}

	return count
}

func RequestCountHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)

	if !ok {
		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
		return
	}

	ctx := r.Context()

	ticker := time.NewTicker(time.Second)

	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return

		case <-ticker.C:

			requestCount := GetRequestCount()

			log.Printf("Sending Request Count: %d\n", requestCount)

			fmt.Fprintf(w, "data: %d\n\n", requestCount)

			flusher.Flush()
		}
	}
}
