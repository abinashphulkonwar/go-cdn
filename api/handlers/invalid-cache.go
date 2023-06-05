package handlers

import (
	"encoding/json"
	"strings"

	"github.com/abinashphulkonwar/go-cdn/api/internal"
	"github.com/abinashphulkonwar/go-cdn/storage"
	"github.com/gofiber/fiber/v2"
)

type requestBody struct {
	Key []string `json:"url"`
}

func InvalidCacheHandler(c *fiber.Ctx, storageSession *storage.Storage, origin string) error {

	body := c.Body()
	if len(body) == 0 {
		return internal.ResponseHandler(c, internal.ERROR, fiber.StatusBadRequest, "request body is empty", nil)
	}

	var reqBody requestBody

	err := json.Unmarshal(body, &reqBody)
	if err != nil {
		return internal.ResponseHandler(c, internal.ERROR, fiber.StatusBadRequest, err.Error(), nil)
	}

	if len(reqBody.Key) == 0 {
		return internal.ResponseHandler(c, internal.SUCCESS, fiber.StatusOK, "", nil)
	}

	processedItems := make([]string, 0)

	for _, key := range reqBody.Key {

		if !strings.HasPrefix(key, "/") {
			key = "/" + key
		}

		err := storageSession.Delete(c.BaseURL() + key)

		if err != nil {
			return internal.ResponseHandler(c, internal.PARTIAL, fiber.StatusInternalServerError, err.Error(), processedItems)
		}

		processedItems = append(processedItems, key)

	}

	return internal.ResponseHandler(c, internal.SUCCESS, fiber.StatusOK, "", processedItems)

}
