package main

import (
	"auth-api/handlers"
	"auth-api/middleware"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create a new Fiber app
	app := fiber.New()

	// Public route for login. Returns a JWT token if credentials are valid.
	app.Post("/login", handlers.Login)

	// Protected route for dashboard. Requires a valid JWT token to access.
	app.Get("/dashboard", middleware.JWTMiddleware, func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to the dashboard!",
		})
	})

	// Start the server on port 3000
	app.Listen(":3000")
}
