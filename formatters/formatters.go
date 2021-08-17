package formatters

import (
	"fmt"
	"strings"

	"github.com/mroach/n64-go/rom"
)

func PrintAll(items []rom.RomFile, outputFormat string) error {
	switch outputFormat {
	case "csv":
		return PrintCsv(items, ',')
	case "tab":
		return PrintCsv(items, '\t')
	case "json":
		return PrintJson(items)
	case "table":
		return PrintTable(items)
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

func PrintOne(item rom.RomFile, outputFormat string) error {
	switch outputFormat {
	case "csv":
		return PrintCsv([]rom.RomFile{item}, ',')
	case "tab":
		return PrintCsv([]rom.RomFile{item}, '\t')
	case "json":
		return PrintJson(item)
	case "table":
		return PrintTable([]rom.RomFile{item})
	case "text":
		return PrintText(item)
	}

	return fmt.Errorf("Invalid output format '%s'", outputFormat)
}
