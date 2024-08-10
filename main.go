package main

import (
	"book-library/db"
	"fmt"
	"log"
	"net/http"
	"os"
)

type SimpleResponse struct {
	Status int    `json:"status"`
	Data   string `json:"data"`
}

func main() {
	dbpool, err := db.NewPostgres()

	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	defer dbpool.Close()

	registerRoutes(dbpool)

	// ListenAndServe always returns an error,
	// since it only returns when an unexpected error occurs.
	// In order to log that error we wrap the
	// function call with log.Fatal
	log.Fatal(http.ListenAndServe(":8080", nil))
}
