package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgres() (*pgxpool.Pool, error) {
	connStr := os.Getenv("DATABASE_URL")
	dbpool, err := pgxpool.New(context.Background(), connStr)

	if err != nil {
		return nil, err
	}

	var greeting string
	err = dbpool.QueryRow(context.Background(), "SELECT 'Hello world!'").Scan(&greeting)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Query row failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)

	return dbpool, nil
}
