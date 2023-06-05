package main

import (
	"github.com/abinashphulkonwar/go-cdn/api"
	"github.com/abinashphulkonwar/go-cdn/service"
	"github.com/abinashphulkonwar/go-cdn/storage"
	"github.com/abinashphulkonwar/go-cdn/worker"
)

func main() {
	configuration := service.Configuration()

	service.InitDir(storage.TempDir)
	service.InitDir(storage.MetaDir)

	app := api.App(configuration)
	app.Listen(":3004")

	go worker.Worker(storage.MetaDir)
}
