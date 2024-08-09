package main

type Book struct {
	Id              int    `json:"id"`
	Title           string `json:"title"`
	Synopsis        string `json:"synopsis"`
	PublicationDate int    `json:"publicationDate"`
	ISBN10          string `json:"isbn10"`
	ISBN13          string `json:"isbn13"`
	// Authors         []Author `json:"authors"`
	// Genres          []Genre  `json:"genres"`
}
