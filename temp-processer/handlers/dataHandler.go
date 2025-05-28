package handlers

import (
	"log"
	"net/http"
	"temp-processer/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

// ReceiveData handles incoming POST requests with JSON-formatted temperature data
func ReceiveData(c *fiber.Ctx) error {
	var reading models.TemperatureReading

	// Parse JSON body into the TemperatureReading struct
	if err := c.BodyParser(&reading); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON data",
		})
	}

	// Add timestamp if missing
	if reading.Timestamp.IsZero() {
		reading.Timestamp = time.Now()
	}

	// Example processing: log high temperature warning
	if reading.Temperature > 40.0 {
		log.Printf("⚠️ High temperature alert! Sensor: %s, Temp: %.2f°C", reading.SensorID, reading.Temperature)
	} else {
		log.Printf("Received: Sensor %s - %.2f°C at %s", reading.SensorID, reading.Temperature, reading.Timestamp)
	}

	// Respond with success
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Data received",
	})
}
