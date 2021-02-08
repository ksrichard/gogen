package cmd

/*
Copyright © 2021 Richard Klavora <klavorasr@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"github.com/ksrichard/gogen/service"
	"github.com/ksrichard/gogen/util"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generating projects/folder structures/files based on a template",
	Long:  `Generating projects/folder structures/files based on a template`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// inputs
		input, _ := cmd.Flags().GetString("input")
		output, _ := cmd.Flags().GetString("output")

		// variables from flag(s) or yaml file
		vars, _ := cmd.Flags().GetStringToString("vars")
		varsConverted := make(map[string]interface{}, len(vars))
		for k, v := range vars {
			varsConverted[k] = v
		}
		varsMap := varsConverted

		// vars from file(s)
		varsFiles, _ := cmd.Flags().GetStringArray("vars-file")
		for _, file := range varsFiles {
			varsFromFile, _ := util.GetVarsFromFile(file)
			if varsFromFile != nil {
				varsMap = util.MergeMaps(varsMap, varsFromFile, varsConverted)
			}
		}

		outputType, _ := cmd.Flags().GetString("output-type")
		err := service.Generate(input, output, outputType, varsMap)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// template folder
	generateCmd.Flags().StringP("input", "i", "", "Template folder/file to be used for generation")
	cobra.MarkFlagDirname(generateCmd.Flags(), "input")
	cobra.MarkFlagRequired(generateCmd.Flags(), "input")

	// output folder
	generateCmd.Flags().StringP("output", "o", "", "Output folder/file where result will be generated\nNot applicable when output type is 'stdout'!")

	generateCmd.Flags().StringP("output-type", "t", "folder", "Output type - file,folder,stdout.\nNot applicable when input is a folder!")

	// variables
	generateCmd.Flags().StringToStringP("vars", "v", map[string]string{}, "Variables for generation")
	generateCmd.Flags().StringArrayP("vars-file", "f", []string{}, "Variables from file (YAML) for generation")

	// if having a piped input
	result, _ := util.ReadStdIn()
	if result != "" {
		// set input automatically
		inputFlag := generateCmd.Flag("input")
		inputFlag.Changed = true
		inputFlag.Value.Set(result)
	}
}
