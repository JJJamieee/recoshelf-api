package routes

import (
	"recoshelf-api/api/handlers"
	"recoshelf-api/api/middlewares"
	"recoshelf-api/pkg/repositories"
	"recoshelf-api/pkg/services"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func MeRouter(app fiber.Router, db *sqlx.DB) {
	releaseRepository := repositories.NewReleaseRepo(db)
	releaseService := services.NewReleaseService(releaseRepository)

	me := app.Group("/me", middlewares.Protected)

	me.Get("/releases", handlers.GetUserReleases(releaseService))
}
