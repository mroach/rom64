package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/mroach/n64-go/formatters"
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

func validateColumns(columns []string) ([]string, error) {
	columns, invalidCols := formatters.ValidateColumnIds(columns)
	if len(invalidCols) > 0 {
		return columns, fmt.Errorf("Invalid columns: %s", strings.Join(invalidCols, ", "))
	}

	return columns, nil
}

func printColumnHelp() {
	fmt.Println("Available columns:")
	fmt.Println(formatters.ColumnHelp())
	fmt.Println()
}
