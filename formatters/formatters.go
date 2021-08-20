package formatters

import (
	"fmt"
	"strings"

	"github.com/mroach/rom64/rom"
)

func DefaultColumns(outputFormat string) []string {
	switch outputFormat {
	case "csv", "tab":
		return DefaultCsvColumns
	case "table":
		return DefaultTableColumns
	}

	return make([]string, 0)
}

func PrintAll(items []rom.RomFile, outputFormat string, columns []string) error {
	switch outputFormat {
	case "csv":
		return PrintCsv(items, ',', columns)
	case "tab":
		return PrintCsv(items, '\t', columns)
	case "json":
		return PrintJson(items)
	case "table":
		return PrintTable(items, columns)
	case "text":
		hr := strings.Repeat("-", 80)
		count := len(items)

		for i, item := range items {
			if err := PrintText(item); err != nil {
				return err
			}
			if i+1 < count {
				fmt.Println(hr)
			}
		}

		return nil
	}

	return fmt.Errorf("Invalid output format '%s'", outputFormat)
}

func PrintOne(item rom.RomFile, outputFormat string, columns []string) error {
	switch outputFormat {
	case "csv":
		return PrintCsv([]rom.RomFile{item}, ',', columns)
	case "tab":
		return PrintCsv([]rom.RomFile{item}, '\t', columns)
	case "json":
		return PrintJson(item)
	case "table":
		return PrintTable([]rom.RomFile{item}, columns)
	case "text":
		return PrintText(item)
	}

	return fmt.Errorf("Invalid output format '%s'", outputFormat)
}
