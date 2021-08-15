package cmd

import (
	"github.com/mroach/n64-go/formatters"
	"github.com/mroach/n64-go/rom"
	"github.com/spf13/cobra"
)

func init() {
	var outputFormat string

	var statCmd = &cobra.Command{
		Use:     "stat",
		Aliases: []string{"info"},
		Short:   "Get ROM file information",
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := args[0]

			info, err := rom.FromPath(path)
			if err != nil {
				return err
			}

			if err := info.AddMD5(); err != nil {
				return err
			}

			return formatters.PrintOne(info, outputFormat)
		},
	}

	statCmd.Flags().StringVarP(&outputFormat, "output", "o", "text", "Output format")

	rootCmd.AddCommand(statCmd)
}
