package util

import (
	"bufio"
	"errors"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
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

//ReadStdIn - read from stdin
func ReadStdIn() (string, error) {
	info, err := os.Stdin.Stat()
	if err != nil {
		return "", err
	}

	// no piping
	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		return "", nil
	}

	// read from stdin
	reader := bufio.NewReader(os.Stdin)
	var output []rune

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	return string(output), nil
}

func GetVarsFromFile(varsFile string) (map[string]interface{}, error) {
	// check if is dir
	if IsDir(varsFile) {
		return nil, errors.New("variables file must NOT be a directory")
	}

	// get file content
	_, statErr := os.Stat(varsFile)
	if statErr != nil {
		return nil, statErr
	}
	buf, readErr := ioutil.ReadFile(varsFile)
	if readErr != nil {
		return nil, readErr
	}

	// read yaml
	var result map[string]interface{}
	readErr = yaml.Unmarshal(buf, &result)
	if readErr != nil {
		return nil, readErr
	}

	return result, nil
}

func MergeMaps(maps ...map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}
