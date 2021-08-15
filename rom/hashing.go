package rom

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

func FileMD5(path string) (string, error) {
	var md5hex string

	file, err := os.Open(path)
	if err != nil {
		return md5hex, err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return md5hex, err
	}

	hashBytes := hash.Sum(nil)[:16]

	return hex.EncodeToString(hashBytes), nil
}

func (romfile *RomFile) AddMD5() error {
	md5hex, err := FileMD5(romfile.File.Path)
	if err != nil {
		return err
	}
	romfile.File.MD5 = md5hex
	return nil
}
