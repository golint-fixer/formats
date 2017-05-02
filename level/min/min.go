// Package min implements access to MIN files.
//
// The MIN file of each level (e.g. "levels/l1data/l1.min") specifies how to
// arrange the frames of the level's CEL file (e.g. "levels/l1data/l1.cel") in
// order to form miniature tiles (henceforth referred to as dungeon pieces, or
// dpieces for short). A dungeon piece consists of either 10 or 16 blocks, where
// each non-empty block is represented by a CEL frame.
//
// Below follows a pseudo-code description of the MIN file format.
//
//    // A MIN file consists of a sequence of dungeon piece definitions.
//    type MIN []DPiece
//
//    // A DPiece consists of a sequence of either 10 or 16 block definitions.
//    type DPiece struct {
//       // nblocks is either 10 (for "l1.min", "l2.min" and "l3.min") or 16
//       // (for "l4.min" and "town.min").
//       blocks [nblocks]Block
//    }
//
//    // A Block stores the CEL frame number and frame type of a dungeon piece
//    // block.
//    //
//    //    frameNum  := block&0x0FFF
//    //    frameType := block&0x7000 >> 12
//    type Block uint16
package min

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"image"
	"image/draw"
	"io"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// A DPiece represents a dungeon piece, which specifies how to arrange frames of
// a level CEL file in order to form a miniature tile. A dungeon piece consists
// of 10 or 16 blocks, where each non-empty block is represented by a CEL frame.
type DPiece struct {
	// Either 10 or 16 blocks constituting the dungeon piece.
	Blocks []Block
}

// A Block specifies the graphics of a single block in a dungeon piece,
// consisting of either 10 or 16 blocks.
type Block struct {
	// Frame number in the level CEL file; or 0 if empty.
	FrameNum int
	// Frame type, specifying the CEL decoding algorithm of the frame.
	FrameType int
}

// Parse parses the given MIN file and returns its dungeon piece definitions.
func Parse(path string) ([]DPiece, error) {
	// Open file for reading.
	fr, err := os.Open(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer fr.Close()
	br := bufio.NewReader(fr)

	// Allocate block buffer.
	var nblocks int
	switch name := filepath.Base(path); name {
	case "l1.min", "l2.min", "l3.min":
		nblocks = 10
	case "l4.min", "town.min":
		nblocks = 16
	default:
		panic(fmt.Errorf("min.Parse: support for MIN file %q not yet implemented", name))
	}
	buf := make([]uint16, nblocks)

	// Decode dungeon pieces.
	var dpieces []DPiece
	for {
		if err := binary.Read(br, binary.LittleEndian, buf); err != nil {
			if err == io.EOF {
				break
			}
			return nil, errors.WithStack(err)
		}
		dpiece := DPiece{
			Blocks: make([]Block, nblocks),
		}
		for i := range buf {
			block := Block{
				FrameNum:  int(buf[i] & 0x0FFF),
				FrameType: int(buf[i] & 0x7000 >> 12),
			}
			dpiece.Blocks[i] = block
		}
		dpieces = append(dpieces, dpiece)
	}

	return dpieces, nil
}

// Image returns an image representation of the dungeon piece, where each non-
// empty block corresponds to a CEL frame from levelFrames.
//
// The size of each block is 32x32 pixels and the blocks are arranged as
// illustrated below to form a miniature tile.
//
//    +----+----+
//    |  0 |  1 |
//    +----+----+
//    |  2 |  3 |
//    +----+----+
//    |  4 |  5 |
//    +----+----+
//    |  6 |  7 |
//    +----+----+
//    |  8 |  9 |
//    +----+----+
//    | 10 | 11 |
//    +----+----+
//    | 12 | 13 |
//    +----+----+
//    | 14 | 15 |
//    +----+----+
func (dpiece DPiece) Image(levelFrames []image.Image) image.Image {
	const (
		blockWidth  = 32
		blockHeight = 32
	)
	width := blockWidth * 2
	height := blockHeight * (len(dpiece.Blocks) / 2)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for blockNum, block := range dpiece.Blocks {
		if block.FrameNum != 0 {
			frame := levelFrames[block.FrameNum-1]
			x := blockWidth * (blockNum % 2)
			y := blockHeight * (blockNum / 2)
			dr := image.Rect(x, y, x+blockWidth, y+blockHeight)
			draw.Draw(img, dr, frame, image.ZP, draw.Src)
		}
	}
	return img
}
