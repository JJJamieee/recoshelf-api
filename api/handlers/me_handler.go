package handlers

import (
	"log/slog"

	"recoshelf-api/api/requests"
	"recoshelf-api/pkg/app_errors"
	"recoshelf-api/pkg/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetUserReleases(releaseService services.ReleaseService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID, ok := c.Locals("userID").(int)
		if !ok {
			slog.Error("Invalid user ID.", "userID", userID)
			return app_errors.InvalidUserIDError
		}

		releases, err := releaseService.GetUserReleases(int64(userID))
		if err != nil {
			slog.Error("Get user releases failed.", "msg", err.Error())
			return app_errors.InternalServerError
		}
		return c.JSON(releases)
	}
}

func CreateUserRelease(releaseService services.ReleaseService, validate *validator.Validate) fiber.Handler {
	return func(c *fiber.Ctx) error {
		releaseRequest := new(requests.ReleaseRequest)
		if err := c.BodyParser(releaseRequest); err != nil {
			return app_errors.InvalidRequestBodyError(err)
		}

		if errs := validate.Struct(releaseRequest); errs != nil {
			return app_errors.InvalidRequestBodyError(errs)
		}

		userID, ok := c.Locals("userID").(int)
		if !ok {
			slog.Error("Invalid user ID.", "userID", userID)
			return app_errors.InvalidUserIDError
		}

		release := releaseRequest.ToEntity()
		err := releaseService.CreateUserRelease(int64(userID), release)
		if err != nil {
			slog.Error("Create user releases failed.", "msg", err.Error())
			return app_errors.InternalServerError
		}

		return c.SendStatus(201)
	}
}
