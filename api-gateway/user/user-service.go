package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.All("/api/v1/users/*", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"source": "User Service",
			"path":   c.OriginalURL(),
		})
	})
	fmt.Println("User Service is running on port 4000")
	app.Listen(":4000")
}
