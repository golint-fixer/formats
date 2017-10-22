// Package til implements access to TIL files.
//
// The TIL file of each level (e.g. "levels/l1data/l1.til") specifies how to
// arrange dungeon pieces - the miniture tiles constructed from the level's MIN
// file - in order to form tiles. A tile consists of 4 dungeon pieces.
//
// Below follows a pseudo-code description of the TIL file format.
//
//    // A TIL file consists of a sequence of tile definitions.
//    type TIL []Tile
//
//    // A Tile consists of four dungeon pieces (top, right, left, bottom),
//    // forming a square.
//    //
//    //     /\        1
//    //    /\/\      3 2
//    //    \/\/       4
//    //     \/
//    type Tile struct {
//       // Dungeon piece ID at the top of the tile.
//       Top uint16
//       // Dungeon piece ID at the right of the tile.
//       Right uint16
//       // Dungeon piece ID at the left of the tile.
//       Left uint16
//       // Dungeon piece ID at the bottom of the tile.
//       Bottom uint16
//    }
package min

import (
	"bufio"
	"encoding/binary"
	"image"
	"image/draw"
	"io"
	"os"

	"github.com/pkg/errors"
	"github.com/sanctuary/formats/level/min"
)

// A Tile consists of four dungeon pieces (top, right, left, bottom), forming a
// square.
//
//
//     /\        1
//    /\/\      3 2
//    \/\/       4
//     \/
type Tile struct {
	// Dungeon piece ID at the top of the tile.
	Top uint16
	// Dungeon piece ID at the right of the tile.
	Right uint16
	// Dungeon piece ID at the left of the tile.
	Left uint16
	// Dungeon piece ID at the bottom of the tile.
	Bottom uint16
}

// Parse parses the given TIL file and returns its tile definitions.
func Parse(path string) ([]Tile, error) {
	// Open file for reading.
	fr, err := os.Open(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer fr.Close()
	br := bufio.NewReader(fr)

	// Decode tiles.
	var tiles []Tile
	for {
		var x [4]uint16
		if err := binary.Read(br, binary.LittleEndian, &x); err != nil {
			if errors.Cause(err) == io.EOF {
				break
			}
			return nil, errors.WithStack(err)
		}
		tile := Tile{
			Top:    uint16(x[0]),
			Right:  uint16(x[1]),
			Left:   uint16(x[2]),
			Bottom: uint16(x[3]),
		}
		tiles = append(tiles, tile)
	}
	return tiles, nil
}

// Image returns an image representation of the tile. The dungeon pieces are
// arranged as illustrated below, forming a square:
//
//           top
//
//            /\
//    left   /\/\   right
//           \/\/
//            \/
//
//          bottom
func (tile Tile) Image(dpieces []min.DPiece, levelFrames []image.Image) image.Image {
	// Generate dungeon piece images.
	top := dpieces[tile.Top].Image(levelFrames)
	right := dpieces[tile.Right].Image(levelFrames)
	left := dpieces[tile.Left].Image(levelFrames)
	bottom := dpieces[tile.Bottom].Image(levelFrames)
	bounds := top.Bounds()

	// The tile is two dungeon pieces in width.
	dpieceWidth := bounds.Dx()
	width := 2 * dpieceWidth

	// The tile is one dungeon piece and one block in height.
	const blockHeight = 32
	height := bounds.Dy() + blockHeight

	// Draw tile image based on dungeon pieces.
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	ptTop := image.Pt(dpieceWidth/2, 0)
	ptRight := image.Pt(dpieceWidth, blockHeight/2)
	ptLeft := image.Pt(0, blockHeight/2)
	ptBottom := image.Pt(dpieceWidth/2, blockHeight)
	draw.Draw(img, bounds.Add(ptTop), top, image.ZP, draw.Over)
	draw.Draw(img, bounds.Add(ptRight), right, image.ZP, draw.Over)
	draw.Draw(img, bounds.Add(ptLeft), left, image.ZP, draw.Over)
	draw.Draw(img, bounds.Add(ptBottom), bottom, image.ZP, draw.Over)
	return img
}
