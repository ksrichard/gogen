![Logo](images/logo.png)

[![Go Report Card](https://goreportcard.com/badge/github.com/ksrichard/gogen)](https://goreportcard.com/report/github.com/ksrichard/gogen)
[![Go Reference](https://pkg.go.dev/badge/github.com/ksrichard/gogen.svg)](https://pkg.go.dev/github.com/ksrichard/gogen)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/ksrichard/gogen.svg)](https://github.com/ksrichard/gogen)
[![GitHub release](https://img.shields.io/github/release/ksrichard/gogen.svg)](https://github.com/ksrichard/gogen/releases/latest/)

# gogen
A simple CLI and library to generate any kind of projects from any kind of templates! (just like maven archetype, but its much easier)
This CLI tool can be a very good replacement of `sed` commands when need to replace several values in configuration files.

There is absolutely no limitation in programming language or whatever you want to generate from a file and directory based template!

## Features
- Generating output from any kind of folder/file structure
- Uses the famous [Mustache](https://mustache.github.io/) templating engine for directory name, file name, file content generation

## Options
```
./gogen generate --help
Generating projects/folder structures based on a template

Usage:
  gogen generate [flags]

Flags:
  -h, --help                  help for generate
  -o, --output-dir string     Output folder where result files will be generated
  -t, --template-dir string   Template folder to be used for generation
  -v, --vars stringToString   Variables for generation (default [])
```
You can use everything in directory/file names and files content that can be done in [Mustache](https://mustache.github.io/) (for loops etc...)

**Note**: Here when you pass any variables to the generator, they will be used for generating both folder/file names and contents of files!

## Installation
You can install the CLI using homebrew (or local build, see next sectioin):
```
brew tap ksrichard/tap
brew install gogen
```

## Example usage of CLI
- Build the CLI using the following command executed in root of the project: `go build`

- Then simply run an example:
```
./gogen generate -t examples/simple-project/ -o output/GoTest -v module=GoTest -v project_name="Go Test"
```

- Finally see the output folder that was generated (and also check for file contents):
![Logo](images/screenshot1.png)

## Example usage of Library
```go
package main

import (
	gogen "github.com/ksrichard/gogen/service"
)

func main() {
	err := gogen.Generate("examples/simple-project", "output/GoTest", map[string]string{
		"module": "GoTest",
		"project_name": "Go Test",
	})
	if err != nil {
		panic(err)
	}
}
```

## TODO (upcoming features)
- [ ] Generate results from single file template
  - [ ] Add support to print result instead of writing to output directory
- [ ] Support more complex variables than just simple string key->value ones



