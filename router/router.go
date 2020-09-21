package router

import (
	"github.com/firmanJS/fiber-with-mongo/handler"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {

	// Middleware
	// api := app.Group("/api", middleware.AuthReq())
	api := app.Group("/api")

	// routes
	api.Get("/", handler.Index)

}
