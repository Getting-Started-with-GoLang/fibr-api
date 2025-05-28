package main

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Proxy route to User Service
	app.All("/api/v1/users/*", func(c *fiber.Ctx) error {
		targetURL := "http://localhost:4000" + c.OriginalURL()

		// Read the body safely once
		body := c.BodyRaw()
		req, err := http.NewRequest(c.Method(), targetURL, bytes.NewReader(body))
		if err != nil {
			return c.Status(500).SendString("Failed to create request")
		}

		// Copy headers to the proxied request
		c.Request().Header.VisitAll(func(key, value []byte) {
			req.Header.Set(string(key), string(value))
		})

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return c.Status(502).SendString("Backend service unreachable")
		}
		defer resp.Body.Close()

		c.Status(resp.StatusCode)
		bodyResp, _ := io.ReadAll(resp.Body)
		return c.Send(bodyResp)
	})

	// Proxy route to Product Service
	app.All("/api/v1/products/*", func(c *fiber.Ctx) error {
		targetURL := "http://localhost:5000" + c.OriginalURL()

		body := c.BodyRaw()
		req, err := http.NewRequest(c.Method(), targetURL, bytes.NewReader(body))
		if err != nil {
			return c.Status(500).SendString("Failed to create request")
		}

		c.Request().Header.VisitAll(func(key, value []byte) {
			req.Header.Set(string(key), string(value))
		})

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return c.Status(502).SendString("Backend service unreachable")
		}
		defer resp.Body.Close()

		c.Status(resp.StatusCode)
		bodyResp, _ := io.ReadAll(resp.Body)
		return c.Send(bodyResp)
	})

	app.Listen(":3000")
}
