package rom

import (
	"os"
	"path/filepath"
	"strings"
)

var Extensions = []string{"bin", "rom", "d64", "n64", "u64", "v64", "z64"}

// Given a directory or file, filter down to paths that are probably
// ROM files based on their extensions.
func FindProbableRomsInPath(path string) ([]string, error) {
	var unfilteredPaths []string

	fi, err := os.Stat(path)
	if err != nil {
		return unfilteredPaths, err
	}

	if fi.IsDir() {
		searchPattern := filepath.Join(path, "*")
		unfilteredPaths, err = filepath.Glob(searchPattern)
		if err != nil {
			return unfilteredPaths, err
		}
	} else {
		unfilteredPaths = []string{path}
	}

	probableRomFiles := make([]string, 0)
	for _, fpath := range unfilteredPaths {
		if HasRomExtension(fpath) {
			probableRomFiles = append(probableRomFiles, fpath)
		}
	}

	return probableRomFiles, nil
}

func HasRomExtension(path string) bool {
	ext := filepath.Ext(path)
	for _, v := range Extensions {
		if strings.EqualFold("."+v, ext) {
			return true
		}
	}
	return false
}
