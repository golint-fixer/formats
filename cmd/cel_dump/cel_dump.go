// The cel_dump tool converts CEL and CL2 files to PNG images (*.cel -> *.png).
//
// Output files are stored to the "_dump_" directory.
package main

import (
	"flag"
	"fmt"
	"image/color"
	"log"
	"os"
	"path/filepath"
	"sort"

	//"github.com/davecheney/profile"
	"github.com/mewkiz/pkg/imgutil"
	"github.com/mewkiz/pkg/pathutil"
	"github.com/mewkiz/pkg/term"
	"github.com/pkg/errors"
	"github.com/sanctuary/formats/image/cel"
	"github.com/sanctuary/formats/image/cel/config"
)

// dbg represents a logger with the "cel_dump:" prefix, which logs debug
// messages to standard error.
var dbg = log.New(os.Stderr, term.CyanBold("cel_dump:")+" ", 0)

func usage() {
	const use = `
Convert CEL and CL2 files to PNG images (*.cel -> *.png).

Usage:

	cel_dump [OPTION]... FILE...

Flags:
`
	fmt.Fprintln(os.Stderr, use[1:])
	flag.PrintDefaults()
}

func main() {
	//defer profile.Start(profile.CPUProfile).Stop()

	// Parse command line flags.
	var (
		// mpqDir specifies the path to an extracted "diabdat.mpq".
		mpqDir string
		// all specifies whether to dump all CEL images.
		all bool
	)
	flag.StringVar(&mpqDir, "mpqdir", "diabdat", `path to extracted "diabdat.mpq"`)
	flag.BoolVar(&all, "a", false, "dump all CEL images")
	flag.Usage = usage
	flag.Parse()
	if !all && flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	// Determine relative CEL paths.
	var relCelPaths []string
	if all {
		for _, relCelPath := range config.RelPaths {
			relCelPaths = append(relCelPaths, relCelPath)
		}
	} else {
		relCelPaths = flag.Args()
	}
	sort.Strings(relCelPaths)

	// Parse CEL and CL2 files.
	for _, relCelPath := range relCelPaths {
		celName := filepath.Base(relCelPath)
		conf, err := config.Get(celName)
		if err != nil {
			log.Fatalf("%+v", err)
		}
		if conf.Nimgs > 0 {
			if err := dumpArchive(mpqDir, relCelPath, conf); err != nil {
				log.Fatalf("%+v", err)
			}
		} else {
			if err := dumpCel(mpqDir, relCelPath, conf); err != nil {
				log.Fatalf("%+v", err)
			}
		}
	}
}

// dumpArchive converts the given CEL archive to a set of PNG images.
func dumpArchive(mpqDir, relCelPath string, conf *config.Config) error {
	dbg.Printf("Extracting %q.", relCelPath)
	// TODO: Remove temporary hack when the config package containing accurate
	// palette descriptions.
	if len(conf.Pals) == 0 {
		conf.Pals = append(conf.Pals, "levels/towndata/town.pal")
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
		celDir := pathutil.TrimExt(relCelPath)
		dstDir := filepath.Join("_dump_", celDir, palDir)

		// Dump CEL image.
		celPath := filepath.Join(mpqDir, relCelPath)
		if err := dumpArchiveWithPal(dstDir, celPath, pal); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

// dumpArchiveWithPal converts the given CEL archive to a set of PNG images,
// using colours from the given palette.
func dumpArchiveWithPal(dstDir, celPath string, pal color.Palette) error {
	archiveImgs, err := cel.DecodeArchive(celPath, pal)
	if err != nil {
		return errors.WithStack(err)
	}
	if err := os.MkdirAll(dstDir, 0755); err != nil {
		return errors.WithStack(err)
	}
	celName := pathutil.FileName(celPath)
	for i, archiveImg := range archiveImgs {
		archiveName := fmt.Sprintf("%s_%d", celName, i)
		archiveDir := filepath.Join(dstDir, archiveName)
		if err := os.MkdirAll(archiveDir, 0755); err != nil {
			return errors.WithStack(err)
		}

		for j, img := range archiveImg {
			pngName := celName + ".png"
			if len(archiveImg) > 1 {
				pngName = fmt.Sprintf("%s_%04d.png", celName, j+1)
			}
			pngPath := filepath.Join(archiveDir, pngName)
			if err := imgutil.WriteFile(pngPath, img); err != nil {
				return errors.WithStack(err)
			}
		}
	}
	return nil
}

// dumpCel converts the given CEL file to a set of PNG images.
func dumpCel(mpqDir, relCelPath string, conf *config.Config) error {
	dbg.Printf("Converting %q.", relCelPath)
	// TODO: Remove temporary hack when the config package containing accurate
	// palette descriptions.
	if len(conf.Pals) == 0 {
		conf.Pals = append(conf.Pals, "levels/towndata/town.pal")
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
		celDir := pathutil.TrimExt(relCelPath)
		dstDir := filepath.Join("_dump_", celDir, palDir)

		// Dump CEL image.
		celPath := filepath.Join(mpqDir, relCelPath)
		if err := dumpCelWithPal(dstDir, celPath, pal); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

// dumpCelWithPal converts the CEL file to a set of PNG images, using colours
// from the given palette.
func dumpCelWithPal(dstDir, celPath string, pal color.Palette) error {
	imgs, err := cel.DecodeAll(celPath, pal)
	if err != nil {
		return errors.WithStack(err)
	}
	if err := os.MkdirAll(dstDir, 0755); err != nil {
		return errors.WithStack(err)
	}
	for i, img := range imgs {
		celName := pathutil.FileName(celPath)
		pngName := celName + ".png"
		if len(imgs) > 1 {
			pngName = fmt.Sprintf("%s_%04d.png", celName, i+1)
		}
		pngPath := filepath.Join(dstDir, pngName)
		if err := imgutil.WriteFile(pngPath, img); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}
