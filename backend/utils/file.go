package utils

import (
	"os"
	"path/filepath"
)

func GetPath(path string, file string) (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	filename := path + file
	fullPath := filepath.Join(dir, filename)

	return fullPath, nil
}
