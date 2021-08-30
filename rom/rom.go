package rom

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"io"
	"math"
	"os"
	"strings"
)

const (
	FormatZ64 = "z64"
	FormatV64 = "v64"
	FormatN64 = "n64"
)

const (
	NTSC = "NTSC"
	PAL  = "PAL"
)

const ROM_HEADER_SIZE = 0x40

var bomZ64 = []byte{0x80, 0x37, 0x12, 0x40}
var bomV64 = []byte{0x37, 0x80, 0x40, 0x12}
var bomN64 = []byte{0x40, 0x12, 0x37, 0x80}

var bootcodeChecksumToCIC = map[uint32]string{
	0x587BD543: "5101",
	0x6170A4A1: "6101",
	0x90BB6CB5: "6102",
	0x0B050EE0: "6103",
	0x98BC2C86: "6105",
	0xACC8580A: "6106",
	0x009E9EA3: "7102",
	0x0E018159: "8303",
}

var MediaFormats = map[string]string{
	"N": "Cartridge",
	"D": "64DD Disk",
	"C": "Cartridge for expandable game",
	"E": "64DD Expansion",
	"Z": "Aleck64 Cartridge",
}

var FileFormats = map[string]string{
	FormatZ64: "Big-endian",
	FormatV64: "Byte-swapped",
	FormatN64: "Little-endian",
}

type Region struct {
	Id          string `json:"id"`
	Short       string `json:"short_name"`
	Description string `json:"description"`
	VideoSystem string `json:"video_system"`
}

var Regions = map[string]Region{
	"7": {"7", "Beta", "Beta", NTSC},
	"A": {"A", "JP/US", "Japan and USA", NTSC},
	"B": {"B", "BRA", "Brazil", NTSC},
	"C": {"C", "CHN", "China", PAL},
	"D": {"D", "GER", "Germany", PAL},
	"E": {"E", "USA", "USA", NTSC},
	"F": {"F", "FRA", "France", PAL},
	"G": {"G", "GW64-N", "Gateway 64 (NTSC)", NTSC},
	"H": {"H", "NED", "Netherlands", PAL},
	"I": {"I", "ITA", "Italy", PAL},
	"J": {"J", "JPN", "Japan", NTSC},
	"K": {"K", "KOR", "South Korea", NTSC},
	"L": {"L", "GW64-P", "Gateway 64 (PAL)", PAL},
	"N": {"N", "CAN", "Canada", NTSC},
	"P": {"P", "EUR", "Europe", PAL},
	"S": {"S", "ESP", "Spain", PAL},
	"U": {"U", "AUS", "Australia", PAL},
	"W": {"W", "NORD", "Scandinavia", PAL},
	"X": {"X", "PAL/X", "PAL Regions", PAL},
	"Y": {"Y", "PAL/Y", "PAL Regions", PAL},
}

// The first 4 bytes of a ROM are used for endianness detection and will already
// have been read by the time this struct is loaded. Noted address are absolute to the file.
// http://en64.shoutwiki.com/wiki/ROM#Cartridge_ROM_Header
type romFileHeader struct {
	ClockRate      uint32   // 0x04
	ProgramCounter uint32   // 0x08
	ReleaseAddress uint32   // 0x0C
	CRC1           uint32   // 0x10
	CRC2           uint32   // 0x14
	Unknown1       [8]byte  // 0x18
	ImageName      [20]byte // 0x20
	Unknown2       [4]byte  // 0x34
	MediaFormat    [4]byte  // 0x38
	CartridgeId    [2]byte  // 0x3C
	RegionCode     [1]byte  // 0x3E
	Version        byte     // 0x3F
}

type CodeDescription struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type FileInfo struct {
	Path   string          `json:"path" xml:"path"`
	Name   string          `json:"name" xml:"name"`
	Format CodeDescription `json:"format" xml:"format"`
	Size   int             `json:"size" xml:"size"`
	MD5    string          `json:"md5" xml:"md5"`
	SHA1   string          `json:"sha1" xml:"sha1"`
}

