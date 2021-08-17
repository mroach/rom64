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

			md5res := make(chan error)
			sha1res := make(chan error)

			go func(errs chan error) {
				errs <- info.AddMD5()
			}(md5res)
			go func(errs chan error) {
				errs <- info.AddSHA1()
			}(sha1res)

			if err := <-md5res; err != nil {
				return err
			}
			if err := <-sha1res; err != nil {
				return err
			}

			return formatters.PrintOne(info, outputFormat)
		},
	}

	statCmd.Flags().StringVarP(&outputFormat, "output", "o", "text", "Output format")

	rootCmd.AddCommand(statCmd)
}
