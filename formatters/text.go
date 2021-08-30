package formatters

import (
	"os"
	"text/template"

	"github.com/mroach/rom64/rom"
)

var textFormat = `File:
  Name:    {{.File.Name}}
  Size:    {{.File.Size}} MB
  Format:  {{.File.Format.Code}} ({{.File.Format.Description}})
  Checksums:
    MD5:     {{if .File.MD5}}{{.File.MD5}}{{else}}Not Calculated{{end}}
    SHA1:    {{if .File.SHA1}}{{.File.SHA1}}{{else}}Not Calculated{{end}}

ROM:
  ID:        {{.Serial}}
  Title:     {{.ImageName}}
  Media:     {{.MediaFormat.Description}}
  Region:    {{.Region.Description}}
  Video:     {{.Region.VideoSystem}}
  Version:   1.{{.Version}}
  CIC:       {{.CIC}}
  CRC 1:     {{.CRC1}}
  CRC 2:     {{.CRC2}}
`

func PrintText(info rom.RomFile) error {
	var defaultTextTemplate = template.Must(template.New("rom").Parse(textFormat))
	return defaultTextTemplate.Execute(os.Stdout, &info)
}
