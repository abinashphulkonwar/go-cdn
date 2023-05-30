package storage_test

import (
	"fmt"
	"testing"

	"github.com/abinashphulkonwar/go-cdn/storage"
)

func TestWriteFile(t *testing.T) {
	client := storage.New("../temp/", "../meta/")

	client.WriteFile("test", []byte("test"))
}

func TestReadFile(t *testing.T) {
	var factor int64 = 1
	var unit string

	_, err := fmt.Sscanf("KB10000000kb", "%d%s", &factor, &unit)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(factor, unit, "🚀")
	storage.ReadChunks()
}
