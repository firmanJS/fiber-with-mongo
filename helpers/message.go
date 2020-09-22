package helpers

import (
	"github.com/gofiber/fiber/v2"
)

type resMessage struct {
	code int
	status  string
	message string
}

func NotFound(c *fiber.Ctx) error {
	resNotFound := &resMessage{
		code:  404,
		status: "not found",
		message: "url not found",
	}
	return c.Status(404).JSON(resNotFound)
}