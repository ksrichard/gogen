package service

import (
	"github.com/cbroglie/mustache"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func Generate(templateDir string, outputDir string, vars map[string]string) error {
	return filepath.Walk(templateDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if templateDir != path {
			relativePath := strings.ReplaceAll(path, templateDir, "")

			// if directory, create them
			if info.IsDir() {
				renderedPaths, renderErr := mustache.Render(relativePath, vars)
				if renderErr != nil {
					return renderErr
				}
				mkdirErr := os.MkdirAll(outputDir + "/" + renderedPaths, os.ModePerm)
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
				renderedFileName, fileNameRenderErr := mustache.Render(relativePath, vars)
				renderedTemplate, renderErr := mustache.Render(string(buf), vars)
				if fileNameRenderErr == nil && renderErr == nil {
					writeErr := ioutil.WriteFile(outputDir + "/" + renderedFileName, []byte(renderedTemplate), os.ModePerm)
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
