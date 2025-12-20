package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func Protected(c *fiber.Ctx) error {
	// TODO replace with JWT auth
	c.Locals("userID", 1)

	return c.Next()
}
