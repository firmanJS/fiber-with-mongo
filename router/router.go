package router

import (
	"github.com/firmanjs/fiber-with-mongo/handler"
	// "github.com/firmanjs/go-restfull-with-fiber/middleware"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {

	// Middleware
	// api := app.Group("/api", middleware.AuthReq())
	api := app.Group("/api")

	// routes
	api.Get("/", handler.index)

}
