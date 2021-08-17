package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var binName string = "rom64"

var rootCmd = &cobra.Command{
	Use:          binName,
	Short:        "Nintendo 64 ROM utility",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := fmt.Println("Use the 'help' command to learn about this application.")
		return err
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
