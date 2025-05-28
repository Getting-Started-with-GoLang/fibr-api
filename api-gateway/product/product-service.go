package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.All("/api/v1/products/*", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"source": "Product Service",
			"path":   c.OriginalURL(),
		})
	})

	fmt.Println("Product Service is running on port 5000")
	app.Listen(":5000")
}
