package service

import (
	"os"
)

func InitDir(path string) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.Mkdir(path, 0755); err != nil {
				println(err.Error())
			}
		}
		panic(err)
	}

}
