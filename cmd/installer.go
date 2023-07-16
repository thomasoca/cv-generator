package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/thomasoca/cv-generator/internal/installer"
)

var installerCmd = &cobra.Command{
	Use:     "install",
	Aliases: []string{"ins"},
	Short:   "Install software dependencies",
	Run: func(cmd *cobra.Command, args []string) {
		runInstallerCmd()
	},
}

func runInstallerCmd() {
	err := installer.InstallPrerequisite()
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("Installation process completed")
}

func init() {
	rootCmd.AddCommand(installerCmd)
}
