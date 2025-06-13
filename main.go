package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tobiadiks/event-api/database"
	"github.com/tobiadiks/event-api/router"
)

func main() {
	app := fiber.New()

	database.ConnectDB()

	app.Get("/", func(c *fiber.Ctx) error {
		err:=c.SendString("Hello world!")
		return err
	})

	router.SetupRoutes(app)

	app.Listen(":8080")
}
