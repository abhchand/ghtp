package main

import (
	"os"
	"path/filepath"
)

func absolutePath(path string) string {

	if string(path[0]) == "/" {
		return path
	}

	pwd, _ := os.Getwd()

	path = filepath.Clean(filepath.Join(pwd, path))
	path, err := filepath.Abs(path)

	if err != nil {
		panic(err)
	}

	return path

}
