package eventRoutes

import (
	"github.com/gofiber/fiber/v2"
	eventHandlers "github.com/tobiadiks/event-api/internal/handlers/event"
)

func SetupEventRoutes(router fiber.Router) {
	event := router.Group("/event")
	// Create a Event
	event.Post("/", eventHandlers.CreateEvents)
	// Read all Events
	event.Get("/", eventHandlers.GetEvents)
	// Read one Event
	event.Get("/:eventId", eventHandlers.GetEvent)
	// Update one Event
	event.Put("/:eventId", eventHandlers.UpdateEvent)
	// Delete one Event
	event.Delete("/:eventId", eventHandlers.DeleteEvent)

}
