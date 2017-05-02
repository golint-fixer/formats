package cel

import (
	"image/color"
	"io/ioutil"

	"github.com/pkg/errors"
)

// ParsePal parses the given PAL file and returns the corresponding palette.
//
// Below follows a pseudo-code description of the PAL file format.
//
//    // A PAL file contains a sequence of colour definitions, representing a
//    // palette.
//    type PAL [256]Color
//
//    // A Color represents a colour specified by red, green and blue intensity
//    // levels.
//    type Color struct {
//       red, green, blue byte
//    }
func ParsePal(path string) (color.Palette, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	const (
		// Number of colours within a palette.
		ncolors = 256
		// The size of each colour in bytes.
		colorSize = 3
	)
	if len(buf) != ncolors*colorSize {
		return nil, errors.Errorf("invalid PAL file size for %q; expected %d, got %d", path, ncolors*colorSize, len(buf))
	}
	pal := make(color.Palette, ncolors)
	for i := range pal {
		pal[i] = color.RGBA{
			R: buf[i*colorSize],
			G: buf[i*colorSize+1],
			B: buf[i*colorSize+2],
			A: 0xFF,
		}
	}
	return pal, nil
}
