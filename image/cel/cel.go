// TODO: Add high-level description of the CEL file format.

//go:generate go run gen.go

// Package cel implements a CEL and CL2 image decoder.
//
// Below follows a pseudo-code description of the CEL file format.
//
//    // A CEL file contains of a file header and a sequence of frames.
//    type CEL struct {
//       // Number of frames.
//       nframes uint32
//       // Offset to each frame.
//       frameOffsets [nframes+1]uint32
//       // Header and pixel data contents of each frame.
//       //
//       //    start: frameOffsets[frameNum]
//       //    end:   frameOffsets[frameNum+1]
//       frames [nframes]Frame
//    }
//
//    // A Frame consists of an optional frame header followed by encoded pixel
//    // data.
//    type Frame struct {
//       // Optional frame header.
//       header []byte
//       // Encoded pixel data.
//       data []byte
//    }
package cel

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"io"
	"os"
	"path/filepath"

	"github.com/mewkiz/pkg/errutil"
	"github.com/sanctuary/formats/image/cel/config"
)

// DecodeAll decodes the given CEL image using colours from the provided
// palette, and returns the sequential frames.
func DecodeAll(path string, pal color.Palette) ([]image.Image, error) {
	// Read the contents of each frame.
	frames, err := readFrames(path)
	if err != nil {
		return nil, errutil.Err(err)
	}

	// Decode frames.
	var imgs []image.Image
	name := filepath.Base(path)
	for frameNum, frame := range frames {
		// Determine decoder type based on file name.
		decode := getDecoder(name, frameNum)

		// Locate image config data.
		relPath, ok := config.RelPaths[name]
		if !ok {
			return nil, errutil.Newf("cel.DecodeAll: unable to locate relative path of %q", name)
		}
		conf, ok := config.Confs[relPath]
		if !ok {
			return nil, errutil.Newf("cel.DecodeAll: unable to locate CEL config for %q", name)
		}
		if conf.Nimgs != 0 {
			// TODO: Implement support for CEL archives.
			panic(fmt.Sprintf("cel.DecodeAll: support for CEL archives not yet implemented; unable to extract %q", name))
		}
		// Use image dimensions for the specific frame number if present.
		w, ok := conf.FrameWidth[frameNum]
		if !ok {
			// Fallback to default frame width.
			w = conf.W
		}
		h, ok := conf.FrameHeight[frameNum]
		if !ok {
			// Fallback to default frame height.
			h = conf.H
		}

		// Decode the frame pixel data.
		data := frame[conf.Header:] // Skip header contents if present.
		img := decode(data, w, h, pal)
		imgs = append(imgs, img)
	}

	return imgs, nil
}

// readFrames returns the contents of each frame width the given CEL image.
func readFrames(path string) (frames [][]byte, err error) {
	// Open file for reading.
	fr, err := os.Open(path)
	if err != nil {
		return nil, errutil.Err(err)
	}
	defer fr.Close()
	br := bufio.NewReader(fr)

	// Read CEL header.
	//
	//    nframes      uint32            // Number of frames.
	//    frameOffsets [nframes+1]uint32 // Offset to each frame.
	var nframes uint32
	if err := binary.Read(br, binary.LittleEndian, &nframes); err != nil {
		return nil, errutil.Err(err)
	}
	frameOffsets := make([]uint32, nframes+1)
	if err := binary.Read(br, binary.LittleEndian, frameOffsets); err != nil {
		return nil, errutil.Err(err)
	}

	// Read the contents of each frame.
	frames = make([][]byte, nframes)
	for i := range frames {
		start, end := frameOffsets[i], frameOffsets[i+1]
		size := end - start
		frames[i] = make([]byte, size)
		if _, err := io.ReadFull(br, frames[i]); err != nil {
			return nil, errutil.Err(err)
		}
	}

	return frames, nil
}

