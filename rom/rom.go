package rom

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"hash/crc32"
	"io"
	"os"
	"strings"
)

var headerZ64 = []byte{0x80, 0x37, 0x12, 0x40}
var headerV64 = []byte{0x37, 0x80, 0x40, 0x12}
var headerN64 = []byte{0x40, 0x12, 0x37, 0x80}

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

var Regions = map[string]string{
	"7": "Beta",
	"A": "NTSC",
	"B": "Brazil",
	"C": "China",
	"D": "Germany",
	"E": "United States",
	"F": "France",
	"G": "Gateway 64 (NTSC)",
	"H": "Netherlands",
	"I": "Italy",
	"J": "Japan",
	"K": "Korea",
	"L": "Gateway 64 (PAL)",
	"N": "Canada",
	"P": "Europe",
	"S": "Spain",
	"U": "Australia",
	"W": "Scandinavia",
	"X": "Europe",
	"Y": "Europe",
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
	Name   string `json:"name"`
	Format string `json:"format"`
	Size   int64  `json:"size"`
}

type RomFile struct {
	CRC1        string          `json:"crc_1"`
	CRC2        string          `json:"crc_2"`
	ImageName   string          `json:"image_name"`
	MediaFormat CodeDescription `json:"media_format"`
	CartridgeId string          `json:"cartridge_id"`
	Region      CodeDescription `json:"region"`
	Version     uint8           `json:"version"`
	CIC         string          `json:"cic"`
	File        FileInfo        `json:"file"`
}

func ParseFile(fh *os.File) (RomFile, error) {
	rominfo, err := ParseIo(fh)
	if err != nil {
		return rominfo, err
	}

	stat, err := fh.Stat()
	if err != nil {
		return rominfo, err
	}

	rominfo.File.Size = stat.Size()
	rominfo.File.Name = stat.Name()

	return rominfo, err
}

func ParseIo(r io.Reader) (RomFile, error) {
	var header romFileHeader
	var info RomFile

	// Read the first 4 bytes to detect the ROM file format
	var endiannessSignature [4]byte
	io.ReadFull(r, endiannessSignature[:])
	romFormat, err := detectRomFormat(endiannessSignature[:])
	if err != nil {
		return info, err
	}

	headerBytes := make([]byte, 0x40 - len(endiannessSignature))
	binary.Read(r, binary.BigEndian, headerBytes)
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
		CRC1:        fmt.Sprintf("%X", header.CRC1),
		CRC2:        fmt.Sprintf("%X", header.CRC2),
		Region: CodeDescription{
			Code:        regionCode,
			Description: Regions[regionCode],
		},
		MediaFormat: CodeDescription{
			Code:        mediaFormatCode,
			Description: MediaFormats[mediaFormatCode],
		},
		File: FileInfo{
			Format: romFormat,
		},
	}

	return info, nil
}

func detectRomFormat(signature []byte) (string, error) {
	if bytes.Equal(signature, headerZ64) {
		return "z64", nil
	}

	if bytes.Equal(signature, headerV64) {
		return "v64", nil
	}

	if bytes.Equal(signature, headerN64) {
		return "n64", nil
	}

	return "", errors.New("Unknown ROM format. Invalid file?")
}

func maybeReverseBytes(bytes []byte, romFormat string) []byte {
	if romFormat == "v64" {
		return reverseBytes(bytes, 2)
	}

	if romFormat == "n64" {
		return reverseBytes(bytes, 4)
	}

	return bytes
}

func reverseBytes(bytes []byte, size int) (reversed []byte) {
	for _, chunk := range chunk(bytes, size) {
		for i := len(chunk) - 1; i >= 0; i = i - 1 {
			reversed = append(reversed, chunk[i])
		}
	}

	return reversed
}

func chunk(bytes []byte, chunkSize int) (chunks [][]byte) {
	for chunkSize < len(bytes) {
		bytes, chunks = bytes[chunkSize:], append(chunks, bytes[0:chunkSize:chunkSize])
	}
	return append(chunks, bytes)
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
