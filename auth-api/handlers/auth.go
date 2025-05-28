package handlers

import (
	"auth-api/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("your_secret_key") // Secret key used to sign JWT

// Login handles user authentication and returns a JWT token
func Login(c *fiber.Ctx) error {
	var user models.User

	// Parse the request body into the user struct
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Dummy credential check (hardcoded for demo)
	if user.Username != "admin" || user.Password != "password" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Create a new JWT token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 2).Unix(), // Token expires in 2 hours
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Return the signed token as JSON
	return c.JSON(fiber.Map{"token": tokenString})
}
