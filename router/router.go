package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {

	api := app.Group("/api", logger.New())

	// routes
	// api.Get("/", handler.Index)

	// UserRoutes(api)
	EmployeRoutes(api)

}
