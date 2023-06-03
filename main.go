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
	api.App(configuration)
	go worker.Worker(storage.MetaDir)
}
