package internal

import "github.com/gofiber/fiber/v2"

const (
	SUCCESS, PARTIAL, ERROR = "SUCCESS", "PARTIAL", "ERROR"
)

func ResponseHandler(c *fiber.Ctx, status string, statusCode int, message string, processedItems []string) error {
	c.Status(statusCode)

	return c.JSON(fiber.Map{
		"status":  status,
		"items":   processedItems,
		"message": message,
	})
}
