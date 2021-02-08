![Logo](images/logo.png)

[![Go Report Card](https://goreportcard.com/badge/github.com/ksrichard/gogen)](https://goreportcard.com/report/github.com/ksrichard/gogen)
[![Go Reference](https://pkg.go.dev/badge/github.com/ksrichard/gogen.svg)](https://pkg.go.dev/github.com/ksrichard/gogen)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/ksrichard/gogen.svg)](https://github.com/ksrichard/gogen)
[![GitHub release](https://img.shields.io/github/release/ksrichard/gogen.svg)](https://github.com/ksrichard/gogen/releases/latest/)

Code coverage: https://gocover.io/github.com/ksrichard/gogen

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
  -h, --help                    help for generate
  -i, --input string            Template folder/file to be used for generation
  -o, --output string           Output folder/file where result will be generated
                                Not applicable when output type is 'stdout'!
  -t, --output-type string      Output type - file,folder,stdout.
                                Not applicable when input is a folder! (default "folder")
  -v, --vars stringToString     Variables for generation (default [])
  -f, --vars-file stringArray   Variables from file (YAML) for generation
```
You can use everything in directory/file names and files content that can be done in [Mustache](https://mustache.github.io/) (for loops etc...)

**Note**: Here when you pass any variables to the generator, they will be used for generating both folder/file names and contents of files!

## Installation
You can install the CLI using homebrew (or local build, see next section):
```
brew tap ksrichard/tap
brew install gogen
```

## Example usages of CLI
- Build the CLI using the following command executed in root of the project: `go build`
**Note:** You can combine any of the input/output, output types parameters to fit to your needs of course!

    - Generate a project/folder structure:
    ```
    ./gogen generate -i examples/simple-project/ -o output/GoTest -v module=GoTest -v project_name="Go Test"
    ```
    See the output folder that was generated (and also check for file contents):
    ![Logo](images/screenshot1.png)
    
    - Generate single file:
    ```
    ./gogen generate -i ./examples/simple-project/main.go -t stdout -v project_name="Test Project"
    ```
    - Generate single file to file output:
    ```
    ./gogen generate -i ./examples/simple-project/main.go -t file -o test.go -v project_name="Test Project"
    ```
    - Generate using variables from file, template from pipe:
    ```bash
    echo "{{testRootValue}} -->> {{#projects}}\"{{name}},{{url}} -> {{#data}}{{value}}{{/data}}\" {{/projects}}" | ./gogen generate -t stdout -f ./examples/test.yaml
    ```
    - Generate using variables from file (properties overwrite from another file), template from pipe:
    ```bash
    echo "{{testRootValue}} -->> {{#projects}}\"{{name}},{{url}} -> {{#data}}{{value}}{{/data}}\" {{/projects}}" | ./gogen generate -t stdout -f ./examples/test.yaml -f ./examples/test2.yaml
    ```
    - Generate using variables from file (properties overwrite from command line option), template from pipe:
    ```bash
    echo "{{testRootValue}} -->> {{#projects}}\"{{name}},{{url}} -> {{#data}}{{value}}{{/data}}\" {{/projects}}" | ./gogen generate -t stdout -f ./examples/test.yaml -v testRootValue="someTest2"
    ```
    - Generate using variables from file (properties overwrite from command line option), template from pipe:
    ```bash
    echo "{{testRootValue}} -->> {{#projects}}\"{{name}},{{url}} -> {{#data}}{{value}}{{/data}}\" {{/projects}}" | ./gogen generate -t stdout -f ./examples/test.yaml -v testRootValue="someTest2"
    ```

## Example usage of Library
```go
package main

import (
	gogen "github.com/ksrichard/gogen/service"
)

func main() {
	err := gogen.Generate("examples/simple-project", "output/GoTest", "folder", map[string]interface{
		"module": "GoTest",
		"project_name": "Go Test",
	})
	if err != nil {
		panic(err)
	}
}
```

## TODO (upcoming features)
- [x] Support input from stdin
- [x] Generate result from single file template
  - [x] Add support to print result instead of writing to file/output directory
- [x] Support more complex variables than just simple string key -> value ones