// decodeType0 decodes the pixel data of a type 0 CEL frame of the specified
// dimensions, using colours from the provided palette.
//
// A type 0 CEL frame corresponds to an unencoded 32x32 image without
// transparency, having pixel data arranged as follows, where 'x' represents an
// explicit regular pixel (colour index into the palette).
//
//    +--------------------------------+
//    |xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx|
//    |xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx|
//    |xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx|
//    |xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx|
//    |xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx|
//    |xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx|
//    |xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx|
//    |xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx|
//    |xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx|
//    |xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx|
//    |xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx|
//    |xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx|
//    |xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx|
//    |xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx|
//    |xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx|
//    |xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx|
//    |xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx|
//    |xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx|
//    |xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx|
//    |xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx|
//    |xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx|
//    |xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx|
//    |xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx|
//    |xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx|
//    |xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx|
//    |xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx|
//    |xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx|
//    |xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx|
//    |xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx|
//    |xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx|
//    |xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx|
//    |xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx|
//    +--------------------------------+
//
func decodeType0(data []byte, w, h int, pal color.Palette) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	drawPixel := pixelDrawer(img, w, h)
	for _, b := range data {
		drawPixel(pal[b])
	}
	return img
}

// TODO: Add high-level description of how type 1 pixel data is encoded.

// decodeType1 decodes the pixel data of a regular (type 1) CEL frame of the
// specified dimensions, using colours from the provided palette.
func decodeType1(data []byte, w, h int, pal color.Palette) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	drawPixel := pixelDrawer(img, w, h)
	for pos := 0; pos < len(data); {
		n := int(int8(data[pos]))
		pos++
		switch {
		case n < 0:
			// Transparent pixels.
			n = -n
			for i := 0; i < n; i++ {
				drawPixel(color.Transparent)
			}
		default:
			// Regular pixels.
			for i := 0; i < n; i++ {
				drawPixel(pal[data[pos]])
				pos++
			}
		}
	}
	return img
}

// decodeType2 decodes the pixel data of a type 2 CEL frame of the specified
// dimensions, using colours from the provided palette.
//
// A type 2 CEL frame corresponds to a 32x32 image of a left-facing triangle,
// having pixel data arranged as follows, where 'x' represents an explicit
// regular pixel (colour index into the palette), '0' an explicit transparent
// pixel, and ' ' an implicit transparent pixel. Note the below illustration
// specifies the pixels as they occur within the file, the output image will be
// upside-down.
//
//    +--------------------------------+
//    |                            00xx|
//    |                            xxxx|
//    |                        00xxxxxx|
//    |                        xxxxxxxx|
//    |                    00xxxxxxxxxx|
//    |                    xxxxxxxxxxxx|
//    |                00xxxxxxxxxxxxxx|
//    |                xxxxxxxxxxxxxxxx|
//    |            00xxxxxxxxxxxxxxxxxx|
//    |            xxxxxxxxxxxxxxxxxxxx|
//    |        00xxxxxxxxxxxxxxxxxxxxxx|
//    |        xxxxxxxxxxxxxxxxxxxxxxxx|
//    |    00xxxxxxxxxxxxxxxxxxxxxxxxxx|
//    |    xxxxxxxxxxxxxxxxxxxxxxxxxxxx|
//    |00xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx|
//    |xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx|
//    |00xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx|
//    |    xxxxxxxxxxxxxxxxxxxxxxxxxxxx|
//    |    00xxxxxxxxxxxxxxxxxxxxxxxxxx|
//    |        xxxxxxxxxxxxxxxxxxxxxxxx|
//    |        00xxxxxxxxxxxxxxxxxxxxxx|
//    |            xxxxxxxxxxxxxxxxxxxx|
//    |            00xxxxxxxxxxxxxxxxxx|
//    |                xxxxxxxxxxxxxxxx|
//    |                00xxxxxxxxxxxxxx|
//    |                    xxxxxxxxxxxx|
//    |                    00xxxxxxxxxx|
//    |                        xxxxxxxx|
//    |                        00xxxxxx|
//    |                            xxxx|
//    |                            00xx|
//    |                                |
//    +--------------------------------+
//
func decodeType2(data []byte, w, h int, pal color.Palette) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	drawPixel := pixelDrawer(img, w, h)
	ns := []int{4, 4, 8, 8, 12, 12, 16, 16, 20, 20, 24, 24, 28, 28, 32, 32, 32, 28, 28, 24, 24, 20, 20, 16, 16, 12, 12, 8, 8, 4, 4, 0}
	pos := 0
	for i, n := range ns {
		if i%2 == 0 {
			// Even lines start with two explicit transparent pixels.
			if data[pos] != 0 || data[pos+1] != 0 {
				panic(fmt.Sprintf("explicit transparent pixel mismatch; expected 0x00 0x00, got 0x%02X 0x%02X", data[pos], data[pos+1]))
			}
			n -= 2 // Skip explicit transparent pixels.
			pos += 2
		}
		for i := n; i < 32; i++ {
			// Transparent pixels.
			drawPixel(color.Transparent)
		}
		for i := 0; i < n; i++ {
			// Regular pixels.
			drawPixel(pal[data[pos]])
			pos++
		}
	}
	return img
}

