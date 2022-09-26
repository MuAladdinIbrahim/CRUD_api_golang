package routes

import (
	handlers "project/pkg/handlers"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(router *mux.Router) {
	s := router.PathPrefix("/products").Subrouter()

	s.HandleFunc("/", handlers.CreateProductHandler).Methods("POST")
	// s.HandleFunc("/{id}").Methods("GET")
	// s.HandleFunc("/{id}").Methods("PUT")
	// s.HandleFunc("/").Methods("GET")
}
