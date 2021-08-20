package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/mroach/rom64/formatters"
	"github.com/spf13/cobra"
)

var binName string = "rom64"

const (
	AnsiReset = "\033[0m"
	FgBlack   = "\033[30m"
	FgRed     = "\033[31m"
	FgGreen   = "\033[32m"
	FgYellow  = "\033[033m"
	FgBlue    = "\033[34m"
)

var asciilogo = func() string {
	asciilogocolors := []string{FgRed, FgGreen, FgBlue, FgYellow, FgRed}
	logoargs := make([]interface{}, 0)
	for _, v := range asciilogocolors {
		logoargs = append(logoargs, v, AnsiReset)
	}
	asciilogo := []string{
		fmt.Sprintf(`%s             %s  %s            %s %s               %s     %s      /\\\\\     %s %s     /\\\        %s`, logoargs...),
		fmt.Sprintf(`%s             %s  %s            %s %s               %s     %s   /\\\\////     %s %s    /\\\\\       %s`, logoargs...),
		fmt.Sprintf(`%s             %s  %s            %s %s               %s     %s /\\\///         %s %s   /\\\/\\\      %s`, logoargs...),
		fmt.Sprintf(`%s/\\/\\\\\\\  %s  %s /\\\\\     %s %s /\\\\\  /\\\\\%s     %s/\\\\\\\\\\\     %s %s  /\\\/\/\\\     %s`, logoargs...),
		fmt.Sprintf(`%s\/\\\/////\\\%s  %s/\\\///\\\  %s %s/\\\///\\\\\///\\\%s  %s/\\\\///////\\\  %s %s /\\\/  \/\\\    %s`, logoargs...),
		fmt.Sprintf(`%s \/\\\   \///%s  %s/\\\  \//\\\%s %s\/\\\ \//\\\  \/\\\%s %s\/\\\      \//\\\%s %s/\\\\\\\\\\\\\\\\%s`, logoargs...),
		fmt.Sprintf(`%s  \/\\\      %s  %s\//\\\  /\\\%s %s \/\\\  \/\\\  \/\\\%s%s \//\\\      /\\\%s %s\///////////\\\//%s`, logoargs...),
		fmt.Sprintf(`%s   \/\\\     %s  %s  \///\\\\\/%s %s  \/\\\  \/\\\  \/\\\%s%s  \///\\\\\\\\\/%s %s           \/\\\ %s`, logoargs...),
		fmt.Sprintf(`%s    \///     %s  %s     \///// %s %s   \///   \///   \///%s%s     \///////// %s %s            \/// %s`, logoargs...),
	}
	return strings.Join(asciilogo, "\n")
}()

var rootCmd = &cobra.Command{
	Use:          binName,
	Short:        "Nintendo 64 ROM utility",
	Long:         asciilogo,
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
