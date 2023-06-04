package middleware

import (
	"strings"

	"github.com/abinashphulkonwar/go-cdn/api/internal"
	"github.com/abinashphulkonwar/go-cdn/service"
	"github.com/gofiber/fiber/v2"
)

func IsAuthenticated(c *fiber.Ctx, config service.Config) error {
	token := c.Get("Authorization")

	if token == "" {
		return internal.ResponseHandler(c, internal.ERROR, fiber.StatusUnauthorized, "Unauthorized", nil)
	}

	tokenArray := strings.Split(token, " ")

	if len(tokenArray) != 2 {
		return internal.ResponseHandler(c, internal.ERROR, fiber.StatusUnauthorized, "Unauthorized", nil)
	}

	if tokenArray[0] != "Bearer" {
		return internal.ResponseHandler(c, internal.ERROR, fiber.StatusUnauthorized, "Unauthorized", nil)
	}

	user, err := service.VerifyToken(tokenArray[1], []byte(config.Cache.Token))

	if err != nil {
		return internal.ResponseHandler(c, internal.ERROR, fiber.StatusUnauthorized, err.Error(), nil)
	}

	c.Locals("user", user)

	return c.Next()
}
