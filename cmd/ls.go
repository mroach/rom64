package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"sync"

	"github.com/mroach/rom64/formatters"
	"github.com/mroach/rom64/rom"
	"github.com/spf13/cobra"
)

func init() {
	var outputFormat string
	var columns []string
	calcMd5 := false
	calcSha := false

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
				return fmt.Errorf("No ROM files found in '%s'", path)
			}

			if len(columns) == 0 {
				columns = formatters.DefaultColumns(outputFormat)
			}

			columns, err := validateColumns(columns)
			if err != nil {
				printColumnHelp()
				return err
			}

			for _, column := range columns {
				if column == "md5" {
					calcMd5 = true
				}
				if column == "sha1" {
					calcSha = true
				}
			}

			results := make(chan rom.RomFile, len(files))
			errs := make(chan error, len(files))

			var wg sync.WaitGroup
			for _, rompath := range files {
				wg.Add(1)
				go func(rompath string) {
					defer wg.Done()
					info, err := rom.FromPath(rompath)
					if err != nil {
						errs <- err
						return
					}
					if calcMd5 {
						if err := info.AddMD5(); err != nil {
							errs <- err
						}
					}
					if calcSha {
						if err := info.AddSHA1(); err != nil {
							errs <- err
						}
					}
					results <- info
				}(rompath)
			}
			wg.Wait()
			close(errs)
			close(results)

			fileInfos := make([]rom.RomFile, 0)
			for info := range results {
				fileInfos = append(fileInfos, info)
			}
			sort.Slice(fileInfos, func(i, j int) bool {
				return fileInfos[i].File.Name < fileInfos[j].File.Name
			})

			if len(errs) > 0 {
				l := log.New(os.Stderr, "", 1)
				l.Println("Some ROMs could not be read:")
				for err := range errs {
					l.Println(err)
				}
			}

			return formatters.PrintAll(fileInfos, outputFormat, columns)
		},
	}

	lsCmd.Flags().StringVarP(&outputFormat, "output", "o", "table",
		fmt.Sprintf("Output format (%s)", strings.Join(formatters.OutputFormats, ", ")))
	lsCmd.Flags().StringSliceVarP(&columns, "columns", "c", make([]string, 0), "Column selection")

	rootCmd.AddCommand(lsCmd)
}
