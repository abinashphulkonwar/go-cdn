package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/abinashphulkonwar/go-cdn/storage"
	"github.com/gofiber/fiber/v2"
)

func ProxyGet(c *fiber.Ctx, storageSession *storage.Storage, origin string) error {

	url := origin + c.OriginalURL()
	if c.Method() != fiber.MethodGet {
		return c.Next()
	}

	req, err := http.NewRequest(c.Method(), url, nil)

	if err != nil {
		return err
	}
	for key, value := range c.GetReqHeaders() {
		if key == fiber.HeaderAccept || key == fiber.HeaderAcceptEncoding || key == fiber.HeaderAcceptLanguage {
			continue
		}
		req.Header.Set(key, value)
	}

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return err
	}

	status := res.StatusCode

	key := c.BaseURL() + c.Path()

	headers := make(map[string]string)
	headers["Content-Type"] = res.Header.Get("Content-Type")
	headers["Content-Length"] = res.Header.Get("Content-Length")
	headers["Date"] = res.Header.Get("Date")
	headers["Content-Encoding"] = res.Header.Get("Content-Disposition")
	headers["Cache-Control"] = res.Header.Get("Cache-Control")

	jsonData, err := json.Marshal(headers)
	if err != nil {
		println(err)
		return err
	}

	isCahced := true

	if res.Header.Get(fiber.HeaderCacheControl) == "no-cache" {
		isCahced = false
	}

	if status == 200 && isCahced {
		ContentType, isFound := headers["Content-Type"]

		if isFound {
			data := strings.Split(ContentType, ";")
			if data[0] != "" {
				ContentType = data[0]
			}

		}

		err = storageSession.WriteFile(key, body, ContentType)
		if err != nil {
			return err
		}
		storageSession.SetMetaData(key, jsonData)

	}
	c.Set("Content-type", headers["Content-Type"])
	c.Response().Header.Del(fiber.HeaderServer)
	return c.Send(body)
}
