package formatters

import (
	"fmt"

	"github.com/mroach/n64-go/rom"
)

var textFormat = `File:
  Name:    %s
  Size:    %d MB
  Format:  %s

ROM:
  ID:        %s%s%s
  Title:     %s
  Media:     %s
  Region:    %s
  Version:   1.%d
  CIC:       %s
  CRC 1:     %s
  CRC 2:     %s
`

func PrintText(info rom.RomFile) error {
	_, err := fmt.Printf(textFormat,
		info.File.Name,
		info.File.Size,
		info.File.Format,
		info.MediaFormat.Code, info.CartridgeId, info.Region.Code,
		info.ImageName,
		info.MediaFormat.Description,
		info.Region.Description,
		info.Version,
		info.CIC,
		info.CRC1,
		info.CRC2,
	)

	return err
}
