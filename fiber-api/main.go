// Activity 1 - Simple RESTful API with CRUD operations

package main

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// User struct represents a user model with ID, Name, and Email fields
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// list to store users
var users = []User{}

// assign unique IDs to users
var idCounter = 1

func main() {
	// Create a new Fiber app
	app := fiber.New()

	// Route to create a new user
	app.Post("/users", func(c *fiber.Ctx) error {
		newUser := new(User)

		// Parse request body into newUser struct
		if err := c.BodyParser(newUser); err != nil {
			return c.Status(400).SendString("Invalid input")
		}

		// Assign ID and add user to the list
		newUser.ID = idCounter
		idCounter++
		users = append(users, *newUser)

		// Return the newly created user as JSON
		return c.JSON(newUser)
	})

	// Route to get all users
	app.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(users)
	})

	// Route to get a single user by ID
	app.Get("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		// Loop through users to find a match
		for _, user := range users {
			if fmt.Sprintf("%d", user.ID) == id {
				return c.JSON(user)
			}
		}
		return c.Status(404).SendString("User not found")
	})

	// Route to update a user by ID
	app.Put("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		updatedUser := new(User)

		// Parse request body into updatedUser struct
		if err := c.BodyParser(updatedUser); err != nil {
			return c.Status(400).SendString("Invalid input")
		}

		// Find user and update the fields
		for i, user := range users {
			if fmt.Sprintf("%d", user.ID) == id {
				users[i].Name = updatedUser.Name
				users[i].Email = updatedUser.Email
				return c.JSON(users[i])
			}
		}
		return c.Status(404).SendString("User not found")
	})

	// Route to delete a user by ID
	app.Delete("/users/:id", func(c *fiber.Ctx) error {
		// Convert ID from string to integer
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).SendString("Invalid user ID")
		}

		// Find and remove the user from the list
		for i, user := range users {
			if user.ID == id {
				users = append(users[:i], users[i+1:]...)
				return c.SendString("User deleted successfully")
			}
		}
		return c.Status(404).SendString("User not found")
	})

	// Start the server on port 3000
	app.Listen(":3000")
}
