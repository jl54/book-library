package main

import (
	"encoding/json"
	"net/http"
)

var books = []Book{
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
	Data []Book `json:"data"`
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

func registerRoutes() {
	http.HandleFunc("/books", getBooks)
}
