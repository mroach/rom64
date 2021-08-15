package cmd

import (
	"fmt"

	"github.com/mroach/n64-go/formatters"
	"github.com/mroach/n64-go/rom"
	"github.com/spf13/cobra"
)

func init() {
	var outputFormat string
	var includeMd5 bool

	var lsCmd = &cobra.Command{
		Use:     "ls",
		Aliases: []string{"list"},
		Short:   "Find and list Nintendo 64 ROMs",
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := args[0]
			files, err := rom.FindProbableRomsInPath(path)
			if err != nil {
				return err
			}

			if len(files) == 0 {
				fmt.Println("No ROM files found in", path)
				return nil
			}

			fileInfos := make([]rom.RomFile, 0)
			for _, romPath := range files {
				info, err := rom.FromPath(romPath)
				if err != nil {
					return err
				}

				if includeMd5 {
					if err := info.AddMD5(); err != nil {
						return err
					}
				}

				fileInfos = append(fileInfos, info)
			}

			return formatters.PrintAll(fileInfos, outputFormat)
		},
	}

	lsCmd.Flags().StringVarP(&outputFormat, "output", "o", "table", "Output format")
	lsCmd.Flags().BoolVarP(&includeMd5, "with-md5", "m", false, "Calculate MD5 hash")

	rootCmd.AddCommand(lsCmd)
}
