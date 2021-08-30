package rom

import (
	"fmt"
	"os"
)

const (
	CRC_CHECKSUM_START  = 0x00001000
	CRC_CHECKSUM_LENGTH = 0x00100000
	CRC_CHECKSUM_END    = CRC_CHECKSUM_START + CRC_CHECKSUM_LENGTH
)

func crcSeedForCIC(cic string) uint32 {
	switch cic {
	case "6101", "6102":
		return 0xF8CA4DDC
	case "6103":
		return 0xA3886759
	case "6105":
		return 0xDF26F436
	case "6106":
		return 0x1FEA617A
	default:
		return 1
	}
}

// Calculate CRC1 (aka "CRC HI") and CRC2 (aka "CRC LO") for the given RomFile
//
// This code is a direct port from the C implementation at http://n64dev.org/n64crc.html
// Since I don't understand the "why" behind this, I can't document it further.
func (rf *RomFile) CalcCRC() error {
	file, err := os.Open(rf.File.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	bytes := make([]byte, CRC_CHECKSUM_END)
	if _, err := file.Read(bytes); err != nil {
		return err
	}
	bytes = maybeReverseBytes(bytes, rf.File.Format.Code)

	cic := rf.CIC
	seed := crcSeedForCIC(cic)
	t1, t2, t3, t4, t5, t6 := seed, seed, seed, seed, seed, seed

	for i := CRC_CHECKSUM_START; i < CRC_CHECKSUM_END; i += 4 {
		d := uint32be(bytes[i : i+4])

		if t6+d < t6 {
			t4++
		}

		t6 += d
		t3 ^= d

		r := rol(d, int(d&0x1F))
		t5 += r

		if t2 > d {
			t2 ^= r
		} else {
			t2 ^= t6 ^ d
		}

		if cic == "6105" {
			extra := ROM_HEADER_SIZE + 0x0710 + (i & 0xFF)
			b := uint32be(bytes[extra : extra+4])
			t1 += b ^ d
		} else {
			t1 += t5 ^ d
		}

	}

	var crc1, crc2 uint32

	switch cic {
	case "6103":
		crc1 = (t6 ^ t4) + t3
		crc2 = (t5 ^ t2) + t1
	case "6106":
		crc1 = (t6 * t4) + t3
		crc2 = (t5 * t2) + t1
	default:
		crc1 = t6 ^ t4 ^ t3
		crc2 = t5 ^ t2 ^ t1
	}

	rf.File.CRC1 = fmt.Sprintf("%08X", crc1)
	rf.File.CRC2 = fmt.Sprintf("%08X", crc2)
	return nil
}

func rol(i uint32, b int) uint32 {
	return (i << b) | (i >> (32 - b))
}

// Read a uint32 from the bytes that are in big-endian order
func uint32be(bytes []byte) uint32 {
	return uint32(bytes[0])<<24 |
		uint32(bytes[1])<<16 |
		uint32(bytes[2])<<8 |
		uint32(bytes[3])
}
