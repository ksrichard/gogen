package util

import (
	"os"
)

func IsDir(dirInput string) bool {
	fi, err := os.Stat(dirInput)
	if err != nil {
		return false
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		return true
	case mode.IsRegular():
		return false
	}

	return false
}

func RemoveCreateDir(folderPath string) error {
	if IsDir(folderPath) {
		err := os.RemoveAll(folderPath)
		if err != nil {
			return err
		}
	}
	return os.MkdirAll(folderPath, os.ModePerm)
}

func GetCurrentDir() (string, error) {
	return os.Getwd()
}


