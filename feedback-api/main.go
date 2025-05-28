package main

import (
	"feedback-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/feedbacks", handlers.CreateFeedback)
	app.Get("feedbacks", handlers.GetFeedbacks)

	app.Listen(":3000")
}
