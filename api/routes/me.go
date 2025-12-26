package routes

import (
	"recoshelf-api/api/handlers"
	"recoshelf-api/api/middlewares"
	"recoshelf-api/pkg/repositories"
	"recoshelf-api/pkg/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func MeRouter(app fiber.Router, db *sqlx.DB, validate *validator.Validate) {
	releaseRepository := repositories.NewReleaseRepo(db)
	releaseService := services.NewReleaseService(releaseRepository)

	me := app.Group("/me", middlewares.Protected)

	me.Get("/releases", handlers.GetUserReleases(releaseService))
	me.Post("/releases", handlers.CreateUserRelease(releaseService, validate))
	me.Delete("/releases/:releaseID", handlers.DeleteUserRelease(releaseService))
}
