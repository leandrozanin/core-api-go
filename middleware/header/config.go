package header

import (
	"github.com/gofiber/fiber/v2"
)

// Config defines the config for middleware.
type Config struct {
	Name   *string
	Domain *string
	Next   func(c *fiber.Ctx) bool
}
