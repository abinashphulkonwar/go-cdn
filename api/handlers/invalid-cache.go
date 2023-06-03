package handlers

import (
	"github.com/abinashphulkonwar/go-cdn/storage"
	"github.com/gofiber/fiber/v2"
)

func InvalidCacheHandler(c *fiber.Ctx, storageSession *storage.Storage, origin string) error {
	return &fiber.Error{
		Code:    404,
		Message: "Invalid cache",
	}
}