type RomFile struct {
	CRC1        string          `json:"crc1" xml:"crc1"`
	CRC2        string          `json:"crc2" xml:"crc2"`
	ImageName   string          `json:"image_name" xml:"image_name"`
	MediaFormat CodeDescription `json:"media_format" xml:"media_format"`
	CartridgeId string          `json:"cartridge_id" xml:"cartridge_id"`
	Region      Region          `json:"region" xml:"region"`
	Version     uint8           `json:"version" xml:"version"`
	CIC         string          `json:"cic" xml:"cic"`
	File        FileInfo        `json:"file" xml:"file"`
}

// 4-char ROM identifier, e.g. NSME = Super Mario 64 (USA), NSMJ = Super Mario 64 (Japan)
func (r *RomFile) Serial() string {
	return r.MediaFormat.Code + r.CartridgeId + r.Region.Id
}

func FromPath(path string) (RomFile, error) {
	var info RomFile

	f, err := os.Open(path)
	if err != nil {
		return info, err
	}
	defer f.Close()

	romfile, err := FromFile(f)
	if err != nil {
		return romfile, err
	}

	romfile.File.Path = path
	return romfile, nil
}

func FromFile(fh *os.File) (RomFile, error) {
	rominfo, err := FromIoReader(fh)
	if err != nil {
		return rominfo, err
	}

	stat, err := fh.Stat()
	if err != nil {
		return rominfo, err
	}

	rominfo.File.Size = romSize(stat.Size())
	rominfo.File.Name = stat.Name()

	return rominfo, err
}

func FromIoReader(r io.Reader) (RomFile, error) {
	var header romFileHeader
	var info RomFile

	// Read the first 4 bytes to detect the ROM file format

	endiannessSignature := make([]byte, 4)
	_, err := io.ReadFull(r, endiannessSignature)
	if err != nil {
		return info, err
	}
	romFormat, err := detectRomFormat(endiannessSignature[:])
	if err != nil {
		return info, err
	}

	headerBytes := make([]byte, ROM_HEADER_SIZE-len(endiannessSignature))
	err = binary.Read(r, binary.BigEndian, headerBytes)
	if err != nil {
		return info, err
	}
	headerBytes = maybeReverseBytes(headerBytes, romFormat)

	err = binary.Read(bytes.NewReader(headerBytes), binary.BigEndian, &header)
	if err != nil {
		return info, err
	}

	bootcode := make([]byte, 4032)
	err = binary.Read(r, binary.BigEndian, &bootcode)
	if err != nil {
		return info, err
	}
	bootcode = maybeReverseBytes(bootcode, romFormat)
	bootcodeHash := crc32.ChecksumIEEE(bootcode)
	cic := bootcodeChecksumToCIC[bootcodeHash]

	mediaFormatCode := bytesToString(header.MediaFormat[3:4])
	regionCode := bytesToString(header.RegionCode[:])

	info = RomFile{
		ImageName:   strings.TrimSpace(bytesToString(header.ImageName[:])),
		CartridgeId: bytesToString(header.CartridgeId[:]),
		CIC:         cic,
		CRC1:        fmt.Sprintf("%08X", header.CRC1),
		CRC2:        fmt.Sprintf("%08X", header.CRC2),
		Version:     header.Version,
		Region:      Regions[regionCode],
		MediaFormat: CodeDescription{
			Code:        mediaFormatCode,
			Description: MediaFormats[mediaFormatCode],
		},
		File: FileInfo{
			Format: CodeDescription{
				Code:        romFormat,
				Description: FileFormats[romFormat],
			},
		},
	}

	return info, nil
}

// ROM sizes are standard and padded-out to be whole-numbers of MB
// 8, 16, 32, 64 are common sizes
func romSize(size int64) int {
	const base float64 = 1024
	fsize := float64(size)
	exp := int(math.Log(fsize) / math.Log(base))
	return int(fsize / math.Pow(base, float64(exp)))
}

func detectRomFormat(signature []byte) (string, error) {
	if bytes.Equal(signature, bomZ64) {
		return FormatZ64, nil
	}

	if bytes.Equal(signature, bomV64) {
		return FormatV64, nil
	}

	if bytes.Equal(signature, bomN64) {
		return FormatN64, nil
	}

	return "", fmt.Errorf("Unknown ROM format. Invalid file?")
}

func bytesToString(bytes []byte) string {
	chars := []rune{}

	for i := range bytes {
		byte := bytes[i]
		if byte > 0 {
			chars = append(chars, rune(byte))
		}
	}

	return string(chars)
}
