package formatters

// Organise the possible columns to return in tables and CSVs.
// The machine keys can be used as selectors.

import (
	"fmt"
	"sort"
	"strings"

	"github.com/mroach/rom64/rom"
)

type columnValue func(rom.RomFile) string

type Column struct {
	Header      string
	Description string
	Generator   columnValue
}

var Columns = map[string]Column{
	"file_name": {
		"File Name",
		"File name on disk",
		func(r rom.RomFile) string { return r.File.Name },
	},
	"file_format": {
		"File Format",
		"File format code. One of: z64, v64, n64",
		func(r rom.RomFile) string { return r.File.Format.Code },
	},
	"file_format_desc": {
		"File Format",
		"File format description. example: Big-endian",
		func(r rom.RomFile) string { return r.File.Format.Description },
	},
	"file_size_mbytes": {
		"Size (MB)",
		"File size in megabytes. Always a whole number. example: 32",
		func(r rom.RomFile) string { return fmt.Sprintf("%d", r.File.Size) },
	},
	"file_size_mbits": {
		"Size (Mb)",
		"File size in megabits. Always a whole number. example: 256",
		func(r rom.RomFile) string { return fmt.Sprintf("%d", r.File.Size*8) },
	},
	"md5": {
		"MD5",
		"MD5 hash/checksum of the file on disk. Lower-case hexadecimal.",
		func(r rom.RomFile) string { return r.File.MD5 },
	},
	"sha1": {
		"SHA1",
		"SHA-1 hash/checksum of the file on disk. Lower-case hexadecimal.",
		func(r rom.RomFile) string { return r.File.SHA1 },
	},
	"image_name": {
		"Image Name",
		"Image name / game title embedded in the ROM.",
		func(r rom.RomFile) string { return r.ImageName },
	},
	"version": {
		"Version",
		"Version of the ROM. One of: 1.0, 1.1, 1.2, or 1.3.",
		func(r rom.RomFile) string { return fmt.Sprintf("1.%d", r.Version) },
	},
	"region": {
		"Region",
		"Region description of the ROM derived from the ROM ID.",
		func(r rom.RomFile) string { return r.Region.Description },
	},
	"region_short": {
		"Region",
		"Region short code",
		func(r rom.RomFile) string { return r.Region.Short },
	},
	"video_system": {
		"Video",
		"Video system derived from the ROM region. NTSC or PAL.",
		func(r rom.RomFile) string { return r.Region.VideoSystem },
	},
	"cic": {
		"CIC",
		"CIC chip type. example: 6102",
		func(r rom.RomFile) string { return r.CIC },
	},
	"crc1": {
		"CRC1",
		"CRC1 checksum of ROM internals. Also known as 'CRC HI'",
		func(r rom.RomFile) string { return r.CRC1 },
	},
	"crc2": {
		"CRC2",
		"CRC2 checksum of ROM internals. Also known as 'CRC LO'",
		func(r rom.RomFile) string { return r.CRC2 },
	},
	"rom_id": {
		"Rom ID",
		"ROM ID / serial. example: NSME for Super Mario 64 (USA)",
		func(r rom.RomFile) string { return r.Serial() },
	},
}

func ValidateColumnIds(column_ids []string) (valid []string, invalid []string) {
	valid = make([]string, 0)
	invalid = make([]string, 0)

	for _, column_id := range column_ids {
		if _, ok := Columns[column_id]; ok {
			valid = append(valid, column_id)
		} else {
			invalid = append(invalid, column_id)
		}
	}

	return valid, invalid
}

func ColumnHelp() string {
	lines := make([]string, 0)
	for colid, col := range Columns {
		lines = append(lines, fmt.Sprintf("  %-20s %s", colid, col.Description))
	}

	sort.Strings(lines)
	return strings.Join(lines, "\n")
}

func ColumnHeaders(column_ids []string) []string {
	headers := make([]string, 0)

	for _, column_id := range column_ids {
		headers = append(headers, Columns[column_id].Header)
	}

	return headers
}

func RomsToRecords(romfiles []rom.RomFile, column_ids []string) [][]string {
	records := make([][]string, 0)

	for _, romfile := range romfiles {
		records = append(records, PluckRomValues(romfile, column_ids))
	}

	return records
}

func PluckRomValues(romfile rom.RomFile, column_ids []string) []string {
	record := make([]string, 0)

	for _, column_id := range column_ids {
		column := Columns[column_id]
		record = append(record, column.Generator(romfile))
	}

	return record
}
