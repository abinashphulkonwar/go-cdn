package main

import (
	"github.com/abinashphulkonwar/go-cdn/api"
	"github.com/abinashphulkonwar/go-cdn/service"
	"github.com/abinashphulkonwar/go-cdn/storage"
)

func main() {
	service.InitDir(storage.TempDir)
	service.InitDir(storage.MetaDir)
	api.App()
}
