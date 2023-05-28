package handlers

import (
	"bytes"
	"encoding/json"

	"github.com/abinashphulkonwar/go-cdn/storage"
	"github.com/gofiber/fiber/v2"
)

func ReadFileHandler(c *fiber.Ctx, storageSession *storage.Storage) error {
	path := c.BaseURL() + c.Path()
	buf, err := storageSession.GetFile(path)
	if err != nil {
		return err
	}
	meta, err := storageSession.GetMetaData(path)
	if err != nil {
		return err
	}

	if len(buf) == 0 || len(meta) == 0 {
		return c.Next()
	}
	metaData := make(map[string]string)

	err = json.Unmarshal(meta, &metaData)

	if err != nil {
		return err
	}

	for key, val := range metaData {
		c.Set(key, val)
	}

	if len(buf) < 148848 {
		return c.Send(buf)
	}

	c.Set(fiber.HeaderContentDisposition, "inline")

	return c.SendStream(bytes.NewReader(buf), len(buf))

}
