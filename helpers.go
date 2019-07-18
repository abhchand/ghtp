package main


import (
	"path/filepath"
	"os"
)

func absolutePath(path string) string {

	pwd, _ := os.Getwd()

	path = filepath.Clean(filepath.Join(pwd, path))
	path, err := filepath.Abs(path)

	if err != nil {
		panic(err)
	}

	return path

}
