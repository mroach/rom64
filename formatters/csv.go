package formatters

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/mroach/n64-go/rom"
)

func PrintCsv(records []rom.RomFile, separator rune) error {
	rows := toCsvRecords(records)

	w := csv.NewWriter(os.Stdout)
	w.Comma = separator
	if err := w.Write(defaultCsvHeaders); err != nil {
		return err
	}
	if err := w.WriteAll(rows); err != nil {
		return err
	}
	return nil
}

var defaultCsvHeaders = []string{
	"file_name", "file_format", "file_size",
	"image_name", "media_format", "cartridge_id",
	"region_name", "video_system", "cic", "crc1", "crc2",
}

func toCsvRecords(romfiles []rom.RomFile) [][]string {
	out := make([][]string, 0)
	for _, romfile := range romfiles {
		out = append(out, toCsvRecord(romfile))
	}
	return out
}

// TODO: Take a list of named fields to customise this
func toCsvRecord(romfile rom.RomFile) []string {
	return []string{
		romfile.File.Name,
		romfile.File.Format,
		fmt.Sprintf("%d", romfile.File.Size),
		romfile.ImageName,
		romfile.MediaFormat.Code,
		romfile.CartridgeId,
		romfile.Region.Description,
		romfile.VideoSystem,
		romfile.CIC,
		romfile.CRC1,
		romfile.CRC2,
	}
}
