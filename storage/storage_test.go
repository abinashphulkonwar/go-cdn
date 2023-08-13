package storage_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/abinashphulkonwar/go-cdn/storage"
)

func TestWriteFile(t *testing.T) {
	client := storage.New("../temp/", "../meta/")

	client.WriteFile("test", []byte("test"), storage.MediaType.CSS)
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

func TestMinifyCss(t *testing.T) {
	// D:\development\web__apps\golang\cdn\storage\
	data, err := os.ReadFile("./4bf0c6aba6185af641b2bec87ff6ab81")
	if err != nil {
		t.Error(err)
	}

	client := storage.New("../temp/", "../meta/")

	err = client.WriteFile("data", data, storage.MediaType.CSS)

	if err != nil {
		t.Error(err)
	}

}
