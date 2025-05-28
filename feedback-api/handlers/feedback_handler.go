package handlers

import (
	"feedback-api/models"
	"github.com/gofiber/fiber/v2"
)

var feedbacks []models.Feedback
var idCounter = 1

func CreateFeedback(c *fiber.Ctx) error {
	var newFeedback models.Feedback

	if err := c.BodyParser(&newFeedback); err != nil {
		return c.Status(400).SendString("Invalid JSON format")
	}

	newFeedback.ID = idCounter
	idCounter++
	feedbacks = append(feedbacks, newFeedback)

	return c.Status(201).JSON(newFeedback)

}

func GetFeedbacks(c *fiber.Ctx) error {
	return c.JSON(feedbacks)
}
