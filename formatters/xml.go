package formatters

import (
	"encoding/xml"
	"fmt"

	"github.com/mroach/rom64/rom"
)

func PrintXml(records []rom.RomFile) error {
	doc := struct {
		Roms    []rom.RomFile `xml:"rom"`
		XMLName struct{}      `xml:"roms"`
	}{Roms: records}

	bytes, err := xml.MarshalIndent(doc, "", "  ")
	if err != nil {
		return err
	}

	_, err = fmt.Printf("%s%s", xml.Header, bytes)
	return err
}
