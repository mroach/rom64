package cmd

import (
	"fmt"

	"github.com/mroach/n64-go/dat"
	"github.com/mroach/n64-go/rom"
	"github.com/spf13/cobra"
)

var datFilePath string

func init() {
	var validateCmd = &cobra.Command{
		Use:   "validate",
		Short: "Validate the hash of a ROM against a known-good list",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := args[0]

			var df dat.DatFile
			var err error

			romfile, err := rom.FromPath(path)
			if err != nil {
				return err
			}
			if err = romfile.AddSHA1(); err != nil {
				return err
			}

			df, err = loadDatfile()
			if err != nil {
				return err
			}

			matches, mismatches, err := romfile.ValidateWithDat(df)
			if err != nil {
				return err
			}

			fmt.Printf("Found %d datfile entries for ROM serial '%s'\n", len(matches), romfile.Serial())

			if len(matches) > 0 {
				for _, match := range matches {
					fmt.Printf("  %-5s \033[32m%-6s\033[0m %40s %s\n", "SHA-1", "MATCH", match.SHA1, match.Name)
				}
				return nil
			}

			fmt.Println("Could not find a checksum match.")
			fmt.Printf("File '%s' has SHA-1 %s", romfile.File.Name, romfile.File.SHA1)
			fmt.Println("The datfile has the following entries for this ROM:")
			for _, mismatch := range mismatches {
				fmt.Printf("  %-5s %50s %s", "SHA-1", mismatch.SHA1, mismatch.Name)
			}
			return fmt.Errorf("%s validation failed", path)
		},
	}

	validateCmd.Flags().StringVarP(&datFilePath, "datfile", "d", "", "Load custom DAT file (XML format)")

	rootCmd.AddCommand(validateCmd)
}

func loadDatfile() (dat.DatFile, error) {
	if datFilePath == "" {
		return dat.IncludedDat, nil
	} else {
		return dat.ReadFromFile(datFilePath)
	}
}
