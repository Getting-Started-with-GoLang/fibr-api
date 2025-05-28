// Activity 5 - File Server with Access Control using Go Fiber

package main

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create a new Fiber app
	app := fiber.New()

	// Serve static files from the "public" directory
	app.Static("/", "./public")

	// Middleware to restrict downloads to only .txt files
	restrictTxtOnly := func(c *fiber.Ctx) error {
		filename := c.Params("filename")

		// Extract the file extension and convert to lowercase
		ext := strings.ToLower(filepath.Ext(filename))

		// Only allow .txt file downloads
		if ext != ".txt" {
			return c.Status(fiber.StatusForbidden).SendString("Only .txt files are allowed")
		}
		return c.Next()
	}

	// Download route with access control middleware
	app.Get("/download/:filename", restrictTxtOnly, func(c *fiber.Ctx) error {
		filename := c.Params("filename")

		// Serve the requested file for download
		return c.Download("./public/" + filename)
	})

	// Start the server on port 3000 and log errors if any
	log.Fatal(app.Listen(":3000"))
}
