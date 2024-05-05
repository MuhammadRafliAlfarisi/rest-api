package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	userRoutes "rest-api/internal/routes"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	userRoutes.SetupUserRoutes(api)
}
