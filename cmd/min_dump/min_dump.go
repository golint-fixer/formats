// The min_dump tool converts MIN files to PNG images (*.min -> *.png).
//
// Output files are stored to the "_dump_/_dpieces_" directory.
package main

import (
	"flag"
	"fmt"
	"image"
	"log"
	"os"
	"path/filepath"
	"sort"

	"github.com/mewkiz/pkg/imgutil"
	"github.com/mewkiz/pkg/pathutil"
	"github.com/mewkiz/pkg/term"
	"github.com/pkg/errors"
	"github.com/sanctuary/formats/image/cel"
	"github.com/sanctuary/formats/image/cel/config"
	"github.com/sanctuary/formats/level/min"
)

// dbg represents a logger with the "min_dump:" prefix, which logs debug
// messages to standard error.
var dbg = log.New(os.Stderr, term.MagentaBold("min_dump:")+" ", 0)

func usage() {
	const use = `
Convert MIN files to PNG images (*.min -> *.png).

Usage:

	min_dump [OPTION]... FILE.min...

Flags:
`
	fmt.Fprintln(os.Stderr, use[1:])
	flag.PrintDefaults()
}

func main() {
	// Parse command line flags.
	var (
		// mpqDir specifies the path to an extracted "diabdat.mpq".
		mpqDir string
		// all specifies whether to dump all MIN files.
		all bool
	)
	flag.StringVar(&mpqDir, "mpqdir", "diabdat", `Path to extracted "diabdat.mpq".`)
	flag.BoolVar(&all, "a", false, "dump all MIN files")
	flag.Usage = usage
	flag.Parse()
	if !all && flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	// Determine relative MIN paths.
	var relMinPaths []string
	if all {
		relMinPaths = []string{
			"levels/l1data/l1.min",
			"levels/l2data/l2.min",
			"levels/l3data/l3.min",
			"levels/l4data/l4.min",
			"levels/towndata/town.min",
		}
	} else {
		relMinPaths = flag.Args()
	}
	sort.Strings(relMinPaths)

	// Parse MIN files.
	for _, relMinPath := range relMinPaths {
		if err := dumpMin(relMinPath, mpqDir); err != nil {
			log.Fatalf("%+v", err)
		}
	}
}

// dumpMin decodes the given MIN file and displays its contents to standard
// output.
func dumpMin(relMinPath, mpqDir string) error {
	dbg.Printf("Converting %q.", relMinPath)

	// Parse MIN file.
	minPath := filepath.Join(mpqDir, relMinPath)
	dpieces, err := min.Parse(minPath)
	if err != nil {
		return errors.WithStack(err)
	}

	// Parse level CEL frames.
	name := pathutil.FileName(relMinPath)
	relCelPath := fmt.Sprintf("levels/%sdata/%s.cel", name, name)
	celName := filepath.Base(relCelPath)
	conf, err := config.Get(celName)
	if err != nil {
		return errors.WithStack(err)
	}
	for _, relPalPath := range conf.Pals {
		// Parse PAL file.
		palPath := filepath.Join(mpqDir, relPalPath)
		pal, err := cel.ParsePal(palPath)
		if err != nil {
			return errors.WithStack(err)
		}

		// Determine destination directory.
		palDir := ""
		if len(conf.Pals) > 1 {
			palDir = filepath.Base(relPalPath)
		}

		// Parse CEL image.
		celPath := filepath.Join(mpqDir, relCelPath)
		levelFrames, err := cel.DecodeAll(celPath, pal)
		if err != nil {
			return errors.WithStack(err)
		}

		// Dump dungeon pieces of MIN file.
		dstDir := filepath.Join("_dump_", "_dpieces_", name, palDir)
		if err := dumpDPieces(dstDir, dpieces, levelFrames); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

// dumpDPieces converts the dungeon pieces of a MIN file to a set of PNG
// images, where each non-empty block corresponds to a CEL frame from
// levelFrames.
func dumpDPieces(dstDir string, dpieces []min.DPiece, levelFrames []image.Image) error {
	if err := os.MkdirAll(dstDir, 0755); err != nil {
		return errors.WithStack(err)
	}
	for i, dpiece := range dpieces {
		dpieceID := i + 1
		pngName := fmt.Sprintf("dpiece_%04d.png", dpieceID)
		pngPath := filepath.Join(dstDir, pngName)
		img := dpiece.Image(levelFrames)
		if err := imgutil.WriteFile(pngPath, img); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}
