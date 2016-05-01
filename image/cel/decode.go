package cel

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"

	"github.com/sanctuary/formats/image/cel/config"
)

// decoders maps CEL frame types to decoder functions.
var decoders = [...]func(data []byte, w, h int, pal color.Palette) image.Image{
	0: decodeType0,
	1: decodeType1,
	2: decodeType2,
	3: decodeType3,
	4: decodeType4,
	5: decodeType5,
	6: decodeType6,
}

// getDecoder returns the CEL frame decoder of the given image config and frame
// number.
func getDecoder(conf *config.Config, frameNum int) func(data []byte, w, h int, pal color.Palette) image.Image {
	return decoders[conf.GetDecoderType(frameNum)]
}

// levelFrameWidth specifies the frame width of level CELs.
const levelFrameWidth = 32

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
	ns := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 30, 28, 26, 24, 22, 20, 18, 16, 14, 12, 10, 8, 6, 4, 2, 0}
	pos := 0
	for i, n := range ns {
		// Even lines end with two explicit transparent pixels.
		if i%2 == 0 {
			for j := 0; j < 2; j++ {
				if data[pos] != 0 {
					panic(fmt.Sprintf("explicit transparent pixel mismatch; expected 0x00, got 0x%02X", data[pos]))
				}
				pos++
			}
		}
		// Transparent pixels.
		for j := n; j < levelFrameWidth; j++ {
			drawPixel(color.Transparent)
		}
		// Regular pixels.
		for j := 0; j < n; j++ {
			drawPixel(pal[data[pos]])
			pos++
		}
	}
	return img
}

// decodeType3 decodes the pixel data of a type 3 CEL frame of the specified
// dimensions, using colours from the provided palette.
//
// A type 3 CEL frame corresponds to a 32x32 image of a right-facing triangle,
// having pixel data arranged as follows, where 'x' represents an explicit
// regular pixel (colour index into the palette), '0' an explicit transparent
// pixel, and ' ' an implicit transparent pixel. Note the below illustration
// specifies the pixels as they occur within the file, the output image will be
// upside-down.
//
//    +--------------------------------+
//    |xx00                            |
//    |xxxx                            |
//    |xxxxxx00                        |
//    |xxxxxxxx                        |
//    |xxxxxxxxxx00                    |
//    |xxxxxxxxxxxx                    |
//    |xxxxxxxxxxxxxx00                |
//    |xxxxxxxxxxxxxxxx                |
//    |xxxxxxxxxxxxxxxxxx00            |
//    |xxxxxxxxxxxxxxxxxxxx            |
//    |xxxxxxxxxxxxxxxxxxxxxx00        |
//    |xxxxxxxxxxxxxxxxxxxxxxxx        |
//    |xxxxxxxxxxxxxxxxxxxxxxxxxx00    |
//    |xxxxxxxxxxxxxxxxxxxxxxxxxxxx    |
//    |xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx00|
//    |xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx|
//    |xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx00|
//    |xxxxxxxxxxxxxxxxxxxxxxxxxxxx    |
//    |xxxxxxxxxxxxxxxxxxxxxxxxxx00    |
//    |xxxxxxxxxxxxxxxxxxxxxxxx        |
//    |xxxxxxxxxxxxxxxxxxxxxx00        |
//    |xxxxxxxxxxxxxxxxxxxx            |
//    |xxxxxxxxxxxxxxxxxx00            |
//    |xxxxxxxxxxxxxxxx                |
//    |xxxxxxxxxxxxxx00                |
//    |xxxxxxxxxxxx                    |
//    |xxxxxxxxxx00                    |
//    |xxxxxxxx                        |
//    |xxxxxx00                        |
//    |xxxx                            |
//    |xx00                            |
//    |                                |
//    +--------------------------------+
//
func decodeType3(data []byte, w, h int, pal color.Palette) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	drawPixel := pixelDrawer(img, w, h)
	ns := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 30, 28, 26, 24, 22, 20, 18, 16, 14, 12, 10, 8, 6, 4, 2, 0}
	pos := 0
	for i, n := range ns {
		// Regular pixels.
		for j := 0; j < n; j++ {
			drawPixel(pal[data[pos]])
			pos++
		}
		// Even lines end with two explicit transparent pixels.
		if i%2 == 0 {
			for j := 0; j < 2; j++ {
				if data[pos] != 0 {
					panic(fmt.Sprintf("explicit transparent pixel mismatch; expected 0x00, got 0x%02X", data[pos]))
				}
				pos++
			}
		}
		// Transparent pixels.
		for j := n; j < levelFrameWidth; j++ {
			drawPixel(color.Transparent)
		}
	}
	return img
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
