package handlers

import (
	"contact-api/models"
	"github.com/gofiber/fiber/v2"
)

// Slice to store contact entries in memory
var contacts []models.Contact

// Counter to assign unique IDs to contacts
var idCounter = 1

// CreateContact handles POST /contacts
// It parses the incoming JSON, assigns an ID, and stores the contact
func CreateContact(c *fiber.Ctx) error {
	var newContact models.Contact

	// Parse the JSON body into newContact
	if err := c.BodyParser(&newContact); err != nil {
		return c.Status(400).SendString("Invalid JSON format")
	}

	// Assign a unique ID to the new contact
	newContact.ID = idCounter
	idCounter++

	// Add the contact to the in-memory slice
	contacts = append(contacts, newContact)

	// Return the created contact as JSON with status 201
	return c.Status(201).JSON(newContact)
}

// GetContacts handles GET /contacts
// It returns all stored contacts as a JSON array
func GetContacts(c *fiber.Ctx) error {
	return c.JSON(contacts)
}
