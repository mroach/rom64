package formatters

import (
	"fmt"

	"github.com/mroach/n64-go/rom"
)

var textFormat = `File:
  Name:    %s
  Size:    %d MB
  Format:  %s

Title:     %s
ROM ID:    %s
Media:     %s (%s)
Version:   1.%d
Region:    %s (%s)
CIC:       %s
CRC 1:     %s
CRC 2:     %s
`

func PrintText(info rom.RomFile) error {
	_, err := fmt.Printf(textFormat,
		info.File.Name,
		info.File.Size,
		info.File.Format,
		info.ImageName,
		info.CartridgeId,
		info.MediaFormat.Code, info.MediaFormat.Description,
		info.Version,
		info.Region.Code, info.Region.Description,
		info.CIC,
		info.CRC1,
		info.CRC2,
	)

	return err
}
