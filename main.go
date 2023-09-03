package main

import (
	"log"
	"net/http"
	"os"

	"github.com/nihankhan/go-todo/internal"
)

func main() {

	logFile, err := os.OpenFile("server.logs", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer logFile.Close()

	log.SetOutput(logFile)

	r := internal.Routers()

	r.Use(internal.RequestLogger)

	log.Println("Server Running on 127.0.0.1:8080")

	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: r,
	}

	done := make(chan bool)

	go func() {
		err := server.ListenAndServe()

		if err != nil {
			log.Fatal(err)
		}

		done <- true
	}()

	<-done
}
