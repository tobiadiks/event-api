package eventHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/tobiadiks/event-api/database"
	"github.com/tobiadiks/event-api/internal/model"
)

func GetEvents(c *fiber.Ctx) error {
	db := database.DB
	var events []model.Event

	// find all events in the database
	db.Find(&events)

	// If no event is present return an error
	if len(events) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No events present", "data": nil})
	}

	// Else return events
	return c.JSON(fiber.Map{"status": "success", "message": "Events Found", "data": events})
}

func CreateEvents(c *fiber.Ctx) error {
	db := database.DB
	event := new(model.Event)

	// Store the body in the event and return error if encountered
	err := c.BodyParser(event)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	// Add a uuid to the event
	event.ID = uuid.New()
	// Create the Event and return error if encountered
	err = db.Create(&event).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create event", "data": err})
	}

	// Return the created event
	return c.JSON(fiber.Map{"status": "success", "message": "Created Event", "data": event})
}

func GetEvent(c *fiber.Ctx) error {
	db := database.DB
	var event model.Event

	// Read the param eventId
	id := c.Params("eventId")

	// Find the event with the given Id
	db.Find(&event, "id = ?", id)

	// If no such event present return an error
	if event.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No event present", "data": nil})
	}

	// Return the event with the Id
	return c.JSON(fiber.Map{"status": "success", "message": "Events Found", "data": event})
}

func UpdateEvent(c *fiber.Ctx) error {
	type updateEvent struct {
		Title    string `json:"title"`
		SubTitle string `json:"sub_title"`
		Date     string `json:"Text"`
	}
	db := database.DB
	var event model.Event

	// Read the param eventId
	id := c.Params("eventId")

	// Find the event with the given Id
	db.Find(&event, "id = ?", id)

	// If no such event present return an error
	if event.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No event present", "data": nil})
	}

	// Store the body containing the updated data and return error if encountered
	var updateEventData updateEvent
	err := c.BodyParser(&updateEventData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Edit the event
	event.Title = updateEventData.Title
	event.SubTitle = updateEventData.SubTitle
	event.Date = updateEventData.Date

	// Save the Changes
	db.Save(&event)

	// Return the updated event
	return c.JSON(fiber.Map{"status": "success", "message": "Events Found", "data": event})
}

func DeleteEvent(c *fiber.Ctx) error {
	db := database.DB
	var event model.Event

	// Read the param eventId
	id := c.Params("eventId")

	// Find the event with the given Id
	db.Find(&event, "id = ?", id)

	// If no such event present return an error
	if event.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No event present", "data": nil})
	}

	// Delete the event and return error if encountered
	err := db.Delete(&event, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete event", "data": nil})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Deleted Event"})
}