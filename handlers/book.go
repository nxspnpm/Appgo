package handlers

import (
	"stream-it-consulting/innovation-team/example-rest-api/models"

	"github.com/gofiber/fiber/v2"
)

// type (
// 	handlers struct {
// 		bookrepo repositories.IBookRepository
// 	}
// 	Handlers interface{
// 		GetBooks(c *fiber.Ctx) error
// 		GetBooksByID(c *fiber.Ctx)error
// 		CreateBook(c *fiber.Ctx)error
// 		Delete(c *fiber.Ctx)error
// 	}
// )

// func NewHandlers(dbConn *sql.DB) handlers{
// 	return handlers{
// 		bookrepo: repositories.NewBookRepository(dbConn),
// 	}
// }

func (h handlers) CreateBook(c *fiber.Ctx) error {
	var book *models.Book

	if err := c.BodyParser(&book); err != nil {
		return err
	}

	newBook, err := h.bookrepo.CreateBook(book)
	if err != nil {
		return err
	}

	return c.JSON(newBook)
}
