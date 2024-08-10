package main

import (
	"book-library/book"
	"book-library/db"
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

var books = []book.Book{
	{
		Id:    1,
		Title: "Mistborn: The Final Empire",
	},
	{
		Id:    2,
		Title: "Mistborn: The Well of Ascension",
	},
	{
		Id:    3,
		Title: "Mistborn: The Hero of Ages",
	},
}

type BookResponse struct {
	Data []book.Book `json:"data"`
}

type BookHandler struct {
	dbpool *pgxpool.Pool
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	response := BookResponse{
		Data: books,
	}

	w.Header().Add("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (bh *BookHandler) createBook(w http.ResponseWriter, r *http.Request) {
	bookStore := db.PgBookStore{}

	var newBook book.Book
	var err error
	err = json.NewDecoder(r.Body).Decode(&newBook)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = bookStore.InsertBookIntoDatabase(bh.dbpool, newBook)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}

func registerRoutes(dbpool *pgxpool.Pool) {
	bookHandler := &BookHandler{
		dbpool: dbpool,
	}

	http.HandleFunc("/books", getBooks)
	http.HandleFunc("/create", bookHandler.createBook)
}
