package helpers

import (
	"github.com/gofiber/fiber/v2"
)

type resMessage struct {
	Code    int         `json:"code,omitempty"`
	Status  bool      `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponseMsg(c *fiber.Ctx, code int, status bool, msg string, data interface{}) error {
	resPonse := &resMessage{
		Code:    code,
		Status:  status,
		Message: msg,
		Data:    data,
	}
	return c.JSON(resPonse)
}
