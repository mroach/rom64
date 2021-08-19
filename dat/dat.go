package dat

import (
	_ "embed"
	"encoding/xml"
	"io"
	"os"
	"strings"
)

//go:embed roms.dat.xml
var embeddedDatFile []byte

type DatFile struct {
	Name    string `xml:"header>name"`
	Version string `xml:"header>version"`
	Roms    []Rom  `xml:"game>rom"`
}

type Rom struct {
	Name   string `xml:"name,attr"`
	Size   int    `xml:"size,attr"`
	Serial string `xml:"serial,attr"`
	CRC32  string `xml:"crc32,attr"`
	MD5    string `xml:"md5,attr"`
	SHA1   string `xml:"sha1,attr"`
	Status string `xml:"status"`
}

func ReadFromIncluded() (df DatFile, err error) {
	df, err = Read(embeddedDatFile)
	if err != nil {
		return df, err
	}
	return df, nil
}

// Read a DatFile from an XML datfile on disk
func ReadFromFile(path string) (df DatFile, err error) {
	f, err := os.Open(path)
	if err != nil {
		return df, err
	}
	defer f.Close()

	bytes, err := io.ReadAll(f)
	if err != nil {
		return df, err
	}

	return Read(bytes)
}

// Read a DatFile from bytes of XML data
func Read(xmlbytes []byte) (df DatFile, err error) {
	if err := xml.Unmarshal(xmlbytes, &df); err != nil {
		return df, err
	}
	return df, nil
}

// Find entries based on a serial, such as NSME or CZLP
func (df *DatFile) FindBySerial(serial string) (results []Rom) {
	for _, rom := range df.Roms {
		if strings.EqualFold(rom.Serial, serial) {
			results = append(results, rom)
		}
	}
	return results
}
