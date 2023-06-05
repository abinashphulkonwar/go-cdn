package api

import (
	"github.com/abinashphulkonwar/go-cdn/api/handlers"
	"github.com/abinashphulkonwar/go-cdn/api/middleware"
	"github.com/abinashphulkonwar/go-cdn/service"
	"github.com/abinashphulkonwar/go-cdn/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func App(config service.Config) *fiber.App {
	app := fiber.New()
	origin := config.Origin
	storageSession := storage.New(storage.TempDir+"/", storage.MetaDir+"/")

	app.Use(logger.New())

	if config.Cache.InvalidationPath != "" {
		app.Post(config.Cache.InvalidationPath, func(c *fiber.Ctx) error {
			return middleware.IsAuthenticated(c, config)
		}, func(c *fiber.Ctx) error {
			return handlers.InvalidCacheHandler(c, storageSession, origin)
		})
	}

	app.All("*",
		func(c *fiber.Ctx) error {
			return middleware.MethodCheckHandler(c, config)
		},
		func(c *fiber.Ctx) error {
			return handlers.ReadFileHandler(c, storageSession)
		}, func(c *fiber.Ctx) error {
			return handlers.ProxyGet(c, storageSession, origin)
		},
		func(c *fiber.Ctx) error {
			return handlers.ProxyHandler(c, storageSession, origin)
		},
	)

	return app
}
