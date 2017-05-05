package cel

import (
	"image/color"
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
func ParseTrn(path string) (*TransitionTable, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	trn := &TransitionTable{}
	for i, b := range buf {
		trn.Indices[i] = b
	}
	return trn, nil
}

// A TransitionTable represents a colour transition table.
type TransitionTable struct {
	// Indices maps from TRN index to palette index.
	Indices [256]uint8
}

// Pal returns a new palette created by resolving colours from the source
// palette using indices from the colour transition table.
func (trn *TransitionTable) Pal(src color.Palette) color.Palette {
	dst := make(color.Palette, len(src))
	for i, t := range trn.Indices {
		dst[i] = src[t]
	}
	return dst
}
