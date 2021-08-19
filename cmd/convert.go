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

			df, err := loadDatfile()
			if err != nil {
				return err
			}

			if err := rom.ConvertRomFormat(inpath, outpath); err != nil {
				return err
			}

			fmt.Printf("Created %s\n", outpath)

			// Read the new ROM info so we can validate it.
			info, err := rom.FromPath(outpath)
			if err != nil {
				return err
			}
			if err = info.AddSHA1(); err != nil {
				return err
			}

			matches, _, err := info.ValidateWithDat(df)
			if err != nil {
				return err
			}

			if len(matches) > 0 {
				fmt.Println("New file's SHA-1 validated against the datfile.")
				fmt.Printf("  OK %s\n", info.File.SHA1)
				fmt.Println("Conversion complete.")
				return nil
			}

			return fmt.Errorf("New file's SHA-1 could not be validated against the datfile.")
		},
	}

	convertCmd.Flags().BoolVarP(&overwrite, "force", "f", false, "Overwrite destination file if it exists")
	convertCmd.Flags().StringVarP(&datFilePath, "datfile", "d", "", "Load custom DAT file (XML format)")
	rootCmd.AddCommand(convertCmd)
}

func basename(filename string) string {
	if pos := strings.LastIndexByte(filename, '.'); pos != -1 {
		return filename[:pos]
	}
	return filename
}
