package handlers

import (
	"database/sql"
	"stream-it-consulting/innovation-team/example-rest-api/repositories"
	"github.com/gofiber/fiber/v2"
)


type (
	handlers struct {
		bookrepo repositories.IBookRepository
	}
	Handlers interface{
		GetBooks(c *fiber.Ctx) error
		GetBooksByID(c *fiber.Ctx)error
		CreateBook(c *fiber.Ctx)error
		Delete(c *fiber.Ctx)error
	}
)

func NewHandlers(dbConn *sql.DB) handlers{
	return handlers{
		bookrepo: repositories.NewBookRepository(dbConn),
	}
}