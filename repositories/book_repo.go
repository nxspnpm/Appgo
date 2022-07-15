package repositories

import (
	"database/sql"
	"stream-it-consulting/innovation-team/example-rest-api/models"
)

type (
	bookRepository struct {
		db *sql.DB
	}
)

func NewBookRepository(dbConn *sql.DB) IBookRepository {
	return bookRepository{db: dbConn}
}

func (r bookRepository) CreateBook(book *models.Book) (*models.Book, error) {
	query := `INSERT INTO books (title, author, year) VALUES ($1, $2, $3)
              RETURNING ID;`
	err := r.db.QueryRow(query, book.Title, book.Author, book.Year).Scan(&book.ID)
	if err != nil {
		return nil, err
	}
	return book, nil
}
