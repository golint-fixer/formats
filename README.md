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
go get github.com/sanctuary/mpq
mpq -dir diabdat -m diabdat.mpq
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
