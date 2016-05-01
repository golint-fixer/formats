package cel

import (
	"image"
	"image/color"

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
