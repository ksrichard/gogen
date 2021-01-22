package util

import (
	"os"
)

// IsDir - Check if input path is a directory
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

//RemoveCreateDir - create a directory structure, if still exist -> delete it before
func RemoveCreateDir(folderPath string) error {
	if IsDir(folderPath) {
		err := os.RemoveAll(folderPath)
		if err != nil {
			return err
		}
	}
	return os.MkdirAll(folderPath, os.ModePerm)
}
