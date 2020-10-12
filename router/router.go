package router

import (
	"github.com/firmanJS/fiber-with-mongo/handler/users"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {

	api := app.Group("/api", logger.New())

	// routes
	api.Get("/", handler.Index)

	users := api.Group("/users")
	users.Get("/", handler.GetUsers)
	users.Post("/", handler.CreateUsers)
	users.Delete("/:id", handler.DeleteUsers)
	users.Get("/:id", handler.GetByIdUsers)
	users.Put("/:id", handler.UpdateUsers)

}
