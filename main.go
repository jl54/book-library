package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SimpleResponse struct {
	Status int    `json:"status"`
	Data   string `json:"data"`
}

func main() {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	defer dbpool.Close()

	var greeting string

	err = dbpool.QueryRow(context.Background(), "SELECT 'Hello World!'").Scan(&greeting)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)

	registerRoutes()

	// ListenAndServe always returns an error,
	// since it only returns when an unexpected error occurs.
	// In order to log that error we wrap the
	// function call with log.Fatal
	log.Fatal(http.ListenAndServe(":8080", nil))
}
