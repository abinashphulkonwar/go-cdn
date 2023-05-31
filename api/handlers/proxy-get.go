package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/abinashphulkonwar/go-cdn/storage"
	"github.com/gofiber/fiber/v2"
)

func ProxyGet(c *fiber.Ctx, storageSession *storage.Storage) error {

	url := "https://departmentofpoliticalsciencehcgdcollege.azurewebsites.net" + c.OriginalURL()

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
	if status == 200 {
		err = storageSession.WriteFile(key, body)
		if err != nil {
			return err
		}
	}
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
	if status == 200 {
		storageSession.SetMetaData(key, jsonData)
	}
	c.Set("Content-type", headers["Content-Type"])
	return c.Send(body)
}
