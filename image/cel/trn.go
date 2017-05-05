package cel

import (
	"io/ioutil"

	"github.com/pkg/errors"
)

// ParseTrn parses the given TRN file and returns the corresponding colour
// transition table.
//
// Below follows a pseudo-code description of the TRN file format.
//
//    // A TRN file contains a sequence of colour transitions, representing
//    // indexes into a palette.
//    type TRN [256]uint8
func ParseTrn(path string) ([256]uint8, error) {
	var trn [256]uint8
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return trn, errors.WithStack(err)
	}
	for i, b := range buf {
		trn[i] = b
	}
	return trn, nil
}
