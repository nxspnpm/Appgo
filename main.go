package main

import (
	"stream-it-consulting/innovation-team/example-rest-api/database"
	"stream-it-consulting/innovation-team/example-rest-api/config"
	"stream-it-consulting/innovation-team/example-rest-api/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadConfig()

	dbConn := database.Initialize()
	defer dbConn.Close()

	app := fiber.New()

	router.HTTPRootRoutes(app)
	router.HTTPBookRoutes(app, dbConn)
	app.Listen(":3000")
}
