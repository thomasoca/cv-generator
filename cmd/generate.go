package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/thomasoca/cv-generator/pkg/app"
)

var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"gen"},
	Short:   "Generate resume directly using terminal",
	Run: func(cmd *cobra.Command, args []string) {
		input, _ := cmd.Flags().GetString("input")
		output, _ := cmd.Flags().GetString("output")
		runAppCmd(input, output)
	},
}

func runAppCmd(input string, output string) {
	a := app.AppCmd{InputPath: input, OutputPath: output}
	o, err := a.GenerateFile()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s' ", err)
		os.Exit(1)
	}
	fmt.Printf("Resume successfully generated on %s", o)
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.PersistentFlags().String("input", "", "Path for the JSON input file")
	generateCmd.PersistentFlags().String("output", "app", "Path for the output pdf and latex files")
	generateCmd.MarkPersistentFlagRequired("input")
}
