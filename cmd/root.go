/*
Copyright © 2023 phaseant
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	// Use:   "PasswordGenerator",
	Use:   "pg",
	Short: "Password Generator is a CLI tool to generate passwords",
	Long:  `Password Generator is a CLI tool to generate passwords. You can use it to generate passwords with different lengths and different characters.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
