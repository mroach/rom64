package formatters

import (
	"encoding/csv"
	"os"

	"github.com/mroach/n64-go/rom"
)

func PrintCsv(records []rom.RomFile, separator rune, column_ids []string) error {
	w := csv.NewWriter(os.Stdout)
	w.Comma = separator

	headers := ColumnHeaders(column_ids)
	if err := w.Write(headers); err != nil {
		return err
	}

	rows := RomsToRecords(records, column_ids)
	if err := w.WriteAll(rows); err != nil {
		return err
	}
	return nil
}

var DefaultCsvColumns = []string{
	"file_name", "file_format", "file_size_mbytes",
	"rom_id", "image_name", "version",
	"region", "video_system", "cic", "crc1", "crc2",
}
