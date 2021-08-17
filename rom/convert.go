package rom

import (
	"errors"
	"io"
	"os"
)

func ConvertRomFormat(inpath string, outpath string) error {
	const bufferSize = 2048

	info, err := FromPath(inpath)
	if err != nil {
		return err
	}

	fileFormat := info.File.Format.Code

	if fileFormat == "z64" {
		return errors.New("File is already in the native Z64 format")
	}

	source, err := os.Open(inpath)
	if err != nil {
		return err
	}

	dest, err := os.Create(outpath)
	if err != nil {
		return err
	}

	for {
		buf := make([]byte, bufferSize)
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		buf = maybeReverseBytes(buf, fileFormat)
		if _, err := dest.Write(buf[:n]); err != nil {
			return err
		}
	}

	return nil
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
