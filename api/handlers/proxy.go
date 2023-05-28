package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/abinashphulkonwar/go-cdn/storage"
	"github.com/gofiber/fiber/v2"
)

func Proxy(c *fiber.Ctx, storageSession *storage.Storage) error {
	//	url := "http://localhost:3001" + c.Path()
	url := "https://abinashphulkonwar.vercel.app/_next/image?url=%2Fbanner%2F1.png&w=1080&q=75"
	res, err := http.Get(url)

	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return err
	}
	key := c.BaseURL() + c.Path()
	err = storageSession.WriteFile(key, body)
	if err != nil {
		return err
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
	storageSession.SetMetaData(key, jsonData)
	c.Set("Content-type", headers["Content-Type"])
	return c.Send(body)
}
