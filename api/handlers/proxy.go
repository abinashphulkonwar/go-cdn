package handlers

import (
	"net/http/httputil"
	"net/url"

	"github.com/abinashphulkonwar/go-cdn/storage"
	"github.com/gofiber/fiber/v2"
)

func ProxyHandler(c *fiber.Ctx, storageSession *storage.Storage) error {
	targetURL := "https://departmentofpoliticalsciencehcgdcollege.azurewebsites.net"

	// Parse the target URL
	target, err := url.Parse(targetURL)
	if err != nil {
		return err
	}

	// Create a new reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(target)

	// Serve the reverse proxy request
	//	proxy.ServeHTTP(c.Response().(http.ResponseWriter), c.Request())

	return nil
}
