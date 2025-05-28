package main

import (
	"dashboard/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Define routes for the dashboard
	app.Get("/admin/users", handlers.GetUsers)
	app.Get("/admin/stats", handlers.GetStats)
	app.Get("/admin/logs", handlers.GetLogs)

	// Start the server
	app.Listen(":3000")

}
