package handlers

import (
	"github.com/abinashphulkonwar/go-cdn/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func ProxyHandler(c *fiber.Ctx, storageSession *storage.Storage) error {
	targetURL := "https://departmentofpoliticalsciencehcgdcollege.azurewebsites.net"
	if c.Method() == fiber.MethodGet {
		return &fiber.Error{
			Message: "ProxyHandler: Method not allowed",
			Code:    fiber.StatusMethodNotAllowed,
		}
	}
	if err := proxy.Do(c, targetURL); err != nil {
		return err
	}

	c.Response().Header.Del(fiber.HeaderServer)

	return nil
}
