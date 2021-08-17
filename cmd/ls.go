package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"
	"sync"

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
				return fmt.Errorf("No ROM files found in '%s'", path)
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
					if includeMd5 {
						if err := info.AddMD5(); err != nil {
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

			return formatters.PrintAll(fileInfos, outputFormat)
		},
	}

	lsCmd.Flags().StringVarP(&outputFormat, "output", "o", "table", "Output format")
	lsCmd.Flags().BoolVarP(&includeMd5, "with-md5", "m", false, "Calculate MD5 hash")

	rootCmd.AddCommand(lsCmd)
}
