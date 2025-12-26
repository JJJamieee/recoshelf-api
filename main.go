package main

import (
	"errors"
	"log/slog"
	"os"

	"recoshelf-api/api/routes"
	"recoshelf-api/pkg/app_errors"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file.", "msg", err.Error())
		os.Exit(1)
	}

	// Database connection
	db, err := sqlx.Connect("mysql", os.Getenv("MYSQL_DSN"))
	if err != nil {
		slog.Error("DB connection failed.", "msg", err.Error())
		os.Exit(1)
	}
	defer db.Close()

	validate := validator.New()

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			var appErr *app_errors.AppError

			if errors.As(err, &appErr) {
				return c.Status(appErr.Status).JSON(fiber.Map{
					"message": appErr.Message,
					"code":    appErr.Code,
				})
			}

			// fallback
			slog.Error("Internal server error", "msg", err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal server error",
				"code":    "INTERNAL_SERVER_ERROR",
			})
		},
	})

	app.Get("/", func(c *fiber.Ctx) error {
		var greeting string
		err := db.QueryRow("SELECT 'Hello, World!'").Scan(&greeting)
		if err != nil {
			return err
		}
		return c.SendString(greeting)
	})

	v1 := app.Group("/v1")
	routes.MeRouter(v1, db, validate)

	app.Listen(":3000")
}
