package handlers

import (
	"bytes"

	"github.com/abinashphulkonwar/go-cdn/storage"
	"github.com/gofiber/fiber/v2"
)

func ReadFileHandler(c *fiber.Ctx, storageSession *storage.Storage) error {
	path := c.BaseURL() + c.OriginalURL()
	buf, err := storageSession.GetFile(path)
	if err != nil {
		return err
	}
	meta, isFound, err := storageSession.GetMetaData(path)
	if err != nil {
		return err
	}

	if len(buf) == 0 || !isFound {
		return c.Next()
	}

	if err != nil {
		return err
	}

	if meta.CacheControl != "" {
		c.Set(fiber.HeaderCacheControl, meta.CacheControl)
	}

	if meta.ContentType != "" {
		c.Set(fiber.HeaderContentType, meta.ContentType)
	}

	if meta.ContentLength != "" {
		c.Set(fiber.HeaderContentLength, meta.ContentLength)
	}

	c.Set(fiber.HeaderContentDisposition, "inline")
	c.Response().Header.Del(fiber.HeaderServer)
	return c.SendStream(bytes.NewReader(buf), len(buf))

}
