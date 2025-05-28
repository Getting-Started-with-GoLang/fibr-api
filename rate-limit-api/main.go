package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func main() {
	// Create a new Fiber app
	app := fiber.New()

	// Apply rate limiter middleware globally
	app.Use(limiter.New(limiter.Config{
		// Allow maximum 5 requests
		Max: 5,
		// Set the time window to 10 seconds
		Duration: 10 * time.Second,
		// Use client IP address as the key for rate limiting
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
	}))

	// Define a GET endpoint at /api
	app.Get("/api", func(c *fiber.Ctx) error {
		// Return a JSON response
		return c.JSON(fiber.Map{
			"message": "Welcome to the rate-limited API!",
		})
	})

	// Start the Fiber app on port 3000
	app.Listen(":3000")
}
