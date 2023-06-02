package worker

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/abinashphulkonwar/go-cdn/storage"
	"github.com/jasonlvhit/gocron"
)

type CacheControl struct {
	MaxAge         int
	NoCache        bool
	NoStore        bool
	NoTransform    bool
	MustRevalidate bool
}

func parseCacheControl(headers string) CacheControl {
	cacheControl := CacheControl{}

	// Get the Cache-Control header value
	cacheControlHeader := headers

	// Split the header value into individual directives
	directives := strings.Split(cacheControlHeader, ",")

	// Parse each directive and set the corresponding field in the struct
	for _, directive := range directives {
		directive = strings.TrimSpace(directive)

		switch {
		case directive == "no-cache":
			cacheControl.NoCache = true
		case directive == "no-store":
			cacheControl.NoStore = true
		case directive == "no-transform":
			cacheControl.NoTransform = true
		case strings.HasPrefix(directive, "max-age="):
			fmt.Sscanf(directive, "max-age=%d", &cacheControl.MaxAge)
		case directive == "must-revalidate":
			cacheControl.MustRevalidate = true
		}
	}

	return cacheControl
}

func ReadFileHandler(fileName string, path string) {
	buf, err := os.ReadFile(path + "/" + fileName)
	if err != nil {

		println(err)
	}
	metaData := storage.MetaData{}
	err = json.Unmarshal(buf, &metaData)

	if err != nil {
		println(err)
		return
	}
	cahceControl := parseCacheControl(metaData.CacheControl)
	expirationTime, err := time.Parse(time.RFC1123, metaData.Date)
	if err != nil {
		println(err.Error())
		return
	}
	expirationTime = expirationTime.Add(time.Duration(cahceControl.MaxAge) * time.Second)

	if time.Now().After(expirationTime) {
		println("Expired")
		println(fileName, metaData.CacheControl, cahceControl.MaxAge)

		err = os.Remove(path + "/" + fileName)
		if err != nil {
			println(err.Error())
		}
		err = os.Remove(strings.Replace(path, "meta", "temp", 1) + "/" + fileName)
		if err != nil {
			println(err.Error())
		}
	}

}

func MetaDataCheck(path string) {
	folder, err := os.Open(path)
	if err != nil {
		println(err)
		panic(err)
	}

	defer folder.Close()

	// Read the folder contents
	fileInfos, err := folder.Readdir(-1)
	if err != nil {
		println("Error reading folder contents:", err)
		return
	}

	// Iterate over the file infos and print the file names
	for _, fileInfo := range fileInfos {

		if fileInfo.Mode().IsRegular() {
			ReadFileHandler(fileInfo.Name(), path)

		}
	}
}

func Worker() {
	gocron.Every(10).Minutes().Do(MetaDataCheck)
}
