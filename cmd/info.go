package cmd

import (
	"github.com/mroach/rom64/formatters"
	"github.com/mroach/rom64/rom"
	"github.com/spf13/cobra"
)

func init() {
	var outputFormat string
	var columns []string

	var infoCmd = &cobra.Command{
		Use:     "info",
		Aliases: []string{"info"},
		Short:   "Get ROM file information",
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := args[0]

			info, err := rom.FromPath(path)
			if err != nil {
				return err
			}

			if len(columns) == 0 {
				columns = formatters.DefaultColumns(outputFormat)
			}

			columns, err := validateColumns(columns)
			if err != nil {
				printColumnHelp()
				return err
			}

			err = info.AddHashes()
			if err != nil {
				return err
			}

			return formatters.PrintOne(info, outputFormat, columns)
		},
	}

	infoCmd.Flags().StringVarP(&outputFormat, "output", "o", "text", "Output format")
	infoCmd.Flags().StringSliceVarP(&columns, "columns", "c", make([]string, 0), "Column selection")

	rootCmd.AddCommand(infoCmd)
}
