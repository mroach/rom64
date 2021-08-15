package formatters

import (
	"fmt"
	"os"

	"github.com/mroach/n64-go/rom"
	"github.com/olekukonko/tablewriter"
)

func PrintTable(romfiles []rom.RomFile) error {
	records := toTableRecords(romfiles)
	printTable(records, defaultTableHeaders)
	return nil
}

var defaultTableHeaders = []string{
	"Title", "Format", "Size", "ROM ID", "Version",
	"Region", "CIC", "CRC1", "CRC2", "MD5", "File Name",
}

func printTable(records [][]string, headers []string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetColWidth(80)
	table.SetAutoFormatHeaders(false)
	table.SetHeader(headers)
	table.AppendBulk(records)
	table.Render()
}

func toTableRecords(romfiles []rom.RomFile) [][]string {
	out := make([][]string, 0)
	for _, romfile := range romfiles {
		out = append(out, toTableRecord(romfile))
	}
	return out
}

// TODO: parameter-driven column selection
func toTableRecord(romfile rom.RomFile) []string {
	return []string{
		romfile.ImageName,
		romfile.File.Format,
		fmt.Sprintf("%d", romfile.File.Size),
		romfile.MediaFormat.Code + romfile.CartridgeId + romfile.Region.Code,
		fmt.Sprintf("1.%d", romfile.Version),
		romfile.Region.Description,
		romfile.CIC,
		romfile.CRC1,
		romfile.CRC2,
		romfile.File.MD5,
		romfile.File.Name,
	}
}
