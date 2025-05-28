package main

import (
	"log"
	"temp-processer/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	// Middleware for logging requests
	app.Use(logger.New())

	// Route to handle incoming temperature data
	app.Post("/data", handlers.ReceiveData)

	// Start the server on port 3000
	log.Println("Temperature Processor is running on port 3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
