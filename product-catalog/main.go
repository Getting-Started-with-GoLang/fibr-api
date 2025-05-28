// Activity 3 - Product Catalog Microservice with MongoDB using Go Fiber

package main

import (
	"product-catalog/database" // Handles MongoDB connection setup
	"product-catalog/handlers" // Contains handler functions for product routes

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Connect to MongoDB and initialize the Product collection
	database.ConnectDB()

	// Create a new Fiber app
	app := fiber.New()

	// Route to create a new product (POST /products)
	app.Post("/products", handlers.CreateProduct)

	// Route to get all products (GET /products)
	app.Get("/products", handlers.GetProducts)

	// Start the server on port 3000
	app.Listen(":3000")
}
