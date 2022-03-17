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
	var quiet bool
	calcMd5 := false
	calcSha := false
	calcCrc := false

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
				if column == "file_md5" {
					calcMd5 = true
				}
				if column == "file_sha1" {
					calcSha = true
				}
				if column == "file_crc1" || column == "file_crc2" {
					calcCrc = true
				}
			}

			results := make(chan rom.RomFile, len(files))
			errs := make(chan struct {
				string
				error
			}, len(files))

			var wg sync.WaitGroup
			for _, rompath := range files {
				wg.Add(1)
				go func(rompath string) {
					defer wg.Done()
					info, err := rom.FromPath(rompath)
					if err != nil {
						sendError(errs, rompath, err)
						return
					}
					if calcMd5 {
						if err := info.AddMD5(); err != nil {
							sendError(errs, rompath, err)
						}
					}
					if calcSha {
						if err := info.AddSHA1(); err != nil {
							sendError(errs, rompath, err)
						}
					}
					if calcCrc {
						if err := info.CalcCRC(); err != nil {
							sendError(errs, rompath, err)
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

			if len(errs) > 0 && !quiet {
				l := log.New(os.Stderr, "", 1)
				l.Println("Errors were encountered while listing some files:")
				for item := range errs {
					l.Printf("%s: %s\n", item.string, item.error)
				}
			}

			return formatters.PrintAll(fileInfos, outputFormat, columns)
		},
	}

	lsCmd.Flags().StringVarP(&outputFormat, "output", "o", "table",
		fmt.Sprintf("Output format (%s)", strings.Join(formatters.OutputFormats, ", ")))
	lsCmd.Flags().StringSliceVarP(&columns, "columns", "c", make([]string, 0), "Column selection")
	lsCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "Quiet mode. Suppress non-fatal errors.")

	rootCmd.AddCommand(lsCmd)
}

func sendError(queue chan struct {
	string
	error
}, path string, err error) {
	queue <- struct {
		string
		error
	}{path, err}
}
