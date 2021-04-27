package router

import (
	"github.com/firmanJS/fiber-with-mongo/handler"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes func
func EmployeRoutes(api fiber.Router) {

	employe := api.Group("/employe")
	employe.Get("/", handler.GetEmploye)

}
