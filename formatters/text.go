package formatters

import (
	"fmt"

	"github.com/mroach/n64-go/rom"
)

var textFormat = `File:
  Name:    %s
  Size:    %d MB
  Format:  %s
  MD5:     %s
  SHA1:    %s

ROM:
  ID:        %s%s%s
  Title:     %s
  Media:     %s
  Region:    %s
  Video:     %s
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
		info.File.MD5,
		info.File.SHA1,
		info.MediaFormat.Code, info.CartridgeId, info.Region.Code,
		info.ImageName,
		info.MediaFormat.Description,
		info.Region.Description,
		info.VideoSystem,
		info.Version,
		info.CIC,
		info.CRC1,
		info.CRC2,
	)

	return err
}
