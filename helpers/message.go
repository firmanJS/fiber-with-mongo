package helpers

import (
	"github.com/gofiber/fiber/v2"
)

type resMessage struct {
	Code    int         `json:"code,omitempty"`
	Status  string      `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NotFound(c *fiber.Ctx) error {
	resNotFound := &resMessage{
		Code:    404,
		Status:  "not found",
		Message: "url not found",
		Data:    "null",
	}
	return c.JSON(resNotFound)
}

func ErrorResponse(msg string, c *fiber.Ctx) error {
	resErr := &resMessage{
		Code:    403,
		Status:  "Error",
		Message: msg,
		Data:    "null",
	}
	return c.JSON(resErr)
}

