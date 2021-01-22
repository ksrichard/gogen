package util

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestIsDirFile(t *testing.T) {
	file, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(file.Name())
	isDir := IsDir(file.Name())
	if isDir {
		t.Errorf("'%s' must NOT be a directory!", file.Name())
	}
}

func TestIsDirFolder(t *testing.T) {
	dirName, err := ioutil.TempDir("", "test")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(dirName)
	isDir := IsDir(dirName)
	if !isDir {
		t.Errorf("'%s' must be a directory!", dirName)
	}
}

func TestIsDirError(t *testing.T) {
	dirName, err := ioutil.TempDir("", "test")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(dirName)
	isDir := IsDir(dirName + "123test")
	if isDir {
		t.Errorf("'%s' must NOT be a directory!", dirName)
	}
}

func TestRemoveCreateDir(t *testing.T) {
	dirName, err := ioutil.TempDir("", "test")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(dirName)
	testFileName := dirName + "/test.txt"
	err = ioutil.WriteFile(testFileName, []byte("Some data"), os.ModePerm)
	if err != nil {
		t.Error(err)
	}
	fExists := fileExists(testFileName)
	if !fExists {
		t.Errorf("File '%s' should exist!", testFileName)
	}
	err = RemoveCreateDir(dirName)
	if err != nil {
		t.Error(err)
	}
	fExists = fileExists(testFileName)
	if fExists {
		t.Errorf("File '%s' should NOT exist!", testFileName)
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
