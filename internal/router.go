package internal

import (
	"github.com/gorilla/mux"

	"github.com/nihankhan/go-todo/internal/api"
)

func Routers() *mux.Router {
	route := mux.NewRouter().StrictSlash(true)

	route.HandleFunc("/", api.Index)
	route.HandleFunc("/add", api.Add).Methods("POST")
	route.HandleFunc("/delete/{id}", api.Delete)
	route.HandleFunc("/complete/{id}", api.Complete)
	route.HandleFunc("/request-count", api.RequestCountHandler)

	return route
}
