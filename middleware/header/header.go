package header

import (
	"github.com/gofiber/fiber/v2"
)

func New(config Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Append("x-link", *config.Name)
		return c.Next()
	}
}
