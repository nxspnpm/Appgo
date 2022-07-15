package repositories

import (
	"stream-it-consulting/innovation-team/example-rest-api/models"
)

type (
	IBookRepository interface{
		CreateBook(book *models.Book) (*models.Book, error)
	}
)
