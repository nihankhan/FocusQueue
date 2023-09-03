package internal

import (
	"net/http"
	"time"

	"github.com/nihankhan/go-todo/config"
)

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		currentTime := time.Now().Format("2006-01-02")

		parsedTime, _ := time.Parse("2006-01-02", currentTime)

		next.ServeHTTP(w, r)

		config.LogRequest(parsedTime)
	})
}
