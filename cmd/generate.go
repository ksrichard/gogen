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
package cmd

import (
	"github.com/ksrichard/gogen/service"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generating projects/folder structures based on a template",
	Long:  `Generating projects/folder structures based on a template`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// inputs
		templateDir, _ := cmd.Flags().GetString("template-dir")
		outputDir, _ := cmd.Flags().GetString("output-dir")
		vars, _ := cmd.Flags().GetStringToString("vars")

		err := service.Generate(templateDir, outputDir, vars)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// template folder
	generateCmd.Flags().StringP("template-dir", "t", "", "Template folder to be used for generation")
	cobra.MarkFlagDirname(generateCmd.Flags(), "template-dir")
	cobra.MarkFlagRequired(generateCmd.Flags(), "template-dir")

	// output folder
	generateCmd.Flags().StringP("output-dir", "o", "", "Output folder where result files will be generated")
	cobra.MarkFlagDirname(generateCmd.Flags(), "output-dir")
	cobra.MarkFlagRequired(generateCmd.Flags(), "output-dir")

	// variables
	generateCmd.Flags().StringToStringP("vars", "v", map[string]string{}, "Variables for generation")
}
