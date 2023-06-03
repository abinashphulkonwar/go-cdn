package handlers

import (
	"github.com/abinashphulkonwar/go-cdn/service"
	"github.com/gofiber/fiber/v2"
)

func MethodCheckHandler(c *fiber.Ctx, config service.Config) error {

	method := c.Method()

	_, isFound := config.Method[method]

	if !isFound {
		return &fiber.Error{
			Code:    404,
			Message: "Method not found",
		}
	}

	return c.Next()

}
