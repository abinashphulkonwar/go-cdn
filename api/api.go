package api

import (
	"crypto/tls"

	"github.com/abinashphulkonwar/go-cdn/api/handlers"
	"github.com/abinashphulkonwar/go-cdn/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func App() {
	app := fiber.New()
	proxy.WithTlsConfig(&tls.Config{
		InsecureSkipVerify: true,
	})
	storageSession := storage.New("temp/", "meta/")
	app.Get("/docs/:id", func(c *fiber.Ctx) error {
		return handlers.ReadFileHandler(c, storageSession)
	}, func(c *fiber.Ctx) error {
		return handlers.Proxy(c, storageSession)
	})

	app.Listen(":3004")
}
