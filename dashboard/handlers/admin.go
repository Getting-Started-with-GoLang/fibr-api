package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	users := []fiber.Map{
		{"id": 1, "name": "John Doe"},
		{"id": 2, "name": "Jane Smith"},
	}
	return c.JSON(users)
}
func GetStats(c *fiber.Ctx) error {
	stats := fiber.Map{
		"total_users":  100,
		"active_users": 75,
	}
	return c.JSON(stats)
}

func GetLogs(c *fiber.Ctx) error {
	logs := []string{
		"User John Doe logged in",
		"User Jane Smith logged out",
	}
	return c.JSON(logs)
}
