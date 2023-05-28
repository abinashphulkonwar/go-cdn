package storage_test

import (
	"testing"

	"github.com/abinashphulkonwar/go-cdn/storage"
)

func TestWriteFile(t *testing.T) {
	client := storage.New("../temp/", "../meta/")

	client.WriteFile("test", []byte("test"))
}