// decodeType3 decodes the pixel data of a type 3 CEL frame of the specified
// dimensions, using colours from the provided palette.
func decodeType3(data []byte, w, h int, pal color.Palette) image.Image {
	panic("cel.decodeType3: not yet implemented")
}

// decodeType4 decodes the pixel data of a type 4 CEL frame of the specified
// dimensions, using colours from the provided palette.
func decodeType4(data []byte, w, h int, pal color.Palette) image.Image {
	panic("cel.decodeType4: not yet implemented")
}

// decodeType5 decodes the pixel data of a type 5 CEL frame of the specified
// dimensions, using colours from the provided palette.
func decodeType5(data []byte, w, h int, pal color.Palette) image.Image {
	panic("cel.decodeType5: not yet implemented")
}

// TODO: Add high-level description of how type 6 pixel data is encoded.

// decodeType6 decodes the pixel data of a regular (type 6) CL2 frame of the
// specified dimensions, using colours from the provided palette.
func decodeType6(data []byte, w, h int, pal color.Palette) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	drawPixel := pixelDrawer(img, w, h)
	for pos := 0; pos < len(data); {
		n := int(int8(data[pos]))
		pos++
		switch {
		case n < -65:
			// Run-length encoded pixels.
			n = -n - 65
			c := pal[data[pos]]
			for i := 0; i < n; i++ {
				drawPixel(c)
			}
			pos++
		case n < 0:
			// Regular pixels.
			n = -n
			for i := 0; i < n; i++ {
				drawPixel(pal[data[pos]])
				pos++
			}
		default:
			// Transparent pixels.
			for i := 0; i < n; i++ {
				drawPixel(color.Transparent)
			}
		}
	}
	return img
}

// pixelDrawer returns a function which may be invoked to incrementally set
// pixels; starting in the lower left corner, going from left to right, and then
// row by row from the bottom to the top of the image.
func pixelDrawer(dst draw.Image, w, h int) func(color.Color) {
	x, y := 0, h-1
	return func(c color.Color) {
		// TODO: Remove sanity check once the cel decoder library has mature.
		if x < 0 || x >= w {
			panic(fmt.Sprintf("cel.pixelDrawer.drawPixel: invalid x; expected 0 <= x < %d, got x=%d", w, x))
		}
		if y < 0 || y >= h {
			panic(fmt.Sprintf("cel.pixelDrawer.drawPixel: invalid y; expected 0 <= y < %d, got y=%d", h, y))
		}
		dst.Set(x, y, c)
		x++
		if x >= w {
			x = 0
			y--
		}
	}
}
