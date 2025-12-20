package handlers

import (
	"log/slog"

	"recoshelf-api/pkg/app_errors"
	"recoshelf-api/pkg/services"

	"github.com/gofiber/fiber/v2"
)

func GetUserReleases(releaseService services.ReleaseService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID, ok := c.Locals("userID").(int)
		if !ok {
			slog.Error("Invalid user ID.", "userID", userID)
			return app_errors.InvalidUserIDError
		}

		releases, err := releaseService.GetUserReleases(userID)
		if err != nil {
			slog.Error("Get user releases failed.", "msg", err.Error())
			return app_errors.InternalServerError
		}
		return c.JSON(releases)
	}
}
