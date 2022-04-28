package helper

import "github.com/gofiber/fiber/v2"

func ResponseSuccess(c *fiber.Ctx, code int, data interface{}) error {
	return c.Status(code).JSON(fiber.Map{
		"status": "success",
		"code":   code,
		"data":   data,
	})
}

func ResponseFailed(c *fiber.Ctx, code int, message error) error {
	return c.Status(code).JSON(fiber.Map{
		"status":  "failed",
		"code":    code,
		"message": message.Error(),
	})
}
