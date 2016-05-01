package cel_test

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"image"
	"image/color"
	"log"
	"path/filepath"
	"testing"

	"github.com/mewkiz/pkg/osutil"
	"github.com/sanctuary/formats/image/cel"
	"github.com/sanctuary/formats/image/cel/config"
)

// palTest specifies the test cases of a given palette.
type palTest struct {
	// Relative PAL path to "diabdat.mpq".
	relPalPath string
	// Expected SHA1 hashsums of the raw pixel data for each frame decoded using
	// the given palette; hashing the pixels from left to right, and top to
	// bottom. The colour of each pixel is represented in RGBA order, using
	// 8-bits for the red, green, blue and alpha channel levels, respectively.
	wants [][sha1.Size]byte
}

func TestDecodeAll(t *testing.T) {
	// mpqDir specifies the path to an extracted "diabdat.mpq".
	mpqDir := "diabdat/"

	// Skip test if extracted "diabdat.mpq" is not present.
	if exist, _ := osutil.Exists(mpqDir); !exist {
		t.Skipf("%q directory not present", mpqDir)
		return
	}

	golden := []struct {
		// Relative CEL paths to "diabdat.mpq".
		relCelPath string
		// Palette tests.
		palTests []palTest
	}{
		{
			relCelPath: "ctrlpan/golddrop.cel",
			palTests: []palTest{
				{
					relPalPath: "levels/towndata/town.pal",
					wants: [][sha1.Size]byte{
						{0x93, 0xAE, 0x68, 0x80, 0x98, 0xCB, 0xB3, 0x48, 0x9D, 0xF0, 0x9B, 0xF0, 0x5C, 0xBB, 0xBC, 0xF9, 0x3D, 0x83, 0x0B, 0xB8},
					},
				},
			},
		},
		{
			relCelPath: "ctrlpan/p8bulbs.cel",
			palTests: []palTest{
				{
					relPalPath: "levels/towndata/town.pal",
					wants: [][sha1.Size]byte{
						{0x49, 0xA7, 0xE3, 0x9E, 0x98, 0xFB, 0xDB, 0x27, 0x53, 0xAE, 0x42, 0x83, 0x7D, 0x68, 0x4A, 0x13, 0x64, 0x71, 0x78, 0x2F},
						{0x8C, 0x42, 0xB9, 0x57, 0xB3, 0x9B, 0x25, 0xCF, 0x8B, 0x86, 0x5A, 0x51, 0xE0, 0x04, 0x12, 0x70, 0x57, 0x3F, 0x04, 0xD1},
					},
				},
			},
		},
	}

	for _, g := range golden {
		// Get CEL file config.
		conf, err := config.Get(filepath.Base(g.relCelPath))
		if err != nil {
			t.Errorf("unable to locate config for %q; %v", g.relCelPath, err)
			continue
		}
		// TODO: Remove temporary hack when the config package containg accurate
		// palette descriptions.
		if len(conf.Pals) < 1 {
			log.Printf(`Adding default palette ("town.pal") to incomplete config for %q.`, g.relCelPath)
			conf.Pals = append(conf.Pals, "levels/towndata/town.pal")
		}
		// TODO: Remove sanity check once the test cases have matured.
		if len(g.palTests) < 1 {
			t.Errorf("%q: incomplete test; no palette test cases", g.relCelPath)
			continue
		}
		if len(conf.Pals) != len(g.palTests) {
			t.Errorf("%q: palette count mismatch; expected %d, got %d", g.relCelPath, len(g.palTests), len(conf.Pals))
			continue
		}

		// Decode CEL frames for each of the given palettes and compare against
		// the expected SHA1 hashsums.
		celPath := filepath.Join(mpqDir, g.relCelPath)
		for _, palTest := range g.palTests {
			// TODO: Remove sanity check once the test cases have matured.
			wants := palTest.wants
			if len(wants) < 1 {
				t.Errorf("%q: incomplete test; no expected SHA1 hashes for the decoded frames of palette %q", g.relCelPath, palTest.relPalPath)
				continue
			}

			// Decode CEL frames for the given palette.
			palPath := filepath.Join(mpqDir, palTest.relPalPath)
			pal, err := cel.ParsePal(palPath)
			if err != nil {
				t.Errorf("%q: unable to parse palette %q; %v", g.relCelPath, palTest.relPalPath, err)
				continue
			}
			imgs, err := cel.DecodeAll(celPath, pal)
			if err != nil {
				t.Errorf("%q: unable to decode CEL frames; %v", g.relCelPath, err)
				continue
			}
			if len(imgs) != len(wants) {
				t.Errorf("%q: frame count mismatch; expected %d, got %d", g.relCelPath, len(wants), len(imgs))
			}

			// Compare the raw pixel data of the decoded frames to the expected
			// SHA1 hashsums.
			for frameNum, img := range imgs {
				if frameNum >= len(wants) {
					t.Errorf("%q: invalid number of frames; expected %d, got %d", g.relCelPath, len(wants), len(imgs))
					break
				}
				got := hashImage(img)
				if want := wants[frameNum]; got != want {
					// TODO: Remove once the test cases have been implemented.
					fmt.Print("{")
					for i, b := range got[:] {
						if i != 0 {
							fmt.Print(", ")
						}
						fmt.Printf("0x%02X", b)
					}
					fmt.Println("},")

					t.Errorf("%q: raw pixel data hash mismatch for frame number %d and palette %q; expected %032X, got %032X", g.relCelPath, frameNum, palTest.relPalPath, want, got)
					continue
				}
			}
		}
	}
}

// hashImage returns a SHA1 hashsum of the raw pixel data for the given image;
// hashing the pixels from left to right, and top to bottom. The colour of each
// pixel is represented in RGBA order, using 8-bits for the red, green, blue and
// alpha channel levels, respectively.
func hashImage(img image.Image) [sha1.Size]byte {
	data := new(bytes.Buffer)
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := img.At(x, y)
			switch c := c.(type) {
			case color.RGBA:
				data.WriteByte(c.R)
				data.WriteByte(c.G)
				data.WriteByte(c.B)
				data.WriteByte(c.A)
			default:
				panic(fmt.Sprintf("cel_test.hashImage: support for %T not yet implemented", c))
			}
		}
	}
	return sha1.Sum(data.Bytes())
}
