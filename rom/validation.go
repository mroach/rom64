package rom

import (
	"fmt"
	"strings"

	"github.com/mroach/rom64/dat"
)

func (r *RomFile) ValidateWithDat(df dat.DatFile) (matches, mismatches []dat.Rom, err error) {
	if r.File.SHA1 == "" {
		return matches, mismatches, fmt.Errorf("ROM file is missing a SHA-1 hash.")
	}

	if r.File.Format.Code != FormatZ64 {
		return matches, mismatches, fmt.Errorf(
			"File must be in z64 (big-endian) format. This file is %s (%s). The `convert` command can help.",
			r.File.Format.Code,
			r.File.Format.Description)
	}

	serial := r.Serial()
	datroms := df.FindBySerial(serial)
	if len(datroms) == 0 {
		return matches, mismatches, fmt.Errorf("Datfile does not contain an entry for %s", serial)
	}

	for _, item := range datroms {
		sha1Match := strings.EqualFold(item.SHA1, r.File.SHA1)

		if sha1Match {
			matches = append(matches, item)
		} else {
			mismatches = append(mismatches, item)
		}
	}

	return matches, mismatches, err
}
