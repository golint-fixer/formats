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
	"image"
	"image/color"
	"io"
	"os"
	"path/filepath"

	"github.com/mewkiz/pkg/errutil"
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
		relPath, ok := relPaths[name]
		if !ok {
			return nil, errutil.Newf("cel.DecodeAll: unable to locate relative path of %q", name)
		}
		conf, ok := confs[relPath]
		if !ok {
			return nil, errutil.Newf("cel.DecodeAll: unable to locate CEL config for %q", name)
		}
		if conf.nimgs != 0 {
			// TODO: Implement support for CEL archives.
			return nil, errutil.Newf("cel.DecodeAll: support for CEL archives not yet implemented; unable to extract %q", name)
		}
		// Use image dimensions for the specific frame number if present.
		w, ok := conf.frameWidth[frameNum]
		if !ok {
			// Fallback to default frame width.
			w = conf.w
		}
		h, ok := conf.frameHeight[frameNum]
		if !ok {
			// Fallback to default frame height.
			h = conf.h
		}

		// Decode the frame pixel data.
		data := frame[conf.header:] // Skip header contents if present.
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

func decodeType0(data []byte, w, h int, pal color.Palette) image.Image {
	panic("cel.decodeType0: not yet implemented")
}

func decodeType1(data []byte, w, h int, pal color.Palette) image.Image {
	panic("cel.decodeType1: not yet implemented")
}

func decodeType2(data []byte, w, h int, pal color.Palette) image.Image {
	panic("cel.decodeType2: not yet implemented")
}

func decodeType3(data []byte, w, h int, pal color.Palette) image.Image {
	panic("cel.decodeType3: not yet implemented")
}

func decodeType4(data []byte, w, h int, pal color.Palette) image.Image {
	panic("cel.decodeType4: not yet implemented")
}

func decodeType5(data []byte, w, h int, pal color.Palette) image.Image {
	panic("cel.decodeType5: not yet implemented")
}

func decodeType6(data []byte, w, h int, pal color.Palette) image.Image {
	panic("cel.decodeType6: not yet implemented")
}
