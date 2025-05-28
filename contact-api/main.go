// Activity 2 - Contact API microservice using Go Fiber

package main

import (
	"contact-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create a new Fiber app instance
	app := fiber.New()

	// Route to handle contact creation
	app.Post("/contacts", handlers.CreateContact)

	// Route to fetch all stored contacts
	app.Get("/contacts", handlers.GetContacts)

	// Start the Fiber server on port 3000
	if err := app.Listen(":3000"); err != nil {

		panic(err)
	}
}
