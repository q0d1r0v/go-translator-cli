// cmd/root.go

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "trans",
	Short: "CLI Translator",
	Long:  `CLI Translator for language translation using API.`,
	Run: func(cmd *cobra.Command, args []string) {
		// default behavior if no subcommand is provided
		fmt.Println("Welcome to the CLI Translator!")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	return rootCmd.Execute()
}
