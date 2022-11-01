package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/mroach/rom64/dat"
	"github.com/mroach/rom64/rom"
	"github.com/spf13/cobra"
)

var datFilePath string
var renameValidated bool

func init() {
	var validateCmd = &cobra.Command{
		Use:   "validate",
		Short: "Validate the hash of a ROM against a known-good list",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			romFilePath := args[0]

			var df dat.DatFile
			var err error

			romfile, err := rom.FromPath(romFilePath)
			if err != nil {
				return err
			}

			if romfile.Serial() == "" {
				return fmt.Errorf("ROM has no serial number which is required to look it up in the DAT file.")
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

			var matchCount = len(matches)

			if matchCount == 0 {
				fmt.Println("Could not find a checksum match.")
				fmt.Printf("File '%s' has SHA-1 %s", romfile.File.Name, romfile.File.SHA1)
				fmt.Println("The datfile has the following entries for this ROM:")
				for _, mismatch := range mismatches {
					fmt.Printf("  %-5s %50s %s", "SHA-1\n", mismatch.SHA1, mismatch.Name)
				}
				return fmt.Errorf("%s validation failed", romFilePath)
			}

			fmt.Printf("Found %d datfile entries for ROM serial '%s'\n", matchCount, romfile.Serial())

			for _, match := range matches {
				fmt.Printf("%-5s \033[32m%-6s\033[0m %40s \"%s\"\n", "SHA-1", "MATCH", match.SHA1, match.Name)
			}

			if matchCount > 1 {
				return fmt.Errorf("Multiple datfile entries found with the same SHA-1 hash. This shouldn't happen. Aborting.")
			}

			var match = matches[0]

			if renameValidated {
				var correctName = match.Name
				if romfile.File.Name == correctName {
					fmt.Printf("ROM file already has the correct name \"%s\"\n", correctName)
					return nil
				}

				var sourcePath = romfile.File.Path
				var newPath = path.Join(path.Dir(sourcePath), correctName)

				fmt.Printf("Renaming \"%s\" => \"%s\"\n", romfile.File.Path, newPath)

				if _, err := os.Stat(newPath); err == nil {
					return fmt.Errorf("Destination file \"%s\" already exists. Aborting.", newPath)
				}

				err = os.Rename(sourcePath, newPath)
				if err != nil {
					return err
				}
			}

			return nil
		},
	}

	validateCmd.Flags().StringVarP(&datFilePath, "datfile", "d", "", "Load custom DAT file (XML format)")
	validateCmd.Flags().BoolVarP(&renameValidated, "rename-validated", "", false, "Rename validated files to match the filename in the datefile.")

	rootCmd.AddCommand(validateCmd)
}

func loadDatfile() (dat.DatFile, error) {
	if datFilePath == "" {
		return dat.ReadFromIncluded()
	} else {
		return dat.ReadFromFile(datFilePath)
	}
}
