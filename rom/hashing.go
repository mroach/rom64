package rom

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"hash"
	"io"
	"os"
)

func FileMD5(path string) (string, error) {
	return hashToHex(path, md5.New())
}

func FileSHA1(path string) (string, error) {
	return hashToHex(path, sha1.New())
}

func hashToHex(path string, hasher hash.Hash) (string, error) {
	var hexHash string

	file, err := os.Open(path)
	if err != nil {
		return hexHash, err
	}
	defer file.Close()

	if _, err := io.Copy(hasher, file); err != nil {
		return hexHash, err
	}

	hashBytes := hasher.Sum(nil)

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

func (romfile *RomFile) AddSHA1() error {
	sha1hex, err := FileSHA1(romfile.File.Path)
	if err != nil {
		return err
	}
	romfile.File.SHA1 = sha1hex
	return nil
}

func (romfile *RomFile) AddHashes() error {
	md5res := make(chan error)
	sha1res := make(chan error)

	go func(errs chan error) {
		errs <- romfile.AddMD5()
	}(md5res)
	go func(errs chan error) {
		errs <- romfile.AddSHA1()
	}(sha1res)

	if err := <-md5res; err != nil {
		return err
	}
	if err := <-sha1res; err != nil {
		return err
	}

	return nil
}
