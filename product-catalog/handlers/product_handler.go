package handlers

import (
	"context"
	"time"

	"product-catalog/database"
	"product-catalog/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateProduct handles POST /products
// It parses the incoming JSON, stores the product in MongoDB, and returns the created product
func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	// Parse the JSON body into the product struct
	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).SendString("Invalid product data")
	}

	// Set a timeout context for MongoDB operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Insert the product into MongoDB
	result, err := database.ProductCollection.InsertOne(ctx, product)
	if err != nil {
		return c.Status(500).SendString("Failed to create product")
	}

	// Assign the generated MongoDB ID to the product
	product.ID = result.InsertedID.(primitive.ObjectID)

	// Return the created product with status 201
	return c.Status(201).JSON(product)
}

// GetProducts handles GET /products
// It retrieves all products from MongoDB and returns them as a JSON array
func GetProducts(c *fiber.Ctx) error {
	// Set a timeout context for MongoDB operation
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Query all documents from the products collection
	cursor, err := database.ProductCollection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(500).SendString("Failed to retrieve products")
	}
	defer cursor.Close(ctx)

	// Decode each document into the product slice
	var products []models.Product
	for cursor.Next(ctx) {
		var product models.Product
		if err := cursor.Decode(&product); err != nil {
			return c.Status(500).SendString("Failed to decode product")
		}
		products = append(products, product)
	}

	// Return the list of products as JSON
	return c.JSON(products)
}
