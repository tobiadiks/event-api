package router

import (
	"github.com/gofiber/fiber/v2"
	eventRoutes "github.com/tobiadiks/event-api/internal/routes/event"
)

// Setups all routes
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	eventRoutes.SetupEventRoutes(api)
}
