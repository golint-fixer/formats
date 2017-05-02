// TODO: Add high-level description of the CEL file format.

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
	"bytes"
	"encoding/binary"
	"image"
	"image/color"
	"io/ioutil"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/sanctuary/formats/image/cel/config"
)

// DecodeArchive decodes the given CEL archive using colours from the provided
// palette, and returns the sequential frames of the embedded CEL images.
func DecodeArchive(path string, pal color.Palette) ([][]image.Image, error) {
	// Locate image config data.
	name := filepath.Base(path)
	conf, err := config.Get(name)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if conf.Nimgs == 0 {
		return nil, errors.Errorf("invalid call for CEL image %q; use cel.DecodeAll instead")
	}

	// Read file contents.
	archive, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Read the contents of each embedded CEL image.
	cels, err := readCELs(archive)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Decode embedded CEL images.
	archiveImgs := make([][]image.Image, len(cels))
	for i, cel := range cels {
		archiveImgs[i], err = decodeAll(cel, pal, conf)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}

	return archiveImgs, nil
}

// DecodeAll decodes the given CEL image using colours from the provided
// palette, and returns the sequential frames.
func DecodeAll(path string, pal color.Palette) ([]image.Image, error) {
	// Locate image config data.
	name := filepath.Base(path)
	conf, err := config.Get(name)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if conf.Nimgs != 0 {
		return nil, errors.Errorf("invalid call cel.DecodeAll for CEL archive %q; use cel.DecodeArchive instead")
	}

	// Read file contents.
	cel, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Decode CEL image frames.
	return decodeAll(cel, pal, conf)
}

// decodeAll decodes the given CEL image using colours from the provided
// palette, and returns the sequential frames.
func decodeAll(cel []byte, pal color.Palette, conf *config.Config) ([]image.Image, error) {
	// Read the contents of each frame.
	frames, err := readFrames(cel)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Decode frames.
	var imgs []image.Image
	for frameNum, frame := range frames {
		// Determine decoder type based on image config and frame number.
		decode := getDecoder(conf, frameNum)

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

// readCELs returns the contents of each embedded CEL image within the given CEL
// archive.
func readCELs(archive []byte) (cels [][]byte, err error) {
	// Read CEL archive header.
	//
	//    celOffsets [8]uint32 // Offset to each embedded CEL image.
	const ncels = 8
	celOffsets := make([]uint32, ncels+1)
	r := bytes.NewReader(archive)
	if err := binary.Read(r, binary.LittleEndian, celOffsets[:ncels]); err != nil {
		return nil, errors.WithStack(err)
	}

	// Append end offset of the last embedded CEL image.
	celOffsets[ncels] = uint32(len(archive))

	// Read the contents of each embedded CEL image.
	cels = make([][]byte, ncels)
	for i := range cels {
		start, end := celOffsets[i], celOffsets[i+1]
		cels[i] = archive[start:end]
	}

	return cels, nil
}

// readFrames returns the contents of each frame within the given CEL image.
func readFrames(cel []byte) (frames [][]byte, err error) {
	// Read CEL header.
	//
	//    nframes      uint32            // Number of frames.
	//    frameOffsets [nframes+1]uint32 // Offset to each frame.
	r := bytes.NewReader(cel)
	var nframes uint32
	if err := binary.Read(r, binary.LittleEndian, &nframes); err != nil {
		return nil, errors.WithStack(err)
	}
	frameOffsets := make([]uint32, nframes+1)
	if err := binary.Read(r, binary.LittleEndian, frameOffsets); err != nil {
		return nil, errors.WithStack(err)
	}

	// Read the contents of each frame.
	frames = make([][]byte, nframes)
	for i := range frames {
		start, end := frameOffsets[i], frameOffsets[i+1]
		frames[i] = cel[start:end]
	}

	return frames, nil
}
