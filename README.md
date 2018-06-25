# formats

[![Build Status](https://travis-ci.org/sanctuary/formats.svg)](https://travis-ci.org/sanctuary/formats)
[![Coverage Status](https://coveralls.io/repos/github/sanctuary/formats/badge.svg)](https://coveralls.io/github/sanctuary/formats)
[![GoDoc](https://godoc.org/github.com/sanctuary/formats?status.svg)](https://godoc.org/github.com/sanctuary/formats)

The aim of this project is to provide open source reference decoders for the file formats of the Diablo 1 game engine.

## Installation

```bash
go get -u github.com/sanctuary/formats/...
```

## Usage

The `cel_dump` and `min_dump` tools search for game assets in the `diabdat/` directory, which should contains the extracted files of `diabdat.mpq`.

### Extract diabdat.mpq

```bash
# Extract DIABDAT.MPQ archive.
go -u get github.com/sanctuary/mpq
mpq -dir diabdat -m diabdat.mpq
```

### Fix broken files in diabdat.mpq

The original `diabdat.mpq` archive contains three broken files, `levels/l1data/banner2.dun`, `monsters/darkmage/dmagew.cl2` and `monsters/unrav/unravw.cel`. All of which can be fixed by running the `mpqfix` tool on the `diabdat/` directory containing the extracted game assets.

```bash
go get -u github.com/mewrnd/blizzconv/cmd/mpqfix
mpqfix -mpqdump diabdat
```

### Dump CEL files

```bash
# Convert all CEL and CL2 files into PNG format.
#
# The command takes ~15 minutes to complete.
cel_dump -a
```

### Dump MIN files

```bash
# Convert all MIN files into PNG format.
#
# The command takes ~1 minute to complete.
min_dump -a
```
