package formatters

import (
	"os"

	"github.com/mroach/n64-go/rom"
	"github.com/olekukonko/tablewriter"
)

func PrintTable(romfiles []rom.RomFile, column_ids []string) error {
	headers := ColumnHeaders(column_ids)
	records := RomsToRecords(romfiles, column_ids)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetColWidth(80)
	table.SetAutoFormatHeaders(false)
	table.SetHeader(headers)
	table.AppendBulk(records)
	table.Render()

	return nil
}

var DefaultTableColumns = []string{
	"image_name", "file_format_desc", "file_size_mbytes",
	"rom_id", "version", "region", "video_system",
	"cic", "crc1", "crc2", "file_name",
}
