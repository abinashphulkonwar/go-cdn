package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/abinashphulkonwar/go-cdn/api"
	"github.com/abinashphulkonwar/go-cdn/service"
	"github.com/gofiber/fiber/v2"
)

type requestBody struct {
	Key []string `json:"url"`
}

func TestInternale(t *testing.T) {

	const secretKey string = "token"
	const path string = "/api/in-validateCache"
	method := make(map[string]string)
	method["POST"] = "POST"
	app := api.App(service.Config{
		Origin: "https://departmentofpoliticalsciencehcgdcollege.azurewebsites.net",
		Method: method,
		Cache: service.CacheConfiguration{
			TTL:              3600,
			InvalidationPath: path,
			Token:            secretKey,
			SecretKey:        []byte(secretKey),
		},
	})

	token, err := service.GetJwtToken(nil, []byte(secretKey))

	if err != nil {
		t.Error(err)
	}

	if err != nil {
		t.Error(err)
	}

	key := make([]string, 0)
	key = append(key, "/0c96a2b89d97b85c12be966190b622a8")
	body := requestBody{
		Key: key,
	}

	data, _ := json.Marshal(body)

	req, err := http.NewRequest("POST", path, bytes.NewReader(data))
	if err != nil {
		t.Error(err)
	}
	req.Header.Add(fiber.HeaderAuthorization, "Bearer "+token)

	res, err := app.Test(req, 10000)
	if err != nil {
		t.Error(err)
	}

	println(res.StatusCode)

}
