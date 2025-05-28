package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Global variable to store the MongoDB collection reference
var ProductCollection *mongo.Collection

// ConnectDB establishes a connection to the MongoDB database
// and initializes the ProductCollection
func ConnectDB() {
	// Set a timeout context for the connection (10 seconds)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Initialize the collection reference for "products"
	ProductCollection = client.Database("product_catalog").Collection("products")
}
