package handler

import (
	"github.com/gofiber/fiber/v2"
)

// GetAllProducts from db
func Index(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"status": "success", 
		"message": "Welcome",
	})
}