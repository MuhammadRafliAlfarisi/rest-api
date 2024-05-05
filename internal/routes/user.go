package userRoutes

import (
	userHandler "rest-api/internal/handler/user"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(router fiber.Router) {
	user := router.Group("/user")
	// Create a user
	user.Post("/", userHandler.Create)
	// Read all users
	user.Get("/", userHandler.Get)
	// // Read one user
	user.Get("/:id_user", userHandler.GetUser)
	// // Update one user
	user.Put("/:id_user", userHandler.Update)
	// // Delete one user
	user.Delete("/:id_user", userHandler.Delete)
}
