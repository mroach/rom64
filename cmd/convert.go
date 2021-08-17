package cmd

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/mroach/n64-go/rom"
	"github.com/spf13/cobra"
)

func init() {
	var overwrite bool
	var convertCmd = &cobra.Command{
		Use:   "convert",
		Short: "Converts a ROM to native Big-Endian Z64 format",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			inpath := args[0]
			dirname, filename := path.Split(inpath)
			baseFilename := basename(filename)
			outFilename := baseFilename + ".z64"
			outpath := path.Join(dirname, outFilename)

			if !overwrite {
				if _, err := os.Stat(outpath); err == nil {
					return fmt.Errorf("Output file already exists: '%s'", outpath)
				}
			}

			if err := rom.ConvertRomFormat(inpath, outpath); err != nil {
				return err
			}

			fmt.Println(inpath, "=>", outpath)
			return nil
		},
	}

	convertCmd.Flags().BoolVarP(&overwrite, "force", "f", false, "Overwrite destination file if it exists")
	rootCmd.AddCommand(convertCmd)
}

func basename(filename string) string {
	if pos := strings.LastIndexByte(filename, '.'); pos != -1 {
		return filename[:pos]
	}
	return filename
}
