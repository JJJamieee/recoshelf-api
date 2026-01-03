package middlewares

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Protected(c *fiber.Ctx) error {
	// TODO replace with JWT auth
	userID := 1
	if headerVal := c.Get("X-User-Id"); headerVal != "" {
		if parsed, err := strconv.Atoi(headerVal); err == nil && parsed > 0 {
			userID = parsed
		}
	}
	c.Locals("userID", userID)

	return c.Next()
}
