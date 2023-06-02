package worker_test

import (
	"testing"

	"github.com/abinashphulkonwar/go-cdn/storage"
	"github.com/abinashphulkonwar/go-cdn/worker"
)

func TestWorker(t *testing.T) {
	worker.MetaDataCheck("../" + storage.MetaDir)
}
