package db

import (
	"book-library/book"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BookStore interface {
	InsertBookIntoDatabase(book.Book) error
	GetBookById(int) (*book.Book, error)
	GetAllBooks() (*[]book.Book, error)
	UpdateBook(int, book.Book) error
	DeleteBook(int) error
}

type PgBookStore struct{}

func (bs *PgBookStore) InsertBookIntoDatabase(dbpool *pgxpool.Pool, book book.Book) error {
	query := `
		INSERT INTO books (title, author) VALUES (@title, @author)
	`

	args := pgx.NamedArgs{
		"title":  book.Title,
		"author": book.Author,
	}

	_, err := dbpool.Exec(context.Background(), query, args)

	if err != nil {
		return err
	}

	return nil
}
