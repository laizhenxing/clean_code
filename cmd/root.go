package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func Execute() {
	addCommands()

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// addCommands add child command
func addCommands() {

}
