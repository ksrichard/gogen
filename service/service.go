package service

import (
	"fmt"
	"github.com/cbroglie/mustache"
	"github.com/ksrichard/gogen/util"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Generate - generating results recursively from a template folder/file to an output folder using variables
func Generate(input string, output string, outputType string, vars map[string]interface{}) error {
	// validate output type
	if !isOutputTypeValid(outputType) {
		return fmt.Errorf("Unknown output type!")
	}

	// if using template dir
	if util.IsDir(input) {
		if output == "" {
			return fmt.Errorf("Output must be set!")
		}

		// create/replace output dir
		err := util.RemoveCreateDir(output)
		if err != nil {
			return err
		}

		return filepath.Walk(input, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if input != path {
				relativePath := strings.ReplaceAll(path, input, "")

				// if directory, create them
				if info.IsDir() {
					renderedPaths, renderErr := mustache.Render(relativePath, vars)
					if renderErr != nil {
						return renderErr
					}
					mkdirErr := os.MkdirAll(output+"/"+renderedPaths, os.ModePerm)
					if mkdirErr != nil {
						return mkdirErr
					}
				}

				// if file, generate file to target dir
				if !info.IsDir() {
					buf, readErr := ioutil.ReadFile(path)
					if readErr != nil {
						return readErr
					}
					renderedFileName, fileNameRenderErr := mustache.Render(info.Name(), vars)
					renderedTemplate, renderErr := mustache.Render(string(buf), vars)
					if fileNameRenderErr == nil && renderErr == nil {
						writeErr := ioutil.WriteFile(output+"/"+renderedFileName, []byte(renderedTemplate), os.ModePerm)
						if writeErr != nil {
							return writeErr
						}
					} else {
						if fileNameRenderErr != nil {
							return fileNameRenderErr
						}
						if renderErr != nil {
							return renderErr
						}
					}
				}

			}
			return nil
		})
	}

	// using file/stdin template
	var inputData string
	var renderedFileName = ""
	var fileNameRenderErr error = nil

	// read input
	_, statErr := os.Stat(input)
	if statErr == nil { // having a file input
		buf, readErr := ioutil.ReadFile(input)
		if readErr != nil {
			return readErr
		}
		inputData = string(buf)
		renderedFileName, fileNameRenderErr = mustache.Render(input, vars)
	} else { // having stdin input
		inputData = input
	}

	renderedTemplate, renderErr := mustache.Render(inputData, vars)
	if fileNameRenderErr == nil && renderErr == nil {
		var outputWriteErr error

		switch strings.ToLower(outputType) {
		case "file":
			if output == "" {
				outputWriteErr = fmt.Errorf("Output must be set!")
				break
			}
			writeErr := ioutil.WriteFile(output, []byte(renderedTemplate), os.ModePerm)
			if writeErr != nil {
				outputWriteErr = writeErr
			}
			break
		case "folder":
			if output == "" {
				outputWriteErr = fmt.Errorf("Output must be set!")
				break
			}
			writeErr := ioutil.WriteFile(output+"/"+renderedFileName, []byte(renderedTemplate), os.ModePerm)
			if writeErr != nil {
				outputWriteErr = writeErr
			}
			break
		case "stdout":
			fmt.Print(renderedTemplate)
			outputWriteErr = nil
			break
		}

		return outputWriteErr
	} else {
		if fileNameRenderErr != nil {
			return fileNameRenderErr
		}
		if renderErr != nil {
			return renderErr
		}
	}

	return nil
}

func isOutputTypeValid(outputType string) bool {
	switch strings.ToLower(outputType) {
	case "file":
		return true
	case "folder":
		return true
	case "stdout":
		return true
	default:
		return false
	}
}
