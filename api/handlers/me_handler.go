package handlers

import (
	"errors"
	"log/slog"
	"strconv"

	"recoshelf-api/api/presenter"
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

func DeleteUserRelease(releaseService services.ReleaseService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID, ok := c.Locals("userID").(int)
		if !ok {
			slog.Error("Invalid user ID.", "userID", userID)
			return app_errors.InvalidUserIDError
		}

		releaseIDParam := c.Params("releaseID")
		releaseID, err := strconv.ParseInt(releaseIDParam, 10, 64)
		if err != nil || releaseID <= 0 {
			if err == nil {
				err = errors.New("releaseID must be a positive integer")
			}
			return app_errors.InvalidRequestBodyError(err)
		}

		err = releaseService.DeleteUserRelease(int64(userID), releaseID)
		if err != nil {
			slog.Error("Delete user releases failed.", "msg", err.Error())
			return app_errors.InternalServerError
		}

		return c.SendStatus(204)
	}
}

func BatchDeleteUserReleases(releaseService services.ReleaseService, validate *validator.Validate) fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(requests.BatchDeleteReleaseRequest)
		if err := c.BodyParser(request); err != nil {
			return app_errors.InvalidRequestBodyError(err)
		}

		if errs := validate.Struct(request); errs != nil {
			return app_errors.InvalidRequestBodyError(errs)
		}

		userID, ok := c.Locals("userID").(int)
		if !ok {
			slog.Error("Invalid user ID.", "userID", userID)
			return app_errors.InvalidUserIDError
		}

		deletedCount, err := releaseService.BatchDeleteUserReleases(int64(userID), request.IDs)
		if err != nil {
			slog.Error("Batch delete user releases failed.", "msg", err.Error())
			return app_errors.InternalServerError
		}

		return c.JSON(presenter.BatchDeleteResponse{
			DeletedCount: deletedCount,
		})
	}
}
