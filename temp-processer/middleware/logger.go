package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

type LoggerMiddleware struct{}

func (l *LoggerMiddleware) Handle(c *fiber.Ctx) error {
	start := time.Now()
	method := c.Method()
	path := c.OriginalURL()

	err := c.Next()
	timeTaken := time.Since(start)

	fmt.Printf("[LOGGER] %s %s - %v\n", method, path, timeTaken)
	return err
}
func New() fiber.Handler {
	logger := &LoggerMiddleware{}
	return logger.Handle

}
