package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads", http.FileServer(http.Dir("uploads"))))

	router.HandleFunc("/api/authors", getAllAuthors).Methods("GET")
	router.HandleFunc("/api/register", register).Methods("POST")

	return router
}
