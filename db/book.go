package db

import (
	"book-library/book"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BookDB struct {
	DbPool *pgxpool.Pool
}

func (bookDb *BookDB) Insert(item book.Book) error {
	query := `
		INSERT INTO books (title, author) VALUES (@title, @author)
	`

	args := pgx.NamedArgs{
		"title":  item.Title,
		"author": item.Author,
	}

	_, err := bookDb.DbPool.Exec(context.Background(), query, args)

	if err != nil {
		return err
	}

	return nil
}

func (bookDB *BookDB) Select(id int) (book.Book, error) {
	query := `
		SELECT title, author FROM books WHERE id = @id
	`

	args := pgx.NamedArgs{
		"id": id,
	}

	var bookFromDb book.Book
	err := bookDB.DbPool.QueryRow(context.Background(), query, args).Scan(&bookFromDb)

	if err != nil {
		return book.Book{}, err
	}

	return bookFromDb, nil
}

func (bookDB *BookDB) Update(id int, item book.Book) error {
	query := `
		UPDATE books SET title = @title, author = @author WHERE id = @id
	`

	args := pgx.NamedArgs{
		"title":  item.Title,
		"author": item.Author,
	}

	_, err := bookDB.DbPool.Exec(context.Background(), query, args)

	if err != nil {
		return err
	}

	return nil
}

func (bookDb *BookDB) Delete(id int) error {
	query := `
		DELETE FROM books WHERE id = @id
	`

	args := pgx.NamedArgs{
		"id": id,
	}

	_, err := bookDb.DbPool.Exec(context.Background(), query, args)

	if err != nil {
		return err
	}

	return nil
}

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
