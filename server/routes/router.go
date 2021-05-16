package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads", http.FileServer(http.Dir("uploads"))))

	router.HandleFunc("/api/authors", getAllAuthors).Methods("GET")
	router.HandleFunc("/api/authors/{id}", getAuthorById).Methods("GET")
	router.HandleFunc("/api/authors/username/{username}", getAuthorByUsername).Methods("GET")
	router.HandleFunc("/api/authors/{id}/fans", getAuthorFans).Methods("GET")
	router.HandleFunc("/api/author", getAuthorByToken).Methods("GET")
	router.HandleFunc("/api/register", register).Methods("POST")
	router.HandleFunc("/api/login", login).Methods("POST")
	router.HandleFunc("/api/authors/{id}", updateAuthor).Methods("PUT")
	router.HandleFunc("/api/authors/{id}/pfp", updateAuthorPfp).Methods("PUT")
	router.HandleFunc("/api/authors/{id}/toggle-fan/{fanId}", toggleAuthorFan).Methods("PUT")
	router.HandleFunc("/api/authors/{id}", deleteAuthor).Methods("DELETE")

	router.HandleFunc("/api/books", getAllBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBookById).Methods("GET")
	router.HandleFunc("/api/books/author/{id}", getBooksByAuthorId).Methods("GET")
	router.HandleFunc("/api/books/liked/{id}", getLikedBooksByAuthorId).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}/add-read", addReadToBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}/toggle-like/{authorId}", toggleBookLike).Methods("PUT")
	router.HandleFunc("/api/books/{id}/cover", updateBookCover).Methods("PUT")
	router.HandleFunc("/api/books/{id}/pdf", updateBookPdf).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	return router
}
