// Activity 4 - Real-time WebSocket communication using Go Fiber

package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func main() {
	// Create a new Fiber app
	app := fiber.New()

	// Middleware to check if the incoming request is a valid WebSocket upgrade
	app.Use("ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return c.SendStatus(fiber.StatusUpgradeRequired)
	})

	// WebSocket endpoint at /ws
	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		var (
			msg []byte
			err error
		)

		// Infinite loop to keep connection open and handle messages
		for {
			// Read message from the client
			if _, msg, err = c.ReadMessage(); err != nil {
				break // Exit on read error or disconnect
			}

			// Handle the message
			if string(msg) == "Hello server" {
				// Respond to known message
				c.WriteMessage(websocket.TextMessage, []byte("Hello Thisara"))
			} else {
				// Respond to any other message
				c.WriteMessage(websocket.TextMessage, []byte("Unknown message"))
			}
		}
	}))

	// Start the WebSocket server on port 3000
	app.Listen(":3000")
}
