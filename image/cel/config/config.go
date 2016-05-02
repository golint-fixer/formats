//go:generate go run gen.go

// Package config specifies the data required for decoding CEL images.
package config

import "github.com/mewkiz/pkg/errutil"

// Get returns the image config data of the given CEL image.
func Get(name string) (*Config, error) {
	relPath, ok := RelPaths[name]
	if !ok {
		return nil, errutil.Newf("config.Get: unable to locate relative path of %q", name)
	}
	conf, ok := confs[relPath]
	if !ok {
		return nil, errutil.Newf("config.Get: unable to locate CEL config for %q", name)
	}
	conf.GetDecoderType = func(frameNum int) int {
		return getDecoderType(name, frameNum)
	}
	return conf, nil
}

// A Config specifies the data required for decoding a given CEL image.
type Config struct {
	// Number of embedded images; a non-zero value implies that the given file is
	// a CEL archive.
	Nimgs int
	// Header size in bytes.
	Header int
	// Default frame dimensions.
	W, H int
	// Specific frame dimensions, mapping from frame number to width or height.
	FrameWidth, FrameHeight map[int]int
	// Palette paths.
	Pals []string
	// Colour transition paths.
	Trns []string
	// GetDecoderType returns the CEL frame decoder type of the given frame
	// number. The decoder type may be one of the following.
	//
	//    (0) cel.decodeType0
	//    (1) cel.decodeType1
	//    (2) cel.decodeType2
	//    (3) cel.decodeType3
	//    (4) cel.decodeType4
	//    (5) cel.decodeType5
	//    (6) cel.decodeType6
	GetDecoderType func(frameNum int) int
}

// TODO: Remove unknown once no longer needed.

// unknown tracks unknown data.
const unknown = 0

// TODO: Add palette paths.

// confs specifies the data required for decoding
var confs = map[string]*Config{
	// CEL files.
	"ctrlpan/golddrop.cel": {
		W: 261, // ref: 0x406B12
		H: 136, // h = npixels/w = 35496/261 = 136
	},
	"ctrlpan/p8bulbs.cel": {
		W: 88, // ref: 0x404707
		H: 88, // h = npixels/w = 7744/88 = 88
	},
	"ctrlpan/p8but2.cel": {
		W: 33, // ref: 0x4049EF
		H: 32, // h = npixels/w = 1056/33 = 32
	},
	"ctrlpan/panel8.cel": {
		W: 640, // ref: 0x4046C7
		H: 144, // h = npixels/w = 92160/640 = 144
	},
	"ctrlpan/panel8bu.cel": {
		W: 71, // ref: 0x404993
		H: 19, // h = npixels/w = 1349/71 = 19
	},
	"ctrlpan/smaltext.cel": {
		W: 13, // ref: 0x404167
		H: 11, // h = npixels/w = 143/13 = 11
	},
	"ctrlpan/spelicon.cel": {
		W: 56, // ref: 0x403E98
		H: 56, // h = npixels/w = 3136/56 = 56
	},
	"ctrlpan/talkbutt.cel": {
		W: 61, // ref: 0x407019
		H: 16, // h = npixels/w = 976/61 = 16
	},
	"ctrlpan/talkpanl.cel": {
		W: 640, // ref: 0x4046C7
		H: 144, // h = npixels/w = 92160/640 = 144
	},
	"data/bigtgold.cel": {
		W: 46, // ref: 0x419EEA
		H: 45, // h = npixels/w = 2070/46 = 45
	},
	"data/char.cel": {
		W: 320, // ref: 0x4056E1
		H: 352, // h = npixels/w = 112640/320 = 352
	},
	"data/charbut.cel": {
		W: 41, // ref: 0x406267
		H: 22, // h = npixels/w = 902/41 = 22
		FrameWidth: map[int]int{
			// NOTE: Unused frame 0?
			0: 95, // w = npixels/h = 2090/22 = 95
		},
	},
	"data/diabsmal.cel": {
		W: 296, // ref: 0x41A0D1
		H: 100, // h = npixels/w = 29600/296 = 100
	},
	"data/inv/inv.cel": {
		W: 320, // ref: 0x41B8D8
		H: 352, // h = npixels/w = 112640/320 = 352
	},
	"data/inv/inv_rog.cel": {
		W: 320, // ref: 0x41B8D8
		H: 352, // h = npixels/w = 112640/320 = 352
	},
	"data/inv/inv_sor.cel": {
		W: 320, // ref: 0x41B8D8
		H: 352, // h = npixels/w = 112640/320 = 352
	},
	"data/inv/objcurs.cel": {
		// The contents of frame 0 at offset 0x2D4 in objcurs.cel starts with the
		// following header: 0A 00 00 00 00 00 00 00 00 00.
		Header: 10,
		// There are 93 occurrences of width 56 in objcurs_frame_width_from_frame.
		W: 56,
		// There are 87 occurrences of height 84 in objcurs_frame_height_from_frame.
		H: 84,
		// ref: https://github.com/sanctuary/notes/blob/master/rdata/cursor.md#objcurs_frame_width_from_frame
		FrameWidth: map[int]int{
			0:  33,
			1:  32,
			2:  32,
			3:  32,
			4:  32,
			5:  32,
			6:  32,
			7:  32,
			8:  32,
			9:  32,
			10: 23,
			11: 28,
			12: 28,
			13: 28,
			14: 28,
			15: 28,
			16: 28,
			17: 28,
			18: 28,
			19: 28,
			20: 28,
			21: 28,
			22: 28,
			23: 28,
			24: 28,
			25: 28,
			26: 28,
			27: 28,
			28: 28,
			29: 28,
			30: 28,
			31: 28,
			32: 28,
			33: 28,
			34: 28,
			35: 28,
			36: 28,
			37: 28,
			38: 28,
			39: 28,
			40: 28,
			41: 28,
			42: 28,
			43: 28,
			44: 28,
			45: 28,
			46: 28,
			47: 28,
			48: 28,
			49: 28,
			50: 28,
			51: 28,
			52: 28,
			53: 28,
			54: 28,
			55: 28,
			56: 28,
			57: 28,
			58: 28,
			59: 28,
			60: 28,
			61: 28,
			62: 28,
			63: 28,
			64: 28,
			65: 28,
			66: 28,
			67: 28,
			68: 28,
			69: 28,
			70: 28,
			71: 28,
			72: 28,
			73: 28,
			74: 28,
			75: 28,
			76: 28,
			77: 28,
			78: 28,
			79: 28,
			80: 28,
			81: 28,
			82: 28,
			83: 28,
			84: 28,
			85: 28,
		},
		// ref: https://github.com/sanctuary/notes/blob/master/rdata/cursor.md#objcurs_frame_height_from_frame
		FrameHeight: map[int]int{
			0:   29,
			1:   32,
			2:   32,
			3:   32,
			4:   32,
			5:   32,
			6:   32,
			7:   32,
			8:   32,
			9:   32,
			10:  35,
			11:  28,
			12:  28,
			13:  28,
			14:  28,
			15:  28,
			16:  28,
			17:  28,
			18:  28,
			19:  28,
			20:  28,
			21:  28,
			22:  28,
			23:  28,
			24:  28,
			25:  28,
			26:  28,
			27:  28,
			28:  28,
			29:  28,
			30:  28,
			31:  28,
			32:  28,
			33:  28,
			34:  28,
			35:  28,
			36:  28,
			37:  28,
			38:  28,
			39:  28,
			40:  28,
			41:  28,
			42:  28,
			43:  28,
			44:  28,
			45:  28,
			46:  28,
			47:  28,
			48:  28,
			49:  28,
			50:  28,
			51:  28,
			52:  28,
			53:  28,
			54:  28,
			55:  28,
			56:  28,
			57:  28,
			58:  28,
			59:  28,
			60:  28,
			61:  56,
			62:  56,
			63:  56,
			64:  56,
			65:  56,
			66:  56,
			86:  56,
			87:  56,
			88:  56,
			89:  56,
			90:  56,
			91:  56,
			92:  56,
			93:  56,
			94:  56,
			95:  56,
			96:  56,
			97:  56,
			98:  56,
			99:  56,
			100: 56,
			101: 56,
			102: 56,
			103: 56,
			104: 56,
			105: 56,
			106: 56,
			107: 56,
			108: 56,
			109: 56,
			110: 56,
		},
	},
	"data/medtexts.cel": {
		W: 22, // ref: 0x4281B2
		H: 22, // h = npixels/w = 484/22 = 22
	},
	"data/optbar.cel": {
		W: 287, // ref: 0x41A161
		H: 32,  // h = npixels/w = 9184/287 = 32
	},
	"data/option.cel": {
		W: 27, // ref: 0x41A1B6
		H: 28, // h = npixels/w = 756/27 = 28
	},
	"data/pentspin.cel": {
		W: 48, // ref: 0x41A204
		H: 48, // h = npixels/w = 2304/48 = 48
	},
	"data/pentspn2.cel": {
		W: 12, // ref: 0x406C14
		H: 12, // h = npixels/w = 144/12 = 12
	},
	"data/quest.cel": {
		W: 320, // ref: 0x4525E7
		H: 352, // h = npixels/w = 112640/320 = 352
	},
	"data/spellbk.cel": {
		W: 320, // ref: 0x406671
		H: 352, // h = npixels/w = 112640/320 = 352
	},
	"data/spellbkb.cel": {
		W: 76, // ref: 0x406697
		H: 29, // h = npixels/w = 2204/76 = 29
	},
	"data/spelli2.cel": {
		W: 37, // ref: 0x40673F
		H: 38, // h = npixels/w = 1406/37 = 38
	},
	"data/square.cel": {
		// The contents of frame 0 at offset 0xC in square.cel starts with the
		// following header: 0A 00 8C 01 AC 01 CC 01 00 00.
		Header: 10,
		W:      64,  // ref: 0x4552A1
		H:      128, // h = npixels/w = 8192/64 = 128
	},
	"data/textbox.cel": {
		W: 591, // ref: 0x428105
		H: 303, // h = npixels/w = 179073/591 = 303
	},
	"data/textbox2.cel": {
		W: 271, // ref: 0x457B79
		H: 303, // h = npixels/w = 82113/271 = 303
	},
	"data/textslid.cel": {
		W: 12, // ref: 0x4180B2
		H: 12, // h = npixels/w = 144/12 = 12
	},
	"gendata/cut2.cel": {
		W: 640, // ref: 0x41B225
		H: 480, // h = npixels/w = 307200/640 = 480
		Pals: []string{
			"gendata/cut2.pal", // ref: 0x41B674
		},
	},
	"gendata/cut3.cel": {
		W: 640, // ref: 0x41B225
		H: 480, // h = npixels/w = 307200/640 = 480
		Pals: []string{
			"gendata/cut3.pal", // ref: 0x41B7AA
		},
	},
	"gendata/cut4.cel": {
		W: 640, // ref: 0x41B225
		H: 480, // h = npixels/w = 307200/640 = 480
		Pals: []string{
			"gendata/cut4.pal", // ref: 0x41B797
		},
	},
	"gendata/cutgate.cel": {
		W: 640, // ref: 0x41B225
		H: 480, // h = npixels/w = 307200/640 = 480
		Pals: []string{
			"gendata/cutgate.pal", // ref: 0x41B65E
		},
	},
	"gendata/cutl1d.cel": {
		W: 640, // ref: 0x41B225
		H: 480, // h = npixels/w = 307200/640 = 480
		Pals: []string{
			"gendata/cutl1d.pal", // ref: 0x41B69E
		},
	},
	"gendata/cutportl.cel": {
		W: 640, // ref: 0x41B225
		H: 480, // h = npixels/w = 307200/640 = 480
		Pals: []string{
			"gendata/cutportl.pal", // ref: 0x41B748
		},
	},
	"gendata/cutportr.cel": {
		W: 640, // ref: 0x41B225
		H: 480, // h = npixels/w = 307200/640 = 480
		Pals: []string{
			"gendata/cutportr.pal", // ref: 0x41B700
		},
	},
	"gendata/cutstart.cel": {
		W: 640, // ref: 0x41B225
		H: 480, // h = npixels/w = 307200/640 = 480
		Pals: []string{
			"gendata/cutstart.pal", // ref: 0x41B75B
		},
	},
	"gendata/cuttt.cel": {
		W: 640, // ref: 0x41B225
		H: 480, // h = npixels/w = 307200/640 = 480
		Pals: []string{
			"gendata/cuttt.pal", // ref: 0x41B7D1
		},
	},
	// NOTE: Unused?
	"gendata/quotes.cel": {
		W: 640,
		H: 480, // h = npixels/w = 307200/640 = 480
		Pals: []string{
			"gendata/quotes.pal",
		},
	},
	"items/armor2.cel": {
		// The contents of frame 0 at offset 0x44 in armor2.cel starts with the
		// following header: 0A 00 8C 00 52 01 72 01 92 01.
		Header: 10,
		W:      96,  // ref: 0x4219E1
		H:      160, // h = npixels/w = 15360/96 = 160
	},
	"items/axe.cel": {
		// The contents of frame 0 at offset 0x3C in axe.cel starts with the
		// following header: 0A 00 2A 00 4A 00 6B 01 97 01.
		Header: 10,
		W:      96,  // ref: 0x4219E1
		H:      160, // h = npixels/w = 15360/96 = 160
	},
	// NOTE: Unused?
	"items/axeflip.cel": {
		// The contents of frame 0 at offset 0x44 in axeflip.cel starts with the
		// following header: 0A 00 6A 01 75 02 95 02 B5 02.
		Header: 10,
		W:      96,
		H:      160, // h = npixels/w = 15360/96 = 160
	},
	"items/bldstn.cel": {
		// The contents of frame 0 at offset 0x3C in bldstn.cel starts with the
		// following header: 0A 00 2A 00 7E 00 1D 01 3D 01.
		Header: 10,
		W:      96,  // ref: 0x4219E1
		H:      160, // h = npixels/w = 15360/96 = 160
	},
	// NOTE: Unused?
	"items/bottle.cel": {
		// The contents of frame 0 at offset 0x44 in bottle.cel starts with the
		// following header: 0A 00 5C 00 E0 00 00 01 20 01.
		Header: 10,
		W:      96,
		H:      160, // h = npixels/w = 15360/96 = 160
	},
	"items/bow.cel": {
		// The contents of frame 0 at offset 0x3C in bow.cel starts with the
		// following header: 0A 00 2A 00 4A 00 A5 00 C5 00.
		Header: 10,
		W:      96,  // ref: 0x4219E1
		H:      160, // h = npixels/w = 15360/96 = 160
	},
	"items/cleaver.cel": {
		// The contents of frame 0 at offset 0x3C in cleaver.cel starts with the
		// following header: 0A 00 2A 00 57 00 4A 01 6A 01.
		Header: 10,
		W:      96,  // ref: 0x4219E1
		H:      160, // h = npixels/w = 15360/96 = 160
	},
	"items/crownf.cel": {
		// The contents of frame 0 at offset 0x3C in crownf.cel starts with the
		// following header: 0A 00 2A 00 4A 00 11 01 00 00.
		Header: 10,
		W:      96,  // ref: 0x4219E1
		H:      128, // h = npixels/w = 12288/96 = 128
	},
	"items/duricons.cel": {
		W: 32, // ref: 0x4064EB
		H: 32, // h = npixels/w = 1024/32 = 32
	},
	"items/fanvil.cel": {
		// The contents of frame 0 at offset 0x3C in fanvil.cel starts with the
		// following header: 0A 00 2A 00 77 00 1E 03 3E 03.
		Header: 10,
		W:      96,  // ref: 0x4219E1
		H:      160, // h = npixels/w = 15360/96 = 160
	},
	"items/fbook.cel": {
		// The contents of frame 0 at offset 0x3C in fbook.cel starts with the
		// following header: 0A 00 2A 00 56 00 43 01 63 01.
		Header: 10,
		W:      96,  // ref: 0x4219E1
		H:      160, // h = npixels/w = 15360/96 = 160
	},
	// NOTE: Unused?
	"items/fbow.cel": {
		// The contents of frame 0 at offset 0x40 in fbow.cel starts with the
		// following header: 0A 00 2A 00 4A 00 6D 00 8D 00.
		Header: 10,
		W:      96,
		H:      160, // h = npixels/w = 15360/96 = 160
	},
	"items/fbrain.cel": {
		// The contents of frame 0 at offset 0x38 in fbrain.cel starts with the
		// following header: 0A 00 2A 00 4A 00 C2 00 E2 00.
		Header: 10,
		W:      96,  // ref: 0x4219E1
		H:      160, // h = npixels/w = 15360/96 = 160
	},
	"items/fbttle.cel": {
		// The contents of frame 0 at offset 0x48 in fbttle.cel starts with the
		// following header: 0A 00 BA 00 DA 00 FA 00 1A 01.
		Header: 10,
		W:      96,  // ref: 0x4219E1
		H:      160, // h = npixels/w = 15360/96 = 160
	},
	"items/fbttlebb.cel": {
		// The contents of frame 0 at offset 0x48 in fbttlebb.cel starts with the
		// following header: 0A 00 BA 00 DA 00 FA 00 1A 01.
		Header: 10,
		W:      96,  // ref: 0x4219E1
		H:      160, // h = npixels/w = 15360/96 = 160
	},
	"items/fbttlebl.cel": {
		// The contents of frame 0 at offset 0x48 in fbttlebl.cel starts with the
		// following header: 0A 00 BA 00 DA 00 FA 00 1A 01.
		Header: 10,
		W:      96,  // ref: 0x4219E1
		H:      160, // h = npixels/w = 15360/96 = 160
	},
	"items/fbttlebr.cel": {
		// The contents of frame 0 at offset 0x48 in fbttlebr.cel starts with the
		// following header: 0A 00 BA 00 DA 00 FA 00 1A 01.
		Header: 10,
		W:      96,  // ref: 0x4219E1
		H:      160, // h = npixels/w = 15360/96 = 160
	},
	"items/fbttleby.cel": {
		// The contents of frame 0 at offset 0x48 in fbttleby.cel starts with the
		// following header: 0A 00 BA 00 DA 00 FA 00 1A 01.
		Header: 10,
		W:      96,  // ref: 0x4219E1
		H:      160, // h = npixels/w = 15360/96 = 160
	},
	"items/fbttledb.cel": {
		// The contents of frame 0 at offset 0x48 in fbttledb.cel starts with the
		// following header: 0A 00 BA 00 DA 00 FA 00 1A 01.
		Header: 10,
		W:      96,  // ref: 0x4219E1
		H:      160, // h = npixels/w = 15360/96 = 160
	},
	"items/fbttledy.cel": {
		// The contents of frame 0 at offset 0x48 in fbttledy.cel starts with the
		// following header: 0A 00 BA 00 DA 00 FA 00 1A 01.
		Header: 10,
		W:      96,  // ref: 0x4219E1
		H:      160, // h = npixels/w = 15360/96 = 160
	},
	"items/fbttleor.cel": {
		// The contents of frame 0 at offset 0x48 in fbttleor.cel starts with the
		// following header: 0A 00 BA 00 DA 00 FA 00 1A 01.
		Header: 10,
		W:      96,  // ref: 0x4219E1
		H:      160, // h = npixels/w = 15360/96 = 160
	},
	"items/fbttlewh.cel": {
		// The contents of frame 0 at offset 0x48 in fbttlewh.cel starts with the
		// following header: 0A 00 BA 00 DA 00 FA 00 1A 01.
		Header: 10,
		W:      96,  // ref: 0x4219E1
		H:      160, // h = npixels/w = 15360/96 = 160
	},
	"items/fear.cel": {
		// The contents of frame 0 at offset 0x3C in fear.cel starts with the
		// following header: 0A 00 2A 00 5F 00 AA 00 00 00.
		Header: 10,
		W:      96,  // ref: 0x4219E1
		H:      128, // h = npixels/w = 12288/96 = 128
	},
	// NOTE: Unused?
	"items/feye.cel": {
		// The contents of frame 0 at offset 0x38 in feye.cel starts with the
		// following header: 0A 00 2A 00 6E 00 D5 00 F5 00.
		Header: 10,
		W:      96,
		H:      160, // h = npixels/w = 15360/96 = 160
	},
	// NOTE: Unused?
	"items/fheart.cel": {
		// The contents of frame 0 at offset 0x38 in fheart.cel starts with the
		// following header: 0A 00 2A 00 70 00 E2 00 02 01.
		Header: 10,
		W:      96,
		H:      160, // h = npixels/w = 15360/96 = 160
	},
	"items/flazstaf.cel": {
		// The contents of frame 0 at offset 0x28 in flazstaf.cel starts with the
		// following header: 0A 00 2A 00 16 01 31 02 73 02.
		Header: 10,
		W:      96,  // ref: 0x4219E1
		H:      160, // h = npixels/w = 15360/96 = 160
	},
	"items/fmush.cel": {
		// The contents of frame 0 at offset 0x38 in fmush.cel starts with the
		// following header: 0A 00 2A 00 4E 00 CD 00 ED 00.
		Header: 10,
		W:      96,  // ref: 0x4219E1
		H:      160, // h = npixels/w = 15360/96 = 160
	},
	"items/food.cel": {
		// The contents of frame 0 at offset 0xC in food.cel starts with the
		// following header: 0A 00 2A 00 66 01 86 01 00 00.
		Header: 10,
		W:      96,  // ref: 0x4219E1
		H:      128, // h = npixels/w = 12288/96 = 128
	},
	"items/fplatear.cel": {
		// The contents of frame 0 at offset 0x3C in fplatear.cel starts with the
		// following header: 0A 00 2A 00 4A 00 BC 02 DC 02.
		Header: 10,
		W:      96,  // ref: 0x4219E1
		H:      160, // h = npixels/w = 15360/96 = 160
	},
	"items/goldflip.cel": {
		// The contents of frame 0 at offset 0x30 in goldflip.cel starts with the
		// following header: 0A 00 2A 00 F3 00 C4 01 E4 01.
		Header: 10,
		W:      96,  // ref: 0x4219E1
		H:      160, // h = npixels/w = 15360/96 = 160
	},
	"items/helmut.cel": {
		// The contents of frame 0 at offset 0x3C in helmut.cel starts with the
		// following header: 0A 00 2A 00 4A 00 44 01 64 01.
		Header: 10,
		W:      96,  // ref: 0x4219E1
		H:      160, // h = npixels/w = 15360/96 = 160
	},
	"items/innsign.cel": {
		// The contents of frame 0 at offset 0x3C in innsign.cel starts with the
		// following header: 0A 00 42 00 09 03 4E 04 6E 04.
		Header: 10,
		W:      96,  // ref: 0x4219E1
		H:      160, // h = npixels/w = 15360/96 = 160
	},
	"items/larmor.cel": {
		// The contents of frame 0 at offset 0x3C in larmor.cel starts with the
		// following header: 0A 00 2A 00 4A 00 9B 02 00 00.
		Header: 10,
		W:      96,  // ref: 0x4219E1
		H:      128, // h = npixels/w = 12288/96 = 128
	},
	"items/mace.cel": {
		// The contents of frame 0 at offset 0x3C in mace.cel starts with the
		// following header: 0A 00 2A 00 69 00 55 01 75 01.
		Header: 10,
		W:      96,  // ref: 0x4219E1
		H:      160, // h = npixels/w = 15360/96 = 160
	},
	// NOTE: Unused?
	"items/manaflip.cel": {
		// The contents of frame 0 at offset 0x44 in manaflip.cel starts with the
		// following header: 0A 00 3B 00 5B 00 7B 00 9B 00.
		Header: 10,
		W:      96,
		H:      160, // h = npixels/w = 15360/96 = 160
	},
	"items/map/mapz0000.cel": {
		W: 640, // ref: 0x40ADBB
		H: 352, // h = npixels/w = 225280/640 = 352
	},
	"items/map/mapz0001.cel": {
		W: 640, // ref: 0x40ADBB
		H: 352, // h = npixels/w = 225280/640 = 352
	},
	"items/map/mapz0002.cel": {
		W: 640, // ref: 0x40ADBB
		H: 352, // h = npixels/w = 225280/640 = 352
	},
	"items/map/mapz0003.cel": {
		W: 640, // ref: 0x40ADBB
		H: 352, // h = npixels/w = 225280/640 = 352
	},
	"items/map/mapz0004.cel": {
		W: 640, // ref: 0x40ADBB
		H: 352, // h = npixels/w = 225280/640 = 352
	},
	"items/map/mapz0005.cel": {
		W: 640, // ref: 0x40ADBB
		H: 352, // h = npixels/w = 225280/640 = 352
	},
	"items/map/mapz0006.cel": {
		W: 640, // ref: 0x40ADBB
		H: 352, // h = npixels/w = 225280/640 = 352
	},
	"items/map/mapz0007.cel": {
		W: 640, // ref: 0x40ADBB
		H: 352, // h = npixels/w = 225280/640 = 352
	},
	"items/map/mapz0008.cel": {
		W: 640, // ref: 0x40ADBB
		H: 352, // h = npixels/w = 225280/640 = 352
	},
	"items/map/mapz0009.cel": {
		W: 640, // ref: 0x40ADBB
		H: 352, // h = npixels/w = 225280/640 = 352
	},
	"items/map/mapz0010.cel": {
		W: 640, // ref: 0x40ADBB
		H: 352, // h = npixels/w = 225280/640 = 352
	},
	"items/map/mapz0011.cel": {
		W: 640, // ref: 0x40ADBB
		H: 352, // h = npixels/w = 225280/640 = 352
	},
	"items/map/mapz0012.cel": {
		W: 640, // ref: 0x40ADBB
		H: 352, // h = npixels/w = 225280/640 = 352
	},
	"items/map/mapz0013.cel": {
		W: 640, // ref: 0x40ADBB
		H: 352, // h = npixels/w = 225280/640 = 352
	},
	"items/map/mapz0014.cel": {
		W: 640, // ref: 0x40ADBB
		H: 352, // h = npixels/w = 225280/640 = 352
	},
	"items/map/mapz0015.cel": {
		W: 640, // ref: 0x40ADBB
		H: 352, // h = npixels/w = 225280/640 = 352
	},
	"items/map/mapz0016.cel": {
		W: 640, // ref: 0x40ADBB
		H: 352, // h = npixels/w = 225280/640 = 352
	},
	"items/map/mapz0017.cel": {
		W: 640, // ref: 0x40ADBB
		H: 352, // h = npixels/w = 225280/640 = 352
	},
	"items/map/mapz0018.cel": {
		W: 640, // ref: 0x40ADBB
		H: 352, // h = npixels/w = 225280/640 = 352
	},
	"items/map/mapz0019.cel": {
		W: 640, // ref: 0x40ADBB
		H: 352, // h = npixels/w = 225280/640 = 352
	},
	"items/map/mapz0020.cel": {
		W: 640, // ref: 0x40ADBB
		H: 352, // h = npixels/w = 225280/640 = 352
	},
	"items/map/mapz0021.cel": {
		W: 640, // ref: 0x40ADBB
		H: 352, // h = npixels/w = 225280/640 = 352
	},
	"items/map/mapz0022.cel": {
		W: 640, // ref: 0x40ADBB
		H: 352, // h = npixels/w = 225280/640 = 352
	},
	"items/map/mapz0023.cel": {
		W: 640, // ref: 0x40ADBB
		H: 352, // h = npixels/w = 225280/640 = 352
	},
	"items/map/mapz0024.cel": {
		W: 640, // ref: 0x40ADBB
		H: 352, // h = npixels/w = 225280/640 = 352
	},
	"items/map/mapz0025.cel": {
		W: 640, // ref: 0x40ADBB
		H: 352, // h = npixels/w = 225280/640 = 352
	},
	"items/map/mapz0026.cel": {
		W: 640, // ref: 0x40ADBB
		H: 352, // h = npixels/w = 225280/640 = 352
	},
	"items/map/mapz0027.cel": {
		W: 640, // ref: 0x40ADBB
		H: 352, // h = npixels/w = 225280/640 = 352
	},
	"items/map/mapz0028.cel": {
		W: 640, // ref: 0x40ADBB
		H: 352, // h = npixels/w = 225280/640 = 352
	},
	"items/map/mapz0029.cel": {
		W: 640, // ref: 0x40ADBB
		H: 352, // h = npixels/w = 225280/640 = 352
	},
	"items/map/mapz0030.cel": {
		W: 640, // ref: 0x40ADBB
		H: 352, // h = npixels/w = 225280/640 = 352
	},
	"items/map/mapzdoom.cel": {
		W: 640, // ref: 0x40ADBB
		H: 352, // h = npixels/w = 225280/640 = 352
	},
	"items/ring.cel": {
		// The contents of frame 0 at offset 0x3C in ring.cel starts with the
		// following header: 0A 00 2A 00 4A 00 A5 00 C5 00.
		Header: 10,
		W:      96,  // ref: 0x4219E1
		H:      160, // h = npixels/w = 15360/96 = 160
	},
	"items/rock.cel": {
		// The contents of frame 0 at offset 0x58 in rock.cel starts with the
		// following header: 0A 00 54 01 74 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4219E1
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"items/scroll.cel": {
		// The contents of frame 0 at offset 0x3C in scroll.cel starts with the
		// following header: 0A 00 2A 00 4A 00 F9 00 19 01.
		Header: 10,
		W:      96,  // ref: 0x4219E1
		H:      160, // h = npixels/w = 15360/96 = 160
	},
	"items/shield.cel": {
		// The contents of frame 0 at offset 0x3C in shield.cel starts with the
		// following header: 0A 00 2A 00 4A 00 70 00 90 00.
		Header: 10,
		W:      96,  // ref: 0x4219E1
		H:      160, // h = npixels/w = 15360/96 = 160
	},
	"items/staff.cel": {
		// The contents of frame 0 at offset 0x3C in staff.cel starts with the
		// following header: 0A 00 2A 00 66 00 0C 01 5E 01.
		Header: 10,
		W:      96,  // ref: 0x4219E1
		H:      160, // h = npixels/w = 15360/96 = 160
	},
	"items/swrdflip.cel": {
		// The contents of frame 0 at offset 0x3C in swrdflip.cel starts with the
		// following header: 0A 00 2A 00 4A 00 0C 01 2C 01.
		Header: 10,
		W:      96,  // ref: 0x4219E1
		H:      160, // h = npixels/w = 15360/96 = 160
	},
	// NOTE: Unused?
	"items/wand.cel": {
		// The contents of frame 0 at offset 0x3C in wand.cel starts with the
		// following header: 0A 00 2A 00 4A 00 AA 00 CA 00.
		Header: 10,
		W:      96,
		H:      160, // h = npixels/w = 15360/96 = 160
	},
	"items/wshield.cel": {
		// The contents of frame 0 at offset 0x3C in wshield.cel starts with the
		// following header: 0A 00 2A 00 4A 00 BD 02 00 00.
		Header: 10,
		W:      96,  // ref: 0x4219E1
		H:      128, // h = npixels/w = 12288/96 = 128
	},
	"levels/l1data/l1.cel": {
		W: 32, // ref: 0x418F12
		H: 32, // h = npixels/w = 1024/32 = 32
	},
	"levels/l1data/l1s.cel": {
		// The contents of frame 0 at offset 0x28 in l1s.cel starts with the
		// following header: 0A 00 D5 00 EC 02 1A 07 8C 09.
		Header: 10,
		W:      64,  // ref: 0x455835
		H:      160, // h = npixels/w = 10240/64 = 160
	},
	"levels/l2data/l2.cel": {
		W: 32, // ref: 0x418F12
		H: 32, // h = npixels/w = 1024/32 = 32
	},
	"levels/l2data/l2s.cel": {
		// The contents of frame 0 at offset 0x20 in l2s.cel starts with the
		// following header: 0A 00 2A 00 57 00 40 01 7D 02.
		Header: 10,
		W:      64,  // ref: 0x455835
		H:      160, // h = npixels/w = 10240/64 = 160
	},
	"levels/l3data/l3.cel": {
		W: 32, // ref: 0x418F12
		H: 32, // h = npixels/w = 1024/32 = 32
	},
	"levels/l4data/l4.cel": {
		W: 32, // ref: 0x418F12
		H: 32, // h = npixels/w = 1024/32 = 32
	},
	"levels/towndata/town.cel": {
		W: 32, // ref: 0x4632D5
		H: 32, // h = npixels/w = 1024/32 = 32
	},
	"levels/towndata/towns.cel": {
		W: 64,  // ref: 0x455835
		H: 224, // h = npixels/w = 14336/64 = 224
	},
	// NOTE: Unused?
	"missiles/flamel1.cel": {
		// The contents of frame 0 at offset 0x24 in flamel1.cel starts with the
		// following header: 0A 00 4D 01 29 03 00 00 00 00.
		Header: 10,
		W:      128,
		H:      96, // h = npixels/w = 12288/128 = 96
	},
	// NOTE: Unused?
	"missiles/flamel2.cel": {
		// The contents of frame 0 at offset 0x24 in flamel2.cel starts with the
		// following header: 0A 00 23 01 35 03 00 00 00 00.
		Header: 10,
		W:      96,
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	// NOTE: Unused?
	"missiles/flamel3.cel": {
		// The contents of frame 0 at offset 0x24 in flamel3.cel starts with the
		// following header: 0A 00 C4 00 50 04 00 00 00 00.
		Header: 10,
		W:      128,
		H:      96, // h = npixels/w = 12288/128 = 96
	},
	// NOTE: Unused?
	"missiles/flamel4.cel": {
		// The contents of frame 0 at offset 0x24 in flamel4.cel starts with the
		// following header: 0A 00 39 00 08 03 00 00 00 00.
		Header: 10,
		W:      96,
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	// NOTE: Unused?
	"missiles/flamel5.cel": {
		// The contents of frame 0 at offset 0x24 in flamel5.cel starts with the
		// following header: 0A 00 45 00 82 05 00 00 00 00.
		Header: 10,
		W:      128,
		H:      96, // h = npixels/w = 12288/128 = 96
	},
	// NOTE: Unused?
	"missiles/flamel6.cel": {
		// The contents of frame 0 at offset 0x24 in flamel6.cel starts with the
		// following header: 0A 00 32 00 01 03 00 00 00 00.
		Header: 10,
		W:      96,
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	// NOTE: Unused?
	"missiles/flamel7.cel": {
		// The contents of frame 0 at offset 0x24 in flamel7.cel starts with the
		// following header: 0A 00 C2 00 25 04 00 00 00 00.
		Header: 10,
		W:      128,
		H:      96, // h = npixels/w = 12288/128 = 96
	},
	// NOTE: Unused?
	"missiles/flamel8.cel": {
		// The contents of frame 0 at offset 0x24 in flamel8.cel starts with the
		// following header: 0A 00 ED 00 23 03 00 00 00 00.
		Header: 10,
		W:      96,
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	// NOTE: Unused?
	"missiles/flamel9.cel": {
		// The contents of frame 0 at offset 0x24 in flamel9.cel starts with the
		// following header: 0A 00 FF 00 1F 03 00 00 00 00.
		Header: 10,
		W:      128,
		H:      96, // h = npixels/w = 12288/128 = 96
	},
	// NOTE: Unused?
	"missiles/flamel10.cel": {
		// The contents of frame 0 at offset 0x24 in flamel10.cel starts with the
		// following header: 0A 00 FA 00 22 03 00 00 00 00.
		Header: 10,
		W:      96,
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	// NOTE: Unused?
	"missiles/flamel11.cel": {
		// The contents of frame 0 at offset 0x24 in flamel11.cel starts with the
		// following header: 0A 00 92 01 24 04 00 00 00 00.
		Header: 10,
		W:      128,
		H:      96, // h = npixels/w = 12288/128 = 96
	},
	// NOTE: Unused?
	"missiles/flamel12.cel": {
		// The contents of frame 0 at offset 0x24 in flamel12.cel starts with the
		// following header: 0A 00 86 00 1C 03 00 00 00 00.
		Header: 10,
		W:      96,
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	// NOTE: Unused?
	"missiles/flamel13.cel": {
		// The contents of frame 0 at offset 0x24 in flamel13.cel starts with the
		// following header: 0A 00 F0 01 83 05 00 00 00 00.
		Header: 10,
		W:      128,
		H:      96, // h = npixels/w = 12288/128 = 96
	},
	// NOTE: Unused?
	"missiles/flamel14.cel": {
		// The contents of frame 0 at offset 0x24 in flamel14.cel starts with the
		// following header: 0A 00 E0 00 29 03 00 00 00 00.
		Header: 10,
		W:      96,
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	// NOTE: Unused?
	"missiles/flamel15.cel": {
		// The contents of frame 0 at offset 0x24 in flamel15.cel starts with the
		// following header: 0A 00 E9 01 3B 04 00 00 00 00.
		Header: 10,
		W:      128,
		H:      96, // h = npixels/w = 12288/128 = 96
	},
	// NOTE: Unused?
	"missiles/flamel16.cel": {
		// The contents of frame 0 at offset 0x24 in flamel16.cel starts with the
		// following header: 0A 00 62 01 4B 03 00 00 00 00.
		Header: 10,
		W:      96,
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	// NOTE: Unused?
	"missiles/flames1.cel": {
		// The contents of frame 0 at offset 0x18 in flames1.cel starts with the
		// following header: 0A 00 2A 00 7D 00 00 00 00 00.
		Header: 10,
		W:      128,
		H:      96, // h = npixels/w = 12288/128 = 96
	},
	// NOTE: Unused?
	"missiles/flames2.cel": {
		// The contents of frame 0 at offset 0x18 in flames2.cel starts with the
		// following header: 0A 00 2A 00 7A 00 00 00 00 00.
		Header: 10,
		W:      96,
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	// NOTE: Unused?
	"missiles/flames3.cel": {
		// The contents of frame 0 at offset 0x18 in flames3.cel starts with the
		// following header: 0A 00 2A 00 A9 00 00 00 00 00.
		Header: 10,
		W:      128,
		H:      96, // h = npixels/w = 12288/128 = 96
	},
	// NOTE: Unused?
	"missiles/flames4.cel": {
		// The contents of frame 0 at offset 0x18 in flames4.cel starts with the
		// following header: 0A 00 2A 00 63 00 00 00 00 00.
		Header: 10,
		W:      96,
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	// NOTE: Unused?
	"missiles/flames5.cel": {
		// The contents of frame 0 at offset 0x18 in flames5.cel starts with the
		// following header: 0A 00 2A 00 AD 00 00 00 00 00.
		Header: 10,
		W:      128,
		H:      96, // h = npixels/w = 12288/128 = 96
	},
	// NOTE: Unused?
	"missiles/flames6.cel": {
		// The contents of frame 0 at offset 0x18 in flames6.cel starts with the
		// following header: 0A 00 2A 00 69 00 00 00 00 00.
		Header: 10,
		W:      96,
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	// NOTE: Unused?
	"missiles/flames7.cel": {
		// The contents of frame 0 at offset 0x18 in flames7.cel starts with the
		// following header: 0A 00 31 00 90 00 00 00 00 00.
		Header: 10,
		W:      128,
		H:      96, // h = npixels/w = 12288/128 = 96
	},
	// NOTE: Unused?
	"missiles/flames8.cel": {
		// The contents of frame 0 at offset 0x18 in flames8.cel starts with the
		// following header: 0A 00 4C 00 86 00 00 00 00 00.
		Header: 10,
		W:      96,
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	// NOTE: Unused?
	"missiles/flames9.cel": {
		// The contents of frame 0 at offset 0x18 in flames9.cel starts with the
		// following header: 0A 00 66 00 98 00 00 00 00 00.
		Header: 10,
		W:      128,
		H:      96, // h = npixels/w = 12288/128 = 96
	},
	// NOTE: Unused?
	"missiles/flames10.cel": {
		// The contents of frame 0 at offset 0x18 in flames10.cel starts with the
		// following header: 0A 00 5E 00 89 00 00 00 00 00.
		Header: 10,
		W:      96,
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	// NOTE: Unused?
	"missiles/flames11.cel": {
		// The contents of frame 0 at offset 0x18 in flames11.cel starts with the
		// following header: 0A 00 7D 00 A0 00 00 00 00 00.
		Header: 10,
		W:      128,
		H:      96, // h = npixels/w = 12288/128 = 96
	},
	// NOTE: Unused?
	"missiles/flames12.cel": {
		// The contents of frame 0 at offset 0x18 in flames12.cel starts with the
		// following header: 0A 00 2D 00 66 00 00 00 00 00.
		Header: 10,
		W:      96,
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	// NOTE: Unused?
	"missiles/flames13.cel": {
		// The contents of frame 0 at offset 0x18 in flames13.cel starts with the
		// following header: 0A 00 54 00 C8 00 00 00 00 00.
		Header: 10,
		W:      128,
		H:      96, // h = npixels/w = 12288/128 = 96
	},
	// NOTE: Unused?
	"missiles/flames14.cel": {
		// The contents of frame 0 at offset 0x18 in flames14.cel starts with the
		// following header: 0A 00 2A 00 6B 00 00 00 00 00.
		Header: 10,
		W:      96,
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	// NOTE: Unused?
	"missiles/flames15.cel": {
		// The contents of frame 0 at offset 0x18 in flames15.cel starts with the
		// following header: 0A 00 2A 00 94 00 00 00 00 00.
		Header: 10,
		W:      128,
		H:      96, // h = npixels/w = 12288/128 = 96
	},
	// NOTE: Unused?
	"missiles/flames16.cel": {
		// The contents of frame 0 at offset 0x18 in flames16.cel starts with the
		// following header: 0A 00 2A 00 80 00 00 00 00 00.
		Header: 10,
		W:      96,
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	// NOTE: Unused?
	"missiles/flaml1.cel": {
		// The contents of frame 0 at offset 0x24 in flaml1.cel starts with the
		// following header: 0A 00 32 00 21 03 6E 03 00 00.
		Header: 10,
		W:      128,
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	// NOTE: Unused?
	"missiles/flaml2.cel": {
		// The contents of frame 0 at offset 0x24 in flaml2.cel starts with the
		// following header: 0A 00 2A 00 9F 04 33 05 00 00.
		Header: 10,
		W:      128,
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	// NOTE: Unused?
	"missiles/flaml3.cel": {
		// The contents of frame 0 at offset 0x24 in flaml3.cel starts with the
		// following header: 0A 00 2A 00 20 06 89 06 00 00.
		Header: 10,
		W:      128,
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	// NOTE: Unused?
	"missiles/flaml4.cel": {
		// The contents of frame 0 at offset 0x24 in flaml4.cel starts with the
		// following header: 0A 00 2A 00 35 04 2B 05 00 00.
		Header: 10,
		W:      128,
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	// NOTE: Unused?
	"missiles/flaml5.cel": {
		// The contents of frame 0 at offset 0x24 in flaml5.cel starts with the
		// following header: 0A 00 2A 00 07 03 72 03 00 00.
		Header: 10,
		W:      128,
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	// NOTE: Unused?
	"missiles/flaml6.cel": {
		// The contents of frame 0 at offset 0x24 in flaml6.cel starts with the
		// following header: 0A 00 3E 00 B1 04 E5 04 00 00.
		Header: 10,
		W:      128,
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	// NOTE: Unused?
	"missiles/flaml7.cel": {
		// The contents of frame 0 at offset 0x24 in flaml7.cel starts with the
		// following header: 0A 00 2A 00 33 06 53 06 00 00.
		Header: 10,
		W:      128,
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	// NOTE: Unused?
	"missiles/flaml8.cel": {
		// The contents of frame 0 at offset 0x24 in flaml8.cel starts with the
		// following header: 0A 00 B7 00 3C 05 67 05 00 00.
		Header: 10,
		W:      128,
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	// NOTE: Unused?
	"missiles/flams1.cel": {
		// The contents of frame 0 at offset 0x3C in flams1.cel starts with the
		// following header: 0A 00 2A 00 69 00 8C 00 00 00.
		Header: 10,
		W:      128,
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	// NOTE: Unused?
	"missiles/flams2.cel": {
		// The contents of frame 0 at offset 0x3C in flams2.cel starts with the
		// following header: 0A 00 2A 00 69 00 99 00 00 00.
		Header: 10,
		W:      128,
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	// NOTE: Unused?
	"missiles/flams3.cel": {
		// The contents of frame 0 at offset 0x3C in flams3.cel starts with the
		// following header: 0A 00 2A 00 88 00 A8 00 00 00.
		Header: 10,
		W:      128,
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	// NOTE: Unused?
	"missiles/flams4.cel": {
		// The contents of frame 0 at offset 0x3C in flams4.cel starts with the
		// following header: 0A 00 2A 00 77 00 97 00 00 00.
		Header: 10,
		W:      128,
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	// NOTE: Unused?
	"missiles/flams5.cel": {
		// The contents of frame 0 at offset 0x3C in flams5.cel starts with the
		// following header: 0A 00 2A 00 7E 00 9E 00 00 00.
		Header: 10,
		W:      128,
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	// NOTE: Unused?
	"missiles/flams6.cel": {
		// The contents of frame 0 at offset 0x3C in flams6.cel starts with the
		// following header: 0A 00 2A 00 7C 00 9C 00 00 00.
		Header: 10,
		W:      128,
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	// NOTE: Unused?
	"missiles/flams7.cel": {
		// The contents of frame 0 at offset 0x3C in flams7.cel starts with the
		// following header: 0A 00 2A 00 92 00 B2 00 00 00.
		Header: 10,
		W:      128,
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	// NOTE: Unused?
	"missiles/flams8.cel": {
		// The contents of frame 0 at offset 0x3C in flams8.cel starts with the
		// following header: 0A 00 2A 00 73 00 93 00 00 00.
		Header: 10,
		W:      128,
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	// NOTE: Unused?
	"missiles/mindmace.cel": {
		// The contents of frame 0 at offset 0x28 in mindmace.cel starts with the
		// following header: 0A 00 2A 00 A6 01 00 00 00 00.
		Header: 10,
		W:      96,
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	// NOTE: Unused?
	"missiles/sentfr.cel": {
		// The contents of frame 0 at offset 0x14 in sentfr.cel starts with the
		// following header: 0A 00 E7 02 34 0C 00 00 00 00.
		Header: 10,
		W:      96,
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	// NOTE: Unused?
	"missiles/sentout.cel": {
		// The contents of frame 0 at offset 0x40 in sentout.cel starts with the
		// following header: 0A 00 4A 00 B2 05 00 00 00 00.
		Header: 10,
		W:      96,
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	// NOTE: Unused?
	"missiles/sentup.cel": {
		// The contents of frame 0 at offset 0x44 in sentup.cel starts with the
		// following header: 0A 00 3F 00 5F 00 00 00 00 00.
		Header: 10,
		W:      96,
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	// NOTE: Unused?
	"monsters/acid/acidpud.cel": {
		// The contents of frame 0 at offset 0x28 in acidpud.cel starts with the
		// following header: 0A 00 43 06 EC 06 00 00 00 00.
		Header: 10,
		W:      128,
		H:      96, // h = npixels/w = 12288/128 = 96
	},
	// NOTE: Unused?
	"monsters/magma/magball1.cel": {
		// The contents of frame 0 at offset 0x48 in magball1.cel starts with the
		// following header: 0A 00 2A 00 4A 00 53 01 00 00.
		Header: 10,
		W:      128,
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	// NOTE: Unused?
	"monsters/magma/magball2.cel": {
		// The contents of frame 0 at offset 0x48 in magball2.cel starts with the
		// following header: 0A 00 2A 00 4A 00 82 01 00 00.
		Header: 10,
		W:      128,
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	// NOTE: Unused?
	"monsters/magma/magball3.cel": {
		// The contents of frame 0 at offset 0x48 in magball3.cel starts with the
		// following header: 0A 00 2A 00 4A 00 88 01 00 00.
		Header: 10,
		W:      128,
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	// NOTE: Unused?
	"monsters/magma/magball4.cel": {
		// The contents of frame 0 at offset 0x48 in magball4.cel starts with the
		// following header: 0A 00 2A 00 4A 00 70 01 00 00.
		Header: 10,
		W:      128,
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	// NOTE: Unused?
	"monsters/magma/magball5.cel": {
		// The contents of frame 0 at offset 0x48 in magball5.cel starts with the
		// following header: 0A 00 2A 00 4A 00 6E 01 00 00.
		Header: 10,
		W:      128,
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	// NOTE: Unused?
	"monsters/magma/magball6.cel": {
		// The contents of frame 0 at offset 0x48 in magball6.cel starts with the
		// following header: 0A 00 2A 00 4A 00 7D 01 00 00.
		Header: 10,
		W:      128,
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	// NOTE: Unused?
	"monsters/magma/magball7.cel": {
		// The contents of frame 0 at offset 0x48 in magball7.cel starts with the
		// following header: 0A 00 2A 00 4A 00 99 01 00 00.
		Header: 10,
		W:      128,
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	// NOTE: Unused?
	"monsters/magma/magball8.cel": {
		// The contents of frame 0 at offset 0x48 in magball8.cel starts with the
		// following header: 0A 00 2A 00 4A 00 81 01 00 00.
		Header: 10,
		W:      128,
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	// NOTE: Unused?
	"monsters/magma/magblos.cel": {
		// The contents of frame 0 at offset 0x30 in magblos.cel starts with the
		// following header: 0A 00 58 00 3B 01 5B 01 00 00.
		Header: 10,
		W:      128,
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	// NOTE: Unused?
	"monsters/rhino/rhinos1.cel": {
		// The contents of frame 0 at offset 0x20 in rhinos1.cel starts with the
		// following header: 0A 00 D0 02 77 09 87 0D 00 00.
		Header: 10,
		W:      160,
		H:      128, // h = npixels/w = 20480/160 = 128
	},
	// NOTE: Unused?
	"monsters/rhino/rhinos2.cel": {
		// The contents of frame 0 at offset 0x20 in rhinos2.cel starts with the
		// following header: 0A 00 9C 03 FB 09 25 0D 00 00.
		Header: 10,
		W:      160,
		H:      128, // h = npixels/w = 20480/160 = 128
	},
	// NOTE: Unused?
	"monsters/rhino/rhinos3.cel": {
		// The contents of frame 0 at offset 0x20 in rhinos3.cel starts with the
		// following header: 0A 00 29 04 0F 08 B6 0B 00 00.
		Header: 10,
		W:      160,
		H:      128, // h = npixels/w = 20480/160 = 128
	},
	// NOTE: Unused?
	"monsters/rhino/rhinos4.cel": {
		// The contents of frame 0 at offset 0x20 in rhinos4.cel starts with the
		// following header: 0A 00 A6 04 A3 0A EF 0E 00 00.
		Header: 10,
		W:      160,
		H:      128, // h = npixels/w = 20480/160 = 128
	},
	// NOTE: Unused?
	"monsters/rhino/rhinos5.cel": {
		// The contents of frame 0 at offset 0x20 in rhinos5.cel starts with the
		// following header: 0A 00 F7 03 AA 0A D5 0E 00 00.
		Header: 10,
		W:      160,
		H:      128, // h = npixels/w = 20480/160 = 128
	},
	// NOTE: Unused?
	"monsters/rhino/rhinos6.cel": {
		// The contents of frame 0 at offset 0x20 in rhinos6.cel starts with the
		// following header: 0A 00 94 04 69 0A A1 0E 00 00.
		Header: 10,
		W:      160,
		H:      128, // h = npixels/w = 20480/160 = 128
	},
	// NOTE: Unused?
	"monsters/rhino/rhinos7.cel": {
		// The contents of frame 0 at offset 0x20 in rhinos7.cel starts with the
		// following header: 0A 00 81 03 05 07 9D 0A 00 00.
		Header: 10,
		W:      160,
		H:      128, // h = npixels/w = 20480/160 = 128
	},
	// NOTE: Unused?
	"monsters/rhino/rhinos8.cel": {
		// The contents of frame 0 at offset 0x20 in rhinos8.cel starts with the
		// following header: 0A 00 A0 03 9D 09 DB 0C 00 00.
		Header: 10,
		W:      160,
		H:      128, // h = npixels/w = 20480/160 = 128
	},
	// NOTE: Unused?
	"monsters/succ/flare.cel": {
		// The contents of frame 0 at offset 0x48 in flare.cel starts with the
		// following header: 0A 00 55 00 21 03 80 03 00 00.
		Header: 10,
		W:      128,
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	// NOTE: Unused?
	"monsters/succ/flarexp.cel": {
		// The contents of frame 0 at offset 0x24 in flarexp.cel starts with the
		// following header: 0A 00 2A 00 4D 00 6D 00 00 00.
		Header: 10,
		W:      128,
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	"monsters/thin/lghning.cel": {
		// The contents of frame 0 at offset 0x28 in lghning.cel starts with the
		// following header: 0A 00 72 00 F1 03 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4219E1
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	// NOTE: Unused?
	"monsters/unrav/unravw.cel": {
		// The contents of the archive header at offset 0 in unravw.cel (after
		// fix, see https://github.com/mewrnd/blizzconv/issues/2):
		//    20 00 00 00  07 27 00 00  C5 49 00 00  A3 6B 00 00
		//    C3 86 00 00  26 A9 00 00  4C D0 00 00  93 ED 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x38 in unravw.cel starts with the
		// following header: 0A 00 B2 01 8C 03 00 00 00 00.
		Header: 10,
		W:      96,
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"objects/altboy.cel": {
		// The contents of frame 0 at offset 0x0C in altboy.cel starts with the
		// following header: 0A 00 2A 06 35 0D 64 0D 00 00.
		Header: 10,
		W:      128, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	"objects/angel.cel": {
		// The contents of frame 0 at offset 0x10 in angel.cel starts with the
		// following header: 0A 00 4E 05 40 09 E8 0E 00 00.
		Header: 10,
		W:      96,  // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      128, // h = npixels/w = 12288/96 = 128
	},
	"objects/armstand.cel": {
		// The contents of frame 0 at offset 0x10 in armstand.cel starts with the
		// following header: 0A 00 A4 01 CE 03 00 00 00 00.
		Header: 10,
		W:      96, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"objects/banner.cel": {
		// The contents of frame 0 at offset 0x14 in banner.cel starts with the
		// following header: 0A 00 C1 02 6F 07 00 00 00 00.
		Header: 10,
		W:      96, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"objects/barrel.cel": {
		// The contents of frame 0 at offset 0x2C in barrel.cel starts with the
		// following header: 0A 00 F2 02 FE 04 00 00 00 00.
		Header: 10,
		W:      96, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"objects/barrelex.cel": {
		// The contents of frame 0 at offset 0x30 in barrelex.cel starts with the
		// following header: 0A 00 F2 02 FE 04 00 00 00 00.
		Header: 10,
		W:      96, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"objects/bcase.cel": {
		// The contents of frame 0 at offset 0x28 in bcase.cel starts with the
		// following header: 0A 00 7E 02 4F 08 8C 0A 00 00.
		Header: 10,
		W:      96,  // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      128, // h = npixels/w = 12288/96 = 128
	},
	"objects/bkslbrnt.cel": {
		// The contents of frame 0 at offset 0x20 in bkslbrnt.cel starts with the
		// following header: 0A 00 99 02 AA 04 00 00 00 00.
		Header: 10,
		W:      96, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	// NOTE: Unused?
	"objects/bkurns.cel": {
		// The contents of frame 0 at offset 0x30 in bkurns.cel starts with the
		// following header: 0A 00 C5 01 98 02 00 00 00 00.
		Header: 10,
		W:      96,
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"objects/bloodfnt.cel": {
		// The contents of frame 0 at offset 0x30 in bloodfnt.cel starts with the
		// following header: 0A 00 BF 05 24 0A 00 00 00 00.
		Header: 10,
		W:      96, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"objects/book1.cel": {
		// The contents of frame 0 at offset 0x20 in book1.cel starts with the
		// following header: 0A 00 77 02 91 04 00 00 00 00.
		Header: 10,
		W:      96, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"objects/book2.cel": {
		// The contents of frame 0 at offset 0x20 in book2.cel starts with the
		// following header: 0A 00 B8 01 2D 03 00 00 00 00.
		Header: 10,
		W:      96, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	// NOTE: Unused?
	"objects/bshelf.cel": {
		// The contents of frame 0 at offset 0x18 in bshelf.cel starts with the
		// following header: 0A 00 DA 02 02 06 22 06 00 00.
		Header: 10,
		W:      96,
		H:      128, // h = npixels/w = 12288/96 = 96
	},
	"objects/burncros.cel": {
		// The contents of frame 0 at offset 0x30 in burncros.cel starts with the
		// following header: 0A 00 FA 01 46 0C A3 12 B3 15.
		Header: 10,
		W:      160, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      160, // h = npixels/w = 25600/160 = 160
	},
	// NOTE: Unused?
	"objects/candlabr.cel": {
		// The contents of frame 0 at offset 0x0C in candlabr.cel starts with the
		// following header: 0A 00 6B 01 72 02 00 00 00 00.
		Header: 10,
		W:      96,
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	// NOTE: Unused?
	"objects/candle.cel": {
		// The contents of frame 0 at offset 0x2C in candle.cel starts with the
		// following header: 0A 00 57 01 2F 02 00 00 00 00.
		Header: 10,
		W:      96,
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"objects/candle2.cel": {
		// The contents of frame 0 at offset 0x18 in candle2.cel starts with the
		// following header: 0A 00 58 01 32 02 00 00 00 00.
		Header: 10,
		W:      96, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"objects/cauldren.cel": {
		// The contents of frame 0 at offset 0x14 in cauldren.cel starts with the
		// following header: 0A 00 50 05 58 09 00 00 00 00.
		Header: 10,
		W:      96, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"objects/chest1.cel": {
		// The contents of frame 0 at offset 0x20 in chest1.cel starts with the
		// following header: 0A 00 6C 02 8C 02 00 00 00 00.
		Header: 10,
		W:      96, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"objects/chest2.cel": {
		// The contents of frame 0 at offset 0x20 in chest2.cel starts with the
		// following header: 0A 00 50 03 B7 03 00 00 00 00.
		Header: 10,
		W:      96, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"objects/chest3.cel": {
		// The contents of frame 0 at offset 0x20 in chest3.cel starts with the
		// following header: 0A 00 9F 04 76 05 00 00 00 00.
		Header: 10,
		W:      96, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"objects/cruxsk1.cel": {
		// The contents of frame 0 at offset 0x44 in cruxsk1.cel starts with the
		// following header: 0A 00 EC 00 8F 03 89 06 00 00.
		Header: 10,
		W:      96,  // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      128, // h = npixels/w = 12288/96 = 128
	},
	"objects/cruxsk2.cel": {
		// The contents of frame 0 at offset 0x44 in cruxsk2.cel starts with the
		// following header: 0A 00 09 01 93 03 8C 06 00 00.
		Header: 10,
		W:      96,  // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      128, // h = npixels/w = 12288/96 = 128
	},
	"objects/cruxsk3.cel": {
		// The contents of frame 0 at offset 0x44 in cruxsk3.cel starts with the
		// following header: 0A 00 01 01 5D 03 5C 06 00 00.
		Header: 10,
		W:      96,  // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      128, // h = npixels/w = 12288/96 = 128
	},
	"objects/decap.cel": {
		// The contents of frame 0 at offset 0x28 in decap.cel starts with the
		// following header: 0A 00 D8 01 0C 03 00 00 00 00.
		Header: 10,
		W:      96, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	// NOTE: Unused?
	"objects/dirtfall.cel": {
		// The contents of frame 0 at offset 0x30 in dirtfall.cel starts with the
		// following header: 0A 00 2A 00 4A 00 95 00 C8 00.
		Header: 10,
		W:      96,
		H:      160, // h = npixels/w = 15360/96 = 160
	},
	// NOTE: Unused?
	"objects/explod1.cel": {
		// The contents of frame 0 at offset 0x34 in explod1.cel starts with the
		// following header: 0A 00 2A 00 EB 00 BF 01 00 00.
		Header: 10,
		W:      128,
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	// NOTE: Unused?
	"objects/explod2.cel": {
		// The contents of frame 0 at offset 0x34 in explod2.cel starts with the
		// following header: 0A 00 2A 00 BB 00 5F 01 00 00.
		Header: 10,
		W:      128,
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	// NOTE: Unused?
	"objects/firewal1.cel": {
		// The contents of frame 0 at offset 0x3C in firewal1.cel starts with the
		// following header: 0A 00 35 07 4B 12 4A 17 00 00.
		Header: 10,
		W:      160,
		H:      128, // h = npixels/w = 20480/160 = 128
	},
	"objects/flame1.cel": {
		// The contents of frame 0 at offset 0x58 in flame1.cel starts with the
		// following header: 0A 00 C5 00 E5 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	// NOTE: Unused?
	"objects/flame3.cel": {
		// The contents of frame 0 at offset 0x60 in flame3.cel starts with the
		// following header: 0A 00 4C 02 9A 06 AD 08 00 00.
		Header: 10,
		W:      96,
		H:      128, // h = npixels/w = 12288/96 = 128
	},
	// NOTE: Unused?
	"objects/ghost.cel": {
		// The contents of frame 0 at offset 0x40 in ghost.cel starts with the
		// following header: 0A 00 3A 02 49 08 F6 0A 00 00.
		Header: 10,
		W:      128,
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	"objects/goatshrn.cel": {
		// The contents of frame 0 at offset 0x30 in goatshrn.cel starts with the
		// following header: 0A 00 30 05 D6 07 00 00 00 00.
		Header: 10,
		W:      96, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"objects/l1braz.cel": {
		// The contents of frame 0 at offset 0x70 in l1braz.cel starts with the
		// following header: 0A 00 2A 00 4A 00 79 02 99 03.
		Header: 10,
		W:      64,  // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      160, // h = npixels/w = 10240/64 = 160
	},
	"objects/l1doors.cel": {
		// The contents of frame 0 at offset 0x18 in l1doors.cel starts with the
		// following header: 0A 00 32 02 78 06 5F 08 7F 08.
		Header: 10,
		W:      64,  // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      160, // h = npixels/w = 10240/64 = 160
	},
	"objects/l2doors.cel": {
		// The contents of frame 0 at offset 0x18 in l2doors.cel starts with the
		// following header: 0A 00 2C 01 2C 04 9D 05 00 00.
		Header: 10,
		W:      64,  // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      128, // h = npixels/w = 8192/64 = 128
	},
	"objects/l3doors.cel": {
		// The contents of frame 0 at offset 0x18 in l3doors.cel starts with the
		// following header: 0A 00 04 02 3F 06 94 08 00 00.
		Header: 10,
		W:      64,  // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      128, // h = npixels/w = 8192/64 = 128
	},
	"objects/lever.cel": {
		// The contents of frame 0 at offset 0x10 in lever.cel starts with the
		// following header: 0A 00 EA 01 70 02 00 00 00 00.
		Header: 10,
		W:      96, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"objects/lshrineg.cel": {
		// The contents of frame 0 at offset 0x60 in lshrineg.cel starts with the
		// following header: 0A 00 D7 01 33 06 F3 0A 00 00.
		Header: 10,
		W:      128, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	"objects/lzstand.cel": {
		// The contents of frame 0 at offset 0x10 in lzstand.cel starts with the
		// following header: 0A 00 6B 03 49 06 B2 06 00 00.
		Header: 10,
		W:      128, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	"objects/mcirl.cel": {
		// The contents of frame 0 at offset 0x18 in mcirl.cel starts with the
		// following header: 0A 00 55 07 D7 09 00 00 00 00.
		Header: 10,
		W:      96, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"objects/mfountn.cel": {
		// The contents of frame 0 at offset 0x30 in mfountn.cel starts with the
		// following header: 0A 00 E9 07 D4 11 F4 11 00 00.
		Header: 10,
		W:      128, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	"objects/miniwatr.cel": {
		// The contents of frame 0 at offset 0x30 in miniwatr.cel starts with the
		// following header: 0A 00 2A 00 5A 00 D7 04 00 00.
		Header: 10,
		W:      64,  // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      128, // h = npixels/w = 8192/64 = 128
	},
	"objects/mushptch.cel": {
		// The contents of frame 0 at offset 0x10 in mushptch.cel starts with the
		// following header: 0A 00 27 02 ED 06 00 00 00 00.
		Header: 10,
		W:      96, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"objects/nude2.cel": {
		// The contents of frame 0 at offset 0x20 in nude2.cel starts with the
		// following header: 0A 00 6D 00 F5 01 DE 03 00 00.
		Header: 10,
		W:      128, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	"objects/pedistl.cel": {
		// The contents of frame 0 at offset 0x18 in pedistl.cel starts with the
		// following header: 0A 00 08 02 0D 04 00 00 00 00.
		Header: 10,
		W:      96, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"objects/pfountn.cel": {
		// The contents of frame 0 at offset 0x30 in pfountn.cel starts with the
		// following header: 0A 00 97 05 EB 0E 0B 0F 00 00.
		Header: 10,
		W:      128, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	// NOTE: Unused?
	"objects/prsrplt1.cel": {
		// The contents of frame 0 at offset 0x30 in prsrplt1.cel starts with the
		// following header: 0A 00 F9 02 19 03 00 00 00 00.
		Header: 10,
		W:      96, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"objects/rockstan.cel": {
		// The contents of frame 0 at offset 0x0C in rockstan.cel starts with the
		// following header: 0A 00 3C 01 1A 03 00 00 00 00.
		Header: 10,
		W:      96, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"objects/rshrineg.cel": {
		// The contents of frame 0 at offset 0x60 in rshrineg.cel starts with the
		// following header: 0A 00 67 01 C8 05 85 0A 00 00.
		Header: 10,
		W:      128, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	"objects/sarc.cel": {
		// The contents of frame 0 at offset 0x1C in sarc.cel starts with the
		// following header: 0A 00 6B 08 12 13 00 00 00 00.
		Header: 10,
		W:      128, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      96,  // h = npixels/w = 12288/128 = 96
	},
	"objects/skulfire.cel": {
		// The contents of frame 0 at offset 0x34 in skulfire.cel starts with the
		// following header: 0A 00 3B 02 16 03 00 00 00 00.
		Header: 10,
		W:      96, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"objects/skulpile.cel": {
		// The contents of frame 0 at offset 0x0C in skulpile.cel starts with the
		// following header: 0A 00 F5 01 15 02 00 00 00 00.
		Header: 10,
		W:      96, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"objects/skulstik.cel": {
		// The contents of frame 0 at offset 0x1C in skulstik.cel starts with the
		// following header: 0A 00 B8 01 69 03 00 00 00 00.
		Header: 10,
		W:      96, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	// NOTE: Unused?
	"objects/switch2.cel": {
		// The contents of frame 0 at offset 0x10 in switch2.cel starts with the
		// following header: 0A 00 4F 01 A9 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	// NOTE: Unused?
	"objects/switch3.cel": {
		// The contents of frame 0 at offset 0x10 in switch4.cel starts with the
		// following header: 0A 00 2B 02 C0 02 00 00 00 00.
		Header: 10,
		W:      96, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"objects/switch4.cel": {
		// The contents of frame 0 at offset 0x10 in switch4.cel starts with the
		// following header: 0A 00 2B 02 C0 02 00 00 00 00.
		Header: 10,
		W:      96, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"objects/tfountn.cel": {
		// The contents of frame 0 at offset 0x28 in tfountn.cel starts with the
		// following header: 0A 00 77 04 1E 0B 00 00 00 00.
		Header: 10,
		W:      128, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      96,  // h = npixels/w = 12288/128 = 96
	},
	"objects/tnudem.cel": {
		// The contents of frame 0 at offset 0x18 in tnudem.cel starts with the
		// following header: 0A 00 86 01 D2 04 FD 05 00 00.
		Header: 10,
		W:      128, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	"objects/tnudew.cel": {
		// The contents of frame 0 at offset 0x14 in tnudew.cel starts with the
		// following header: 0A 00 7A 02 33 05 AD 06 00 00.
		Header: 10,
		W:      128, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	"objects/traphole.cel": {
		// The contents of frame 0 at offset 0x10 in traphole.cel starts with the
		// following header: 0A 00 2A 00 C4 00 E4 00 04 01.
		Header: 10,
		W:      64,  // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      144, // h = npixels/w = 9216/64 = 144
	},
	"objects/tsoul.cel": {
		// The contents of frame 0 at offset 0x20 in tsoul.cel starts with the
		// following header: 0A 00 FE 01 78 04 00 00 00 00.
		Header: 10,
		W:      128, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      96,  // h = npixels/w = 12288/128 = 96
	},
	// NOTE: Unused?
	"objects/vapor1.cel": {
		// The contents of frame 0 at offset 0x3C in vapor1.cel starts with the
		// following header: 0A 00 3B 07 50 11 75 15 00 00.
		Header: 10,
		W:      128,
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	// NOTE: Unused?
	"objects/water.cel": {
		// The contents of frame 0 at offset 0x30 in water.cel starts with the
		// following header: 0A 00 2A 00 2A 05 B1 10 76 11.
		Header: 10,
		W:      128,
		H:      160, // h = npixels/w = 20480/128 = 160
	},
	// NOTE: Unused?
	"objects/waterjug.cel": {
		// The contents of frame 0 at offset 0x18 in waterjug.cel starts with the
		// following header: 0A 00 3C 02 05 03 00 00 00 00.
		Header: 10,
		W:      96, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"objects/weapstnd.cel": {
		// The contents of frame 0 at offset 0x18 in weapstnd.cel starts with the
		// following header: 0A 00 8D 05 4F 0B 00 00 00 00.
		Header: 10,
		W:      96, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"objects/wtorch1.cel": {
		// The contents of frame 0 at offset 0x2C in wtorch1.cel starts with the
		// following header: 0A 00 2A 00 49 01 D8 02 00 00.
		Header: 10,
		W:      96,  // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      128, // h = npixels/w = 12288/96 = 128
	},
	"objects/wtorch2.cel": {
		// The contents of frame 0 at offset 0x2C in wtorch2.cel starts with the
		// following header: 0A 00 2A 00 4A 01 E3 02 00 00.
		Header: 10,
		W:      96,  // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      128, // h = npixels/w = 12288/96 = 128
	},
	"objects/wtorch3.cel": {
		// The contents of frame 0 at offset 0x2C in wtorch3.cel starts with the
		// following header: 0A 00 2A 00 E7 00 6F 02 00 00.
		Header: 10,
		W:      96,  // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      128, // h = npixels/w = 12288/96 = 128
	},
	"objects/wtorch4.cel": {
		// The contents of frame 0 at offset 0x2C in wtorch4.cel starts with the
		// following header: 0A 00 2A 00 EB 00 6D 02 00 00.
		Header: 10,
		W:      96, // ref: cross-referencing 0x49F450 and 0x4A0554
		H:      128,
	},
	"towners/animals/cow.cel": {
		// The contents of the archive header at offset 0 in cow.cel:
		//    20 00 00 00  D7 49 00 00  F1 C2 00 00  3B 48 01 00
		//    9F C1 01 00  28 1B 02 00  A6 90 02 00  8C 10 03 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x38 in cow.cel starts with the
		// following header: 0A 00 94 00 0C 04 D5 05 00 00.
		Header: 10,
		W:      128, // ref: 0x460861
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	"towners/butch/deadguy.cel": {
		// The contents of frame 0 at offset 0x28 in deadguy.cel starts with the
		// following header: 0A 00 FC 03 55 04 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x46044A
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"towners/drunk/twndrunk.cel": {
		// The contents of frame 0 at offset 0x50 in twndrunk.cel starts with the
		// following header: 0A 00 8B 02 8D 03 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4607AC
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"towners/healer/healer.cel": {
		// The contents of frame 0 at offset 0x58 in healer.cel starts with the
		// following header: 0A 00 BB 01 1A 04 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x46068D
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	// NOTE: Unused?
	"towners/priest/priest8.cel": {
		// The contents of frame 0 at offset 0x8C in priest8.cel starts with the
		// following header: 0A 00 E6 01 A7 04 00 00 00 00.
		Header: 10,
		W:      96,
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"towners/smith/smithn.cel": {
		// The contents of frame 0 at offset 0x48 in smithn.cel starts with the
		// following header: 0A 00 21 02 7D 05 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x460324
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	// NOTE: Unused?
	"towners/smith/smithw.cel": {
		// The contents of the archive header at offset 0 in smithw.cel:
		//    20 00 00 00  6B 38 00 00  03 6B 00 00  DE 98 00 00
		//    01 CA 00 00  93 FE 00 00  C5 2F 01 00  B5 5C 01 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x28 in smithw.cel starts with the
		// following header: 0A 00 65 02 2F 06 00 00 00 00.
		Header: 10,
		W:      96,
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"towners/strytell/strytell.cel": {
		// The contents of frame 0 at offset 0x6C in strytell.cel starts with the
		// following header: 0A 00 E6 01 54 05 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x46071D
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"towners/townboy/pegkid1.cel": {
		// The contents of frame 0 at offset 0x58 in pegkid1.cel starts with the
		// following header: 0A 00 4E 01 00 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4605FA
		H:      64, // h = npixels/w = 6144/96 = 64
	},
	"towners/townwmn1/witch.cel": {
		// The contents of frame 0 at offset 0x54 in witch.cel starts with the
		// following header: 0A 00 9D 01 23 04 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4604DA
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"towners/townwmn1/wmnn.cel": {
		// The contents of frame 0 at offset 0x50 in wmnn.cel starts with the
		// following header: 0A 00 80 01 9C 03 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x460569
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	// NOTE: Unused?
	"towners/townwmn1/wmnw.cel": {
		// The contents of the archive header at offset 0 in wmnw.cel:
		//    20 00 00 00  0E 24 00 00  6B 44 00 00  51 60 00 00
		//    13 80 00 00  D1 A3 00 00  0F C3 00 00  A6 DE 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x28 in wmnw.cel starts with the
		// following header: 0A 00 9D 01 DF 03 00 00 00 00.
		Header: 10,
		W:      96,
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"towners/twnf/twnfn.cel": {
		// The contents of frame 0 at offset 0x48 in twnfn.cel starts with the
		// following header: 0A 00 9B 01 46 04 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4603BA
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	// NOTE: Unused?
	"towners/twnf/twnfw.cel": {
		// The contents of the archive header at offset 0 in twnfw.cel:
		//    20 00 00 00  B3 2B 00 00  96 53 00 00  A3 75 00 00
		//    BC 9C 00 00  1A C7 00 00  E9 ED 00 00  47 0F 01 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x28 in twnfw.cel starts with the
		// following header: 0A 00 E0 01 A8 04 00 00 00 00.
		Header: 10,
		W:      96,
		H:      96, // h = npixels/w = 9216/96 = 96
	},

	// CL2 files.
	"missiles/acidbf1.cl2": {
		// The contents of frame 0 at offset 0x2C in acidbf1.cl2 starts with the
		// following header: 0A 00 59 00 B8 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/acidbf2.cl2": {
		// The contents of frame 0 at offset 0x28 in acidbf2.cl2 starts with the
		// following header: 0A 00 61 00 CA 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/acidbf3.cl2": {
		// The contents of frame 0 at offset 0x28 in acidbf3.cl2 starts with the
		// following header: 0A 00 6C 00 DD 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/acidbf4.cl2": {
		// The contents of frame 0 at offset 0x28 in acidbf4.cl2 starts with the
		// following header: 0A 00 76 00 E3 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/acidbf5.cl2": {
		// The contents of frame 0 at offset 0x28 in acidbf5.cl2 starts with the
		// following header: 0A 00 7C 00 EF 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/acidbf6.cl2": {
		// The contents of frame 0 at offset 0x28 in acidbf6.cl2 starts with the
		// following header: 0A 00 8E 00 F0 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/acidbf7.cl2": {
		// The contents of frame 0 at offset 0x28 in acidbf7.cl2 starts with the
		// following header: 0A 00 8B 00 E1 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/acidbf8.cl2": {
		// The contents of frame 0 at offset 0x28 in acidbf8.cl2 starts with the
		// following header: 0A 00 82 00 D5 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/acidbf9.cl2": {
		// The contents of frame 0 at offset 0x28 in acidbf9.cl2 starts with the
		// following header: 0A 00 79 00 D0 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/acidbf10.cl2": {
		// The contents of frame 0 at offset 0x2C in acidbf10.cl2 starts with the
		// following header: 0A 00 79 00 C8 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/acidbf11.cl2": {
		// The contents of frame 0 at offset 0x2C in acidbf11.cl2 starts with the
		// following header: 0A 00 7D 00 D0 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/acidbf12.cl2": {
		// The contents of frame 0 at offset 0x28 in acidbf12.cl2 starts with the
		// following header: 0A 00 74 00 D1 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/acidbf13.cl2": {
		// The contents of frame 0 at offset 0x28 in acidbf13.cl2 starts with the
		// following header: 0A 00 72 00 E9 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/acidbf14.cl2": {
		// The contents of frame 0 at offset 0x28 in acidbf14.cl2 starts with the
		// following header: 0A 00 69 00 E5 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/acidbf15.cl2": {
		// The contents of frame 0 at offset 0x28 in acidbf15.cl2 starts with the
		// following header: 0A 00 62 00 D8 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/acidbf16.cl2": {
		// The contents of frame 0 at offset 0x28 in acidbf16.cl2 starts with the
		// following header: 0A 00 5D 00 CA 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/acidpud1.cl2": {
		// The contents of frame 0 at offset 0x2C in acidpud1.cl2 starts with the
		// following header: 0A 00 3B 01 54 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/acidpud2.cl2": {
		// The contents of frame 0 at offset 0x18 in acidpud2.cl2 starts with the
		// following header: 0A 00 B0 01 C9 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/acidspla.cl2": {
		// The contents of frame 0 at offset 0x2C in acidspla.cl2 starts with the
		// following header: 0A 00 B3 00 44 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/arrows.cl2": {
		// The contents of frame 0 at offset 0x48 in arrows.cl2 starts with the
		// following header: 0A 00 23 00 6E 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/bigexp.cl2": {
		// The contents of frame 0 at offset 0x48 in bigexp.cl2 starts with the
		// following header: 0A 00 0D 01 F1 02 1A 03 43 03.
		Header: 10,
		W:      160, // ref: 0x4950B8
		H:      160, // h = npixels/w = 25600/160 = 160
	},
	// NOTE: Unused?
	"missiles/blodbur0.cl2": {
		// The contents of frame 0 at offset 0x28 in blodbur0.cl2 starts with the
		// following header: 0A 00 2B 00 F6 01 85 02 00 00.
		Header: 10,
		W:      128, // TODO: Verify frame width.
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	"missiles/blodbur1.cl2": {
		// The contents of frame 0 at offset 0x28 in blodbur1.cl2 starts with the
		// following header: 0A 00 33 00 C1 01 E2 01 00 00.
		Header: 10,
		W:      128, // ref: 0x4950B8
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	"missiles/blodbur2.cl2": {
		// The contents of frame 0 at offset 0x28 in blodbur2.cl2 starts with the
		// following header: 0A 00 2B 00 F6 01 85 02 00 00.
		Header: 10,
		W:      128, // ref: 0x4950B8
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	// NOTE: Unused?
	"missiles/blodburs.cl2": {
		// The contents of frame 0 at offset 0x28 in blodburs.cl2 starts with the
		// following header: 0A 00 33 00 C1 01 E2 01 00 00.
		Header: 10,
		W:      128, // TODO: Verify frame width.
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	"missiles/blood1.cl2": {
		// The contents of frame 0 at offset 0x44 in blood1.cl2 starts with the
		// following header: 0A 00 23 00 52 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/blood2.cl2": {
		// The contents of frame 0 at offset 0x28 in blood2.cl2 starts with the
		// following header: 0A 00 33 00 C1 01 E2 01 00 00.
		Header: 10,
		W:      128, // ref: 0x4950B8
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	"missiles/blood3.cl2": {
		// The contents of frame 0 at offset 0x28 in blood3.cl2 starts with the
		// following header: 0A 00 2B 00 F6 01 85 02 00 00.
		Header: 10,
		W:      128, // ref: 0x4950B8
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	"missiles/blood4.cl2": {
		// The contents of frame 0 at offset 0x44 in blood4.cl2 starts with the
		// following header: 0A 00 2B 00 93 00 00 00 00 00.
		Header: 10,
		W:      128, // ref: 0x4950B8
		H:      96,  // h = npixels/w = 12288/128 = 96
	},
	"missiles/bluexbk.cl2": {
		// The contents of frame 0 at offset 0x54 in bluexbk.cl2 starts with the
		// following header: 0A 00 33 00 62 00 8B 00 00 00.
		Header: 10,
		W:      160, // ref: 0x4950B8
		H:      128, // h = npixels/w = 20480/160 = 128
	},
	"missiles/bluexfr.cl2": {
		// The contents of frame 0 at offset 0x54 in bluexfr.cl2 starts with the
		// following header: 0A 00 33 00 62 00 8B 00 00 00.
		Header: 10,
		W:      160, // ref: 0x4950B8
		H:      128, // h = npixels/w = 20480/160 = 128
	},
	"missiles/bone1.cl2": {
		// The contents of frame 0 at offset 0x28 in bone1.cl2 starts with the
		// following header: 0A 00 2B 00 1B 01 45 01 00 00.
		Header: 10,
		W:      128, // ref: 0x4950B8
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	"missiles/bone2.cl2": {
		// The contents of frame 0 at offset 0x28 in bone2.cl2 starts with the
		// following header: 0A 00 6D 00 3B 01 5E 01 00 00.
		Header: 10,
		W:      128, // ref: 0x4950B8
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	"missiles/bone3.cl2": {
		// The contents of frame 0 at offset 0x28 in bone3.cl2 starts with the
		// following header: 0A 00 6D 00 8C 01 B6 01 00 00.
		Header: 10,
		W:      128, // ref: 0x4950B8
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	"missiles/doom1.cl2": {
		// The contents of frame 0 at offset 0x44 in doom1.cl2 starts with the
		// following header: 0A 00 23 00 30 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/doom2.cl2": {
		// The contents of frame 0 at offset 0x44 in doom2.cl2 starts with the
		// following header: 0A 00 23 00 0B 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/doom3.cl2": {
		// The contents of frame 0 at offset 0x44 in doom3.cl2 starts with the
		// following header: 0A 00 23 00 36 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/doom4.cl2": {
		// The contents of frame 0 at offset 0x44 in doom4.cl2 starts with the
		// following header: 0A 00 23 00 FF 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/doom5.cl2": {
		// The contents of frame 0 at offset 0x44 in doom5.cl2 starts with the
		// following header: 0A 00 23 00 DB 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/doom6.cl2": {
		// The contents of frame 0 at offset 0x44 in doom6.cl2 starts with the
		// following header: 0A 00 23 00 15 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/doom7.cl2": {
		// The contents of frame 0 at offset 0x44 in doom7.cl2 starts with the
		// following header: 0A 00 23 00 32 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/doom8.cl2": {
		// The contents of frame 0 at offset 0x44 in doom8.cl2 starts with the
		// following header: 0A 00 23 00 13 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/doom9.cl2": {
		// The contents of frame 0 at offset 0x44 in doom9.cl2 starts with the
		// following header: 0A 00 62 00 85 00 9E 00 00 00.
		Header: 10,
		W:      96,  // ref: 0x4950B8
		H:      128, // h = npixels/w = 12288/96 = 128
	},
	// NOTE: Unused?
	"missiles/doomexp.cl2": {
		// The contents of frame 0 at offset 0x44 in doomexp.cl2 starts with the
		// following header: 0A 00 62 00 85 00 9E 00 00 00.
		Header: 10,
		W:      96,  // TODO: Verify frame width.
		H:      128, // h = npixels/w = 12288/96 = 128
	},
	"missiles/ethrshld.cl2": {
		// The contents of frame 0 at offset 0x0C in ethrshld.cl2 starts with the
		// following header: 0A 00 23 00 3C 00 81 00 00 00.
		Header: 10,
		W:      96,  // ref: 0x4950B8
		H:      128, // h = npixels/w = 12288/96 = 128
	},
	"missiles/farrow1.cl2": {
		// The contents of frame 0 at offset 0x18 in farrow1.cl2 starts with the
		// following header: 0A 00 23 00 8C 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/farrow2.cl2": {
		// The contents of frame 0 at offset 0x18 in farrow2.cl2 starts with the
		// following header: 0A 00 23 00 93 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/farrow3.cl2": {
		// The contents of frame 0 at offset 0x18 in farrow3.cl2 starts with the
		// following header: 0A 00 23 00 96 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/farrow4.cl2": {
		// The contents of frame 0 at offset 0x18 in farrow4.cl2 starts with the
		// following header: 0A 00 23 00 95 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/farrow5.cl2": {
		// The contents of frame 0 at offset 0x18 in farrow5.cl2 starts with the
		// following header: 0A 00 23 00 70 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/farrow6.cl2": {
		// The contents of frame 0 at offset 0x18 in farrow6.cl2 starts with the
		// following header: 0A 00 23 00 49 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/farrow7.cl2": {
		// The contents of frame 0 at offset 0x18 in farrow7.cl2 starts with the
		// following header: 0A 00 23 00 4B 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/farrow8.cl2": {
		// The contents of frame 0 at offset 0x18 in farrow8.cl2 starts with the
		// following header: 0A 00 23 00 4B 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/farrow9.cl2": {
		// The contents of frame 0 at offset 0x18 in farrow9.cl2 starts with the
		// following header: 0A 00 23 00 4C 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/farrow10.cl2": {
		// The contents of frame 0 at offset 0x18 in farrow10.cl2 starts with the
		// following header: 0A 00 23 00 54 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/farrow11.cl2": {
		// The contents of frame 0 at offset 0x18 in farrow11.cl2 starts with the
		// following header: 0A 00 23 00 59 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/farrow12.cl2": {
		// The contents of frame 0 at offset 0x18 in farrow12.cl2 starts with the
		// following header: 0A 00 23 00 83 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/farrow13.cl2": {
		// The contents of frame 0 at offset 0x18 in farrow13.cl2 starts with the
		// following header: 0A 00 23 00 A2 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/farrow14.cl2": {
		// The contents of frame 0 at offset 0x18 in farrow14.cl2 starts with the
		// following header: 0A 00 23 00 AA 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/farrow15.cl2": {
		// The contents of frame 0 at offset 0x18 in farrow15.cl2 starts with the
		// following header: 0A 00 23 00 9E 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/farrow16.cl2": {
		// The contents of frame 0 at offset 0x18 in farrow16.cl2 starts with the
		// following header: 0A 00 23 00 8E 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/firarwex.cl2": {
		// The contents of frame 0 at offset 0x20 in firarwex.cl2 starts with the
		// following header: 0A 00 DA 00 00 00 00 00 00 00.
		Header: 10,
		W:      64, // ref: 0x4950B8
		H:      64, // h = npixels/w = 4096/64 = 64
	},
	"missiles/fireba1.cl2": {
		// The contents of frame 0 at offset 0x40 in fireba1.cl2 starts with the
		// following header: 0A 00 81 00 3A 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/fireba2.cl2": {
		// The contents of frame 0 at offset 0x40 in fireba2.cl2 starts with the
		// following header: 0A 00 82 00 3F 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/fireba3.cl2": {
		// The contents of frame 0 at offset 0x40 in fireba3.cl2 starts with the
		// following header: 0A 00 8C 00 3B 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/fireba4.cl2": {
		// The contents of frame 0 at offset 0x40 in fireba4.cl2 starts with the
		// following header: 0A 00 6F 00 BC 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/fireba5.cl2": {
		// The contents of frame 0 at offset 0x40 in fireba5.cl2 starts with the
		// following header: 0A 00 A7 00 AF 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/fireba6.cl2": {
		// The contents of frame 0 at offset 0x40 in fireba6.cl2 starts with the
		// following header: 0A 00 8A 00 AD 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/fireba7.cl2": {
		// The contents of frame 0 at offset 0x40 in fireba7.cl2 starts with the
		// following header: 0A 00 1C 01 A0 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/fireba8.cl2": {
		// The contents of frame 0 at offset 0x40 in fireba8.cl2 starts with the
		// following header: 0A 00 E7 00 D1 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/fireba9.cl2": {
		// The contents of frame 0 at offset 0x40 in fireba9.cl2 starts with the
		// following header: 0A 00 C1 00 C5 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/fireba10.cl2": {
		// The contents of frame 0 at offset 0x40 in fireba10.cl2 starts with the
		// following header: 0A 00 E7 00 D1 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/fireba11.cl2": {
		// The contents of frame 0 at offset 0x40 in fireba11.cl2 starts with the
		// following header: 0A 00 1C 01 A0 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/fireba12.cl2": {
		// The contents of frame 0 at offset 0x40 in fireba12.cl2 starts with the
		// following header: 0A 00 0F 01 AD 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/fireba13.cl2": {
		// The contents of frame 0 at offset 0x40 in fireba13.cl2 starts with the
		// following header: 0A 00 DF 00 AF 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/fireba14.cl2": {
		// The contents of frame 0 at offset 0x40 in fireba14.cl2 starts with the
		// following header: 0A 00 11 01 BB 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/fireba15.cl2": {
		// The contents of frame 0 at offset 0x40 in fireba15.cl2 starts with the
		// following header: 0A 00 6A 00 3A 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/fireba16.cl2": {
		// The contents of frame 0 at offset 0x40 in fireba16.cl2 starts with the
		// following header: 0A 00 82 00 3F 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/fireplar.cl2": {
		// The contents of frame 0 at offset 0x4C in fireplar.cl2 starts with the
		// following header: 0A 00 DD 02 11 05 3A 05 63 05.
		Header: 10,
		W:      160, // ref: 0x4950B8
		H:      160, // h = npixels/w = 25600/160 = 160
	},
	"missiles/firerun1.cl2": {
		// The contents of frame 0 at offset 0x38 in firerun1.cl2 starts with the
		// following header: 0A 00 13 01 FA 03 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/firerun2.cl2": {
		// The contents of frame 0 at offset 0x38 in firerun2.cl2 starts with the
		// following header: 0A 00 2D 01 DC 03 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/firerun3.cl2": {
		// The contents of frame 0 at offset 0x38 in firerun3.cl2 starts with the
		// following header: 0A 00 4F 01 D3 03 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/firerun4.cl2": {
		// The contents of frame 0 at offset 0x38 in firerun4.cl2 starts with the
		// following header: 0A 00 0E 01 70 03 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/firerun5.cl2": {
		// The contents of frame 0 at offset 0x38 in firerun5.cl2 starts with the
		// following header: 0A 00 E7 00 50 03 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/firerun6.cl2": {
		// The contents of frame 0 at offset 0x38 in firerun6.cl2 starts with the
		// following header: 0A 00 1F 01 93 03 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/firerun7.cl2": {
		// The contents of frame 0 at offset 0x38 in firerun7.cl2 starts with the
		// following header: 0A 00 40 01 9F 03 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/firerun8.cl2": {
		// The contents of frame 0 at offset 0x38 in firerun8.cl2 starts with the
		// following header: 0A 00 49 01 B4 03 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/firewal1.cl2": {
		// The contents of frame 0 at offset 0x3C in firewal1.cl2 starts with the
		// following header: 0A 00 2F 00 50 00 71 00 00 00.
		Header: 10,
		W:      128, // ref: 0x4950B8
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	"missiles/firewal2.cl2": {
		// The contents of frame 0 at offset 0x3C in firewal2.cl2 starts with the
		// following header: 0A 00 CE 04 E7 0C 61 0E 00 00.
		Header: 10,
		W:      128, // ref: 0x4950B8
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	"missiles/flare.cl2": {
		// The contents of frame 0 at offset 0x48 in flare.cl2 starts with the
		// following header: 0A 00 4B 00 F6 02 4D 03 00 00.
		Header: 10,
		W:      128, // ref: 0x4950B8
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	"missiles/flareexp.cl2": {
		// The contents of frame 0 at offset 0x24 in flareexp.cl2 starts with the
		// following header: 0A 00 2B 00 4E 00 6F 00 00 00.
		Header: 10,
		W:      128, // ref: 0x4950B8
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	"missiles/guard1.cl2": {
		// The contents of frame 0 at offset 0x44 in guard1.cl2 starts with the
		// following header: 0A 00 10 01 2B 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/guard2.cl2": {
		// The contents of frame 0 at offset 0x40 in guard2.cl2 starts with the
		// following header: 0A 00 18 04 0E 08 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/guard3.cl2": {
		// The contents of frame 0 at offset 0x14 in guard3.cl2 starts with the
		// following header: 0A 00 23 04 68 08 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/holy1.cl2": {
		// The contents of frame 0 at offset 0x40 in holy1.cl2 starts with the
		// following header: 0A 00 80 00 33 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/holy2.cl2": {
		// The contents of frame 0 at offset 0x40 in holy2.cl2 starts with the
		// following header: 0A 00 81 00 3A 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/holy3.cl2": {
		// The contents of frame 0 at offset 0x40 in holy3.cl2 starts with the
		// following header: 0A 00 89 00 33 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/holy4.cl2": {
		// The contents of frame 0 at offset 0x40 in holy4.cl2 starts with the
		// following header: 0A 00 64 00 A0 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/holy5.cl2": {
		// The contents of frame 0 at offset 0x40 in holy5.cl2 starts with the
		// following header: 0A 00 94 00 95 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/holy6.cl2": {
		// The contents of frame 0 at offset 0x40 in holy6.cl2 starts with the
		// following header: 0A 00 75 00 96 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/holy7.cl2": {
		// The contents of frame 0 at offset 0x40 in holy7.cl2 starts with the
		// following header: 0A 00 08 01 8C 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/holy8.cl2": {
		// The contents of frame 0 at offset 0x40 in holy8.cl2 starts with the
		// following header: 0A 00 E2 00 CC 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/holy9.cl2": {
		// The contents of frame 0 at offset 0x40 in holy9.cl2 starts with the
		// following header: 0A 00 C1 00 C3 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/holy10.cl2": {
		// The contents of frame 0 at offset 0x40 in holy10.cl2 starts with the
		// following header: 0A 00 E2 00 CC 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/holy11.cl2": {
		// The contents of frame 0 at offset 0x40 in holy11.cl2 starts with the
		// following header: 0A 00 08 01 8C 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/holy12.cl2": {
		// The contents of frame 0 at offset 0x40 in holy12.cl2 starts with the
		// following header: 0A 00 FA 00 96 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/holy13.cl2": {
		// The contents of frame 0 at offset 0x40 in holy13.cl2 starts with the
		// following header: 0A 00 CB 00 95 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/holy14.cl2": {
		// The contents of frame 0 at offset 0x40 in holy14.cl2 starts with the
		// following header: 0A 00 FD 00 9F 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/holy15.cl2": {
		// The contents of frame 0 at offset 0x40 in holy15.cl2 starts with the
		// following header: 0A 00 67 00 32 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/holy16.cl2": {
		// The contents of frame 0 at offset 0x40 in holy16.cl2 starts with the
		// following header: 0A 00 81 00 3A 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/holyexpl.cl2": {
		// The contents of frame 0 at offset 0x28 in holyexpl.cl2 starts with the
		// following header: 0A 00 33 00 22 01 4B 01 74 01.
		Header: 10,
		W:      160, // ref: 0x4950B8
		H:      160, // h = npixels/w = 25600/160 = 160
	},
	"missiles/inferno.cl2": {
		// The contents of frame 0 at offset 0x58 in inferno.cl2 starts with the
		// following header: 0A 00 31 00 A3 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/krull.cl2": {
		// The contents of frame 0 at offset 0x40 in krull.cl2 starts with the
		// following header: 0A 00 59 00 72 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/larrow1.cl2": {
		// The contents of frame 0 at offset 0x18 in larrow1.cl2 starts with the
		// following header: 0A 00 23 00 9F 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/larrow2.cl2": {
		// The contents of frame 0 at offset 0x18 in larrow2.cl2 starts with the
		// following header: 0A 00 23 00 A2 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/larrow3.cl2": {
		// The contents of frame 0 at offset 0x18 in larrow3.cl2 starts with the
		// following header: 0A 00 23 00 B0 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/larrow4.cl2": {
		// The contents of frame 0 at offset 0x18 in larrow4.cl2 starts with the
		// following header: 0A 00 23 00 BC 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/larrow5.cl2": {
		// The contents of frame 0 at offset 0x18 in larrow5.cl2 starts with the
		// following header: 0A 00 23 00 A1 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/larrow6.cl2": {
		// The contents of frame 0 at offset 0x18 in larrow6.cl2 starts with the
		// following header: 0A 00 23 00 60 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/larrow7.cl2": {
		// The contents of frame 0 at offset 0x18 in larrow7.cl2 starts with the
		// following header: 0A 00 23 00 51 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/larrow8.cl2": {
		// The contents of frame 0 at offset 0x18 in larrow8.cl2 starts with the
		// following header: 0A 00 23 00 4C 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/larrow9.cl2": {
		// The contents of frame 0 at offset 0x18 in larrow9.cl2 starts with the
		// following header: 0A 00 23 00 4C 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/larrow10.cl2": {
		// The contents of frame 0 at offset 0x18 in larrow10.cl2 starts with the
		// following header: 0A 00 23 00 48 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/larrow11.cl2": {
		// The contents of frame 0 at offset 0x18 in larrow11.cl2 starts with the
		// following header: 0A 00 23 00 54 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/larrow12.cl2": {
		// The contents of frame 0 at offset 0x18 in larrow12.cl2 starts with the
		// following header: 0A 00 23 00 96 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/larrow13.cl2": {
		// The contents of frame 0 at offset 0x18 in larrow13.cl2 starts with the
		// following header: 0A 00 23 00 C8 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/larrow14.cl2": {
		// The contents of frame 0 at offset 0x18 in larrow14.cl2 starts with the
		// following header: 0A 00 23 00 C6 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/larrow15.cl2": {
		// The contents of frame 0 at offset 0x18 in larrow15.cl2 starts with the
		// following header: 0A 00 23 00 B0 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/larrow16.cl2": {
		// The contents of frame 0 at offset 0x18 in larrow16.cl2 starts with the
		// following header: 0A 00 23 00 84 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/lghning.cl2": {
		// The contents of frame 0 at offset 0x28 in lghning.cl2 starts with the
		// following header: 0A 00 7F 00 87 04 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/magball1.cl2": {
		// The contents of frame 0 at offset 0x48 in magball1.cl2 starts with the
		// following header: 0A 00 2B 00 4C 00 41 01 00 00.
		Header: 10,
		W:      128, // ref: 0x4950B8
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	"missiles/magball2.cl2": {
		// The contents of frame 0 at offset 0x48 in magball2.cl2 starts with the
		// following header: 0A 00 2B 00 4C 00 72 01 00 00.
		Header: 10,
		W:      128, // ref: 0x4950B8
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	"missiles/magball3.cl2": {
		// The contents of frame 0 at offset 0x48 in magball3.cl2 starts with the
		// following header: 0A 00 2B 00 4C 00 76 01 00 00.
		Header: 10,
		W:      128, // ref: 0x4950B8
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	"missiles/magball4.cl2": {
		// The contents of frame 0 at offset 0x48 in magball4.cl2 starts with the
		// following header: 0A 00 2B 00 4C 00 5F 01 00 00.
		Header: 10,
		W:      128, // ref: 0x4950B8
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	"missiles/magball5.cl2": {
		// The contents of frame 0 at offset 0x48 in magball5.cl2 starts with the
		// following header: 0A 00 2B 00 4C 00 4F 01 00 00.
		Header: 10,
		W:      128, // ref: 0x4950B8
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	"missiles/magball6.cl2": {
		// The contents of frame 0 at offset 0x48 in magball6.cl2 starts with the
		// following header: 0A 00 2B 00 4C 00 69 01 00 00.
		Header: 10,
		W:      128, // ref: 0x4950B8
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	"missiles/magball7.cl2": {
		// The contents of frame 0 at offset 0x48 in magball7.cl2 starts with the
		// following header: 0A 00 2B 00 4C 00 7A 01 00 00.
		Header: 10,
		W:      128, // ref: 0x4950B8
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	"missiles/magball8.cl2": {
		// The contents of frame 0 at offset 0x48 in magball8.cl2 starts with the
		// following header: 0A 00 2B 00 4C 00 6E 01 00 00.
		Header: 10,
		W:      128, // ref: 0x4950B8
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	"missiles/magblos.cl2": {
		// The contents of frame 0 at offset 0x30 in magblos.cl2 starts with the
		// following header: 0A 00 52 00 29 01 4A 01 00 00.
		Header: 10,
		W:      128, // ref: 0x4950B8
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	"missiles/manashld.cl2": {
		// The contents of frame 0 at offset 0x0C in manashld.cl2 starts with the
		// following header: 0A 00 23 00 3C 00 55 00 00 00.
		Header: 10,
		W:      96,  // ref: 0x4950B8
		H:      128, // h = npixels/w = 12288/96 = 128
	},
	"missiles/metlhit1.cl2": {
		// The contents of frame 0 at offset 0x30 in metlhit1.cl2 starts with the
		// following header: 0A 00 23 00 61 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/metlhit2.cl2": {
		// The contents of frame 0 at offset 0x30 in metlhit2.cl2 starts with the
		// following header: 0A 00 23 00 71 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/metlhit3.cl2": {
		// The contents of frame 0 at offset 0x30 in metlhit3.cl2 starts with the
		// following header: 0A 00 23 00 8F 00 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/miniltng.cl2": {
		// The contents of frame 0 at offset 0x28 in miniltng.cl2 starts with the
		// following header: 0A 00 E7 00 00 00 00 00 00 00.
		Header: 10,
		W:      64, // ref: 0x4950B8
		H:      64, // h = npixels/w = 4096/64 = 64
	},
	"missiles/newexp.cl2": {
		// The contents of frame 0 at offset 0x44 in newexp.cl2 starts with the
		// following header: 0A 00 3C 00 55 00 6E 00 87 00.
		Header: 10,
		W:      96,  // ref: 0x4950B8
		H:      160, // h = npixels/w = 15360/96 = 160
	},
	// NOTE: Unused?
	"missiles/portal.cl2": {
		// The contents of frame 0 at offset 0x48 in portal.cl2 starts with the
		// following header: 0A 00 96 02 0F 06 7C 09 00 00.
		Header: 10,
		W:      96,  // TODO: Verify frame width.
		H:      128, // h = npixels/w = 12288/96 = 128
	},
	"missiles/portal1.cl2": {
		// The contents of frame 0 at offset 0x48 in portal1.cl2 starts with the
		// following header: 0A 00 73 00 FE 00 5F 01 00 00.
		Header: 10,
		W:      96,  // ref: 0x4950B8
		H:      128, // h = npixels/w = 12288/96 = 128
	},
	"missiles/portal2.cl2": {
		// The contents of frame 0 at offset 0x48 in portal2.cl2 starts with the
		// following header: 0A 00 96 02 0F 06 7C 09 00 00.
		Header: 10,
		W:      96,  // ref: 0x4950B8
		H:      128, // h = npixels/w = 12288/96 = 128
	},
	// NOTE: Unused?
	"missiles/portalu.cl2": {
		// The contents of frame 0 at offset 0x48 in portalu.cl2 starts with the
		// following header: 0A 00 73 00 FE 00 5F 01 00 00.
		Header: 10,
		W:      96,  // TODO: Verify frame width.
		H:      128, // h = npixels/w = 12288/96 = 128
	},
	"missiles/ressur1.cl2": {
		// The contents of frame 0 at offset 0x48 in ressur1.cl2 starts with the
		// following header: 0A 00 23 00 3C 00 55 00 6E 00.
		Header: 10,
		W:      96,  // ref: 0x4950B8
		H:      160, // h = npixels/w = 15360/96 = 160
	},
	"missiles/rportal1.cl2": {
		// The contents of frame 0 at offset 0x48 in rportal1.cl2 starts with the
		// following header: 0A 00 4B 00 AC 00 0D 01 00 00.
		Header: 10,
		W:      96,  // ref: 0x4950B8
		H:      128, // h = npixels/w = 12288/96 = 128
	},
	"missiles/rportal2.cl2": {
		// The contents of frame 0 at offset 0x48 in rportal2.cl2 starts with the
		// following header: 0A 00 4F 02 6F 05 81 08 00 00.
		Header: 10,
		W:      96,  // ref: 0x4950B8
		H:      128, // h = npixels/w = 12288/96 = 128
	},
	"missiles/scbsexpb.cl2": {
		// The contents of frame 0 at offset 0x24 in scbsexpb.cl2 starts with the
		// following header: 0A 00 2B 00 4E 00 6F 00 00 00.
		Header: 10,
		W:      128, // ref: 0x4950B8
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	"missiles/scbsexpc.cl2": {
		// The contents of frame 0 at offset 0x24 in scbsexpc.cl2 starts with the
		// following header: 0A 00 2B 00 4E 00 6F 00 00 00.
		Header: 10,
		W:      128, // ref: 0x4950B8
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	"missiles/scbsexpd.cl2": {
		// The contents of frame 0 at offset 0x24 in scbsexpd.cl2 starts with the
		// following header: 0A 00 2B 00 4E 00 6F 00 00 00.
		Header: 10,
		W:      128, // ref: 0x4950B8
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	"missiles/scubmisb.cl2": {
		// The contents of frame 0 at offset 0x48 in scubmisb.cl2 starts with the
		// following header: 0A 00 27 01 FA 02 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/scubmisc.cl2": {
		// The contents of frame 0 at offset 0x48 in scubmisc.cl2 starts with the
		// following header: 0A 00 1D 01 BF 02 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/scubmisd.cl2": {
		// The contents of frame 0 at offset 0x48 in scubmisd.cl2 starts with the
		// following header: 0A 00 27 01 FA 02 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/shatter1.cl2": {
		// The contents of frame 0 at offset 0x38 in shatter1.cl2 starts with the
		// following header: 0A 00 2B 00 02 02 D6 03 00 00.
		Header: 10,
		W:      128, // ref: 0x4950B8
		H:      128, // h = npixels/w = 16384/128 = 128
	},
	"missiles/sklball1.cl2": {
		// The contents of frame 0 at offset 0x48 in sklball1.cl2 starts with the
		// following header: 0A 00 23 00 F3 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/sklball2.cl2": {
		// The contents of frame 0 at offset 0x48 in sklball2.cl2 starts with the
		// following header: 0A 00 23 00 04 02 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/sklball3.cl2": {
		// The contents of frame 0 at offset 0x48 in sklball3.cl2 starts with the
		// following header: 0A 00 23 00 37 02 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/sklball4.cl2": {
		// The contents of frame 0 at offset 0x48 in sklball4.cl2 starts with the
		// following header: 0A 00 23 00 EF 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/sklball5.cl2": {
		// The contents of frame 0 at offset 0x48 in sklball5.cl2 starts with the
		// following header: 0A 00 23 00 0A 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/sklball6.cl2": {
		// The contents of frame 0 at offset 0x48 in sklball6.cl2 starts with the
		// following header: 0A 00 23 00 60 02 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/sklball7.cl2": {
		// The contents of frame 0 at offset 0x48 in sklball7.cl2 starts with the
		// following header: 0A 00 23 00 9E 02 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/sklball8.cl2": {
		// The contents of frame 0 at offset 0x48 in sklball8.cl2 starts with the
		// following header: 0A 00 23 00 66 02 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/sklball9.cl2": {
		// The contents of frame 0 at offset 0x28 in sklball9.cl2 starts with the
		// following header: 0A 00 23 00 CC 01 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"missiles/thinlght.cl2": {
		// The contents of frame 0 at offset 0x28 in thinlght.cl2 starts with the
		// following header: 0A 00 66 00 B9 03 00 00 00 00.
		Header: 10,
		W:      96, // ref: 0x4950B8
		H:      96, // h = npixels/w = 9216/96 = 96
	},
	"monsters/acid/acida.cl2": {
		// The contents of the archive header at offset 0 in acida.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in acida.cl2 starts with the
		// following header: 0A 00 3B 02 61 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/acid/acidd.cl2": {
		// The contents of the archive header at offset 0 in acidd.cl2:
		//    20 00 00 00  88 00 00 00  F0 00 00 00  58 01 00 00
		//    C0 01 00 00  28 02 00 00  90 02 00 00  F8 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x340 in acidd.cl2 starts with the
		// following header: 0A 00 17 02 95 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/acid/acidh.cl2": {
		// The contents of the archive header at offset 0 in acidh.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in acidh.cl2 starts with the
		// following header: 0A 00 F7 01 74 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/acid/acidn.cl2": {
		// The contents of the archive header at offset 0 in acidn.cl2:
		//    20 00 00 00  5C 00 00 00  98 00 00 00  D4 00 00 00
		//    10 01 00 00  4C 01 00 00  88 01 00 00  C4 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1E0 in acidn.cl2 starts with the
		// following header: 0A 00 46 02 7D 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/acid/acids.cl2": {
		// The contents of the archive header at offset 0 in acids.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in acids.cl2 starts with the
		// following header: 0A 00 3D 02 6F 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/acid/acidw.cl2": {
		// The contents of the archive header at offset 0 in acidw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in acidw.cl2 starts with the
		// following header: 0A 00 4B 02 81 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/bat/bata.cl2": {
		// The contents of the archive header at offset 0 in bata.cl2:
		//    20 00 00 00  50 00 00 00  80 00 00 00  B0 00 00 00
		//    E0 00 00 00  10 01 00 00  40 01 00 00  70 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x180 in bata.cl2 starts with the
		// following header: 0A 00 5B 00 46 01 0B 03 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/bat/batd.cl2": {
		// The contents of the archive header at offset 0 in batd.cl2:
		//    20 00 00 00  5C 00 00 00  98 00 00 00  D4 00 00 00
		//    10 01 00 00  4C 01 00 00  88 01 00 00  C4 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1E0 in batd.cl2 starts with the
		// following header: 0A 00 2A 00 CE 00 9F 02 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/bat/bath.cl2": {
		// The contents of the archive header at offset 0 in bath.cl2:
		//    20 00 00 00  4C 00 00 00  78 00 00 00  A4 00 00 00
		//    D0 00 00 00  FC 00 00 00  28 01 00 00  54 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x160 in bath.cl2 starts with the
		// following header: 0A 00 3C 00 08 01 AC 02 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/bat/batn.cl2": {
		// The contents of the archive header at offset 0 in batn.cl2:
		//    20 00 00 00  4C 00 00 00  78 00 00 00  A4 00 00 00
		//    D0 00 00 00  FC 00 00 00  28 01 00 00  54 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x160 in batn.cl2 starts with the
		// following header: 0A 00 36 00 D4 00 8A 02 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/bat/batw.cl2": {
		// The contents of the archive header at offset 0 in batw.cl2:
		//    20 00 00 00  5C 00 00 00  98 00 00 00  D4 00 00 00
		//    10 01 00 00  4C 01 00 00  88 01 00 00  C4 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1E0 in batw.cl2 starts with the
		// following header: 0A 00 29 00 BE 00 34 02 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/bigfall/fallga.cl2": {
		// The contents of the archive header at offset 0 in fallga.cl2:
		//    20 00 00 00  54 00 00 00  88 00 00 00  BC 00 00 00
		//    F0 00 00 00  24 01 00 00  58 01 00 00  8C 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1A0 in fallga.cl2 starts with the
		// following header: 0A 00 DB 02 AE 05 CF 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/bigfall/fallgd.cl2": {
		// The contents of the archive header at offset 0 in fallgd.cl2:
		//    20 00 00 00  6C 00 00 00  B8 00 00 00  04 01 00 00
		//    50 01 00 00  9C 01 00 00  E8 01 00 00  34 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x260 in fallgd.cl2 starts with the
		// following header: 0A 00 2F 03 CC 05 ED 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/bigfall/fallgh.cl2": {
		// The contents of the archive header at offset 0 in fallgh.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in fallgh.cl2 starts with the
		// following header: 0A 00 E9 02 BF 05 E0 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/bigfall/fallgn.cl2": {
		// The contents of the archive header at offset 0 in fallgn.cl2:
		//    20 00 00 00  50 00 00 00  80 00 00 00  B0 00 00 00
		//    E0 00 00 00  10 01 00 00  40 01 00 00  70 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x180 in fallgn.cl2 starts with the
		// following header: 0A 00 E1 02 B8 05 D9 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/bigfall/fallgw.cl2": {
		// The contents of the archive header at offset 0 in fallgw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in fallgw.cl2 starts with the
		// following header: 0A 00 F2 02 A8 05 C9 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/black/blacka.cl2": {
		// The contents of the archive header at offset 0 in blacka.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in blacka.cl2 starts with the
		// following header: 0A 00 E6 01 A4 05 82 08 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 20480/unknown = baz
	},
	"monsters/black/blackd.cl2": {
		// The contents of the archive header at offset 0 in blackd.cl2:
		//    20 00 00 00  88 00 00 00  F0 00 00 00  58 01 00 00
		//    C0 01 00 00  28 02 00 00  90 02 00 00  F8 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x340 in blackd.cl2 starts with the
		// following header: 0A 00 16 03 0A 07 42 0B F1 0B.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 25600/unknown = baz
	},
	"monsters/black/blackh.cl2": {
		// The contents of the archive header at offset 0 in blackh.cl2:
		//    20 00 00 00  38 00 00 00  50 00 00 00  68 00 00 00
		//    80 00 00 00  98 00 00 00  B0 00 00 00  C8 00 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0xC0 in blackh.cl2 starts with the
		// following header: 0A 00 E4 01 AA 05 94 08 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 20480/unknown = baz
	},
	"monsters/black/blackn.cl2": {
		// The contents of the archive header at offset 0 in blackn.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in blackn.cl2 starts with the
		// following header: 0A 00 E6 01 A5 05 89 08 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 20480/unknown = baz
	},
	"monsters/black/blackw.cl2": {
		// The contents of the archive header at offset 0 in blackw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in blackw.cl2 starts with the
		// following header: 0A 00 FB 01 11 06 2F 09 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 20480/unknown = baz
	},
	"monsters/darkmage/dmagea.cl2": {
		// The contents of the archive header at offset 0 in dmagea.cl2:
		//    20 00 00 00  7C 00 00 00  D8 00 00 00  34 01 00 00
		//    90 01 00 00  EC 01 00 00  48 02 00 00  A4 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2E0 in dmagea.cl2 starts with the
		// following header: 0A 00 26 01 F9 03 AF 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/darkmage/dmaged.cl2": {
		// The contents of the archive header at offset 0 in dmaged.cl2:
		//    20 00 00 00  84 00 00 00  E8 00 00 00  4C 01 00 00
		//    B0 01 00 00  14 02 00 00  78 02 00 00  DC 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x320 in dmaged.cl2 starts with the
		// following header: 0A 00 0D 01 0B 03 5C 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/darkmage/dmageh.cl2": {
		// The contents of the archive header at offset 0 in dmageh.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in dmageh.cl2 starts with the
		// following header: 0A 00 8A 01 7C 04 F3 06 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/darkmage/dmagen.cl2": {
		// The contents of the archive header at offset 0 in dmagen.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in dmagen.cl2 starts with the
		// following header: 0A 00 3B 01 14 04 C7 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/darkmage/dmages.cl2": {
		// The contents of the archive header at offset 0 in dmages.cl2:
		//    20 00 00 00  70 00 00 00  C0 00 00 00  10 01 00 00
		//    60 01 00 00  B0 01 00 00  00 02 00 00  50 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x280 in dmages.cl2 starts with the
		// following header: 0A 00 3E 01 10 04 A6 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/darkmage/dmagew.cl2": {
		// The contents of the archive header at offset 0 in dmagew.cl2 (after
		// fix, see https://github.com/mewrnd/blizzconv/issues/2):
		//    20 00 00 00  28 00 00 00  30 00 00 00  38 00 00 00
		//    40 00 00 00  48 00 00 00  50 00 00 00  58 00 00 00
		Nimgs: 8,
	},
	"monsters/demskel/demskla.cl2": {
		// The contents of the archive header at offset 0 in demskla.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in demskla.cl2 starts with the
		// following header: 0A 00 EC 01 C1 04 A2 07 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/demskel/demskld.cl2": {
		// The contents of the archive header at offset 0 in demskld.cl2:
		//    20 00 00 00  88 00 00 00  F0 00 00 00  58 01 00 00
		//    C0 01 00 00  28 02 00 00  90 02 00 00  F8 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x340 in demskld.cl2 starts with the
		// following header: 0A 00 53 02 B3 05 6B 08 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/demskel/demsklh.cl2": {
		// The contents of the archive header at offset 0 in demsklh.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in demsklh.cl2 starts with the
		// following header: 0A 00 BC 01 62 04 7B 07 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/demskel/demskln.cl2": {
		// The contents of the archive header at offset 0 in demskln.cl2:
		//    20 00 00 00  50 00 00 00  80 00 00 00  B0 00 00 00
		//    E0 00 00 00  10 01 00 00  40 01 00 00  70 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x180 in demskln.cl2 starts with the
		// following header: 0A 00 18 02 27 05 55 08 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/demskel/demskls.cl2": {
		// The contents of the archive header at offset 0 in demskls.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in demskls.cl2 starts with the
		// following header: 0A 00 EA 01 C0 04 A3 07 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/demskel/demsklw.cl2": {
		// The contents of the archive header at offset 0 in demsklw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in demsklw.cl2 starts with the
		// following header: 0A 00 D3 01 8D 04 83 07 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/diablo/diabloa.cl2": {
		// The contents of the archive header at offset 0 in diabloa.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in diabloa.cl2 starts with the
		// following header: 0A 00 E8 02 72 07 42 0F 5B 15.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 25600/unknown = baz
	},
	"monsters/diablo/diablod.cl2": {
		// The contents of the archive header at offset 0 in diablod.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in diablod.cl2 starts with the
		// following header: 0A 00 D0 04 55 0B E9 13 83 1B.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 25600/unknown = baz
	},
	"monsters/diablo/diabloh.cl2": {
		// The contents of the archive header at offset 0 in diabloh.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in diabloh.cl2 starts with the
		// following header: 0A 00 E2 02 7C 07 38 0F 58 15.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 25600/unknown = baz
	},
	"monsters/diablo/diablon.cl2": {
		// The contents of the archive header at offset 0 in diablon.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in diablon.cl2 starts with the
		// following header: 0A 00 CD 02 5C 07 1C 0F 41 15.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 25600/unknown = baz
	},
	"monsters/diablo/diablos.cl2": {
		// The contents of the archive header at offset 0 in diablos.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in diablos.cl2 starts with the
		// following header: 0A 00 D2 02 61 07 3D 0F 64 15.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 25600/unknown = baz
	},
	"monsters/diablo/diablow.cl2": {
		// The contents of the archive header at offset 0 in diablow.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in diablow.cl2 starts with the
		// following header: 0A 00 C1 02 55 08 E5 10 3C 15.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 25600/unknown = baz
	},
	"monsters/falspear/phalla.cl2": {
		// The contents of the archive header at offset 0 in phalla.cl2:
		//    20 00 00 00  5C 00 00 00  98 00 00 00  D4 00 00 00
		//    10 01 00 00  4C 01 00 00  88 01 00 00  C4 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1E0 in phalla.cl2 starts with the
		// following header: 0A 00 AC 01 30 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/falspear/phalld.cl2": {
		// The contents of the archive header at offset 0 in phalld.cl2:
		//    20 00 00 00  70 00 00 00  C0 00 00 00  10 01 00 00
		//    60 01 00 00  B0 01 00 00  00 02 00 00  50 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x280 in phalld.cl2 starts with the
		// following header: 0A 00 A5 01 2A 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/falspear/phallh.cl2": {
		// The contents of the archive header at offset 0 in phallh.cl2:
		//    20 00 00 00  54 00 00 00  88 00 00 00  BC 00 00 00
		//    F0 00 00 00  24 01 00 00  58 01 00 00  8C 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1A0 in phallh.cl2 starts with the
		// following header: 0A 00 A7 01 26 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/falspear/phalln.cl2": {
		// The contents of the archive header at offset 0 in phalln.cl2:
		//    20 00 00 00  54 00 00 00  88 00 00 00  BC 00 00 00
		//    F0 00 00 00  24 01 00 00  58 01 00 00  8C 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1A0 in phalln.cl2 starts with the
		// following header: 0A 00 A7 01 24 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/falspear/phalls.cl2": {
		// The contents of the archive header at offset 0 in phalls.cl2:
		//    20 00 00 00  5C 00 00 00  98 00 00 00  D4 00 00 00
		//    10 01 00 00  4C 01 00 00  88 01 00 00  C4 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1E0 in phalls.cl2 starts with the
		// following header: 0A 00 A8 01 2A 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/falspear/phallw.cl2": {
		// The contents of the archive header at offset 0 in phallw.cl2:
		//    20 00 00 00  54 00 00 00  88 00 00 00  BC 00 00 00
		//    F0 00 00 00  24 01 00 00  58 01 00 00  8C 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1A0 in phallw.cl2 starts with the
		// following header: 0A 00 A7 01 24 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/falsword/falla.cl2": {
		// The contents of the archive header at offset 0 in falla.cl2:
		//    20 00 00 00  5C 00 00 00  98 00 00 00  D4 00 00 00
		//    10 01 00 00  4C 01 00 00  88 01 00 00  C4 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1E0 in falla.cl2 starts with the
		// following header: 0A 00 AB 02 FE 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/falsword/falld.cl2": {
		// The contents of the archive header at offset 0 in falld.cl2:
		//    20 00 00 00  60 00 00 00  A0 00 00 00  E0 00 00 00
		//    20 01 00 00  60 01 00 00  A0 01 00 00  E0 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x200 in falld.cl2 starts with the
		// following header: 0A 00 9A 02 EB 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/falsword/fallh.cl2": {
		// The contents of the archive header at offset 0 in fallh.cl2:
		//    20 00 00 00  54 00 00 00  88 00 00 00  BC 00 00 00
		//    F0 00 00 00  24 01 00 00  58 01 00 00  8C 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1A0 in fallh.cl2 starts with the
		// following header: 0A 00 AC 02 00 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/falsword/falln.cl2": {
		// The contents of the archive header at offset 0 in falln.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in falln.cl2 starts with the
		// following header: 0A 00 AC 02 FF 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/falsword/falls.cl2": {
		// The contents of the archive header at offset 0 in falls.cl2:
		//    20 00 00 00  64 00 00 00  A8 00 00 00  EC 00 00 00
		//    30 01 00 00  74 01 00 00  B8 01 00 00  FC 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x220 in falls.cl2 starts with the
		// following header: 0A 00 AB 02 FE 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/falsword/fallw.cl2": {
		// The contents of the archive header at offset 0 in fallw.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in fallw.cl2 starts with the
		// following header: 0A 00 96 02 01 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/fatc/fatca.cl2": {
		// The contents of the archive header at offset 0 in fatca.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in fatca.cl2 starts with the
		// following header: 0A 00 69 02 1E 08 44 09 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/fatc/fatcd.cl2": {
		// The contents of the archive header at offset 0 in fatcd.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in fatcd.cl2 starts with the
		// following header: 0A 00 5F 02 DA 07 6A 09 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/fatc/fatch.cl2": {
		// The contents of the archive header at offset 0 in fatch.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in fatch.cl2 starts with the
		// following header: 0A 00 61 02 12 08 6C 09 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/fatc/fatcn.cl2": {
		// The contents of the archive header at offset 0 in fatcn.cl2:
		//    20 00 00 00  50 00 00 00  80 00 00 00  B0 00 00 00
		//    E0 00 00 00  10 01 00 00  40 01 00 00  70 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x180 in fatcn.cl2 starts with the
		// following header: 0A 00 70 02 16 08 39 09 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/fatc/fatcw.cl2": {
		// The contents of the archive header at offset 0 in fatcw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in fatcw.cl2 starts with the
		// following header: 0A 00 38 02 B5 07 AC 08 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/fat/fata.cl2": {
		// The contents of the archive header at offset 0 in fata.cl2:
		//    20 00 00 00  64 00 00 00  A8 00 00 00  EC 00 00 00
		//    30 01 00 00  74 01 00 00  B8 01 00 00  FC 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x220 in fata.cl2 starts with the
		// following header: 0A 00 40 02 37 07 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/fat/fatd.cl2": {
		// The contents of the archive header at offset 0 in fatd.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in fatd.cl2 starts with the
		// following header: 0A 00 39 02 54 07 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/fat/fath.cl2": {
		// The contents of the archive header at offset 0 in fath.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in fath.cl2 starts with the
		// following header: 0A 00 2C 02 43 07 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/fat/fatn.cl2": {
		// The contents of the archive header at offset 0 in fatn.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in fatn.cl2 starts with the
		// following header: 0A 00 42 02 85 07 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/fat/fats.cl2": {
		// The contents of the archive header at offset 0 in fats.cl2:
		//    20 00 00 00  50 00 00 00  80 00 00 00  B0 00 00 00
		//    E0 00 00 00  10 01 00 00  40 01 00 00  70 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x180 in fats.cl2 starts with the
		// following header: 0A 00 2E 02 35 07 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/fat/fatw.cl2": {
		// The contents of the archive header at offset 0 in fatw.cl2:
		//    20 00 00 00  50 00 00 00  80 00 00 00  B0 00 00 00
		//    E0 00 00 00  10 01 00 00  40 01 00 00  70 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x180 in fatw.cl2 starts with the
		// following header: 0A 00 44 02 9E 07 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/fireman/firema.cl2": {
		// The contents of the archive header at offset 0 in firema.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in firema.cl2 starts with the
		// following header: 0A 00 E4 01 72 04 93 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/fireman/firemd.cl2": {
		// The contents of the archive header at offset 0 in firemd.cl2:
		//    20 00 00 00  60 00 00 00  A0 00 00 00  E0 00 00 00
		//    20 01 00 00  60 01 00 00  A0 01 00 00  E0 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x200 in firemd.cl2 starts with the
		// following header: 0A 00 2B 00 94 01 74 04 BA 04.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 20480/unknown = baz
	},
	"monsters/fireman/firemh.cl2": {
		// The contents of the archive header at offset 0 in firemh.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in firemh.cl2 starts with the
		// following header: 0A 00 2B 00 7E 01 3E 04 82 04.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 20480/unknown = baz
	},
	"monsters/fireman/firemn.cl2": {
		// The contents of the archive header at offset 0 in firemn.cl2:
		//    20 00 00 00  60 00 00 00  A0 00 00 00  E0 00 00 00
		//    20 01 00 00  60 01 00 00  A0 01 00 00  E0 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x200 in firemn.cl2 starts with the
		// following header: 0A 00 77 01 50 04 98 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/fireman/firems.cl2": {
		// The contents of the archive header at offset 0 in firems.cl2:
		//    20 00 00 00  84 00 00 00  E8 00 00 00  4C 01 00 00
		//    B0 01 00 00  14 02 00 00  78 02 00 00  DC 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x320 in firems.cl2 starts with the
		// following header: 0A 00 D5 00 F6 00 17 01 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/fireman/firemw.cl2": {
		// The contents of the archive header at offset 0 in firemw.cl2:
		//    20 00 00 00  74 00 00 00  C8 00 00 00  1C 01 00 00
		//    70 01 00 00  C4 01 00 00  18 02 00 00  6C 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2A0 in firemw.cl2 starts with the
		// following header: 0A 00 00 01 21 01 42 01 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/gargoyle/gargoa.cl2": {
		// The contents of the archive header at offset 0 in gargoa.cl2:
		//    20 00 00 00  60 00 00 00  A0 00 00 00  E0 00 00 00
		//    20 01 00 00  60 01 00 00  A0 01 00 00  E0 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x200 in gargoa.cl2 starts with the
		// following header: 0A 00 9B 00 0E 03 89 05 35 07.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 25600/unknown = baz
	},
	"monsters/gargoyle/gargod.cl2": {
		// The contents of the archive header at offset 0 in gargod.cl2:
		//    20 00 00 00  70 00 00 00  C0 00 00 00  10 01 00 00
		//    60 01 00 00  B0 01 00 00  00 02 00 00  50 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x280 in gargod.cl2 starts with the
		// following header: 0A 00 9C 00 03 03 78 05 2B 07.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 25600/unknown = baz
	},
	"monsters/gargoyle/gargoh.cl2": {
		// The contents of the archive header at offset 0 in gargoh.cl2:
		//    20 00 00 00  50 00 00 00  80 00 00 00  B0 00 00 00
		//    E0 00 00 00  10 01 00 00  40 01 00 00  70 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x180 in gargoh.cl2 starts with the
		// following header: 0A 00 B6 00 07 03 FF 05 58 07.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 25600/unknown = baz
	},
	"monsters/gargoyle/gargon.cl2": {
		// The contents of the archive header at offset 0 in gargon.cl2:
		//    20 00 00 00  60 00 00 00  A0 00 00 00  E0 00 00 00
		//    20 01 00 00  60 01 00 00  A0 01 00 00  E0 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x200 in gargon.cl2 starts with the
		// following header: 0A 00 88 00 E5 02 5A 05 01 07.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 25600/unknown = baz
	},
	"monsters/gargoyle/gargos.cl2": {
		// The contents of the archive header at offset 0 in gargos.cl2:
		//    20 00 00 00  60 00 00 00  A0 00 00 00  E0 00 00 00
		//    20 01 00 00  60 01 00 00  A0 01 00 00  E0 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x200 in gargos.cl2 starts with the
		// following header: 0A 00 FF 02 1A 05 64 05 8D 05.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 25600/unknown = baz
	},
	"monsters/gargoyle/gargow.cl2": {
		// The contents of the archive header at offset 0 in gargow.cl2:
		//    20 00 00 00  60 00 00 00  A0 00 00 00  E0 00 00 00
		//    20 01 00 00  60 01 00 00  A0 01 00 00  E0 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x200 in gargow.cl2 starts with the
		// following header: 0A 00 86 00 BD 02 15 05 76 06.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 25600/unknown = baz
	},
	"monsters/goatbow/goatba.cl2": {
		// The contents of the archive header at offset 0 in goatba.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in goatba.cl2 starts with the
		// following header: 0A 00 38 01 61 04 0B 06 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/goatbow/goatbd.cl2": {
		// The contents of the archive header at offset 0 in goatbd.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in goatbd.cl2 starts with the
		// following header: 0A 00 FA 00 DA 03 84 06 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/goatbow/goatbh.cl2": {
		// The contents of the archive header at offset 0 in goatbh.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in goatbh.cl2 starts with the
		// following header: 0A 00 42 01 94 04 C4 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/goatbow/goatbn.cl2": {
		// The contents of the archive header at offset 0 in goatbn.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in goatbn.cl2 starts with the
		// following header: 0A 00 42 01 99 04 C9 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/goatbow/goatbw.cl2": {
		// The contents of the archive header at offset 0 in goatbw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in goatbw.cl2 starts with the
		// following header: 0A 00 26 01 95 04 83 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/goatlord/goatla.cl2": {
		// The contents of the archive header at offset 0 in goatla.cl2:
		//    20 00 00 00  60 00 00 00  A0 00 00 00  E0 00 00 00
		//    20 01 00 00  60 01 00 00  A0 01 00 00  E0 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x200 in goatla.cl2 starts with the
		// following header: 0A 00 38 01 C5 05 73 0A 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 20480/unknown = baz
	},
	"monsters/goatlord/goatld.cl2": {
		// The contents of the archive header at offset 0 in goatld.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in goatld.cl2 starts with the
		// following header: 0A 00 44 01 CE 05 35 0A 9C 0B.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 25600/unknown = baz
	},
	"monsters/goatlord/goatlh.cl2": {
		// The contents of the archive header at offset 0 in goatlh.cl2:
		//    20 00 00 00  4C 00 00 00  78 00 00 00  A4 00 00 00
		//    D0 00 00 00  FC 00 00 00  28 01 00 00  54 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x160 in goatlh.cl2 starts with the
		// following header: 0A 00 37 01 C0 05 6F 0A 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 20480/unknown = baz
	},
	"monsters/goatlord/goatln.cl2": {
		// The contents of the archive header at offset 0 in goatln.cl2:
		//    20 00 00 00  5C 00 00 00  98 00 00 00  D4 00 00 00
		//    10 01 00 00  4C 01 00 00  88 01 00 00  C4 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1E0 in goatln.cl2 starts with the
		// following header: 0A 00 3D 01 CE 05 9B 0A 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 20480/unknown = baz
	},
	"monsters/goatlord/goatlw.cl2": {
		// The contents of the archive header at offset 0 in goatlw.cl2:
		//    20 00 00 00  5C 00 00 00  98 00 00 00  D4 00 00 00
		//    10 01 00 00  4C 01 00 00  88 01 00 00  C4 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1E0 in goatlw.cl2 starts with the
		// following header: 0A 00 48 01 25 05 07 09 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 20480/unknown = baz
	},
	"monsters/goatmace/goata.cl2": {
		// The contents of the archive header at offset 0 in goata.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in goata.cl2 starts with the
		// following header: 0A 00 73 01 2E 05 31 06 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/goatmace/goatd.cl2": {
		// The contents of the archive header at offset 0 in goatd.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in goatd.cl2 starts with the
		// following header: 0A 00 74 01 45 05 B4 06 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/goatmace/goath.cl2": {
		// The contents of the archive header at offset 0 in goath.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in goath.cl2 starts with the
		// following header: 0A 00 73 01 2D 05 2F 06 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/goatmace/goatn.cl2": {
		// The contents of the archive header at offset 0 in goatn.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in goatn.cl2 starts with the
		// following header: 0A 00 7E 01 3D 05 3F 06 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/goatmace/goats.cl2": {
		// The contents of the archive header at offset 0 in goats.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in goats.cl2 starts with the
		// following header: 0A 00 79 01 48 05 4D 06 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/goatmace/goatw.cl2": {
		// The contents of the archive header at offset 0 in goatw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in goatw.cl2 starts with the
		// following header: 0A 00 2F 01 C4 04 AA 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/golem/golema.cl2": {
		// The contents of the archive header at offset 0 in golema.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in golema.cl2 starts with the
		// following header: 0A 00 BE 01 48 05 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"monsters/golem/golemd.cl2": {
		// The contents of frame 0 at offset 0x38 in golemd.cl2 starts with the
		// following header: 0A 00 A7 00 C3 02 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"monsters/golem/golems.cl2": {
		// The contents of frame 0 at offset 0x58 in golems.cl2 starts with the
		// following header: 0A 00 F0 02 09 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"monsters/golem/golemw.cl2": {
		// The contents of the archive header at offset 0 in golemw.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in golemw.cl2 starts with the
		// following header: 0A 00 70 01 5D 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"monsters/mage/magea.cl2": {
		// The contents of the archive header at offset 0 in magea.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in magea.cl2 starts with the
		// following header: 0A 00 B5 01 DE 04 0C 07 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/mage/maged.cl2": {
		// The contents of the archive header at offset 0 in maged.cl2:
		//    20 00 00 00  98 00 00 00  10 01 00 00  88 01 00 00
		//    00 02 00 00  78 02 00 00  F0 02 00 00  68 03 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x3C0 in maged.cl2 starts with the
		// following header: 0A 00 E0 01 3A 05 81 07 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/mage/mageh.cl2": {
		// The contents of the archive header at offset 0 in mageh.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in mageh.cl2 starts with the
		// following header: 0A 00 A9 01 D9 04 0D 07 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/mage/magen.cl2": {
		// The contents of the archive header at offset 0 in magen.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in magen.cl2 starts with the
		// following header: 0A 00 AF 01 DB 04 0C 07 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/mage/mages.cl2": {
		// The contents of the archive header at offset 0 in mages.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in mages.cl2 starts with the
		// following header: 0A 00 59 02 85 05 BA 07 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/mage/magew.cl2": {
		// The contents of the archive header at offset 0 in magew.cl2:
		//    20 00 00 00  2C 00 00 00  38 00 00 00  44 00 00 00
		//    50 00 00 00  5C 00 00 00  68 00 00 00  74 00 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x60 in magew.cl2 starts with the
		// following header: 0A 00 30 00 51 00 72 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/magma/magmaa.cl2": {
		// The contents of the archive header at offset 0 in magmaa.cl2:
		//    20 00 00 00  60 00 00 00  A0 00 00 00  E0 00 00 00
		//    20 01 00 00  60 01 00 00  A0 01 00 00  E0 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x200 in magmaa.cl2 starts with the
		// following header: 0A 00 05 02 00 07 50 08 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/magma/magmad.cl2": {
		// The contents of the archive header at offset 0 in magmad.cl2:
		//    20 00 00 00  70 00 00 00  C0 00 00 00  10 01 00 00
		//    60 01 00 00  B0 01 00 00  00 02 00 00  50 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x280 in magmad.cl2 starts with the
		// following header: 0A 00 05 02 67 06 7C 07 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/magma/magmah.cl2": {
		// The contents of the archive header at offset 0 in magmah.cl2:
		//    20 00 00 00  44 00 00 00  68 00 00 00  8C 00 00 00
		//    B0 00 00 00  D4 00 00 00  F8 00 00 00  1C 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x120 in magmah.cl2 starts with the
		// following header: 0A 00 0B 02 0C 07 63 08 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/magma/magman.cl2": {
		// The contents of the archive header at offset 0 in magman.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in magman.cl2 starts with the
		// following header: 0A 00 14 02 17 07 45 08 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/magma/magmas.cl2": {
		// The contents of the archive header at offset 0 in magmas.cl2:
		//    20 00 00 00  70 00 00 00  C0 00 00 00  10 01 00 00
		//    60 01 00 00  B0 01 00 00  00 02 00 00  50 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x280 in magmas.cl2 starts with the
		// following header: 0A 00 07 02 07 07 54 08 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/magma/magmaw.cl2": {
		// The contents of the archive header at offset 0 in magmaw.cl2:
		//    20 00 00 00  60 00 00 00  A0 00 00 00  E0 00 00 00
		//    20 01 00 00  60 01 00 00  A0 01 00 00  E0 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x200 in magmaw.cl2 starts with the
		// following header: 0A 00 04 02 F8 06 4F 08 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/mega/megaa.cl2": {
		// The contents of the archive header at offset 0 in megaa.cl2:
		//    20 00 00 00  60 00 00 00  A0 00 00 00  E0 00 00 00
		//    20 01 00 00  60 01 00 00  A0 01 00 00  E0 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x200 in megaa.cl2 starts with the
		// following header: 0A 00 ED 02 F9 08 DC 11 3B 13.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 25600/unknown = baz
	},
	"monsters/mega/megad.cl2": {
		// The contents of the archive header at offset 0 in megad.cl2:
		//    20 00 00 00  88 00 00 00  F0 00 00 00  58 01 00 00
		//    C0 01 00 00  28 02 00 00  90 02 00 00  F8 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x340 in megad.cl2 starts with the
		// following header: 0A 00 33 03 D3 09 DF 13 60 15.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 25600/unknown = baz
	},
	"monsters/mega/megah.cl2": {
		// The contents of the archive header at offset 0 in megah.cl2:
		//    20 00 00 00  2C 00 00 00  38 00 00 00  44 00 00 00
		//    50 00 00 00  5C 00 00 00  68 00 00 00  74 00 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x60 in megah.cl2 starts with the
		// following header: 0A 00 5D 02 94 07 F3 0F 1E 11.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 25600/unknown = baz
	},
	"monsters/mega/megan.cl2": {
		// The contents of the archive header at offset 0 in megan.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in megan.cl2 starts with the
		// following header: 0A 00 5E 02 19 08 28 0F 24 11.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 25600/unknown = baz
	},
	"monsters/mega/megas.cl2": {
		// The contents of the archive header at offset 0 in megas.cl2:
		//    20 00 00 00  3C 00 00 00  58 00 00 00  74 00 00 00
		//    90 00 00 00  AC 00 00 00  C8 00 00 00  E4 00 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0xE0 in megas.cl2 starts with the
		// following header: 0A 00 5E 02 2E 08 0D 11 26 13.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 25600/unknown = baz
	},
	"monsters/mega/megaw.cl2": {
		// The contents of the archive header at offset 0 in megaw.cl2:
		//    20 00 00 00  44 00 00 00  68 00 00 00  8C 00 00 00
		//    B0 00 00 00  D4 00 00 00  F8 00 00 00  1C 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x120 in megaw.cl2 starts with the
		// following header: 0A 00 3E 02 42 07 14 0F A2 11.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 25600/unknown = baz
	},
	"monsters/rhino/rhinoa.cl2": {
		// The contents of the archive header at offset 0 in rhinoa.cl2:
		//    20 00 00 00  60 00 00 00  A0 00 00 00  E0 00 00 00
		//    20 01 00 00  60 01 00 00  A0 01 00 00  E0 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x200 in rhinoa.cl2 starts with the
		// following header: 0A 00 E1 01 53 07 E2 0B 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 20480/unknown = baz
	},
	"monsters/rhino/rhinod.cl2": {
		// The contents of the archive header at offset 0 in rhinod.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rhinod.cl2 starts with the
		// following header: 0A 00 A6 01 8B 05 A9 0B 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 20480/unknown = baz
	},
	"monsters/rhino/rhinoh.cl2": {
		// The contents of the archive header at offset 0 in rhinoh.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in rhinoh.cl2 starts with the
		// following header: 0A 00 E4 01 4C 07 EC 0B 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 20480/unknown = baz
	},
	"monsters/rhino/rhinon.cl2": {
		// The contents of the archive header at offset 0 in rhinon.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rhinon.cl2 starts with the
		// following header: 0A 00 E2 01 4F 07 E4 0B 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 20480/unknown = baz
	},
	"monsters/rhino/rhinos.cl2": {
		// The contents of the archive header at offset 0 in rhinos.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in rhinos.cl2 starts with the
		// following header: 0A 00 E0 01 2B 07 DB 0A 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 20480/unknown = baz
	},
	"monsters/rhino/rhinow.cl2": {
		// The contents of the archive header at offset 0 in rhinow.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rhinow.cl2 starts with the
		// following header: 0A 00 E3 01 54 07 58 0C 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 20480/unknown = baz
	},
	"monsters/scav/scava.cl2": {
		// The contents of the archive header at offset 0 in scava.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in scava.cl2 starts with the
		// following header: 0A 00 70 01 2B 02 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/scav/scavd.cl2": {
		// The contents of the archive header at offset 0 in scavd.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in scavd.cl2 starts with the
		// following header: 0A 00 6B 01 25 02 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/scav/scavh.cl2": {
		// The contents of the archive header at offset 0 in scavh.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in scavh.cl2 starts with the
		// following header: 0A 00 6D 01 2A 02 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/scav/scavn.cl2": {
		// The contents of the archive header at offset 0 in scavn.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in scavn.cl2 starts with the
		// following header: 0A 00 70 01 30 02 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/scav/scavs.cl2": {
		// The contents of the archive header at offset 0 in scavs.cl2:
		//    20 00 00 00  54 00 00 00  88 00 00 00  BC 00 00 00
		//    F0 00 00 00  24 01 00 00  58 01 00 00  8C 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1A0 in scavs.cl2 starts with the
		// following header: 0A 00 7F 01 54 02 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/scav/scavw.cl2": {
		// The contents of the archive header at offset 0 in scavw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in scavw.cl2 starts with the
		// following header: 0A 00 4E 01 32 02 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/skelaxe/sklaxa.cl2": {
		// The contents of the archive header at offset 0 in sklaxa.cl2:
		//    20 00 00 00  5C 00 00 00  98 00 00 00  D4 00 00 00
		//    10 01 00 00  4C 01 00 00  88 01 00 00  C4 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1E0 in sklaxa.cl2 starts with the
		// following header: 0A 00 7F 01 CC 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/skelaxe/sklaxd.cl2": {
		// The contents of the archive header at offset 0 in sklaxd.cl2:
		//    20 00 00 00  6C 00 00 00  B8 00 00 00  04 01 00 00
		//    50 01 00 00  9C 01 00 00  E8 01 00 00  34 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x260 in sklaxd.cl2 starts with the
		// following header: 0A 00 B6 00 97 02 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/skelaxe/sklaxh.cl2": {
		// The contents of the archive header at offset 0 in sklaxh.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in sklaxh.cl2 starts with the
		// following header: 0A 00 93 01 E2 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/skelaxe/sklaxn.cl2": {
		// The contents of the archive header at offset 0 in sklaxn.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in sklaxn.cl2 starts with the
		// following header: 0A 00 72 01 B7 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/skelaxe/sklaxs.cl2": {
		// The contents of the archive header at offset 0 in sklaxs.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in sklaxs.cl2 starts with the
		// following header: 0A 00 16 02 37 02 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/skelaxe/sklaxw.cl2": {
		// The contents of the archive header at offset 0 in sklaxw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in sklaxw.cl2 starts with the
		// following header: 0A 00 97 01 E5 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/skelbow/sklbwa.cl2": {
		// The contents of the archive header at offset 0 in sklbwa.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in sklbwa.cl2 starts with the
		// following header: 0A 00 35 01 58 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/skelbow/sklbwd.cl2": {
		// The contents of the archive header at offset 0 in sklbwd.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in sklbwd.cl2 starts with the
		// following header: 0A 00 31 01 4F 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/skelbow/sklbwh.cl2": {
		// The contents of the archive header at offset 0 in sklbwh.cl2:
		//    20 00 00 00  3C 00 00 00  58 00 00 00  74 00 00 00
		//    90 00 00 00  AC 00 00 00  C8 00 00 00  E4 00 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0xE0 in sklbwh.cl2 starts with the
		// following header: 0A 00 4C 01 2A 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/skelbow/sklbwn.cl2": {
		// The contents of the archive header at offset 0 in sklbwn.cl2:
		//    20 00 00 00  4C 00 00 00  78 00 00 00  A4 00 00 00
		//    D0 00 00 00  FC 00 00 00  28 01 00 00  54 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x160 in sklbwn.cl2 starts with the
		// following header: 0A 00 31 01 58 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/skelbow/sklbws.cl2": {
		// The contents of the archive header at offset 0 in sklbws.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in sklbws.cl2 starts with the
		// following header: 0A 00 07 02 28 02 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/skelbow/sklbww.cl2": {
		// The contents of the archive header at offset 0 in sklbww.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in sklbww.cl2 starts with the
		// following header: 0A 00 5B 01 06 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/skelsd/sklsra.cl2": {
		// The contents of the archive header at offset 0 in sklsra.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in sklsra.cl2 starts with the
		// following header: 0A 00 9C 01 DB 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/skelsd/sklsrd.cl2": {
		// The contents of the archive header at offset 0 in sklsrd.cl2:
		//    20 00 00 00  64 00 00 00  A8 00 00 00  EC 00 00 00
		//    30 01 00 00  74 01 00 00  B8 01 00 00  FC 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x220 in sklsrd.cl2 starts with the
		// following header: 0A 00 B8 00 B8 02 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/skelsd/sklsrh.cl2": {
		// The contents of the archive header at offset 0 in sklsrh.cl2:
		//    20 00 00 00  44 00 00 00  68 00 00 00  8C 00 00 00
		//    B0 00 00 00  D4 00 00 00  F8 00 00 00  1C 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x120 in sklsrh.cl2 starts with the
		// following header: 0A 00 79 01 C0 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/skelsd/sklsrn.cl2": {
		// The contents of the archive header at offset 0 in sklsrn.cl2:
		//    20 00 00 00  5C 00 00 00  98 00 00 00  D4 00 00 00
		//    10 01 00 00  4C 01 00 00  88 01 00 00  C4 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1E0 in sklsrn.cl2 starts with the
		// following header: 0A 00 8B 01 D0 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/skelsd/sklsrs.cl2": {
		// The contents of the archive header at offset 0 in sklsrs.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in sklsrs.cl2 starts with the
		// following header: 0A 00 16 02 37 02 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/skelsd/sklsrw.cl2": {
		// The contents of the archive header at offset 0 in sklsrw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in sklsrw.cl2 starts with the
		// following header: 0A 00 57 01 B5 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/sking/skinga.cl2": {
		// The contents of the archive header at offset 0 in skinga.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in skinga.cl2 starts with the
		// following header: 0A 00 A8 01 EB 03 83 07 2C 08.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 25600/unknown = baz
	},
	"monsters/sking/skingd.cl2": {
		// The contents of the archive header at offset 0 in skingd.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in skingd.cl2 starts with the
		// following header: 0A 00 4C 01 08 04 85 08 4E 09.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 25600/unknown = baz
	},
	"monsters/sking/skingh.cl2": {
		// The contents of the archive header at offset 0 in skingh.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in skingh.cl2 starts with the
		// following header: 0A 00 83 01 E0 03 F9 07 F1 08.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 25600/unknown = baz
	},
	"monsters/sking/skingn.cl2": {
		// The contents of the archive header at offset 0 in skingn.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in skingn.cl2 starts with the
		// following header: 0A 00 A4 01 EC 03 8B 07 36 08.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 25600/unknown = baz
	},
	"monsters/sking/skings.cl2": {
		// The contents of the archive header at offset 0 in skings.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in skings.cl2 starts with the
		// following header: 0A 00 A7 01 EC 03 82 07 2C 08.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 25600/unknown = baz
	},
	"monsters/sking/skingw.cl2": {
		// The contents of the archive header at offset 0 in skingw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in skingw.cl2 starts with the
		// following header: 0A 00 50 01 98 03 61 07 F3 07.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 25600/unknown = baz
	},
	"monsters/snake/snakea.cl2": {
		// The contents of the archive header at offset 0 in snakea.cl2:
		//    20 00 00 00  5C 00 00 00  98 00 00 00  D4 00 00 00
		//    10 01 00 00  4C 01 00 00  88 01 00 00  C4 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1E0 in snakea.cl2 starts with the
		// following header: 0A 00 17 03 B1 06 05 07 2E 07.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 25600/unknown = baz
	},
	"monsters/snake/snaked.cl2": {
		// The contents of the archive header at offset 0 in snaked.cl2:
		//    20 00 00 00  70 00 00 00  C0 00 00 00  10 01 00 00
		//    60 01 00 00  B0 01 00 00  00 02 00 00  50 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x280 in snaked.cl2 starts with the
		// following header: 0A 00 31 02 C1 05 ED 08 5A 09.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 25600/unknown = baz
	},
	"monsters/snake/snakeh.cl2": {
		// The contents of the archive header at offset 0 in snakeh.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in snakeh.cl2 starts with the
		// following header: 0A 00 2A 02 DF 05 5B 09 C2 09.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 25600/unknown = baz
	},
	"monsters/snake/snaken.cl2": {
		// The contents of the archive header at offset 0 in snaken.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in snaken.cl2 starts with the
		// following header: 0A 00 24 02 D1 05 45 09 A8 09.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 25600/unknown = baz
	},
	"monsters/snake/snakes.cl2": {
		// The contents of the archive header at offset 0 in snakes.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in snakes.cl2 starts with the
		// following header: 0A 00 28 02 CD 05 4C 09 AC 09.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 25600/unknown = baz
	},
	"monsters/snake/snakew.cl2": {
		// The contents of the archive header at offset 0 in snakew.cl2:
		//    20 00 00 00  54 00 00 00  88 00 00 00  BC 00 00 00
		//    F0 00 00 00  24 01 00 00  58 01 00 00  8C 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1A0 in snakew.cl2 starts with the
		// following header: 0A 00 DA 03 9B 06 E6 06 0F 07.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 25600/unknown = baz
	},
	"monsters/sneak/sneaka.cl2": {
		// The contents of the archive header at offset 0 in sneaka.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in sneaka.cl2 starts with the
		// following header: 0A 00 32 02 FA 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/sneak/sneakd.cl2": {
		// The contents of the archive header at offset 0 in sneakd.cl2:
		//    20 00 00 00  88 00 00 00  F0 00 00 00  58 01 00 00
		//    C0 01 00 00  28 02 00 00  90 02 00 00  F8 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x340 in sneakd.cl2 starts with the
		// following header: 0A 00 F8 01 95 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/sneak/sneakh.cl2": {
		// The contents of the archive header at offset 0 in sneakh.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in sneakh.cl2 starts with the
		// following header: 0A 00 2F 02 00 05 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/sneak/sneakn.cl2": {
		// The contents of the archive header at offset 0 in sneakn.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in sneakn.cl2 starts with the
		// following header: 0A 00 1B 02 ED 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/sneak/sneaks.cl2": {
		// The contents of the archive header at offset 0 in sneaks.cl2:
		//    20 00 00 00  64 00 00 00  A8 00 00 00  EC 00 00 00
		//    30 01 00 00  74 01 00 00  B8 01 00 00  FC 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x220 in sneaks.cl2 starts with the
		// following header: 0A 00 6D 00 BD 00 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/sneak/sneakw.cl2": {
		// The contents of the archive header at offset 0 in sneakw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in sneakw.cl2 starts with the
		// following header: 0A 00 EF 01 4B 05 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/succ/scbsa.cl2": {
		// The contents of the archive header at offset 0 in scbsa.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in scbsa.cl2 starts with the
		// following header: 0A 00 71 01 89 04 FB 06 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/succ/scbsd.cl2": {
		// The contents of the archive header at offset 0 in scbsd.cl2:
		//    20 00 00 00  88 00 00 00  F0 00 00 00  58 01 00 00
		//    C0 01 00 00  28 02 00 00  90 02 00 00  F8 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x340 in scbsd.cl2 starts with the
		// following header: 0A 00 4E 01 74 04 D9 06 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/succ/scbsh.cl2": {
		// The contents of the archive header at offset 0 in scbsh.cl2:
		//    20 00 00 00  44 00 00 00  68 00 00 00  8C 00 00 00
		//    B0 00 00 00  D4 00 00 00  F8 00 00 00  1C 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x120 in scbsh.cl2 starts with the
		// following header: 0A 00 74 01 91 04 08 07 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/succ/scbsn.cl2": {
		// The contents of the archive header at offset 0 in scbsn.cl2:
		//    20 00 00 00  60 00 00 00  A0 00 00 00  E0 00 00 00
		//    20 01 00 00  60 01 00 00  A0 01 00 00  E0 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x200 in scbsn.cl2 starts with the
		// following header: 0A 00 6F 01 85 04 EC 06 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/succ/scbsw.cl2": {
		// The contents of the archive header at offset 0 in scbsw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in scbsw.cl2 starts with the
		// following header: 0A 00 12 01 88 03 D7 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/thin/thina.cl2": {
		// The contents of the archive header at offset 0 in thina.cl2:
		//    20 00 00 00  70 00 00 00  C0 00 00 00  10 01 00 00
		//    60 01 00 00  B0 01 00 00  00 02 00 00  50 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x280 in thina.cl2 starts with the
		// following header: 0A 00 29 02 6A 07 71 0A 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 20480/unknown = baz
	},
	"monsters/thin/thind.cl2": {
		// The contents of the archive header at offset 0 in thind.cl2:
		//    20 00 00 00  6C 00 00 00  B8 00 00 00  04 01 00 00
		//    50 01 00 00  9C 01 00 00  E8 01 00 00  34 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x260 in thind.cl2 starts with the
		// following header: 0A 00 C1 01 2F 05 23 08 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 20480/unknown = baz
	},
	"monsters/thin/thinh.cl2": {
		// The contents of the archive header at offset 0 in thinh.cl2:
		//    20 00 00 00  38 00 00 00  50 00 00 00  68 00 00 00
		//    80 00 00 00  98 00 00 00  B0 00 00 00  C8 00 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0xC0 in thinh.cl2 starts with the
		// following header: 0A 00 DE 01 E1 06 C6 0A 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 20480/unknown = baz
	},
	"monsters/thin/thinn.cl2": {
		// The contents of the archive header at offset 0 in thinn.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in thinn.cl2 starts with the
		// following header: 0A 00 40 02 47 07 E5 0A 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 20480/unknown = baz
	},
	"monsters/thin/thins.cl2": {
		// The contents of the archive header at offset 0 in thins.cl2:
		//    20 00 00 00  60 00 00 00  A0 00 00 00  E0 00 00 00
		//    20 01 00 00  60 01 00 00  A0 01 00 00  E0 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x200 in thins.cl2 starts with the
		// following header: 0A 00 1F 02 55 07 59 0B 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 20480/unknown = baz
	},
	"monsters/thin/thinw.cl2": {
		// The contents of the archive header at offset 0 in thinw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in thinw.cl2 starts with the
		// following header: 0A 00 A9 01 37 07 03 0A 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 20480/unknown = baz
	},
	"monsters/tsneak/tsneaka.cl2": {
		// The contents of the archive header at offset 0 in tsneaka.cl2:
		//    20 00 00 00  64 00 00 00  A8 00 00 00  EC 00 00 00
		//    30 01 00 00  74 01 00 00  B8 01 00 00  FC 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x220 in tsneaka.cl2 starts with the
		// following header: 0A 00 AF 01 24 05 47 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/tsneak/tsneakd.cl2": {
		// The contents of the archive header at offset 0 in tsneakd.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in tsneakd.cl2 starts with the
		// following header: 0A 00 AD 01 25 05 48 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/tsneak/tsneakh.cl2": {
		// The contents of the archive header at offset 0 in tsneakh.cl2:
		//    20 00 00 00  54 00 00 00  88 00 00 00  BC 00 00 00
		//    F0 00 00 00  24 01 00 00  58 01 00 00  8C 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1A0 in tsneakh.cl2 starts with the
		// following header: 0A 00 B0 01 22 05 45 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/tsneak/tsneakn.cl2": {
		// The contents of the archive header at offset 0 in tsneakn.cl2:
		//    20 00 00 00  5C 00 00 00  98 00 00 00  D4 00 00 00
		//    10 01 00 00  4C 01 00 00  88 01 00 00  C4 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1E0 in tsneakn.cl2 starts with the
		// following header: 0A 00 B9 01 2F 05 52 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/tsneak/tsneakw.cl2": {
		// The contents of the archive header at offset 0 in tsneakw.cl2:
		//    20 00 00 00  5C 00 00 00  98 00 00 00  D4 00 00 00
		//    10 01 00 00  4C 01 00 00  88 01 00 00  C4 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1E0 in tsneakw.cl2 starts with the
		// following header: 0A 00 B4 01 A0 04 CB 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"monsters/unrav/unrava.cl2": {
		// The contents of the archive header at offset 0 in unrava.cl2:
		//    20 00 00 00  70 00 00 00  C0 00 00 00  10 01 00 00
		//    60 01 00 00  B0 01 00 00  00 02 00 00  50 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x280 in unrava.cl2 starts with the
		// following header: 0A 00 99 01 62 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"monsters/unrav/unravd.cl2": {
		// The contents of the archive header at offset 0 in unravd.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in unravd.cl2 starts with the
		// following header: 0A 00 7A 00 52 02 CC 03 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/unrav/unravh.cl2": {
		// The contents of the archive header at offset 0 in unravh.cl2:
		//    20 00 00 00  3C 00 00 00  58 00 00 00  74 00 00 00
		//    90 00 00 00  AC 00 00 00  C8 00 00 00  E4 00 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0xE0 in unravh.cl2 starts with the
		// following header: 0A 00 53 00 C6 01 6B 03 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/unrav/unravn.cl2": {
		// The contents of the archive header at offset 0 in unravn.cl2:
		//    20 00 00 00  50 00 00 00  80 00 00 00  B0 00 00 00
		//    E0 00 00 00  10 01 00 00  40 01 00 00  70 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x180 in unravn.cl2 starts with the
		// following header: 0A 00 67 00 FD 01 BD 03 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/unrav/unravs.cl2": {
		// The contents of the archive header at offset 0 in unravs.cl2:
		//    20 00 00 00  5C 00 00 00  98 00 00 00  D4 00 00 00
		//    10 01 00 00  4C 01 00 00  88 01 00 00  C4 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1E0 in unravs.cl2 starts with the
		// following header: 0A 00 63 00 0D 02 B1 03 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/zombie/zombiea.cl2": {
		// The contents of the archive header at offset 0 in zombiea.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in zombiea.cl2 starts with the
		// following header: 0A 00 CC 01 38 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/zombie/zombied.cl2": {
		// The contents of the archive header at offset 0 in zombied.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in zombied.cl2 starts with the
		// following header: 0A 00 AA 01 23 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/zombie/zombieh.cl2": {
		// The contents of the archive header at offset 0 in zombieh.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in zombieh.cl2 starts with the
		// following header: 0A 00 A6 01 E5 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/zombie/zombien.cl2": {
		// The contents of the archive header at offset 0 in zombien.cl2:
		//    20 00 00 00  54 00 00 00  88 00 00 00  BC 00 00 00
		//    F0 00 00 00  24 01 00 00  58 01 00 00  8C 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1A0 in zombien.cl2 starts with the
		// following header: 0A 00 B0 01 16 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/zombie/zombies.cl2": {
		// The contents of the archive header at offset 0 in zombies.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in zombies.cl2 starts with the
		// following header: 0A 00 FA 00 73 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"monsters/zombie/zombiew.cl2": {
		// The contents of the archive header at offset 0 in zombiew.cl2:
		//    20 00 00 00  88 00 00 00  F0 00 00 00  58 01 00 00
		//    C0 01 00 00  28 02 00 00  90 02 00 00  F8 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x340 in zombiew.cl2 starts with the
		// following header: 0A 00 B5 01 0F 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"plrgfx/rogue/rha/rhaas.cl2": {
		// The contents of the archive header at offset 0 in rhaas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rhaas.cl2 starts with the
		// following header: 0A 00 18 01 D9 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rha/rhaat.cl2": {
		// The contents of the archive header at offset 0 in rhaat.cl2:
		//    20 00 00 00  80 00 00 00  E0 00 00 00  40 01 00 00
		//    A0 01 00 00  00 02 00 00  60 02 00 00  C0 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x300 in rhaat.cl2 starts with the
		// following header: 0A 00 17 01 DB 03 71 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/rogue/rha/rhaaw.cl2": {
		// The contents of the archive header at offset 0 in rhaaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rhaaw.cl2 starts with the
		// following header: 0A 00 FD 00 EE 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rha/rhafm.cl2": {
		// The contents of the archive header at offset 0 in rhafm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rhafm.cl2 starts with the
		// following header: 0A 00 1D 01 DB 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rha/rhaht.cl2": {
		// The contents of the archive header at offset 0 in rhaht.cl2:
		//    20 00 00 00  44 00 00 00  68 00 00 00  8C 00 00 00
		//    B0 00 00 00  D4 00 00 00  F8 00 00 00  1C 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x120 in rhaht.cl2 starts with the
		// following header: 0A 00 1C 01 23 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rha/rhalm.cl2": {
		// The contents of the archive header at offset 0 in rhalm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rhalm.cl2 starts with the
		// following header: 0A 00 17 01 D9 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rha/rhaqm.cl2": {
		// The contents of the archive header at offset 0 in rhaqm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rhaqm.cl2 starts with the
		// following header: 0A 00 20 01 E6 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rha/rhast.cl2": {
		// The contents of the archive header at offset 0 in rhast.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in rhast.cl2 starts with the
		// following header: 0A 00 18 01 BF 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rha/rhawl.cl2": {
		// The contents of the archive header at offset 0 in rhawl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rhawl.cl2 starts with the
		// following header: 0A 00 F0 00 6C 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhb/rhbas.cl2": {
		// The contents of the archive header at offset 0 in rhbas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rhbas.cl2 starts with the
		// following header: 0A 00 46 01 9B 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhb/rhbat.cl2": {
		// The contents of the archive header at offset 0 in rhbat.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in rhbat.cl2 starts with the
		// following header: 0A 00 56 01 AB 03 32 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/rogue/rhb/rhbaw.cl2": {
		// The contents of the archive header at offset 0 in rhbaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rhbaw.cl2 starts with the
		// following header: 0A 00 EB 00 3B 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhb/rhbfm.cl2": {
		// The contents of the archive header at offset 0 in rhbfm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rhbfm.cl2 starts with the
		// following header: 0A 00 07 01 46 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhb/rhbht.cl2": {
		// The contents of the archive header at offset 0 in rhbht.cl2:
		//    20 00 00 00  44 00 00 00  68 00 00 00  8C 00 00 00
		//    B0 00 00 00  D4 00 00 00  F8 00 00 00  1C 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x120 in rhbht.cl2 starts with the
		// following header: 0A 00 4B 01 35 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhb/rhblm.cl2": {
		// The contents of the archive header at offset 0 in rhblm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rhblm.cl2 starts with the
		// following header: 0A 00 07 01 4D 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhb/rhbqm.cl2": {
		// The contents of the archive header at offset 0 in rhbqm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rhbqm.cl2 starts with the
		// following header: 0A 00 01 01 42 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhb/rhbst.cl2": {
		// The contents of the archive header at offset 0 in rhbst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in rhbst.cl2 starts with the
		// following header: 0A 00 0F 01 A2 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhb/rhbwl.cl2": {
		// The contents of the archive header at offset 0 in rhbwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rhbwl.cl2 starts with the
		// following header: 0A 00 F0 00 58 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhd/rhdas.cl2": {
		// The contents of the archive header at offset 0 in rhdas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rhdas.cl2 starts with the
		// following header: 0A 00 44 01 18 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhd/rhdat.cl2": {
		// The contents of the archive header at offset 0 in rhdat.cl2:
		//    20 00 00 00  70 00 00 00  C0 00 00 00  10 01 00 00
		//    60 01 00 00  B0 01 00 00  00 02 00 00  50 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x280 in rhdat.cl2 starts with the
		// following header: 0A 00 34 01 2B 04 C0 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/rogue/rhd/rhdaw.cl2": {
		// The contents of the archive header at offset 0 in rhdaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rhdaw.cl2 starts with the
		// following header: 0A 00 F5 00 98 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhd/rhdbl.cl2": {
		// The contents of the archive header at offset 0 in rhdbl.cl2:
		//    20 00 00 00  38 00 00 00  50 00 00 00  68 00 00 00
		//    80 00 00 00  98 00 00 00  B0 00 00 00  C8 00 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0xC0 in rhdbl.cl2 starts with the
		// following header: 0A 00 F0 00 9F 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhd/rhdfm.cl2": {
		// The contents of the archive header at offset 0 in rhdfm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rhdfm.cl2 starts with the
		// following header: 0A 00 2B 01 F2 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhd/rhdht.cl2": {
		// The contents of the archive header at offset 0 in rhdht.cl2:
		//    20 00 00 00  44 00 00 00  68 00 00 00  8C 00 00 00
		//    B0 00 00 00  D4 00 00 00  F8 00 00 00  1C 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x120 in rhdht.cl2 starts with the
		// following header: 0A 00 E3 00 59 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhd/rhdlm.cl2": {
		// The contents of the archive header at offset 0 in rhdlm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rhdlm.cl2 starts with the
		// following header: 0A 00 37 01 F0 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhd/rhdqm.cl2": {
		// The contents of the archive header at offset 0 in rhdqm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rhdqm.cl2 starts with the
		// following header: 0A 00 31 01 E5 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhd/rhdst.cl2": {
		// The contents of the archive header at offset 0 in rhdst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in rhdst.cl2 starts with the
		// following header: 0A 00 08 01 D2 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhd/rhdwl.cl2": {
		// The contents of the archive header at offset 0 in rhdwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rhdwl.cl2 starts with the
		// following header: 0A 00 EF 00 9B 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhh/rhhas.cl2": {
		// The contents of the archive header at offset 0 in rhhas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rhhas.cl2 starts with the
		// following header: 0A 00 11 01 15 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhh/rhhat.cl2": {
		// The contents of the archive header at offset 0 in rhhat.cl2:
		//    20 00 00 00  70 00 00 00  C0 00 00 00  10 01 00 00
		//    60 01 00 00  B0 01 00 00  00 02 00 00  50 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x280 in rhhat.cl2 starts with the
		// following header: 0A 00 2A 01 25 04 BA 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/rogue/rhh/rhhaw.cl2": {
		// The contents of the archive header at offset 0 in rhhaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rhhaw.cl2 starts with the
		// following header: 0A 00 15 01 C1 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhh/rhhbl.cl2": {
		// The contents of the archive header at offset 0 in rhhbl.cl2:
		//    20 00 00 00  38 00 00 00  50 00 00 00  68 00 00 00
		//    80 00 00 00  98 00 00 00  B0 00 00 00  C8 00 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0xC0 in rhhbl.cl2 starts with the
		// following header: 0A 00 E3 00 41 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhh/rhhfm.cl2": {
		// The contents of the archive header at offset 0 in rhhfm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rhhfm.cl2 starts with the
		// following header: 0A 00 37 01 03 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhh/rhhht.cl2": {
		// The contents of the archive header at offset 0 in rhhht.cl2:
		//    20 00 00 00  44 00 00 00  68 00 00 00  8C 00 00 00
		//    B0 00 00 00  D4 00 00 00  F8 00 00 00  1C 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x120 in rhhht.cl2 starts with the
		// following header: 0A 00 35 01 CA 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhh/rhhlm.cl2": {
		// The contents of the archive header at offset 0 in rhhlm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rhhlm.cl2 starts with the
		// following header: 0A 00 3C 01 01 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhh/rhhqm.cl2": {
		// The contents of the archive header at offset 0 in rhhqm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rhhqm.cl2 starts with the
		// following header: 0A 00 3D 01 F7 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhh/rhhst.cl2": {
		// The contents of the archive header at offset 0 in rhhst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in rhhst.cl2 starts with the
		// following header: 0A 00 19 01 F7 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhh/rhhwl.cl2": {
		// The contents of the archive header at offset 0 in rhhwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rhhwl.cl2 starts with the
		// following header: 0A 00 0E 01 D3 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhm/rhmas.cl2": {
		// The contents of the archive header at offset 0 in rhmas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rhmas.cl2 starts with the
		// following header: 0A 00 20 01 CF 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhm/rhmat.cl2": {
		// The contents of the archive header at offset 0 in rhmat.cl2:
		//    20 00 00 00  70 00 00 00  C0 00 00 00  10 01 00 00
		//    60 01 00 00  B0 01 00 00  00 02 00 00  50 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x280 in rhmat.cl2 starts with the
		// following header: 0A 00 35 01 D7 03 6D 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/rogue/rhm/rhmaw.cl2": {
		// The contents of the archive header at offset 0 in rhmaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rhmaw.cl2 starts with the
		// following header: 0A 00 16 01 99 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhm/rhmfm.cl2": {
		// The contents of the archive header at offset 0 in rhmfm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rhmfm.cl2 starts with the
		// following header: 0A 00 37 01 B0 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhm/rhmht.cl2": {
		// The contents of the archive header at offset 0 in rhmht.cl2:
		//    20 00 00 00  44 00 00 00  68 00 00 00  8C 00 00 00
		//    B0 00 00 00  D4 00 00 00  F8 00 00 00  1C 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x120 in rhmht.cl2 starts with the
		// following header: 0A 00 33 01 88 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhm/rhmlm.cl2": {
		// The contents of the archive header at offset 0 in rhmlm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rhmlm.cl2 starts with the
		// following header: 0A 00 47 01 CF 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhm/rhmqm.cl2": {
		// The contents of the archive header at offset 0 in rhmqm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rhmqm.cl2 starts with the
		// following header: 0A 00 39 01 B9 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhm/rhmst.cl2": {
		// The contents of the archive header at offset 0 in rhmst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in rhmst.cl2 starts with the
		// following header: 0A 00 1F 01 C7 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhm/rhmwl.cl2": {
		// The contents of the archive header at offset 0 in rhmwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rhmwl.cl2 starts with the
		// following header: 0A 00 0C 01 93 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhn/rhnas.cl2": {
		// The contents of the archive header at offset 0 in rhnas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rhnas.cl2 starts with the
		// following header: 0A 00 27 01 00 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhn/rhnat.cl2": {
		// The contents of the archive header at offset 0 in rhnat.cl2:
		//    20 00 00 00  70 00 00 00  C0 00 00 00  10 01 00 00
		//    60 01 00 00  B0 01 00 00  00 02 00 00  50 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x280 in rhnat.cl2 starts with the
		// following header: 0A 00 21 01 B3 03 49 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/rogue/rhn/rhnaw.cl2": {
		// The contents of the archive header at offset 0 in rhnaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rhnaw.cl2 starts with the
		// following header: 0A 00 F3 00 74 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhn/rhndt.cl2": {
		// The contents of the archive header at offset 0 in rhndt.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in rhndt.cl2 starts with the
		// following header: 0A 00 2E 01 8D 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"plrgfx/rogue/rhn/rhnfm.cl2": {
		// The contents of the archive header at offset 0 in rhnfm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rhnfm.cl2 starts with the
		// following header: 0A 00 1E 01 81 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhn/rhnht.cl2": {
		// The contents of the archive header at offset 0 in rhnht.cl2:
		//    20 00 00 00  44 00 00 00  68 00 00 00  8C 00 00 00
		//    B0 00 00 00  D4 00 00 00  F8 00 00 00  1C 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x120 in rhnht.cl2 starts with the
		// following header: 0A 00 17 01 E6 02 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhn/rhnlm.cl2": {
		// The contents of the archive header at offset 0 in rhnlm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rhnlm.cl2 starts with the
		// following header: 0A 00 F7 00 24 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhn/rhnqm.cl2": {
		// The contents of the archive header at offset 0 in rhnqm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rhnqm.cl2 starts with the
		// following header: 0A 00 F7 00 27 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhn/rhnst.cl2": {
		// The contents of the archive header at offset 0 in rhnst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in rhnst.cl2 starts with the
		// following header: 0A 00 0E 01 A5 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhn/rhnwl.cl2": {
		// The contents of the archive header at offset 0 in rhnwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rhnwl.cl2 starts with the
		// following header: 0A 00 F1 00 4D 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhs/rhsas.cl2": {
		// The contents of the archive header at offset 0 in rhsas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rhsas.cl2 starts with the
		// following header: 0A 00 3E 01 F0 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhs/rhsat.cl2": {
		// The contents of the archive header at offset 0 in rhsat.cl2:
		//    20 00 00 00  70 00 00 00  C0 00 00 00  10 01 00 00
		//    60 01 00 00  B0 01 00 00  00 02 00 00  50 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x280 in rhsat.cl2 starts with the
		// following header: 0A 00 23 01 A2 03 3A 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/rogue/rhs/rhsaw.cl2": {
		// The contents of the archive header at offset 0 in rhsaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rhsaw.cl2 starts with the
		// following header: 0A 00 F0 00 70 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhs/rhsfm.cl2": {
		// The contents of the archive header at offset 0 in rhsfm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rhsfm.cl2 starts with the
		// following header: 0A 00 1B 01 90 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhs/rhsht.cl2": {
		// The contents of the archive header at offset 0 in rhsht.cl2:
		//    20 00 00 00  44 00 00 00  68 00 00 00  8C 00 00 00
		//    B0 00 00 00  D4 00 00 00  F8 00 00 00  1C 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x120 in rhsht.cl2 starts with the
		// following header: 0A 00 FC 00 1D 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhs/rhslm.cl2": {
		// The contents of the archive header at offset 0 in rhslm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rhslm.cl2 starts with the
		// following header: 0A 00 25 01 A9 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhs/rhsqm.cl2": {
		// The contents of the archive header at offset 0 in rhsqm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rhsqm.cl2 starts with the
		// following header: 0A 00 16 01 7F 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhs/rhsst.cl2": {
		// The contents of the archive header at offset 0 in rhsst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in rhsst.cl2 starts with the
		// following header: 0A 00 18 01 B2 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhs/rhswl.cl2": {
		// The contents of the archive header at offset 0 in rhswl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rhswl.cl2 starts with the
		// following header: 0A 00 F1 00 6C 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rht/rhtas.cl2": {
		// The contents of the archive header at offset 0 in rhtas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rhtas.cl2 starts with the
		// following header: 0A 00 1F 01 3B 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rht/rhtat.cl2": {
		// The contents of the archive header at offset 0 in rhtat.cl2:
		//    20 00 00 00  70 00 00 00  C0 00 00 00  10 01 00 00
		//    60 01 00 00  B0 01 00 00  00 02 00 00  50 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x280 in rhtat.cl2 starts with the
		// following header: 0A 00 1B 01 38 03 BF 03 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/rogue/rht/rhtaw.cl2": {
		// The contents of the archive header at offset 0 in rhtaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rhtaw.cl2 starts with the
		// following header: 0A 00 FF 00 24 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rht/rhtfm.cl2": {
		// The contents of the archive header at offset 0 in rhtfm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rhtfm.cl2 starts with the
		// following header: 0A 00 E5 00 EC 02 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rht/rhtht.cl2": {
		// The contents of the archive header at offset 0 in rhtht.cl2:
		//    20 00 00 00  44 00 00 00  68 00 00 00  8C 00 00 00
		//    B0 00 00 00  D4 00 00 00  F8 00 00 00  1C 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x120 in rhtht.cl2 starts with the
		// following header: 0A 00 FA 00 57 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rht/rhtlm.cl2": {
		// The contents of the archive header at offset 0 in rhtlm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rhtlm.cl2 starts with the
		// following header: 0A 00 F6 00 FF 02 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rht/rhtqm.cl2": {
		// The contents of the archive header at offset 0 in rhtqm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rhtqm.cl2 starts with the
		// following header: 0A 00 02 01 09 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rht/rhtst.cl2": {
		// The contents of the archive header at offset 0 in rhtst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in rhtst.cl2 starts with the
		// following header: 0A 00 47 01 BA 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rht/rhtwl.cl2": {
		// The contents of the archive header at offset 0 in rhtwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rhtwl.cl2 starts with the
		// following header: 0A 00 14 01 A2 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhu/rhuas.cl2": {
		// The contents of the archive header at offset 0 in rhuas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rhuas.cl2 starts with the
		// following header: 0A 00 15 01 AA 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhu/rhuat.cl2": {
		// The contents of the archive header at offset 0 in rhuat.cl2:
		//    20 00 00 00  70 00 00 00  C0 00 00 00  10 01 00 00
		//    60 01 00 00  B0 01 00 00  00 02 00 00  50 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x280 in rhuat.cl2 starts with the
		// following header: 0A 00 22 01 B7 03 4C 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/rogue/rhu/rhuaw.cl2": {
		// The contents of the archive header at offset 0 in rhuaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rhuaw.cl2 starts with the
		// following header: 0A 00 F3 00 8F 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhu/rhubl.cl2": {
		// The contents of the archive header at offset 0 in rhubl.cl2:
		//    20 00 00 00  38 00 00 00  50 00 00 00  68 00 00 00
		//    80 00 00 00  98 00 00 00  B0 00 00 00  C8 00 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0xC0 in rhubl.cl2 starts with the
		// following header: 0A 00 DA 00 2A 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhu/rhufm.cl2": {
		// The contents of the archive header at offset 0 in rhufm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rhufm.cl2 starts with the
		// following header: 0A 00 F4 00 E3 02 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhu/rhuht.cl2": {
		// The contents of the archive header at offset 0 in rhuht.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rhuht.cl2 starts with the
		// following header: 0A 00 ED 00 1C 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhu/rhulm.cl2": {
		// The contents of the archive header at offset 0 in rhulm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rhulm.cl2 starts with the
		// following header: 0A 00 02 01 E9 02 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhu/rhuqm.cl2": {
		// The contents of the archive header at offset 0 in rhuqm.cl2:
		//    20 00 00 00  6C 00 00 00  B8 00 00 00  04 01 00 00
		//    50 01 00 00  9C 01 00 00  E8 01 00 00  34 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x260 in rhuqm.cl2 starts with the
		// following header: 0A 00 F3 00 DC 02 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhu/rhust.cl2": {
		// The contents of the archive header at offset 0 in rhust.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in rhust.cl2 starts with the
		// following header: 0A 00 13 01 EA 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rhu/rhuwl.cl2": {
		// The contents of the archive header at offset 0 in rhuwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rhuwl.cl2 starts with the
		// following header: 0A 00 F2 00 8F 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rla/rlaas.cl2": {
		// The contents of the archive header at offset 0 in rlaas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rlaas.cl2 starts with the
		// following header: 0A 00 4A 01 07 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rla/rlaat.cl2": {
		// The contents of the archive header at offset 0 in rlaat.cl2:
		//    20 00 00 00  80 00 00 00  E0 00 00 00  40 01 00 00
		//    A0 01 00 00  00 02 00 00  60 02 00 00  C0 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x300 in rlaat.cl2 starts with the
		// following header: 0A 00 4A 01 09 04 6C 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/rogue/rla/rlaaw.cl2": {
		// The contents of the archive header at offset 0 in rlaaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rlaaw.cl2 starts with the
		// following header: 0A 00 10 01 F9 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rla/rlafm.cl2": {
		// The contents of the archive header at offset 0 in rlafm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rlafm.cl2 starts with the
		// following header: 0A 00 42 01 F5 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rla/rlaht.cl2": {
		// The contents of the archive header at offset 0 in rlaht.cl2:
		//    20 00 00 00  44 00 00 00  68 00 00 00  8C 00 00 00
		//    B0 00 00 00  D4 00 00 00  F8 00 00 00  1C 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x120 in rlaht.cl2 starts with the
		// following header: 0A 00 32 01 41 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rla/rlalm.cl2": {
		// The contents of the archive header at offset 0 in rlalm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rlalm.cl2 starts with the
		// following header: 0A 00 4B 01 06 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rla/rlaqm.cl2": {
		// The contents of the archive header at offset 0 in rlaqm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rlaqm.cl2 starts with the
		// following header: 0A 00 52 01 F1 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rla/rlast.cl2": {
		// The contents of the archive header at offset 0 in rlast.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in rlast.cl2 starts with the
		// following header: 0A 00 3B 01 E8 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rla/rlawl.cl2": {
		// The contents of the archive header at offset 0 in rlawl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rlawl.cl2 starts with the
		// following header: 0A 00 09 01 7D 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rlb/rlbas.cl2": {
		// The contents of the archive header at offset 0 in rlbas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rlbas.cl2 starts with the
		// following header: 0A 00 34 01 8A 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rlb/rlbat.cl2": {
		// The contents of the archive header at offset 0 in rlbat.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in rlbat.cl2 starts with the
		// following header: 0A 00 5B 01 AA 03 0A 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/rogue/rlb/rlbaw.cl2": {
		// The contents of the archive header at offset 0 in rlbaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rlbaw.cl2 starts with the
		// following header: 0A 00 05 01 49 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rlb/rlbfm.cl2": {
		// The contents of the archive header at offset 0 in rlbfm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rlbfm.cl2 starts with the
		// following header: 0A 00 36 01 8B 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rlb/rlbht.cl2": {
		// The contents of the archive header at offset 0 in rlbht.cl2:
		//    20 00 00 00  44 00 00 00  68 00 00 00  8C 00 00 00
		//    B0 00 00 00  D4 00 00 00  F8 00 00 00  1C 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x120 in rlbht.cl2 starts with the
		// following header: 0A 00 62 01 51 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rlb/rlblm.cl2": {
		// The contents of the archive header at offset 0 in rlblm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rlblm.cl2 starts with the
		// following header: 0A 00 33 01 89 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rlb/rlbqm.cl2": {
		// The contents of the archive header at offset 0 in rlbqm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rlbqm.cl2 starts with the
		// following header: 0A 00 16 01 68 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rlb/rlbst.cl2": {
		// The contents of the archive header at offset 0 in rlbst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in rlbst.cl2 starts with the
		// following header: 0A 00 2F 01 D2 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rlb/rlbwl.cl2": {
		// The contents of the archive header at offset 0 in rlbwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rlbwl.cl2 starts with the
		// following header: 0A 00 09 01 56 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rld/rldas.cl2": {
		// The contents of the archive header at offset 0 in rldas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rldas.cl2 starts with the
		// following header: 0A 00 45 01 B7 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rld/rldat.cl2": {
		// The contents of the archive header at offset 0 in rldat.cl2:
		//    20 00 00 00  70 00 00 00  C0 00 00 00  10 01 00 00
		//    60 01 00 00  B0 01 00 00  00 02 00 00  50 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x280 in rldat.cl2 starts with the
		// following header: 0A 00 3E 01 B9 03 1C 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/rogue/rld/rldaw.cl2": {
		// The contents of the archive header at offset 0 in rldaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rldaw.cl2 starts with the
		// following header: 0A 00 FE 00 83 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rld/rldbl.cl2": {
		// The contents of the archive header at offset 0 in rldbl.cl2:
		//    20 00 00 00  38 00 00 00  50 00 00 00  68 00 00 00
		//    80 00 00 00  98 00 00 00  B0 00 00 00  C8 00 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0xC0 in rldbl.cl2 starts with the
		// following header: 0A 00 43 01 97 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rld/rldfm.cl2": {
		// The contents of the archive header at offset 0 in rldfm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rldfm.cl2 starts with the
		// following header: 0A 00 3D 01 B6 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rld/rldht.cl2": {
		// The contents of the archive header at offset 0 in rldht.cl2:
		//    20 00 00 00  44 00 00 00  68 00 00 00  8C 00 00 00
		//    B0 00 00 00  D4 00 00 00  F8 00 00 00  1C 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x120 in rldht.cl2 starts with the
		// following header: 0A 00 39 01 55 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rld/rldlm.cl2": {
		// The contents of the archive header at offset 0 in rldlm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rldlm.cl2 starts with the
		// following header: 0A 00 41 01 B1 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rld/rldqm.cl2": {
		// The contents of the archive header at offset 0 in rldqm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rldqm.cl2 starts with the
		// following header: 0A 00 2A 01 83 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rld/rldst.cl2": {
		// The contents of the archive header at offset 0 in rldst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in rldst.cl2 starts with the
		// following header: 0A 00 36 01 DA 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rld/rldwl.cl2": {
		// The contents of the archive header at offset 0 in rldwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rldwl.cl2 starts with the
		// following header: 0A 00 05 01 6A 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rlh/rlhas.cl2": {
		// The contents of the archive header at offset 0 in rlhas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rlhas.cl2 starts with the
		// following header: 0A 00 64 01 E2 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rlh/rlhat.cl2": {
		// The contents of the archive header at offset 0 in rlhat.cl2:
		//    20 00 00 00  70 00 00 00  C0 00 00 00  10 01 00 00
		//    60 01 00 00  B0 01 00 00  00 02 00 00  50 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x280 in rlhat.cl2 starts with the
		// following header: 0A 00 61 01 F2 03 57 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/rogue/rlh/rlhaw.cl2": {
		// The contents of the archive header at offset 0 in rlhaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rlhaw.cl2 starts with the
		// following header: 0A 00 25 01 B1 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rlh/rlhbl.cl2": {
		// The contents of the archive header at offset 0 in rlhbl.cl2:
		//    20 00 00 00  38 00 00 00  50 00 00 00  68 00 00 00
		//    80 00 00 00  98 00 00 00  B0 00 00 00  C8 00 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0xC0 in rlhbl.cl2 starts with the
		// following header: 0A 00 4B 01 BA 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rlh/rlhfm.cl2": {
		// The contents of the archive header at offset 0 in rlhfm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rlhfm.cl2 starts with the
		// following header: 0A 00 46 01 B0 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rlh/rlhht.cl2": {
		// The contents of the archive header at offset 0 in rlhht.cl2:
		//    20 00 00 00  44 00 00 00  68 00 00 00  8C 00 00 00
		//    B0 00 00 00  D4 00 00 00  F8 00 00 00  1C 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x120 in rlhht.cl2 starts with the
		// following header: 0A 00 46 01 59 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rlh/rlhlm.cl2": {
		// The contents of the archive header at offset 0 in rlhlm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rlhlm.cl2 starts with the
		// following header: 0A 00 47 01 AF 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rlh/rlhqm.cl2": {
		// The contents of the archive header at offset 0 in rlhqm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rlhqm.cl2 starts with the
		// following header: 0A 00 49 01 AB 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rlh/rlhst.cl2": {
		// The contents of the archive header at offset 0 in rlhst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in rlhst.cl2 starts with the
		// following header: 0A 00 3D 01 DF 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rlh/rlhwl.cl2": {
		// The contents of the archive header at offset 0 in rlhwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rlhwl.cl2 starts with the
		// following header: 0A 00 27 01 B4 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rlm/rlmas.cl2": {
		// The contents of the archive header at offset 0 in rlmas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rlmas.cl2 starts with the
		// following header: 0A 00 45 01 EE 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rlm/rlmat.cl2": {
		// The contents of the archive header at offset 0 in rlmat.cl2:
		//    20 00 00 00  70 00 00 00  C0 00 00 00  10 01 00 00
		//    60 01 00 00  B0 01 00 00  00 02 00 00  50 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x280 in rlmat.cl2 starts with the
		// following header: 0A 00 51 01 F9 03 5D 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/rogue/rlm/rlmaw.cl2": {
		// The contents of the archive header at offset 0 in rlmaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rlmaw.cl2 starts with the
		// following header: 0A 00 21 01 AA 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rlm/rlmfm.cl2": {
		// The contents of the archive header at offset 0 in rlmfm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rlmfm.cl2 starts with the
		// following header: 0A 00 5D 01 EC 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rlm/rlmht.cl2": {
		// The contents of the archive header at offset 0 in rlmht.cl2:
		//    20 00 00 00  44 00 00 00  68 00 00 00  8C 00 00 00
		//    B0 00 00 00  D4 00 00 00  F8 00 00 00  1C 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x120 in rlmht.cl2 starts with the
		// following header: 0A 00 52 01 BB 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rlm/rlmlm.cl2": {
		// The contents of the archive header at offset 0 in rlmlm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rlmlm.cl2 starts with the
		// following header: 0A 00 60 01 F8 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rlm/rlmqm.cl2": {
		// The contents of the archive header at offset 0 in rlmqm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rlmqm.cl2 starts with the
		// following header: 0A 00 43 01 C8 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rlm/rlmst.cl2": {
		// The contents of the archive header at offset 0 in rlmst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in rlmst.cl2 starts with the
		// following header: 0A 00 43 01 ED 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rlm/rlmwl.cl2": {
		// The contents of the archive header at offset 0 in rlmwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rlmwl.cl2 starts with the
		// following header: 0A 00 29 01 A3 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rln/rlnas.cl2": {
		// The contents of the archive header at offset 0 in rlnas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rlnas.cl2 starts with the
		// following header: 0A 00 3B 01 1B 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rln/rlnat.cl2": {
		// The contents of the archive header at offset 0 in rlnat.cl2:
		//    20 00 00 00  70 00 00 00  C0 00 00 00  10 01 00 00
		//    60 01 00 00  B0 01 00 00  00 02 00 00  50 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x280 in rlnat.cl2 starts with the
		// following header: 0A 00 48 01 D3 03 35 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/rogue/rln/rlnaw.cl2": {
		// The contents of the archive header at offset 0 in rlnaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rlnaw.cl2 starts with the
		// following header: 0A 00 F5 00 6B 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rln/rlndt.cl2": {
		// The contents of the archive header at offset 0 in rlndt.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in rlndt.cl2 starts with the
		// following header: 0A 00 57 01 CA 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"plrgfx/rogue/rln/rlnfm.cl2": {
		// The contents of the archive header at offset 0 in rlnfm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rlnfm.cl2 starts with the
		// following header: 0A 00 37 01 16 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rln/rlnht.cl2": {
		// The contents of the archive header at offset 0 in rlnht.cl2:
		//    20 00 00 00  44 00 00 00  68 00 00 00  8C 00 00 00
		//    B0 00 00 00  D4 00 00 00  F8 00 00 00  1C 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x120 in rlnht.cl2 starts with the
		// following header: 0A 00 42 01 0C 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rln/rlnlm.cl2": {
		// The contents of the archive header at offset 0 in rlnlm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rlnlm.cl2 starts with the
		// following header: 0A 00 3B 01 17 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rln/rlnqm.cl2": {
		// The contents of the archive header at offset 0 in rlnqm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rlnqm.cl2 starts with the
		// following header: 0A 00 27 01 F8 02 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rln/rlnst.cl2": {
		// The contents of the archive header at offset 0 in rlnst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in rlnst.cl2 starts with the
		// following header: 0A 00 30 01 D3 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rln/rlnwl.cl2": {
		// The contents of the archive header at offset 0 in rlnwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rlnwl.cl2 starts with the
		// following header: 0A 00 F7 00 2D 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rls/rlsas.cl2": {
		// The contents of the archive header at offset 0 in rlsas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rlsas.cl2 starts with the
		// following header: 0A 00 3F 01 D5 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rls/rlsat.cl2": {
		// The contents of the archive header at offset 0 in rlsat.cl2:
		//    20 00 00 00  70 00 00 00  C0 00 00 00  10 01 00 00
		//    60 01 00 00  B0 01 00 00  00 02 00 00  50 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x280 in rlsat.cl2 starts with the
		// following header: 0A 00 3D 01 D0 03 31 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/rogue/rls/rlsaw.cl2": {
		// The contents of the archive header at offset 0 in rlsaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rlsaw.cl2 starts with the
		// following header: 0A 00 F9 00 79 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rls/rlsfm.cl2": {
		// The contents of the archive header at offset 0 in rlsfm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rlsfm.cl2 starts with the
		// following header: 0A 00 40 01 DA 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rls/rlsht.cl2": {
		// The contents of the archive header at offset 0 in rlsht.cl2:
		//    20 00 00 00  44 00 00 00  68 00 00 00  8C 00 00 00
		//    B0 00 00 00  D4 00 00 00  F8 00 00 00  1C 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x120 in rlsht.cl2 starts with the
		// following header: 0A 00 54 01 AA 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rls/rlslm.cl2": {
		// The contents of the archive header at offset 0 in rlslm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rlslm.cl2 starts with the
		// following header: 0A 00 01 01 C6 02 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rls/rlsqm.cl2": {
		// The contents of the archive header at offset 0 in rlsqm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rlsqm.cl2 starts with the
		// following header: 0A 00 54 01 CA 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rls/rlsst.cl2": {
		// The contents of the archive header at offset 0 in rlsst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in rlsst.cl2 starts with the
		// following header: 0A 00 36 01 D8 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rls/rlswl.cl2": {
		// The contents of the archive header at offset 0 in rlswl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rlswl.cl2 starts with the
		// following header: 0A 00 FB 00 53 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rlt/rltas.cl2": {
		// The contents of the archive header at offset 0 in rltas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rltas.cl2 starts with the
		// following header: 0A 00 2A 01 3E 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rlt/rltat.cl2": {
		// The contents of the archive header at offset 0 in rltat.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rltat.cl2 starts with the
		// following header: 0A 00 2F 01 44 03 A6 03 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/rogue/rlt/rltaw.cl2": {
		// The contents of the archive header at offset 0 in rltaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rltaw.cl2 starts with the
		// following header: 0A 00 12 01 33 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rlt/rltfm.cl2": {
		// The contents of the archive header at offset 0 in rltfm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rltfm.cl2 starts with the
		// following header: 0A 00 28 01 40 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rlt/rltht.cl2": {
		// The contents of the archive header at offset 0 in rltht.cl2:
		//    20 00 00 00  44 00 00 00  68 00 00 00  8C 00 00 00
		//    B0 00 00 00  D4 00 00 00  F8 00 00 00  1C 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x120 in rltht.cl2 starts with the
		// following header: 0A 00 14 01 79 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rlt/rltlm.cl2": {
		// The contents of the archive header at offset 0 in rltlm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rltlm.cl2 starts with the
		// following header: 0A 00 26 01 39 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rlt/rltqm.cl2": {
		// The contents of the archive header at offset 0 in rltqm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rltqm.cl2 starts with the
		// following header: 0A 00 1A 01 21 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rlt/rltst.cl2": {
		// The contents of the archive header at offset 0 in rltst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in rltst.cl2 starts with the
		// following header: 0A 00 4E 01 C5 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rlt/rltwl.cl2": {
		// The contents of the archive header at offset 0 in rltwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rltwl.cl2 starts with the
		// following header: 0A 00 2C 01 99 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rlu/rluas.cl2": {
		// The contents of the archive header at offset 0 in rluas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rluas.cl2 starts with the
		// following header: 0A 00 3A 01 78 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rlu/rluat.cl2": {
		// The contents of the archive header at offset 0 in rluat.cl2:
		//    20 00 00 00  70 00 00 00  C0 00 00 00  10 01 00 00
		//    60 01 00 00  B0 01 00 00  00 02 00 00  50 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x280 in rluat.cl2 starts with the
		// following header: 0A 00 3B 01 79 03 DD 03 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/rogue/rlu/rluaw.cl2": {
		// The contents of the archive header at offset 0 in rluaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rluaw.cl2 starts with the
		// following header: 0A 00 FB 00 77 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rlu/rlubl.cl2": {
		// The contents of the archive header at offset 0 in rlubl.cl2:
		//    20 00 00 00  38 00 00 00  50 00 00 00  68 00 00 00
		//    80 00 00 00  98 00 00 00  B0 00 00 00  C8 00 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0xC0 in rlubl.cl2 starts with the
		// following header: 0A 00 37 01 6C 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rlu/rlufm.cl2": {
		// The contents of the archive header at offset 0 in rlufm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rlufm.cl2 starts with the
		// following header: 0A 00 32 01 66 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rlu/rluht.cl2": {
		// The contents of the archive header at offset 0 in rluht.cl2:
		//    20 00 00 00  44 00 00 00  68 00 00 00  8C 00 00 00
		//    B0 00 00 00  D4 00 00 00  F8 00 00 00  1C 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x120 in rluht.cl2 starts with the
		// following header: 0A 00 4F 01 6F 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rlu/rlulm.cl2": {
		// The contents of the archive header at offset 0 in rlulm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rlulm.cl2 starts with the
		// following header: 0A 00 3E 01 7E 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rlu/rluqm.cl2": {
		// The contents of the archive header at offset 0 in rluqm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rluqm.cl2 starts with the
		// following header: 0A 00 27 01 59 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rlu/rlust.cl2": {
		// The contents of the archive header at offset 0 in rlust.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in rlust.cl2 starts with the
		// following header: 0A 00 37 01 DD 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rlu/rluwl.cl2": {
		// The contents of the archive header at offset 0 in rluwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rluwl.cl2 starts with the
		// following header: 0A 00 06 01 6F 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rma/rmaas.cl2": {
		// The contents of the archive header at offset 0 in rmaas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rmaas.cl2 starts with the
		// following header: 0A 00 48 01 11 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rma/rmaat.cl2": {
		// The contents of the archive header at offset 0 in rmaat.cl2:
		//    20 00 00 00  80 00 00 00  E0 00 00 00  40 01 00 00
		//    A0 01 00 00  00 02 00 00  60 02 00 00  C0 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x300 in rmaat.cl2 starts with the
		// following header: 0A 00 4C 01 16 04 7E 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/rogue/rma/rmaaw.cl2": {
		// The contents of the archive header at offset 0 in rmaaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rmaaw.cl2 starts with the
		// following header: 0A 00 0D 01 09 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rma/rmafm.cl2": {
		// The contents of the archive header at offset 0 in rmafm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rmafm.cl2 starts with the
		// following header: 0A 00 46 01 18 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rma/rmaht.cl2": {
		// The contents of the archive header at offset 0 in rmaht.cl2:
		//    20 00 00 00  44 00 00 00  68 00 00 00  8C 00 00 00
		//    B0 00 00 00  D4 00 00 00  F8 00 00 00  1C 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x120 in rmaht.cl2 starts with the
		// following header: 0A 00 31 01 30 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rma/rmalm.cl2": {
		// The contents of the archive header at offset 0 in rmalm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rmalm.cl2 starts with the
		// following header: 0A 00 4B 01 20 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rma/rmaqm.cl2": {
		// The contents of the archive header at offset 0 in rmaqm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rmaqm.cl2 starts with the
		// following header: 0A 00 45 01 11 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rma/rmast.cl2": {
		// The contents of the archive header at offset 0 in rmast.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in rmast.cl2 starts with the
		// following header: 0A 00 38 01 F7 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rma/rmawl.cl2": {
		// The contents of the archive header at offset 0 in rmawl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rmawl.cl2 starts with the
		// following header: 0A 00 FF 00 87 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmb/rmbas.cl2": {
		// The contents of the archive header at offset 0 in rmbas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rmbas.cl2 starts with the
		// following header: 0A 00 4F 01 9C 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmb/rmbat.cl2": {
		// The contents of the archive header at offset 0 in rmbat.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in rmbat.cl2 starts with the
		// following header: 0A 00 6F 01 B5 03 15 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/rogue/rmb/rmbaw.cl2": {
		// The contents of the archive header at offset 0 in rmbaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rmbaw.cl2 starts with the
		// following header: 0A 00 0A 01 52 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmb/rmbfm.cl2": {
		// The contents of the archive header at offset 0 in rmbfm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rmbfm.cl2 starts with the
		// following header: 0A 00 3A 01 90 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmb/rmbht.cl2": {
		// The contents of the archive header at offset 0 in rmbht.cl2:
		//    20 00 00 00  44 00 00 00  68 00 00 00  8C 00 00 00
		//    B0 00 00 00  D4 00 00 00  F8 00 00 00  1C 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x120 in rmbht.cl2 starts with the
		// following header: 0A 00 68 01 59 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmb/rmblm.cl2": {
		// The contents of the archive header at offset 0 in rmblm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rmblm.cl2 starts with the
		// following header: 0A 00 3F 01 8E 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmb/rmbqm.cl2": {
		// The contents of the archive header at offset 0 in rmbqm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rmbqm.cl2 starts with the
		// following header: 0A 00 3A 01 80 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmb/rmbst.cl2": {
		// The contents of the archive header at offset 0 in rmbst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in rmbst.cl2 starts with the
		// following header: 0A 00 26 01 D2 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmb/rmbwl.cl2": {
		// The contents of the archive header at offset 0 in rmbwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rmbwl.cl2 starts with the
		// following header: 0A 00 FF 00 76 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmd/rmdas.cl2": {
		// The contents of the archive header at offset 0 in rmdas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rmdas.cl2 starts with the
		// following header: 0A 00 67 01 5F 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmd/rmdat.cl2": {
		// The contents of the archive header at offset 0 in rmdat.cl2:
		//    20 00 00 00  70 00 00 00  C0 00 00 00  10 01 00 00
		//    60 01 00 00  B0 01 00 00  00 02 00 00  50 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x280 in rmdat.cl2 starts with the
		// following header: 0A 00 56 01 2E 04 92 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/rogue/rmd/rmdaw.cl2": {
		// The contents of the archive header at offset 0 in rmdaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rmdaw.cl2 starts with the
		// following header: 0A 00 F8 00 9C 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmd/rmdbl.cl2": {
		// The contents of the archive header at offset 0 in rmdbl.cl2:
		//    20 00 00 00  38 00 00 00  50 00 00 00  68 00 00 00
		//    80 00 00 00  98 00 00 00  B0 00 00 00  C8 00 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0xC0 in rmdbl.cl2 starts with the
		// following header: 0A 00 09 01 BD 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmd/rmdfm.cl2": {
		// The contents of the archive header at offset 0 in rmdfm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rmdfm.cl2 starts with the
		// following header: 0A 00 60 01 2C 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmd/rmdht.cl2": {
		// The contents of the archive header at offset 0 in rmdht.cl2:
		//    20 00 00 00  44 00 00 00  68 00 00 00  8C 00 00 00
		//    B0 00 00 00  D4 00 00 00  F8 00 00 00  1C 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x120 in rmdht.cl2 starts with the
		// following header: 0A 00 FF 00 83 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmd/rmdlm.cl2": {
		// The contents of the archive header at offset 0 in rmdlm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rmdlm.cl2 starts with the
		// following header: 0A 00 5C 01 2E 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmd/rmdqm.cl2": {
		// The contents of the archive header at offset 0 in rmdqm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rmdqm.cl2 starts with the
		// following header: 0A 00 5B 01 1D 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmd/rmdst.cl2": {
		// The contents of the archive header at offset 0 in rmdst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in rmdst.cl2 starts with the
		// following header: 0A 00 2F 01 05 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmd/rmdwl.cl2": {
		// The contents of the archive header at offset 0 in rmdwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rmdwl.cl2 starts with the
		// following header: 0A 00 03 01 9B 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmh/rmhas.cl2": {
		// The contents of the archive header at offset 0 in rmhas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rmhas.cl2 starts with the
		// following header: 0A 00 3B 01 17 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmh/rmhat.cl2": {
		// The contents of the archive header at offset 0 in rmhat.cl2:
		//    20 00 00 00  70 00 00 00  C0 00 00 00  10 01 00 00
		//    60 01 00 00  B0 01 00 00  00 02 00 00  50 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x280 in rmhat.cl2 starts with the
		// following header: 0A 00 51 01 27 04 8E 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/rogue/rmh/rmhaw.cl2": {
		// The contents of the archive header at offset 0 in rmhaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rmhaw.cl2 starts with the
		// following header: 0A 00 21 01 D6 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmh/rmhbl.cl2": {
		// The contents of the archive header at offset 0 in rmhbl.cl2:
		//    20 00 00 00  38 00 00 00  50 00 00 00  68 00 00 00
		//    80 00 00 00  98 00 00 00  B0 00 00 00  C8 00 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0xC0 in rmhbl.cl2 starts with the
		// following header: 0A 00 FA 00 38 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmh/rmhfm.cl2": {
		// The contents of the archive header at offset 0 in rmhfm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rmhfm.cl2 starts with the
		// following header: 0A 00 5E 01 11 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmh/rmhht.cl2": {
		// The contents of the archive header at offset 0 in rmhht.cl2:
		//    20 00 00 00  44 00 00 00  68 00 00 00  8C 00 00 00
		//    B0 00 00 00  D4 00 00 00  F8 00 00 00  1C 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x120 in rmhht.cl2 starts with the
		// following header: 0A 00 5F 01 E2 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmh/rmhlm.cl2": {
		// The contents of the archive header at offset 0 in rmhlm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rmhlm.cl2 starts with the
		// following header: 0A 00 63 01 0D 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmh/rmhqm.cl2": {
		// The contents of the archive header at offset 0 in rmhqm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rmhqm.cl2 starts with the
		// following header: 0A 00 62 01 0F 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmh/rmhst.cl2": {
		// The contents of the archive header at offset 0 in rmhst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in rmhst.cl2 starts with the
		// following header: 0A 00 47 01 1F 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmh/rmhwl.cl2": {
		// The contents of the archive header at offset 0 in rmhwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rmhwl.cl2 starts with the
		// following header: 0A 00 20 01 D0 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmm/rmmas.cl2": {
		// The contents of the archive header at offset 0 in rmmas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rmmas.cl2 starts with the
		// following header: 0A 00 3F 01 F1 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmm/rmmat.cl2": {
		// The contents of the archive header at offset 0 in rmmat.cl2:
		//    20 00 00 00  70 00 00 00  C0 00 00 00  10 01 00 00
		//    60 01 00 00  B0 01 00 00  00 02 00 00  50 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x280 in rmmat.cl2 starts with the
		// following header: 0A 00 54 01 05 04 69 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/rogue/rmm/rmmaw.cl2": {
		// The contents of the archive header at offset 0 in rmmaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rmmaw.cl2 starts with the
		// following header: 0A 00 1D 01 AB 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmm/rmmfm.cl2": {
		// The contents of the archive header at offset 0 in rmmfm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rmmfm.cl2 starts with the
		// following header: 0A 00 65 01 07 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmm/rmmht.cl2": {
		// The contents of the archive header at offset 0 in rmmht.cl2:
		//    20 00 00 00  44 00 00 00  68 00 00 00  8C 00 00 00
		//    B0 00 00 00  D4 00 00 00  F8 00 00 00  1C 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x120 in rmmht.cl2 starts with the
		// following header: 0A 00 5C 01 D8 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmm/rmmlm.cl2": {
		// The contents of the archive header at offset 0 in rmmlm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rmmlm.cl2 starts with the
		// following header: 0A 00 63 01 01 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmm/rmmqm.cl2": {
		// The contents of the archive header at offset 0 in rmmqm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rmmqm.cl2 starts with the
		// following header: 0A 00 51 01 C2 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmm/rmmst.cl2": {
		// The contents of the archive header at offset 0 in rmmst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in rmmst.cl2 starts with the
		// following header: 0A 00 40 01 F4 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmm/rmmwl.cl2": {
		// The contents of the archive header at offset 0 in rmmwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rmmwl.cl2 starts with the
		// following header: 0A 00 25 01 B8 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmn/rmnas.cl2": {
		// The contents of the archive header at offset 0 in rmnas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rmnas.cl2 starts with the
		// following header: 0A 00 40 01 1C 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmn/rmnat.cl2": {
		// The contents of the archive header at offset 0 in rmnat.cl2:
		//    20 00 00 00  70 00 00 00  C0 00 00 00  10 01 00 00
		//    60 01 00 00  B0 01 00 00  00 02 00 00  50 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x280 in rmnat.cl2 starts with the
		// following header: 0A 00 47 01 DF 03 4A 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/rogue/rmn/rmnaw.cl2": {
		// The contents of the archive header at offset 0 in rmnaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rmnaw.cl2 starts with the
		// following header: 0A 00 F6 00 77 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmn/rmndt.cl2": {
		// The contents of the archive header at offset 0 in rmndt.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in rmndt.cl2 starts with the
		// following header: 0A 00 54 01 E1 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"plrgfx/rogue/rmn/rmnfm.cl2": {
		// The contents of the archive header at offset 0 in rmnfm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rmnfm.cl2 starts with the
		// following header: 0A 00 3E 01 1F 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmn/rmnht.cl2": {
		// The contents of the archive header at offset 0 in rmnht.cl2:
		//    20 00 00 00  44 00 00 00  68 00 00 00  8C 00 00 00
		//    B0 00 00 00  D4 00 00 00  F8 00 00 00  1C 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x120 in rmnht.cl2 starts with the
		// following header: 0A 00 3C 01 0A 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmn/rmnlm.cl2": {
		// The contents of the archive header at offset 0 in rmnlm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rmnlm.cl2 starts with the
		// following header: 0A 00 43 01 1B 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmn/rmnqm.cl2": {
		// The contents of the archive header at offset 0 in rmnqm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rmnqm.cl2 starts with the
		// following header: 0A 00 3A 01 17 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmn/rmnst.cl2": {
		// The contents of the archive header at offset 0 in rmnst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in rmnst.cl2 starts with the
		// following header: 0A 00 31 01 DD 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmn/rmnwl.cl2": {
		// The contents of the archive header at offset 0 in rmnwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rmnwl.cl2 starts with the
		// following header: 0A 00 F7 00 5A 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rms/rmsas.cl2": {
		// The contents of the archive header at offset 0 in rmsas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rmsas.cl2 starts with the
		// following header: 0A 00 5F 01 06 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rms/rmsat.cl2": {
		// The contents of the archive header at offset 0 in rmsat.cl2:
		//    20 00 00 00  70 00 00 00  C0 00 00 00  10 01 00 00
		//    60 01 00 00  B0 01 00 00  00 02 00 00  50 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x280 in rmsat.cl2 starts with the
		// following header: 0A 00 5B 01 FD 03 64 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/rogue/rms/rmsaw.cl2": {
		// The contents of the archive header at offset 0 in rmsaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rmsaw.cl2 starts with the
		// following header: 0A 00 FA 00 83 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rms/rmsfm.cl2": {
		// The contents of the archive header at offset 0 in rmsfm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rmsfm.cl2 starts with the
		// following header: 0A 00 41 01 F9 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rms/rmsht.cl2": {
		// The contents of the archive header at offset 0 in rmsht.cl2:
		//    20 00 00 00  44 00 00 00  68 00 00 00  8C 00 00 00
		//    B0 00 00 00  D4 00 00 00  F8 00 00 00  1C 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x120 in rmsht.cl2 starts with the
		// following header: 0A 00 07 01 1C 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rms/rmslm.cl2": {
		// The contents of the archive header at offset 0 in rmslm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rmslm.cl2 starts with the
		// following header: 0A 00 3B 01 EC 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rms/rmsqm.cl2": {
		// The contents of the archive header at offset 0 in rmsqm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rmsqm.cl2 starts with the
		// following header: 0A 00 39 01 E2 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rms/rmsst.cl2": {
		// The contents of the archive header at offset 0 in rmsst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in rmsst.cl2 starts with the
		// following header: 0A 00 2C 01 E3 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rms/rmswl.cl2": {
		// The contents of the archive header at offset 0 in rmswl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rmswl.cl2 starts with the
		// following header: 0A 00 FD 00 81 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmt/rmtas.cl2": {
		// The contents of the archive header at offset 0 in rmtas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rmtas.cl2 starts with the
		// following header: 0A 00 2A 01 43 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmt/rmtat.cl2": {
		// The contents of the archive header at offset 0 in rmtat.cl2:
		//    20 00 00 00  70 00 00 00  C0 00 00 00  10 01 00 00
		//    60 01 00 00  B0 01 00 00  00 02 00 00  50 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x280 in rmtat.cl2 starts with the
		// following header: 0A 00 27 01 38 03 99 03 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/rogue/rmt/rmtaw.cl2": {
		// The contents of the archive header at offset 0 in rmtaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rmtaw.cl2 starts with the
		// following header: 0A 00 06 01 33 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmt/rmtfm.cl2": {
		// The contents of the archive header at offset 0 in rmtfm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rmtfm.cl2 starts with the
		// following header: 0A 00 2F 01 4D 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmt/rmtht.cl2": {
		// The contents of the archive header at offset 0 in rmtht.cl2:
		//    20 00 00 00  44 00 00 00  68 00 00 00  8C 00 00 00
		//    B0 00 00 00  D4 00 00 00  F8 00 00 00  1C 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x120 in rmtht.cl2 starts with the
		// following header: 0A 00 1B 01 92 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmt/rmtlm.cl2": {
		// The contents of the archive header at offset 0 in rmtlm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rmtlm.cl2 starts with the
		// following header: 0A 00 2A 01 39 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmt/rmtqm.cl2": {
		// The contents of the archive header at offset 0 in rmtqm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rmtqm.cl2 starts with the
		// following header: 0A 00 23 01 37 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmt/rmtst.cl2": {
		// The contents of the archive header at offset 0 in rmtst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in rmtst.cl2 starts with the
		// following header: 0A 00 5A 01 E0 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmt/rmtwl.cl2": {
		// The contents of the archive header at offset 0 in rmtwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rmtwl.cl2 starts with the
		// following header: 0A 00 23 01 A9 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmu/rmuas.cl2": {
		// The contents of the archive header at offset 0 in rmuas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rmuas.cl2 starts with the
		// following header: 0A 00 3F 01 2E 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmu/rmuat.cl2": {
		// The contents of the archive header at offset 0 in rmuat.cl2:
		//    20 00 00 00  70 00 00 00  C0 00 00 00  10 01 00 00
		//    60 01 00 00  B0 01 00 00  00 02 00 00  50 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x280 in rmuat.cl2 starts with the
		// following header: 0A 00 3A 01 B7 03 1E 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/rogue/rmu/rmuaw.cl2": {
		// The contents of the archive header at offset 0 in rmuaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rmuaw.cl2 starts with the
		// following header: 0A 00 F8 00 9C 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmu/rmubl.cl2": {
		// The contents of the archive header at offset 0 in rmubl.cl2:
		//    20 00 00 00  38 00 00 00  50 00 00 00  68 00 00 00
		//    80 00 00 00  98 00 00 00  B0 00 00 00  C8 00 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0xC0 in rmubl.cl2 starts with the
		// following header: 0A 00 34 01 65 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmu/rmufm.cl2": {
		// The contents of the archive header at offset 0 in rmufm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rmufm.cl2 starts with the
		// following header: 0A 00 3B 01 B8 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmu/rmuht.cl2": {
		// The contents of the archive header at offset 0 in rmuht.cl2:
		//    20 00 00 00  44 00 00 00  68 00 00 00  8C 00 00 00
		//    B0 00 00 00  D4 00 00 00  F8 00 00 00  1C 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x120 in rmuht.cl2 starts with the
		// following header: 0A 00 3A 01 61 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmu/rmulm.cl2": {
		// The contents of the archive header at offset 0 in rmulm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rmulm.cl2 starts with the
		// following header: 0A 00 30 01 AC 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmu/rmuqm.cl2": {
		// The contents of the archive header at offset 0 in rmuqm.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in rmuqm.cl2 starts with the
		// following header: 0A 00 35 01 B6 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmu/rmust.cl2": {
		// The contents of the archive header at offset 0 in rmust.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in rmust.cl2 starts with the
		// following header: 0A 00 2E 01 FD 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/rogue/rmu/rmuwl.cl2": {
		// The contents of the archive header at offset 0 in rmuwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in rmuwl.cl2 starts with the
		// following header: 0A 00 04 01 A5 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/sha/shaas.cl2": {
		// The contents of the archive header at offset 0 in shaas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in shaas.cl2 starts with the
		// following header: 0A 00 8E 01 CE 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/sha/shaat.cl2": {
		// The contents of the archive header at offset 0 in shaat.cl2:
		//    20 00 00 00  88 00 00 00  F0 00 00 00  58 01 00 00
		//    C0 01 00 00  28 02 00 00  90 02 00 00  F8 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x340 in shaat.cl2 starts with the
		// following header: 0A 00 8D 01 CD 04 71 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/sha/shaaw.cl2": {
		// The contents of the archive header at offset 0 in shaaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in shaaw.cl2 starts with the
		// following header: 0A 00 A3 01 0B 05 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/sha/shafm.cl2": {
		// The contents of the archive header at offset 0 in shafm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in shafm.cl2 starts with the
		// following header: 0A 00 A0 01 BC 04 5B 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/sha/shaht.cl2": {
		// The contents of the archive header at offset 0 in shaht.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in shaht.cl2 starts with the
		// following header: 0A 00 83 01 F8 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/sha/shalm.cl2": {
		// The contents of the archive header at offset 0 in shalm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in shalm.cl2 starts with the
		// following header: 0A 00 9A 01 C3 04 68 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/sha/shaqm.cl2": {
		// The contents of the archive header at offset 0 in shaqm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in shaqm.cl2 starts with the
		// following header: 0A 00 98 01 CD 04 78 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/sha/shast.cl2": {
		// The contents of the archive header at offset 0 in shast.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in shast.cl2 starts with the
		// following header: 0A 00 73 01 76 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/sha/shawl.cl2": {
		// The contents of the archive header at offset 0 in shawl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in shawl.cl2 starts with the
		// following header: 0A 00 4E 01 60 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/shb/shbas.cl2": {
		// The contents of the archive header at offset 0 in shbas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in shbas.cl2 starts with the
		// following header: 0A 00 8B 01 77 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/shb/shbat.cl2": {
		// The contents of the archive header at offset 0 in shbat.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in shbat.cl2 starts with the
		// following header: 0A 00 90 01 82 04 32 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/shb/shbaw.cl2": {
		// The contents of the archive header at offset 0 in shbaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in shbaw.cl2 starts with the
		// following header: 0A 00 AD 01 E4 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/shb/shbfm.cl2": {
		// The contents of the archive header at offset 0 in shbfm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in shbfm.cl2 starts with the
		// following header: 0A 00 B1 01 7F 04 27 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/shb/shbht.cl2": {
		// The contents of the archive header at offset 0 in shbht.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in shbht.cl2 starts with the
		// following header: 0A 00 8F 01 EC 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/shb/shblm.cl2": {
		// The contents of the archive header at offset 0 in shblm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in shblm.cl2 starts with the
		// following header: 0A 00 97 01 A0 04 4F 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/shb/shbqm.cl2": {
		// The contents of the archive header at offset 0 in shbqm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in shbqm.cl2 starts with the
		// following header: 0A 00 A0 01 A9 04 57 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/shb/shbst.cl2": {
		// The contents of the archive header at offset 0 in shbst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in shbst.cl2 starts with the
		// following header: 0A 00 76 01 81 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/shb/shbwl.cl2": {
		// The contents of the archive header at offset 0 in shbwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in shbwl.cl2 starts with the
		// following header: 0A 00 4A 01 33 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/shd/shdas.cl2": {
		// The contents of the archive header at offset 0 in shdas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in shdas.cl2 starts with the
		// following header: 0A 00 E0 01 7B 05 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/shd/shdat.cl2": {
		// The contents of the archive header at offset 0 in shdat.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in shdat.cl2 starts with the
		// following header: 0A 00 E0 01 94 05 1C 06 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/shd/shdaw.cl2": {
		// The contents of the archive header at offset 0 in shdaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in shdaw.cl2 starts with the
		// following header: 0A 00 BE 01 72 05 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/shd/shdbl.cl2": {
		// The contents of the archive header at offset 0 in shdbl.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in shdbl.cl2 starts with the
		// following header: 0A 00 81 01 D7 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/shd/shdfm.cl2": {
		// The contents of the archive header at offset 0 in shdfm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in shdfm.cl2 starts with the
		// following header: 0A 00 EA 01 71 05 E3 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/shd/shdht.cl2": {
		// The contents of the archive header at offset 0 in shdht.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in shdht.cl2 starts with the
		// following header: 0A 00 95 01 81 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/shd/shdlm.cl2": {
		// The contents of the archive header at offset 0 in shdlm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in shdlm.cl2 starts with the
		// following header: 0A 00 DF 01 72 05 E9 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/shd/shdqm.cl2": {
		// The contents of the archive header at offset 0 in shdqm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in shdqm.cl2 starts with the
		// following header: 0A 00 E6 01 8C 05 04 06 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/shd/shdst.cl2": {
		// The contents of the archive header at offset 0 in shdst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in shdst.cl2 starts with the
		// following header: 0A 00 78 01 B6 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/shd/shdwl.cl2": {
		// The contents of the archive header at offset 0 in shdwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in shdwl.cl2 starts with the
		// following header: 0A 00 53 01 7E 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/shh/shhas.cl2": {
		// The contents of the archive header at offset 0 in shhas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in shhas.cl2 starts with the
		// following header: 0A 00 C4 01 59 05 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/shh/shhat.cl2": {
		// The contents of the archive header at offset 0 in shhat.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in shhat.cl2 starts with the
		// following header: 0A 00 C0 01 67 05 EE 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/shh/shhaw.cl2": {
		// The contents of the archive header at offset 0 in shhaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in shhaw.cl2 starts with the
		// following header: 0A 00 AC 01 4D 05 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/shh/shhbl.cl2": {
		// The contents of the archive header at offset 0 in shhbl.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in shhbl.cl2 starts with the
		// following header: 0A 00 86 01 86 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/shh/shhfm.cl2": {
		// The contents of the archive header at offset 0 in shhfm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in shhfm.cl2 starts with the
		// following header: 0A 00 C3 01 49 05 CB 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/shh/shhht.cl2": {
		// The contents of the archive header at offset 0 in shhht.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in shhht.cl2 starts with the
		// following header: 0A 00 8C 01 72 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/shh/shhlm.cl2": {
		// The contents of the archive header at offset 0 in shhlm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in shhlm.cl2 starts with the
		// following header: 0A 00 BE 01 5C 05 E7 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/shh/shhqm.cl2": {
		// The contents of the archive header at offset 0 in shhqm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in shhqm.cl2 starts with the
		// following header: 0A 00 C3 01 5C 05 E4 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/shh/shhst.cl2": {
		// The contents of the archive header at offset 0 in shhst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in shhst.cl2 starts with the
		// following header: 0A 00 81 01 C0 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/shh/shhwl.cl2": {
		// The contents of the archive header at offset 0 in shhwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in shhwl.cl2 starts with the
		// following header: 0A 00 56 01 94 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/shm/shmas.cl2": {
		// The contents of the archive header at offset 0 in shmas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in shmas.cl2 starts with the
		// following header: 0A 00 BE 01 F3 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/shm/shmat.cl2": {
		// The contents of the archive header at offset 0 in shmat.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in shmat.cl2 starts with the
		// following header: 0A 00 B9 01 FF 04 85 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/shm/shmaw.cl2": {
		// The contents of the archive header at offset 0 in shmaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in shmaw.cl2 starts with the
		// following header: 0A 00 A9 01 ED 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/shm/shmfm.cl2": {
		// The contents of the archive header at offset 0 in shmfm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in shmfm.cl2 starts with the
		// following header: 0A 00 CB 01 F5 04 76 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/shm/shmht.cl2": {
		// The contents of the archive header at offset 0 in shmht.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in shmht.cl2 starts with the
		// following header: 0A 00 88 01 27 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/shm/shmlm.cl2": {
		// The contents of the archive header at offset 0 in shmlm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in shmlm.cl2 starts with the
		// following header: 0A 00 BC 01 04 05 8E 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/shm/shmqm.cl2": {
		// The contents of the archive header at offset 0 in shmqm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in shmqm.cl2 starts with the
		// following header: 0A 00 B9 01 F8 04 7D 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/shm/shmst.cl2": {
		// The contents of the archive header at offset 0 in shmst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in shmst.cl2 starts with the
		// following header: 0A 00 77 01 89 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/shm/shmwl.cl2": {
		// The contents of the archive header at offset 0 in shmwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in shmwl.cl2 starts with the
		// following header: 0A 00 41 01 37 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/shn/shnas.cl2": {
		// The contents of the archive header at offset 0 in shnas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in shnas.cl2 starts with the
		// following header: 0A 00 AC 01 9A 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/shn/shnat.cl2": {
		// The contents of the archive header at offset 0 in shnat.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in shnat.cl2 starts with the
		// following header: 0A 00 A9 01 A1 04 32 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/shn/shnaw.cl2": {
		// The contents of the archive header at offset 0 in shnaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in shnaw.cl2 starts with the
		// following header: 0A 00 A7 01 AD 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/shn/shndt.cl2": {
		// The contents of the archive header at offset 0 in shndt.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in shndt.cl2 starts with the
		// following header: 0A 00 2D 01 2D 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"plrgfx/sorceror/shn/shnfm.cl2": {
		// The contents of the archive header at offset 0 in shnfm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in shnfm.cl2 starts with the
		// following header: 0A 00 B2 01 98 04 1B 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/shn/shnht.cl2": {
		// The contents of the archive header at offset 0 in shnht.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in shnht.cl2 starts with the
		// following header: 0A 00 7B 01 D4 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/shn/shnlm.cl2": {
		// The contents of the archive header at offset 0 in shnlm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in shnlm.cl2 starts with the
		// following header: 0A 00 AE 01 9C 04 25 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/shn/shnqm.cl2": {
		// The contents of the archive header at offset 0 in shnqm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in shnqm.cl2 starts with the
		// following header: 0A 00 B4 01 AE 04 37 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/shn/shnst.cl2": {
		// The contents of the archive header at offset 0 in shnst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in shnst.cl2 starts with the
		// following header: 0A 00 71 01 7B 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/shn/shnwl.cl2": {
		// The contents of the archive header at offset 0 in shnwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in shnwl.cl2 starts with the
		// following header: 0A 00 4B 01 39 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/shs/shsas.cl2": {
		// The contents of the archive header at offset 0 in shsas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in shsas.cl2 starts with the
		// following header: 0A 00 E0 01 4B 05 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/shs/shsat.cl2": {
		// The contents of the archive header at offset 0 in shsat.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in shsat.cl2 starts with the
		// following header: 0A 00 DF 01 59 05 E0 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/shs/shsaw.cl2": {
		// The contents of the archive header at offset 0 in shsaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in shsaw.cl2 starts with the
		// following header: 0A 00 C0 01 46 05 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/shs/shsfm.cl2": {
		// The contents of the archive header at offset 0 in shsfm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in shsfm.cl2 starts with the
		// following header: 0A 00 E6 01 44 05 C6 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/shs/shsht.cl2": {
		// The contents of the archive header at offset 0 in shsht.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in shsht.cl2 starts with the
		// following header: 0A 00 92 01 56 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/shs/shslm.cl2": {
		// The contents of the archive header at offset 0 in shslm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in shslm.cl2 starts with the
		// following header: 0A 00 DA 01 43 05 CA 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/shs/shsqm.cl2": {
		// The contents of the archive header at offset 0 in shsqm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in shsqm.cl2 starts with the
		// following header: 0A 00 E4 01 57 05 DF 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/shs/shsst.cl2": {
		// The contents of the archive header at offset 0 in shsst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in shsst.cl2 starts with the
		// following header: 0A 00 75 01 85 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/shs/shswl.cl2": {
		// The contents of the archive header at offset 0 in shswl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in shswl.cl2 starts with the
		// following header: 0A 00 38 01 40 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/sht/shtas.cl2": {
		// The contents of the archive header at offset 0 in shtas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in shtas.cl2 starts with the
		// following header: 0A 00 8C 01 B3 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/sht/shtat.cl2": {
		// The contents of the archive header at offset 0 in shtat.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in shtat.cl2 starts with the
		// following header: 0A 00 8D 01 BE 04 62 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/sht/shtaw.cl2": {
		// The contents of the archive header at offset 0 in shtaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in shtaw.cl2 starts with the
		// following header: 0A 00 A7 01 06 05 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/sht/shtfm.cl2": {
		// The contents of the archive header at offset 0 in shtfm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in shtfm.cl2 starts with the
		// following header: 0A 00 97 01 A8 04 4B 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/sht/shtht.cl2": {
		// The contents of the archive header at offset 0 in shtht.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in shtht.cl2 starts with the
		// following header: 0A 00 53 01 2C 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/sht/shtlm.cl2": {
		// The contents of the archive header at offset 0 in shtlm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in shtlm.cl2 starts with the
		// following header: 0A 00 8D 01 AF 04 56 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/sht/shtqm.cl2": {
		// The contents of the archive header at offset 0 in shtqm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in shtqm.cl2 starts with the
		// following header: 0A 00 88 01 B8 04 5F 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/sht/shtst.cl2": {
		// The contents of the archive header at offset 0 in shtst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in shtst.cl2 starts with the
		// following header: 0A 00 C4 01 CA 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/sht/shtwl.cl2": {
		// The contents of the archive header at offset 0 in shtwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in shtwl.cl2 starts with the
		// following header: 0A 00 5C 01 54 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/shu/shuas.cl2": {
		// The contents of the archive header at offset 0 in shuas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in shuas.cl2 starts with the
		// following header: 0A 00 BA 01 08 05 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/shu/shuat.cl2": {
		// The contents of the archive header at offset 0 in shuat.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in shuat.cl2 starts with the
		// following header: 0A 00 B6 01 0A 05 91 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/shu/shuaw.cl2": {
		// The contents of the archive header at offset 0 in shuaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in shuaw.cl2 starts with the
		// following header: 0A 00 AA 01 D1 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/shu/shubl.cl2": {
		// The contents of the archive header at offset 0 in shubl.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in shubl.cl2 starts with the
		// following header: 0A 00 79 01 31 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/shu/shufm.cl2": {
		// The contents of the archive header at offset 0 in shufm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in shufm.cl2 starts with the
		// following header: 0A 00 BF 01 00 05 72 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/shu/shuht.cl2": {
		// The contents of the archive header at offset 0 in shuht.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in shuht.cl2 starts with the
		// following header: 0A 00 7D 01 A5 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/shu/shulm.cl2": {
		// The contents of the archive header at offset 0 in shulm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in shulm.cl2 starts with the
		// following header: 0A 00 BB 01 08 05 7A 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/shu/shuqm.cl2": {
		// The contents of the archive header at offset 0 in shuqm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in shuqm.cl2 starts with the
		// following header: 0A 00 B6 01 06 05 7A 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/shu/shust.cl2": {
		// The contents of the archive header at offset 0 in shust.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in shust.cl2 starts with the
		// following header: 0A 00 77 01 AA 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/shu/shuwl.cl2": {
		// The contents of the archive header at offset 0 in shuwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in shuwl.cl2 starts with the
		// following header: 0A 00 49 01 48 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/sla/slaas.cl2": {
		// The contents of the archive header at offset 0 in slaas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in slaas.cl2 starts with the
		// following header: 0A 00 8E 01 C3 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/sla/slaat.cl2": {
		// The contents of the archive header at offset 0 in slaat.cl2:
		//    20 00 00 00  88 00 00 00  F0 00 00 00  58 01 00 00
		//    C0 01 00 00  28 02 00 00  90 02 00 00  F8 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x340 in slaat.cl2 starts with the
		// following header: 0A 00 92 01 C8 04 51 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/sla/slaaw.cl2": {
		// The contents of the archive header at offset 0 in slaaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in slaaw.cl2 starts with the
		// following header: 0A 00 A5 01 0C 05 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/sla/slafm.cl2": {
		// The contents of the archive header at offset 0 in slafm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in slafm.cl2 starts with the
		// following header: 0A 00 DC 01 17 05 8C 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/sla/slaht.cl2": {
		// The contents of the archive header at offset 0 in slaht.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in slaht.cl2 starts with the
		// following header: 0A 00 71 01 D5 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/sla/slalm.cl2": {
		// The contents of the archive header at offset 0 in slalm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in slalm.cl2 starts with the
		// following header: 0A 00 BE 01 01 05 7A 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/sla/slaqm.cl2": {
		// The contents of the archive header at offset 0 in slaqm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in slaqm.cl2 starts with the
		// following header: 0A 00 B7 01 EB 04 62 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/sla/slast.cl2": {
		// The contents of the archive header at offset 0 in slast.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in slast.cl2 starts with the
		// following header: 0A 00 7A 01 8B 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/sla/slawl.cl2": {
		// The contents of the archive header at offset 0 in slawl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in slawl.cl2 starts with the
		// following header: 0A 00 54 01 6C 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/slb/slbas.cl2": {
		// The contents of the archive header at offset 0 in slbas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in slbas.cl2 starts with the
		// following header: 0A 00 A1 01 8E 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/slb/slbat.cl2": {
		// The contents of the archive header at offset 0 in slbat.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in slbat.cl2 starts with the
		// following header: 0A 00 A5 01 8F 04 2A 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/slb/slbaw.cl2": {
		// The contents of the archive header at offset 0 in slbaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in slbaw.cl2 starts with the
		// following header: 0A 00 A1 01 D2 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/slb/slbfm.cl2": {
		// The contents of the archive header at offset 0 in slbfm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in slbfm.cl2 starts with the
		// following header: 0A 00 96 01 80 04 1D 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/slb/slbht.cl2": {
		// The contents of the archive header at offset 0 in slbht.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in slbht.cl2 starts with the
		// following header: 0A 00 8F 01 E1 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/slb/slblm.cl2": {
		// The contents of the archive header at offset 0 in slblm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in slblm.cl2 starts with the
		// following header: 0A 00 A5 01 A2 04 3D 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/slb/slbqm.cl2": {
		// The contents of the archive header at offset 0 in slbqm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in slbqm.cl2 starts with the
		// following header: 0A 00 AB 01 B6 04 50 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/slb/slbst.cl2": {
		// The contents of the archive header at offset 0 in slbst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in slbst.cl2 starts with the
		// following header: 0A 00 7C 01 93 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/slb/slbwl.cl2": {
		// The contents of the archive header at offset 0 in slbwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in slbwl.cl2 starts with the
		// following header: 0A 00 53 01 68 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/sld/sldas.cl2": {
		// The contents of the archive header at offset 0 in sldas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in sldas.cl2 starts with the
		// following header: 0A 00 E3 01 52 05 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/sld/sldat.cl2": {
		// The contents of the archive header at offset 0 in sldat.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in sldat.cl2 starts with the
		// following header: 0A 00 E4 01 50 05 BF 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/sld/sldaw.cl2": {
		// The contents of the archive header at offset 0 in sldaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in sldaw.cl2 starts with the
		// following header: 0A 00 A3 01 2F 05 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/sld/sldbl.cl2": {
		// The contents of the archive header at offset 0 in sldbl.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in sldbl.cl2 starts with the
		// following header: 0A 00 AE 01 B4 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/sld/sldfm.cl2": {
		// The contents of the archive header at offset 0 in sldfm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in sldfm.cl2 starts with the
		// following header: 0A 00 15 02 99 05 08 06 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/sld/sldht.cl2": {
		// The contents of the archive header at offset 0 in sldht.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in sldht.cl2 starts with the
		// following header: 0A 00 87 01 33 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/sld/sldlm.cl2": {
		// The contents of the archive header at offset 0 in sldlm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in sldlm.cl2 starts with the
		// following header: 0A 00 FD 01 80 05 EF 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/sld/sldqm.cl2": {
		// The contents of the archive header at offset 0 in sldqm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in sldqm.cl2 starts with the
		// following header: 0A 00 F6 01 70 05 DC 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/sld/sldst.cl2": {
		// The contents of the archive header at offset 0 in sldst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in sldst.cl2 starts with the
		// following header: 0A 00 7D 01 A7 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/sld/sldwl.cl2": {
		// The contents of the archive header at offset 0 in sldwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in sldwl.cl2 starts with the
		// following header: 0A 00 41 01 70 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/slh/slhas.cl2": {
		// The contents of the archive header at offset 0 in slhas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in slhas.cl2 starts with the
		// following header: 0A 00 BC 01 07 05 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/slh/slhat.cl2": {
		// The contents of the archive header at offset 0 in slhat.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in slhat.cl2 starts with the
		// following header: 0A 00 C1 01 0A 05 78 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/slh/slhaw.cl2": {
		// The contents of the archive header at offset 0 in slhaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in slhaw.cl2 starts with the
		// following header: 0A 00 A6 01 07 05 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/slh/slhbl.cl2": {
		// The contents of the archive header at offset 0 in slhbl.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in slhbl.cl2 starts with the
		// following header: 0A 00 98 01 63 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/slh/slhfm.cl2": {
		// The contents of the archive header at offset 0 in slhfm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in slhfm.cl2 starts with the
		// following header: 0A 00 AA 01 C0 04 28 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/slh/slhht.cl2": {
		// The contents of the archive header at offset 0 in slhht.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in slhht.cl2 starts with the
		// following header: 0A 00 87 01 19 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/slh/slhlm.cl2": {
		// The contents of the archive header at offset 0 in slhlm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in slhlm.cl2 starts with the
		// following header: 0A 00 CC 01 2F 05 9D 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/slh/slhqm.cl2": {
		// The contents of the archive header at offset 0 in slhqm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in slhqm.cl2 starts with the
		// following header: 0A 00 98 01 D2 04 41 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/slh/slhst.cl2": {
		// The contents of the archive header at offset 0 in slhst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in slhst.cl2 starts with the
		// following header: 0A 00 84 01 B2 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/slh/slhwl.cl2": {
		// The contents of the archive header at offset 0 in slhwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in slhwl.cl2 starts with the
		// following header: 0A 00 5C 01 94 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/slm/slmas.cl2": {
		// The contents of the archive header at offset 0 in slmas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in slmas.cl2 starts with the
		// following header: 0A 00 BC 01 EC 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/slm/slmat.cl2": {
		// The contents of the archive header at offset 0 in slmat.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in slmat.cl2 starts with the
		// following header: 0A 00 C6 01 F7 04 66 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/slm/slmaw.cl2": {
		// The contents of the archive header at offset 0 in slmaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in slmaw.cl2 starts with the
		// following header: 0A 00 A6 01 E5 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/slm/slmfm.cl2": {
		// The contents of the archive header at offset 0 in slmfm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in slmfm.cl2 starts with the
		// following header: 0A 00 E2 01 50 05 BE 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/slm/slmht.cl2": {
		// The contents of the archive header at offset 0 in slmht.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in slmht.cl2 starts with the
		// following header: 0A 00 74 01 09 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/slm/slmlm.cl2": {
		// The contents of the archive header at offset 0 in slmlm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in slmlm.cl2 starts with the
		// following header: 0A 00 CA 01 2A 05 98 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/slm/slmqm.cl2": {
		// The contents of the archive header at offset 0 in slmqm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in slmqm.cl2 starts with the
		// following header: 0A 00 CF 01 31 05 9D 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/slm/slmst.cl2": {
		// The contents of the archive header at offset 0 in slmst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in slmst.cl2 starts with the
		// following header: 0A 00 86 01 AB 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/slm/slmwl.cl2": {
		// The contents of the archive header at offset 0 in slmwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in slmwl.cl2 starts with the
		// following header: 0A 00 55 01 81 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/sln/slnas.cl2": {
		// The contents of the archive header at offset 0 in slnas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in slnas.cl2 starts with the
		// following header: 0A 00 B4 01 9D 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/sln/slnat.cl2": {
		// The contents of the archive header at offset 0 in slnat.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in slnat.cl2 starts with the
		// following header: 0A 00 B4 01 A1 04 0E 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/sln/slnaw.cl2": {
		// The contents of the archive header at offset 0 in slnaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in slnaw.cl2 starts with the
		// following header: 0A 00 6B 01 84 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/sln/slndt.cl2": {
		// The contents of the archive header at offset 0 in slndt.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in slndt.cl2 starts with the
		// following header: 0A 00 09 01 01 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"plrgfx/sorceror/sln/slnfm.cl2": {
		// The contents of the archive header at offset 0 in slnfm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in slnfm.cl2 starts with the
		// following header: 0A 00 BA 01 A9 04 16 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/sln/slnht.cl2": {
		// The contents of the archive header at offset 0 in slnht.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in slnht.cl2 starts with the
		// following header: 0A 00 6D 01 C5 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/sln/slnlm.cl2": {
		// The contents of the archive header at offset 0 in slnlm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in slnlm.cl2 starts with the
		// following header: 0A 00 B1 01 90 04 FC 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/sln/slnqm.cl2": {
		// The contents of the archive header at offset 0 in slnqm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in slnqm.cl2 starts with the
		// following header: 0A 00 B9 01 A7 04 15 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/sln/slnst.cl2": {
		// The contents of the archive header at offset 0 in slnst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in slnst.cl2 starts with the
		// following header: 0A 00 7B 01 92 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/sln/slnwl.cl2": {
		// The contents of the archive header at offset 0 in slnwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in slnwl.cl2 starts with the
		// following header: 0A 00 4E 01 5D 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/sls/slsas.cl2": {
		// The contents of the archive header at offset 0 in slsas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in slsas.cl2 starts with the
		// following header: 0A 00 E8 01 4A 05 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/sls/slsat.cl2": {
		// The contents of the archive header at offset 0 in slsat.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in slsat.cl2 starts with the
		// following header: 0A 00 E8 01 45 05 B3 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/sls/slsaw.cl2": {
		// The contents of the archive header at offset 0 in slsaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in slsaw.cl2 starts with the
		// following header: 0A 00 BE 01 3E 05 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/sls/slsfm.cl2": {
		// The contents of the archive header at offset 0 in slsfm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in slsfm.cl2 starts with the
		// following header: 0A 00 C3 01 F5 04 61 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/sls/slsht.cl2": {
		// The contents of the archive header at offset 0 in slsht.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in slsht.cl2 starts with the
		// following header: 0A 00 8A 01 41 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/sls/slslm.cl2": {
		// The contents of the archive header at offset 0 in slslm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in slslm.cl2 starts with the
		// following header: 0A 00 E1 01 41 05 B0 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/sls/slsqm.cl2": {
		// The contents of the archive header at offset 0 in slsqm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in slsqm.cl2 starts with the
		// following header: 0A 00 C5 01 10 05 7C 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/sls/slsst.cl2": {
		// The contents of the archive header at offset 0 in slsst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in slsst.cl2 starts with the
		// following header: 0A 00 80 01 9B 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/sls/slswl.cl2": {
		// The contents of the archive header at offset 0 in slswl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in slswl.cl2 starts with the
		// following header: 0A 00 51 01 74 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/slt/sltas.cl2": {
		// The contents of the archive header at offset 0 in sltas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in sltas.cl2 starts with the
		// following header: 0A 00 83 01 72 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/slt/sltat.cl2": {
		// The contents of the archive header at offset 0 in sltat.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in sltat.cl2 starts with the
		// following header: 0A 00 8C 01 6F 04 FD 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/slt/sltaw.cl2": {
		// The contents of the archive header at offset 0 in sltaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in sltaw.cl2 starts with the
		// following header: 0A 00 9B 01 D7 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/slt/sltfm.cl2": {
		// The contents of the archive header at offset 0 in sltfm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in sltfm.cl2 starts with the
		// following header: 0A 00 77 01 4D 04 D9 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/slt/sltht.cl2": {
		// The contents of the archive header at offset 0 in sltht.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in sltht.cl2 starts with the
		// following header: 0A 00 50 01 FA 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/slt/sltlm.cl2": {
		// The contents of the archive header at offset 0 in sltlm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in sltlm.cl2 starts with the
		// following header: 0A 00 81 01 67 04 F4 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/slt/sltqm.cl2": {
		// The contents of the archive header at offset 0 in sltqm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in sltqm.cl2 starts with the
		// following header: 0A 00 82 01 71 04 FE 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/slt/sltst.cl2": {
		// The contents of the archive header at offset 0 in sltst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in sltst.cl2 starts with the
		// following header: 0A 00 9B 01 BE 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/slt/sltwl.cl2": {
		// The contents of the archive header at offset 0 in sltwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in sltwl.cl2 starts with the
		// following header: 0A 00 59 01 84 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/slu/sluas.cl2": {
		// The contents of the archive header at offset 0 in sluas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in sluas.cl2 starts with the
		// following header: 0A 00 B2 01 B8 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/slu/sluat.cl2": {
		// The contents of the archive header at offset 0 in sluat.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in sluat.cl2 starts with the
		// following header: 0A 00 B4 01 BA 04 28 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/slu/sluaw.cl2": {
		// The contents of the archive header at offset 0 in sluaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in sluaw.cl2 starts with the
		// following header: 0A 00 91 01 2C 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/slu/slubl.cl2": {
		// The contents of the archive header at offset 0 in slubl.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in slubl.cl2 starts with the
		// following header: 0A 00 95 01 1A 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/slu/slufm.cl2": {
		// The contents of the archive header at offset 0 in slufm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in slufm.cl2 starts with the
		// following header: 0A 00 95 01 7F 04 ED 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/slu/sluht.cl2": {
		// The contents of the archive header at offset 0 in sluht.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in sluht.cl2 starts with the
		// following header: 0A 00 6A 01 99 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/slu/slulm.cl2": {
		// The contents of the archive header at offset 0 in slulm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in slulm.cl2 starts with the
		// following header: 0A 00 95 01 91 04 FF 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/slu/sluqm.cl2": {
		// The contents of the archive header at offset 0 in sluqm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in sluqm.cl2 starts with the
		// following header: 0A 00 9E 01 A9 04 16 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/slu/slust.cl2": {
		// The contents of the archive header at offset 0 in slust.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in slust.cl2 starts with the
		// following header: 0A 00 74 01 93 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/slu/sluwl.cl2": {
		// The contents of the archive header at offset 0 in sluwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in sluwl.cl2 starts with the
		// following header: 0A 00 4B 01 6B 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/sma/smaas.cl2": {
		// The contents of the archive header at offset 0 in smaas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in smaas.cl2 starts with the
		// following header: 0A 00 89 01 A2 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/sma/smaat.cl2": {
		// The contents of the archive header at offset 0 in smaat.cl2:
		//    20 00 00 00  88 00 00 00  F0 00 00 00  58 01 00 00
		//    C0 01 00 00  28 02 00 00  90 02 00 00  F8 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x340 in smaat.cl2 starts with the
		// following header: 0A 00 90 01 AF 04 37 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/sma/smaaw.cl2": {
		// The contents of the archive header at offset 0 in smaaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in smaaw.cl2 starts with the
		// following header: 0A 00 97 01 EC 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/sma/smafm.cl2": {
		// The contents of the archive header at offset 0 in smafm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in smafm.cl2 starts with the
		// following header: 0A 00 BB 01 C4 04 3C 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/sma/smaht.cl2": {
		// The contents of the archive header at offset 0 in smaht.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in smaht.cl2 starts with the
		// following header: 0A 00 67 01 C0 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/sma/smalm.cl2": {
		// The contents of the archive header at offset 0 in smalm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in smalm.cl2 starts with the
		// following header: 0A 00 B9 01 E0 04 59 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/sma/smaqm.cl2": {
		// The contents of the archive header at offset 0 in smaqm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in smaqm.cl2 starts with the
		// following header: 0A 00 B9 01 DE 04 5B 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/sma/smast.cl2": {
		// The contents of the archive header at offset 0 in smast.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in smast.cl2 starts with the
		// following header: 0A 00 77 01 7E 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/sma/smawl.cl2": {
		// The contents of the archive header at offset 0 in smawl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in smawl.cl2 starts with the
		// following header: 0A 00 26 01 1A 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/smb/smbas.cl2": {
		// The contents of the archive header at offset 0 in smbas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in smbas.cl2 starts with the
		// following header: 0A 00 93 01 5E 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/smb/smbat.cl2": {
		// The contents of the archive header at offset 0 in smbat.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in smbat.cl2 starts with the
		// following header: 0A 00 A5 01 7B 04 10 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/smb/smbaw.cl2": {
		// The contents of the archive header at offset 0 in smbaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in smbaw.cl2 starts with the
		// following header: 0A 00 A1 01 A6 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/smb/smbfm.cl2": {
		// The contents of the archive header at offset 0 in smbfm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in smbfm.cl2 starts with the
		// following header: 0A 00 A5 01 57 04 E8 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/smb/smbht.cl2": {
		// The contents of the archive header at offset 0 in smbht.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in smbht.cl2 starts with the
		// following header: 0A 00 8A 01 C3 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/smb/smblm.cl2": {
		// The contents of the archive header at offset 0 in smblm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in smblm.cl2 starts with the
		// following header: 0A 00 A4 01 93 04 26 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/smb/smbqm.cl2": {
		// The contents of the archive header at offset 0 in smbqm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in smbqm.cl2 starts with the
		// following header: 0A 00 A8 01 91 04 2C 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/smb/smbst.cl2": {
		// The contents of the archive header at offset 0 in smbst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in smbst.cl2 starts with the
		// following header: 0A 00 69 01 75 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/smb/smbwl.cl2": {
		// The contents of the archive header at offset 0 in smbwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in smbwl.cl2 starts with the
		// following header: 0A 00 28 01 0E 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/smd/smdas.cl2": {
		// The contents of the archive header at offset 0 in smdas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in smdas.cl2 starts with the
		// following header: 0A 00 DF 01 6B 05 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/smd/smdat.cl2": {
		// The contents of the archive header at offset 0 in smdat.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in smdat.cl2 starts with the
		// following header: 0A 00 DF 01 6F 05 DF 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/smd/smdaw.cl2": {
		// The contents of the archive header at offset 0 in smdaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in smdaw.cl2 starts with the
		// following header: 0A 00 B2 01 54 05 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/smd/smdbl.cl2": {
		// The contents of the archive header at offset 0 in smdbl.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in smdbl.cl2 starts with the
		// following header: 0A 00 7A 01 C0 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/smd/smdfm.cl2": {
		// The contents of the archive header at offset 0 in smdfm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in smdfm.cl2 starts with the
		// following header: 0A 00 E3 01 5E 05 C6 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/smd/smdht.cl2": {
		// The contents of the archive header at offset 0 in smdht.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in smdht.cl2 starts with the
		// following header: 0A 00 8D 01 53 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/smd/smdlm.cl2": {
		// The contents of the archive header at offset 0 in smdlm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in smdlm.cl2 starts with the
		// following header: 0A 00 E0 01 6B 05 D7 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/smd/smdqm.cl2": {
		// The contents of the archive header at offset 0 in smdqm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in smdqm.cl2 starts with the
		// following header: 0A 00 D7 01 66 05 D5 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/smd/smdst.cl2": {
		// The contents of the archive header at offset 0 in smdst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in smdst.cl2 starts with the
		// following header: 0A 00 73 01 AC 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/smd/smdwl.cl2": {
		// The contents of the archive header at offset 0 in smdwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in smdwl.cl2 starts with the
		// following header: 0A 00 2E 01 60 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/smh/smhas.cl2": {
		// The contents of the archive header at offset 0 in smhas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in smhas.cl2 starts with the
		// following header: 0A 00 BF 01 43 05 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/smh/smhat.cl2": {
		// The contents of the archive header at offset 0 in smhat.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in smhat.cl2 starts with the
		// following header: 0A 00 C0 01 42 05 B3 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/smh/smhaw.cl2": {
		// The contents of the archive header at offset 0 in smhaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in smhaw.cl2 starts with the
		// following header: 0A 00 A2 01 2C 05 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/smh/smhbl.cl2": {
		// The contents of the archive header at offset 0 in smhbl.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in smhbl.cl2 starts with the
		// following header: 0A 00 7B 01 5C 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/smh/smhfm.cl2": {
		// The contents of the archive header at offset 0 in smhfm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in smhfm.cl2 starts with the
		// following header: 0A 00 CB 01 49 05 B4 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/smh/smhht.cl2": {
		// The contents of the archive header at offset 0 in smhht.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in smhht.cl2 starts with the
		// following header: 0A 00 82 01 48 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/smh/smhlm.cl2": {
		// The contents of the archive header at offset 0 in smhlm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in smhlm.cl2 starts with the
		// following header: 0A 00 C2 01 4D 05 BE 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/smh/smhqm.cl2": {
		// The contents of the archive header at offset 0 in smhqm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in smhqm.cl2 starts with the
		// following header: 0A 00 C2 01 45 05 BC 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/smh/smhst.cl2": {
		// The contents of the archive header at offset 0 in smhst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in smhst.cl2 starts with the
		// following header: 0A 00 7F 01 C3 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/smh/smhwl.cl2": {
		// The contents of the archive header at offset 0 in smhwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in smhwl.cl2 starts with the
		// following header: 0A 00 39 01 76 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/smm/smmas.cl2": {
		// The contents of the archive header at offset 0 in smmas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in smmas.cl2 starts with the
		// following header: 0A 00 BB 01 DD 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/smm/smmat.cl2": {
		// The contents of the archive header at offset 0 in smmat.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in smmat.cl2 starts with the
		// following header: 0A 00 BF 01 EA 04 60 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/smm/smmaw.cl2": {
		// The contents of the archive header at offset 0 in smmaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in smmaw.cl2 starts with the
		// following header: 0A 00 A0 01 CB 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/smm/smmfm.cl2": {
		// The contents of the archive header at offset 0 in smmfm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in smmfm.cl2 starts with the
		// following header: 0A 00 CA 01 E0 04 4B 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/smm/smmht.cl2": {
		// The contents of the archive header at offset 0 in smmht.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in smmht.cl2 starts with the
		// following header: 0A 00 77 01 F5 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/smm/smmlm.cl2": {
		// The contents of the archive header at offset 0 in smmlm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in smmlm.cl2 starts with the
		// following header: 0A 00 C8 01 2A 05 9A 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/smm/smmqm.cl2": {
		// The contents of the archive header at offset 0 in smmqm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in smmqm.cl2 starts with the
		// following header: 0A 00 C0 01 EC 04 5E 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/smm/smmst.cl2": {
		// The contents of the archive header at offset 0 in smmst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in smmst.cl2 starts with the
		// following header: 0A 00 79 01 94 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/smm/smmwl.cl2": {
		// The contents of the archive header at offset 0 in smmwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in smmwl.cl2 starts with the
		// following header: 0A 00 2B 01 24 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/smn/smnas.cl2": {
		// The contents of the archive header at offset 0 in smnas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in smnas.cl2 starts with the
		// following header: 0A 00 AE 01 87 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/smn/smnat.cl2": {
		// The contents of the archive header at offset 0 in smnat.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in smnat.cl2 starts with the
		// following header: 0A 00 AD 01 AE 04 1C 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/smn/smnaw.cl2": {
		// The contents of the archive header at offset 0 in smnaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in smnaw.cl2 starts with the
		// following header: 0A 00 9D 01 96 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/smn/smndt.cl2": {
		// The contents of the archive header at offset 0 in smndt.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in smndt.cl2 starts with the
		// following header: 0A 00 25 01 0C 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"plrgfx/sorceror/smn/smnfm.cl2": {
		// The contents of the archive header at offset 0 in smnfm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in smnfm.cl2 starts with the
		// following header: 0A 00 B8 01 85 04 F1 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/smn/smnht.cl2": {
		// The contents of the archive header at offset 0 in smnht.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in smnht.cl2 starts with the
		// following header: 0A 00 6A 01 9F 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/smn/smnlm.cl2": {
		// The contents of the archive header at offset 0 in smnlm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in smnlm.cl2 starts with the
		// following header: 0A 00 B6 01 9C 04 0B 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/smn/smnqm.cl2": {
		// The contents of the archive header at offset 0 in smnqm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in smnqm.cl2 starts with the
		// following header: 0A 00 AF 01 9A 04 0D 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/smn/smnst.cl2": {
		// The contents of the archive header at offset 0 in smnst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in smnst.cl2 starts with the
		// following header: 0A 00 74 01 6E 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/smn/smnwl.cl2": {
		// The contents of the archive header at offset 0 in smnwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in smnwl.cl2 starts with the
		// following header: 0A 00 3B 01 09 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/sms/smsas.cl2": {
		// The contents of the archive header at offset 0 in smsas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in smsas.cl2 starts with the
		// following header: 0A 00 E3 01 3C 05 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/sms/smsat.cl2": {
		// The contents of the archive header at offset 0 in smsat.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in smsat.cl2 starts with the
		// following header: 0A 00 DB 01 31 05 A4 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/sms/smsaw.cl2": {
		// The contents of the archive header at offset 0 in smsaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in smsaw.cl2 starts with the
		// following header: 0A 00 AE 01 22 05 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/sms/smsfm.cl2": {
		// The contents of the archive header at offset 0 in smsfm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in smsfm.cl2 starts with the
		// following header: 0A 00 F2 01 32 05 9D 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/sms/smsht.cl2": {
		// The contents of the archive header at offset 0 in smsht.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in smsht.cl2 starts with the
		// following header: 0A 00 8D 01 3B 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/sms/smslm.cl2": {
		// The contents of the archive header at offset 0 in smslm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in smslm.cl2 starts with the
		// following header: 0A 00 DF 01 3A 05 AF 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/sms/smsqm.cl2": {
		// The contents of the archive header at offset 0 in smsqm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in smsqm.cl2 starts with the
		// following header: 0A 00 DF 01 37 05 AA 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/sms/smsst.cl2": {
		// The contents of the archive header at offset 0 in smsst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in smsst.cl2 starts with the
		// following header: 0A 00 73 01 81 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/sms/smswl.cl2": {
		// The contents of the archive header at offset 0 in smswl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in smswl.cl2 starts with the
		// following header: 0A 00 22 01 12 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/smt/smtas.cl2": {
		// The contents of the archive header at offset 0 in smtas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in smtas.cl2 starts with the
		// following header: 0A 00 91 01 A6 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/smt/smtat.cl2": {
		// The contents of the archive header at offset 0 in smtat.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in smtat.cl2 starts with the
		// following header: 0A 00 90 01 9F 04 29 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/smt/smtaw.cl2": {
		// The contents of the archive header at offset 0 in smtaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in smtaw.cl2 starts with the
		// following header: 0A 00 9C 01 D1 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/smt/smtfm.cl2": {
		// The contents of the archive header at offset 0 in smtfm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in smtfm.cl2 starts with the
		// following header: 0A 00 93 01 6E 04 F7 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/smt/smtht.cl2": {
		// The contents of the archive header at offset 0 in smtht.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in smtht.cl2 starts with the
		// following header: 0A 00 48 01 FF 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/smt/smtlm.cl2": {
		// The contents of the archive header at offset 0 in smtlm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in smtlm.cl2 starts with the
		// following header: 0A 00 86 01 8A 04 14 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/smt/smtqm.cl2": {
		// The contents of the archive header at offset 0 in smtqm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in smtqm.cl2 starts with the
		// following header: 0A 00 81 01 8A 04 17 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/smt/smtst.cl2": {
		// The contents of the archive header at offset 0 in smtst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in smtst.cl2 starts with the
		// following header: 0A 00 BB 01 C9 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/smt/smtwl.cl2": {
		// The contents of the archive header at offset 0 in smtwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in smtwl.cl2 starts with the
		// following header: 0A 00 27 01 0A 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/smu/smuas.cl2": {
		// The contents of the archive header at offset 0 in smuas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in smuas.cl2 starts with the
		// following header: 0A 00 B0 01 F1 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/smu/smuat.cl2": {
		// The contents of the archive header at offset 0 in smuat.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in smuat.cl2 starts with the
		// following header: 0A 00 B4 01 FF 04 6F 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/smu/smuaw.cl2": {
		// The contents of the archive header at offset 0 in smuaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in smuaw.cl2 starts with the
		// following header: 0A 00 9A 01 B5 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/smu/smubl.cl2": {
		// The contents of the archive header at offset 0 in smubl.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in smubl.cl2 starts with the
		// following header: 0A 00 69 01 04 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/smu/smufm.cl2": {
		// The contents of the archive header at offset 0 in smufm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in smufm.cl2 starts with the
		// following header: 0A 00 B8 01 E1 04 4D 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/smu/smuht.cl2": {
		// The contents of the archive header at offset 0 in smuht.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in smuht.cl2 starts with the
		// following header: 0A 00 72 01 73 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/smu/smulm.cl2": {
		// The contents of the archive header at offset 0 in smulm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in smulm.cl2 starts with the
		// following header: 0A 00 AE 01 F4 04 68 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/smu/smuqm.cl2": {
		// The contents of the archive header at offset 0 in smuqm.cl2:
		//    20 00 00 00  58 00 00 00  90 00 00 00  C8 00 00 00
		//    00 01 00 00  38 01 00 00  70 01 00 00  A8 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x1C0 in smuqm.cl2 starts with the
		// following header: 0A 00 B8 01 00 05 75 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/sorceror/smu/smust.cl2": {
		// The contents of the archive header at offset 0 in smust.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in smust.cl2 starts with the
		// following header: 0A 00 7A 01 A3 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/sorceror/smu/smuwl.cl2": {
		// The contents of the archive header at offset 0 in smuwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in smuwl.cl2 starts with the
		// following header: 0A 00 4C 01 63 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wha/whaas.cl2": {
		// The contents of the archive header at offset 0 in whaas.cl2:
		//    20 00 00 00  50 00 00 00  80 00 00 00  B0 00 00 00
		//    E0 00 00 00  10 01 00 00  40 01 00 00  70 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x180 in whaas.cl2 starts with the
		// following header: 0A 00 42 01 0E 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wha/whaat.cl2": {
		// The contents of the archive header at offset 0 in whaat.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in whaat.cl2 starts with the
		// following header: 0A 00 49 01 2F 04 C6 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/warrior/wha/whaaw.cl2": {
		// The contents of the archive header at offset 0 in whaaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in whaaw.cl2 starts with the
		// following header: 0A 00 68 01 3F 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wha/whafm.cl2": {
		// The contents of the archive header at offset 0 in whafm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in whafm.cl2 starts with the
		// following header: 0A 00 4E 01 50 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wha/whaht.cl2": {
		// The contents of the archive header at offset 0 in whaht.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in whaht.cl2 starts with the
		// following header: 0A 00 49 01 27 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wha/whalm.cl2": {
		// The contents of the archive header at offset 0 in whalm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in whalm.cl2 starts with the
		// following header: 0A 00 43 01 29 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wha/whaqm.cl2": {
		// The contents of the archive header at offset 0 in whaqm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in whaqm.cl2 starts with the
		// following header: 0A 00 BB 01 8D 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wha/whast.cl2": {
		// The contents of the archive header at offset 0 in whast.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in whast.cl2 starts with the
		// following header: 0A 00 5B 01 C9 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wha/whawl.cl2": {
		// The contents of the archive header at offset 0 in whawl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in whawl.cl2 starts with the
		// following header: 0A 00 48 01 8B 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whb/whbas.cl2": {
		// The contents of the archive header at offset 0 in whbas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in whbas.cl2 starts with the
		// following header: 0A 00 1B 01 F0 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whb/whbat.cl2": {
		// The contents of the archive header at offset 0 in whbat.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in whbat.cl2 starts with the
		// following header: 0A 00 2A 02 B6 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"plrgfx/warrior/whb/whbaw.cl2": {
		// The contents of the archive header at offset 0 in whbaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in whbaw.cl2 starts with the
		// following header: 0A 00 61 01 3F 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whb/whbfm.cl2": {
		// The contents of the archive header at offset 0 in whbfm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in whbfm.cl2 starts with the
		// following header: 0A 00 3C 01 71 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whb/whbht.cl2": {
		// The contents of the archive header at offset 0 in whbht.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in whbht.cl2 starts with the
		// following header: 0A 00 3A 01 17 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whb/whblm.cl2": {
		// The contents of the archive header at offset 0 in whblm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in whblm.cl2 starts with the
		// following header: 0A 00 26 01 40 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whb/whbqm.cl2": {
		// The contents of the archive header at offset 0 in whbqm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in whbqm.cl2 starts with the
		// following header: 0A 00 B7 01 91 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whb/whbst.cl2": {
		// The contents of the archive header at offset 0 in whbst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in whbst.cl2 starts with the
		// following header: 0A 00 4A 01 9E 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whb/whbwl.cl2": {
		// The contents of the archive header at offset 0 in whbwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in whbwl.cl2 starts with the
		// following header: 0A 00 2C 01 7E 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whd/whdas.cl2": {
		// The contents of the archive header at offset 0 in whdas.cl2:
		//    20 00 00 00  50 00 00 00  80 00 00 00  B0 00 00 00
		//    E0 00 00 00  10 01 00 00  40 01 00 00  70 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x180 in whdas.cl2 starts with the
		// following header: 0A 00 75 01 96 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whd/whdat.cl2": {
		// The contents of the archive header at offset 0 in whdat.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in whdat.cl2 starts with the
		// following header: 0A 00 81 01 E8 04 E4 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/warrior/whd/whdaw.cl2": {
		// The contents of the archive header at offset 0 in whdaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in whdaw.cl2 starts with the
		// following header: 0A 00 5C 01 32 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whd/whdbl.cl2": {
		// The contents of the archive header at offset 0 in whdbl.cl2:
		//    20 00 00 00  30 00 00 00  40 00 00 00  50 00 00 00
		//    60 00 00 00  70 00 00 00  80 00 00 00  90 00 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x80 in whdbl.cl2 starts with the
		// following header: 0A 00 84 01 CF 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whd/whdfm.cl2": {
		// The contents of the archive header at offset 0 in whdfm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in whdfm.cl2 starts with the
		// following header: 0A 00 82 01 A8 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whd/whdht.cl2": {
		// The contents of the archive header at offset 0 in whdht.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in whdht.cl2 starts with the
		// following header: 0A 00 3A 01 3D 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whd/whdlm.cl2": {
		// The contents of the archive header at offset 0 in whdlm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in whdlm.cl2 starts with the
		// following header: 0A 00 60 01 83 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whd/whdqm.cl2": {
		// The contents of the archive header at offset 0 in whdqm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in whdqm.cl2 starts with the
		// following header: 0A 00 C5 01 E3 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whd/whdst.cl2": {
		// The contents of the archive header at offset 0 in whdst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in whdst.cl2 starts with the
		// following header: 0A 00 53 01 D9 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whd/whdwl.cl2": {
		// The contents of the archive header at offset 0 in whdwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in whdwl.cl2 starts with the
		// following header: 0A 00 0A 01 75 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whh/whhas.cl2": {
		// The contents of the archive header at offset 0 in whhas.cl2:
		//    20 00 00 00  50 00 00 00  80 00 00 00  B0 00 00 00
		//    E0 00 00 00  10 01 00 00  40 01 00 00  70 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x180 in whhas.cl2 starts with the
		// following header: 0A 00 70 01 51 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whh/whhat.cl2": {
		// The contents of the archive header at offset 0 in whhat.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in whhat.cl2 starts with the
		// following header: 0A 00 7A 01 96 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"plrgfx/warrior/whh/whhaw.cl2": {
		// The contents of the archive header at offset 0 in whhaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in whhaw.cl2 starts with the
		// following header: 0A 00 61 01 01 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whh/whhbl.cl2": {
		// The contents of the archive header at offset 0 in whhbl.cl2:
		//    20 00 00 00  30 00 00 00  40 00 00 00  50 00 00 00
		//    60 00 00 00  70 00 00 00  80 00 00 00  90 00 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x80 in whhbl.cl2 starts with the
		// following header: 0A 00 73 01 82 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whh/whhfm.cl2": {
		// The contents of the archive header at offset 0 in whhfm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in whhfm.cl2 starts with the
		// following header: 0A 00 69 01 6C 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whh/whhht.cl2": {
		// The contents of the archive header at offset 0 in whhht.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in whhht.cl2 starts with the
		// following header: 0A 00 39 01 21 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whh/whhlm.cl2": {
		// The contents of the archive header at offset 0 in whhlm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in whhlm.cl2 starts with the
		// following header: 0A 00 5B 01 4C 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whh/whhqm.cl2": {
		// The contents of the archive header at offset 0 in whhqm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in whhqm.cl2 starts with the
		// following header: 0A 00 CB 01 BD 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whh/whhst.cl2": {
		// The contents of the archive header at offset 0 in whhst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in whhst.cl2 starts with the
		// following header: 0A 00 94 01 2C 05 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whh/whhwl.cl2": {
		// The contents of the archive header at offset 0 in whhwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in whhwl.cl2 starts with the
		// following header: 0A 00 2A 01 B5 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whm/whmas.cl2": {
		// The contents of the archive header at offset 0 in whmas.cl2:
		//    20 00 00 00  50 00 00 00  80 00 00 00  B0 00 00 00
		//    E0 00 00 00  10 01 00 00  40 01 00 00  70 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x180 in whmas.cl2 starts with the
		// following header: 0A 00 44 01 95 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whm/whmat.cl2": {
		// The contents of the archive header at offset 0 in whmat.cl2:
		//    20 00 00 00  6C 00 00 00  B8 00 00 00  04 01 00 00
		//    50 01 00 00  9C 01 00 00  E8 01 00 00  34 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x260 in whmat.cl2 starts with the
		// following header: 0A 00 40 01 8C 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"plrgfx/warrior/whm/whmaw.cl2": {
		// The contents of the archive header at offset 0 in whmaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in whmaw.cl2 starts with the
		// following header: 0A 00 68 01 EC 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whm/whmfm.cl2": {
		// The contents of the archive header at offset 0 in whmfm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in whmfm.cl2 starts with the
		// following header: 0A 00 49 01 AE 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whm/whmht.cl2": {
		// The contents of the archive header at offset 0 in whmht.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in whmht.cl2 starts with the
		// following header: 0A 00 38 01 9E 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whm/whmlm.cl2": {
		// The contents of the archive header at offset 0 in whmlm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in whmlm.cl2 starts with the
		// following header: 0A 00 40 01 97 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whm/whmqm.cl2": {
		// The contents of the archive header at offset 0 in whmqm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in whmqm.cl2 starts with the
		// following header: 0A 00 B9 01 F3 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whm/whmst.cl2": {
		// The contents of the archive header at offset 0 in whmst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in whmst.cl2 starts with the
		// following header: 0A 00 85 01 D3 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whm/whmwl.cl2": {
		// The contents of the archive header at offset 0 in whmwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in whmwl.cl2 starts with the
		// following header: 0A 00 21 01 64 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whn/whnas.cl2": {
		// The contents of the archive header at offset 0 in whnas.cl2:
		//    20 00 00 00  50 00 00 00  80 00 00 00  B0 00 00 00
		//    E0 00 00 00  10 01 00 00  40 01 00 00  70 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x180 in whnas.cl2 starts with the
		// following header: 0A 00 29 01 29 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whn/whnat.cl2": {
		// The contents of the archive header at offset 0 in whnat.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in whnat.cl2 starts with the
		// following header: 0A 00 35 01 58 03 D7 03 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/warrior/whn/whnaw.cl2": {
		// The contents of the archive header at offset 0 in whnaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in whnaw.cl2 starts with the
		// following header: 0A 00 5E 01 3A 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whn/whndt.cl2": {
		// The contents of the archive header at offset 0 in whndt.cl2:
		//    20 00 00 00  64 00 00 00  A8 00 00 00  EC 00 00 00
		//    30 01 00 00  74 01 00 00  B8 01 00 00  FC 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x220 in whndt.cl2 starts with the
		// following header: 0A 00 CD 01 86 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"plrgfx/warrior/whn/whnfm.cl2": {
		// The contents of the archive header at offset 0 in whnfm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in whnfm.cl2 starts with the
		// following header: 0A 00 2A 01 38 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whn/whnht.cl2": {
		// The contents of the archive header at offset 0 in whnht.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in whnht.cl2 starts with the
		// following header: 0A 00 30 01 62 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whn/whnlm.cl2": {
		// The contents of the archive header at offset 0 in whnlm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in whnlm.cl2 starts with the
		// following header: 0A 00 32 01 53 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whn/whnqm.cl2": {
		// The contents of the archive header at offset 0 in whnqm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in whnqm.cl2 starts with the
		// following header: 0A 00 7E 01 4C 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whn/whnst.cl2": {
		// The contents of the archive header at offset 0 in whnst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in whnst.cl2 starts with the
		// following header: 0A 00 43 01 6D 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whn/whnwl.cl2": {
		// The contents of the archive header at offset 0 in whnwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in whnwl.cl2 starts with the
		// following header: 0A 00 F8 00 0F 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whs/whsas.cl2": {
		// The contents of the archive header at offset 0 in whsas.cl2:
		//    20 00 00 00  50 00 00 00  80 00 00 00  B0 00 00 00
		//    E0 00 00 00  10 01 00 00  40 01 00 00  70 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x180 in whsas.cl2 starts with the
		// following header: 0A 00 3F 01 C8 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whs/whsat.cl2": {
		// The contents of the archive header at offset 0 in whsat.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in whsat.cl2 starts with the
		// following header: 0A 00 46 01 DB 03 DB 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/warrior/whs/whsaw.cl2": {
		// The contents of the archive header at offset 0 in whsaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in whsaw.cl2 starts with the
		// following header: 0A 00 65 01 33 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whs/whsfm.cl2": {
		// The contents of the archive header at offset 0 in whsfm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in whsfm.cl2 starts with the
		// following header: 0A 00 50 01 FA 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whs/whsht.cl2": {
		// The contents of the archive header at offset 0 in whsht.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in whsht.cl2 starts with the
		// following header: 0A 00 36 01 C4 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whs/whslm.cl2": {
		// The contents of the archive header at offset 0 in whslm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in whslm.cl2 starts with the
		// following header: 0A 00 4B 01 E2 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whs/whsqm.cl2": {
		// The contents of the archive header at offset 0 in whsqm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in whsqm.cl2 starts with the
		// following header: 0A 00 B1 01 2D 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whs/whsst.cl2": {
		// The contents of the archive header at offset 0 in whsst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in whsst.cl2 starts with the
		// following header: 0A 00 49 01 78 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whs/whswl.cl2": {
		// The contents of the archive header at offset 0 in whswl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in whswl.cl2 starts with the
		// following header: 0A 00 21 01 56 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wht/whtas.cl2": {
		// The contents of the archive header at offset 0 in whtas.cl2:
		//    20 00 00 00  50 00 00 00  80 00 00 00  B0 00 00 00
		//    E0 00 00 00  10 01 00 00  40 01 00 00  70 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x180 in whtas.cl2 starts with the
		// following header: 0A 00 44 01 2C 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wht/whtat.cl2": {
		// The contents of the archive header at offset 0 in whtat.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in whtat.cl2 starts with the
		// following header: 0A 00 4C 01 11 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"plrgfx/warrior/wht/whtaw.cl2": {
		// The contents of the archive header at offset 0 in whtaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in whtaw.cl2 starts with the
		// following header: 0A 00 69 01 8A 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wht/whtfm.cl2": {
		// The contents of the archive header at offset 0 in whtfm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in whtfm.cl2 starts with the
		// following header: 0A 00 7E 01 51 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wht/whtht.cl2": {
		// The contents of the archive header at offset 0 in whtht.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in whtht.cl2 starts with the
		// following header: 0A 00 51 01 DD 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wht/whtlm.cl2": {
		// The contents of the archive header at offset 0 in whtlm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in whtlm.cl2 starts with the
		// following header: 0A 00 64 01 1B 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wht/whtqm.cl2": {
		// The contents of the archive header at offset 0 in whtqm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in whtqm.cl2 starts with the
		// following header: 0A 00 DC 01 84 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wht/whtst.cl2": {
		// The contents of the archive header at offset 0 in whtst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in whtst.cl2 starts with the
		// following header: 0A 00 B6 01 07 05 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wht/whtwl.cl2": {
		// The contents of the archive header at offset 0 in whtwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in whtwl.cl2 starts with the
		// following header: 0A 00 56 01 A2 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whu/whuas.cl2": {
		// The contents of the archive header at offset 0 in whuas.cl2:
		//    20 00 00 00  50 00 00 00  80 00 00 00  B0 00 00 00
		//    E0 00 00 00  10 01 00 00  40 01 00 00  70 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x180 in whuas.cl2 starts with the
		// following header: 0A 00 7A 01 60 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whu/whuat.cl2": {
		// The contents of the archive header at offset 0 in whuat.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in whuat.cl2 starts with the
		// following header: 0A 00 59 01 37 04 B8 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/warrior/whu/whuaw.cl2": {
		// The contents of the archive header at offset 0 in whuaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in whuaw.cl2 starts with the
		// following header: 0A 00 3F 01 E9 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whu/whubl.cl2": {
		// The contents of the archive header at offset 0 in whubl.cl2:
		//    20 00 00 00  30 00 00 00  40 00 00 00  50 00 00 00
		//    60 00 00 00  70 00 00 00  80 00 00 00  90 00 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x80 in whubl.cl2 starts with the
		// following header: 0A 00 7A 01 96 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whu/whufm.cl2": {
		// The contents of the archive header at offset 0 in whufm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in whufm.cl2 starts with the
		// following header: 0A 00 31 01 AE 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whu/whuht.cl2": {
		// The contents of the archive header at offset 0 in whuht.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in whuht.cl2 starts with the
		// following header: 0A 00 33 01 0F 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whu/whulm.cl2": {
		// The contents of the archive header at offset 0 in whulm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in whulm.cl2 starts with the
		// following header: 0A 00 56 01 30 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whu/whuqm.cl2": {
		// The contents of the archive header at offset 0 in whuqm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in whuqm.cl2 starts with the
		// following header: 0A 00 B6 01 6E 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whu/whust.cl2": {
		// The contents of the archive header at offset 0 in whust.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in whust.cl2 starts with the
		// following header: 0A 00 56 01 D1 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/whu/whuwl.cl2": {
		// The contents of the archive header at offset 0 in whuwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in whuwl.cl2 starts with the
		// following header: 0A 00 10 01 76 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wla/wlaas.cl2": {
		// The contents of the archive header at offset 0 in wlaas.cl2:
		//    20 00 00 00  50 00 00 00  80 00 00 00  B0 00 00 00
		//    E0 00 00 00  10 01 00 00  40 01 00 00  70 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x180 in wlaas.cl2 starts with the
		// following header: 0A 00 30 01 8E 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wla/wlaat.cl2": {
		// The contents of the archive header at offset 0 in wlaat.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wlaat.cl2 starts with the
		// following header: 0A 00 32 01 89 03 15 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/warrior/wla/wlaaw.cl2": {
		// The contents of the archive header at offset 0 in wlaaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in wlaaw.cl2 starts with the
		// following header: 0A 00 5D 01 F6 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wla/wlafm.cl2": {
		// The contents of the archive header at offset 0 in wlafm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wlafm.cl2 starts with the
		// following header: 0A 00 34 01 03 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wla/wlaht.cl2": {
		// The contents of the archive header at offset 0 in wlaht.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in wlaht.cl2 starts with the
		// following header: 0A 00 1C 01 94 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wla/wlalm.cl2": {
		// The contents of the archive header at offset 0 in wlalm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wlalm.cl2 starts with the
		// following header: 0A 00 24 01 87 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wla/wlaqm.cl2": {
		// The contents of the archive header at offset 0 in wlaqm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wlaqm.cl2 starts with the
		// following header: 0A 00 A2 01 3A 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wla/wlast.cl2": {
		// The contents of the archive header at offset 0 in wlast.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wlast.cl2 starts with the
		// following header: 0A 00 84 01 A1 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wla/wlawl.cl2": {
		// The contents of the archive header at offset 0 in wlawl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in wlawl.cl2 starts with the
		// following header: 0A 00 12 01 FD 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wlb/wlbas.cl2": {
		// The contents of the archive header at offset 0 in wlbas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in wlbas.cl2 starts with the
		// following header: 0A 00 1C 01 A3 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wlb/wlbat.cl2": {
		// The contents of the archive header at offset 0 in wlbat.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in wlbat.cl2 starts with the
		// following header: 0A 00 44 02 8A 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"plrgfx/warrior/wlb/wlbaw.cl2": {
		// The contents of the archive header at offset 0 in wlbaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in wlbaw.cl2 starts with the
		// following header: 0A 00 4D 01 EA 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wlb/wlbfm.cl2": {
		// The contents of the archive header at offset 0 in wlbfm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wlbfm.cl2 starts with the
		// following header: 0A 00 30 01 31 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wlb/wlbht.cl2": {
		// The contents of the archive header at offset 0 in wlbht.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in wlbht.cl2 starts with the
		// following header: 0A 00 24 01 99 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wlb/wlblm.cl2": {
		// The contents of the archive header at offset 0 in wlblm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wlblm.cl2 starts with the
		// following header: 0A 00 2A 01 CA 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wlb/wlbqm.cl2": {
		// The contents of the archive header at offset 0 in wlbqm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wlbqm.cl2 starts with the
		// following header: 0A 00 BC 01 8C 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wlb/wlbst.cl2": {
		// The contents of the archive header at offset 0 in wlbst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wlbst.cl2 starts with the
		// following header: 0A 00 62 01 80 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wlb/wlbwl.cl2": {
		// The contents of the archive header at offset 0 in wlbwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in wlbwl.cl2 starts with the
		// following header: 0A 00 F7 00 E1 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wld/wldas.cl2": {
		// The contents of the archive header at offset 0 in wldas.cl2:
		//    20 00 00 00  50 00 00 00  80 00 00 00  B0 00 00 00
		//    E0 00 00 00  10 01 00 00  40 01 00 00  70 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x180 in wldas.cl2 starts with the
		// following header: 0A 00 68 01 8E 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wld/wldat.cl2": {
		// The contents of the archive header at offset 0 in wldat.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in wldat.cl2 starts with the
		// following header: 0A 00 76 01 E4 04 9B 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/warrior/wld/wldaw.cl2": {
		// The contents of the archive header at offset 0 in wldaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in wldaw.cl2 starts with the
		// following header: 0A 00 59 01 C7 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wld/wldbl.cl2": {
		// The contents of the archive header at offset 0 in wldbl.cl2:
		//    20 00 00 00  30 00 00 00  40 00 00 00  50 00 00 00
		//    60 00 00 00  70 00 00 00  80 00 00 00  90 00 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x80 in wldbl.cl2 starts with the
		// following header: 0A 00 77 01 BD 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wld/wldfm.cl2": {
		// The contents of the archive header at offset 0 in wldfm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wldfm.cl2 starts with the
		// following header: 0A 00 73 01 E1 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wld/wldht.cl2": {
		// The contents of the archive header at offset 0 in wldht.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in wldht.cl2 starts with the
		// following header: 0A 00 62 01 95 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wld/wldlm.cl2": {
		// The contents of the archive header at offset 0 in wldlm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wldlm.cl2 starts with the
		// following header: 0A 00 69 01 D2 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wld/wldqm.cl2": {
		// The contents of the archive header at offset 0 in wldqm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wldqm.cl2 starts with the
		// following header: 0A 00 D9 01 4C 05 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wld/wldst.cl2": {
		// The contents of the archive header at offset 0 in wldst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wldst.cl2 starts with the
		// following header: 0A 00 52 01 8B 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wld/wldwl.cl2": {
		// The contents of the archive header at offset 0 in wldwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in wldwl.cl2 starts with the
		// following header: 0A 00 FC 00 0C 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wlh/wlhas.cl2": {
		// The contents of the archive header at offset 0 in wlhas.cl2:
		//    20 00 00 00  50 00 00 00  80 00 00 00  B0 00 00 00
		//    E0 00 00 00  10 01 00 00  40 01 00 00  70 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x180 in wlhas.cl2 starts with the
		// following header: 0A 00 21 01 4A 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wlh/wlhat.cl2": {
		// The contents of the archive header at offset 0 in wlhat.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in wlhat.cl2 starts with the
		// following header: 0A 00 1F 01 50 04 F6 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/warrior/wlh/wlhaw.cl2": {
		// The contents of the archive header at offset 0 in wlhaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in wlhaw.cl2 starts with the
		// following header: 0A 00 49 01 84 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wlh/wlhbl.cl2": {
		// The contents of the archive header at offset 0 in wlhbl.cl2:
		//    20 00 00 00  30 00 00 00  40 00 00 00  50 00 00 00
		//    60 00 00 00  70 00 00 00  80 00 00 00  90 00 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x80 in wlhbl.cl2 starts with the
		// following header: 0A 00 09 01 40 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wlh/wlhfm.cl2": {
		// The contents of the archive header at offset 0 in wlhfm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wlhfm.cl2 starts with the
		// following header: 0A 00 2B 01 7E 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wlh/wlhht.cl2": {
		// The contents of the archive header at offset 0 in wlhht.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in wlhht.cl2 starts with the
		// following header: 0A 00 16 01 F2 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wlh/wlhlm.cl2": {
		// The contents of the archive header at offset 0 in wlhlm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wlhlm.cl2 starts with the
		// following header: 0A 00 28 01 4E 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wlh/wlhqm.cl2": {
		// The contents of the archive header at offset 0 in wlhqm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wlhqm.cl2 starts with the
		// following header: 0A 00 CB 01 2C 05 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wlh/wlhst.cl2": {
		// The contents of the archive header at offset 0 in wlhst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wlhst.cl2 starts with the
		// following header: 0A 00 68 01 C6 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wlh/wlhwl.cl2": {
		// The contents of the archive header at offset 0 in wlhwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in wlhwl.cl2 starts with the
		// following header: 0A 00 18 01 22 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wlm/wlmas.cl2": {
		// The contents of the archive header at offset 0 in wlmas.cl2:
		//    20 00 00 00  50 00 00 00  80 00 00 00  B0 00 00 00
		//    E0 00 00 00  10 01 00 00  40 01 00 00  70 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x180 in wlmas.cl2 starts with the
		// following header: 0A 00 26 01 09 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wlm/wlmat.cl2": {
		// The contents of the archive header at offset 0 in wlmat.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in wlmat.cl2 starts with the
		// following header: 0A 00 26 01 FF 02 A4 03 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/warrior/wlm/wlmaw.cl2": {
		// The contents of the archive header at offset 0 in wlmaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in wlmaw.cl2 starts with the
		// following header: 0A 00 4E 01 98 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wlm/wlmfm.cl2": {
		// The contents of the archive header at offset 0 in wlmfm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wlmfm.cl2 starts with the
		// following header: 0A 00 3A 01 61 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wlm/wlmht.cl2": {
		// The contents of the archive header at offset 0 in wlmht.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in wlmht.cl2 starts with the
		// following header: 0A 00 16 01 14 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wlm/wlmlm.cl2": {
		// The contents of the archive header at offset 0 in wlmlm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wlmlm.cl2 starts with the
		// following header: 0A 00 20 01 07 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wlm/wlmqm.cl2": {
		// The contents of the archive header at offset 0 in wlmqm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wlmqm.cl2 starts with the
		// following header: 0A 00 A3 01 98 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wlm/wlmst.cl2": {
		// The contents of the archive header at offset 0 in wlmst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wlmst.cl2 starts with the
		// following header: 0A 00 62 01 8C 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wlm/wlmwl.cl2": {
		// The contents of the archive header at offset 0 in wlmwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in wlmwl.cl2 starts with the
		// following header: 0A 00 10 01 D2 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wln/wlnas.cl2": {
		// The contents of the archive header at offset 0 in wlnas.cl2:
		//    20 00 00 00  50 00 00 00  80 00 00 00  B0 00 00 00
		//    E0 00 00 00  10 01 00 00  40 01 00 00  70 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x180 in wlnas.cl2 starts with the
		// following header: 0A 00 0F 01 ED 02 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wln/wlnat.cl2": {
		// The contents of the archive header at offset 0 in wlnat.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in wlnat.cl2 starts with the
		// following header: 0A 00 15 01 E8 02 4C 03 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/warrior/wln/wlnaw.cl2": {
		// The contents of the archive header at offset 0 in wlnaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in wlnaw.cl2 starts with the
		// following header: 0A 00 47 01 EF 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wln/wlndt.cl2": {
		// The contents of the archive header at offset 0 in wlndt.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wlndt.cl2 starts with the
		// following header: 0A 00 0A 01 CE 02 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"plrgfx/warrior/wln/wlnfm.cl2": {
		// The contents of the archive header at offset 0 in wlnfm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wlnfm.cl2 starts with the
		// following header: 0A 00 15 01 F6 02 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wln/wlnht.cl2": {
		// The contents of the archive header at offset 0 in wlnht.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in wlnht.cl2 starts with the
		// following header: 0A 00 14 01 D0 02 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wln/wlnlm.cl2": {
		// The contents of the archive header at offset 0 in wlnlm.cl2:
		//    20 00 00 00  7C 00 00 00  D8 00 00 00  34 01 00 00
		//    90 01 00 00  EC 01 00 00  48 02 00 00  A4 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2E0 in wlnlm.cl2 starts with the
		// following header: 0A 00 1C 01 D5 02 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wln/wlnqm.cl2": {
		// The contents of the archive header at offset 0 in wlnqm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wlnqm.cl2 starts with the
		// following header: 0A 00 84 01 26 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wln/wlnst.cl2": {
		// The contents of the archive header at offset 0 in wlnst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wlnst.cl2 starts with the
		// following header: 0A 00 5B 01 3E 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wln/wlnwl.cl2": {
		// The contents of the archive header at offset 0 in wlnwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in wlnwl.cl2 starts with the
		// following header: 0A 00 EC 00 97 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wls/wlsas.cl2": {
		// The contents of the archive header at offset 0 in wlsas.cl2:
		//    20 00 00 00  50 00 00 00  80 00 00 00  B0 00 00 00
		//    E0 00 00 00  10 01 00 00  40 01 00 00  70 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x180 in wlsas.cl2 starts with the
		// following header: 0A 00 71 01 94 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wls/wlsat.cl2": {
		// The contents of the archive header at offset 0 in wlsat.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in wlsat.cl2 starts with the
		// following header: 0A 00 6B 01 8E 03 43 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/warrior/wls/wlsaw.cl2": {
		// The contents of the archive header at offset 0 in wlsaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in wlsaw.cl2 starts with the
		// following header: 0A 00 60 01 19 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wls/wlsfm.cl2": {
		// The contents of the archive header at offset 0 in wlsfm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wlsfm.cl2 starts with the
		// following header: 0A 00 72 01 BB 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wls/wlsht.cl2": {
		// The contents of the archive header at offset 0 in wlsht.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in wlsht.cl2 starts with the
		// following header: 0A 00 5F 01 5D 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wls/wlslm.cl2": {
		// The contents of the archive header at offset 0 in wlslm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wlslm.cl2 starts with the
		// following header: 0A 00 71 01 9B 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wls/wlsqm.cl2": {
		// The contents of the archive header at offset 0 in wlsqm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wlsqm.cl2 starts with the
		// following header: 0A 00 DD 01 07 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wls/wlsst.cl2": {
		// The contents of the archive header at offset 0 in wlsst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wlsst.cl2 starts with the
		// following header: 0A 00 58 01 5F 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wls/wlswl.cl2": {
		// The contents of the archive header at offset 0 in wlswl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in wlswl.cl2 starts with the
		// following header: 0A 00 EB 00 C0 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wlt/wltas.cl2": {
		// The contents of the archive header at offset 0 in wltas.cl2:
		//    20 00 00 00  50 00 00 00  80 00 00 00  B0 00 00 00
		//    E0 00 00 00  10 01 00 00  40 01 00 00  70 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x180 in wltas.cl2 starts with the
		// following header: 0A 00 48 01 90 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wlt/wltat.cl2": {
		// The contents of the archive header at offset 0 in wltat.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in wltat.cl2 starts with the
		// following header: 0A 00 53 01 A8 03 39 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/warrior/wlt/wltaw.cl2": {
		// The contents of the archive header at offset 0 in wltaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in wltaw.cl2 starts with the
		// following header: 0A 00 52 01 23 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wlt/wltfm.cl2": {
		// The contents of the archive header at offset 0 in wltfm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wltfm.cl2 starts with the
		// following header: 0A 00 7D 01 19 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wlt/wltht.cl2": {
		// The contents of the archive header at offset 0 in wltht.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in wltht.cl2 starts with the
		// following header: 0A 00 46 01 83 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wlt/wltlm.cl2": {
		// The contents of the archive header at offset 0 in wltlm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wltlm.cl2 starts with the
		// following header: 0A 00 72 01 C5 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wlt/wltqm.cl2": {
		// The contents of the archive header at offset 0 in wltqm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wltqm.cl2 starts with the
		// following header: 0A 00 EF 01 50 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wlt/wltst.cl2": {
		// The contents of the archive header at offset 0 in wltst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wltst.cl2 starts with the
		// following header: 0A 00 B7 01 AF 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wlt/wltwl.cl2": {
		// The contents of the archive header at offset 0 in wltwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in wltwl.cl2 starts with the
		// following header: 0A 00 48 01 43 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wlu/wluas.cl2": {
		// The contents of the archive header at offset 0 in wluas.cl2:
		//    20 00 00 00  50 00 00 00  80 00 00 00  B0 00 00 00
		//    E0 00 00 00  10 01 00 00  40 01 00 00  70 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x180 in wluas.cl2 starts with the
		// following header: 0A 00 1E 01 30 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wlu/wluat.cl2": {
		// The contents of the archive header at offset 0 in wluat.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in wluat.cl2 starts with the
		// following header: 0A 00 1A 01 41 04 A2 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/warrior/wlu/wluaw.cl2": {
		// The contents of the archive header at offset 0 in wluaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in wluaw.cl2 starts with the
		// following header: 0A 00 37 01 66 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wlu/wlubl.cl2": {
		// The contents of the archive header at offset 0 in wlubl.cl2:
		//    20 00 00 00  30 00 00 00  40 00 00 00  50 00 00 00
		//    60 00 00 00  70 00 00 00  80 00 00 00  90 00 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x80 in wlubl.cl2 starts with the
		// following header: 0A 00 06 01 2A 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wlu/wlufm.cl2": {
		// The contents of the archive header at offset 0 in wlufm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wlufm.cl2 starts with the
		// following header: 0A 00 24 01 6D 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wlu/wluht.cl2": {
		// The contents of the archive header at offset 0 in wluht.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in wluht.cl2 starts with the
		// following header: 0A 00 0F 01 D4 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wlu/wlulm.cl2": {
		// The contents of the archive header at offset 0 in wlulm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wlulm.cl2 starts with the
		// following header: 0A 00 29 01 54 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wlu/wluqm.cl2": {
		// The contents of the archive header at offset 0 in wluqm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wluqm.cl2 starts with the
		// following header: 0A 00 A9 01 D8 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wlu/wlust.cl2": {
		// The contents of the archive header at offset 0 in wlust.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wlust.cl2 starts with the
		// following header: 0A 00 60 01 8B 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wlu/wluwl.cl2": {
		// The contents of the archive header at offset 0 in wluwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in wluwl.cl2 starts with the
		// following header: 0A 00 F1 00 D1 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wma/wmaas.cl2": {
		// The contents of the archive header at offset 0 in wmaas.cl2:
		//    20 00 00 00  50 00 00 00  80 00 00 00  B0 00 00 00
		//    E0 00 00 00  10 01 00 00  40 01 00 00  70 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x180 in wmaas.cl2 starts with the
		// following header: 0A 00 45 01 19 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wma/wmaat.cl2": {
		// The contents of the archive header at offset 0 in wmaat.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wmaat.cl2 starts with the
		// following header: 0A 00 45 01 2A 04 DB 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/warrior/wma/wmaaw.cl2": {
		// The contents of the archive header at offset 0 in wmaaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in wmaaw.cl2 starts with the
		// following header: 0A 00 64 01 3F 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wma/wmafm.cl2": {
		// The contents of the archive header at offset 0 in wmafm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wmafm.cl2 starts with the
		// following header: 0A 00 49 01 5E 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wma/wmaht.cl2": {
		// The contents of the archive header at offset 0 in wmaht.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in wmaht.cl2 starts with the
		// following header: 0A 00 47 01 32 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wma/wmalm.cl2": {
		// The contents of the archive header at offset 0 in wmalm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wmalm.cl2 starts with the
		// following header: 0A 00 49 01 35 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wma/wmaqm.cl2": {
		// The contents of the archive header at offset 0 in wmaqm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wmaqm.cl2 starts with the
		// following header: 0A 00 B2 01 8C 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wma/wmast.cl2": {
		// The contents of the archive header at offset 0 in wmast.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wmast.cl2 starts with the
		// following header: 0A 00 7A 01 EA 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wma/wmawl.cl2": {
		// The contents of the archive header at offset 0 in wmawl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in wmawl.cl2 starts with the
		// following header: 0A 00 2C 01 54 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmb/wmbas.cl2": {
		// The contents of the archive header at offset 0 in wmbas.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in wmbas.cl2 starts with the
		// following header: 0A 00 1E 01 FD 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmb/wmbat.cl2": {
		// The contents of the archive header at offset 0 in wmbat.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in wmbat.cl2 starts with the
		// following header: 0A 00 85 02 0A 05 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"plrgfx/warrior/wmb/wmbaw.cl2": {
		// The contents of the archive header at offset 0 in wmbaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in wmbaw.cl2 starts with the
		// following header: 0A 00 52 01 24 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmb/wmbfm.cl2": {
		// The contents of the archive header at offset 0 in wmbfm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wmbfm.cl2 starts with the
		// following header: 0A 00 29 01 6F 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmb/wmbht.cl2": {
		// The contents of the archive header at offset 0 in wmbht.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in wmbht.cl2 starts with the
		// following header: 0A 00 40 01 08 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmb/wmblm.cl2": {
		// The contents of the archive header at offset 0 in wmblm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wmblm.cl2 starts with the
		// following header: 0A 00 34 01 62 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmb/wmbqm.cl2": {
		// The contents of the archive header at offset 0 in wmbqm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wmbqm.cl2 starts with the
		// following header: 0A 00 C0 01 C3 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmb/wmbst.cl2": {
		// The contents of the archive header at offset 0 in wmbst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wmbst.cl2 starts with the
		// following header: 0A 00 52 01 B3 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmb/wmbwl.cl2": {
		// The contents of the archive header at offset 0 in wmbwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in wmbwl.cl2 starts with the
		// following header: 0A 00 08 01 37 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmd/wmdas.cl2": {
		// The contents of the archive header at offset 0 in wmdas.cl2:
		//    20 00 00 00  50 00 00 00  80 00 00 00  B0 00 00 00
		//    E0 00 00 00  10 01 00 00  40 01 00 00  70 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x180 in wmdas.cl2 starts with the
		// following header: 0A 00 A3 01 12 05 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmd/wmdat.cl2": {
		// The contents of the archive header at offset 0 in wmdat.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in wmdat.cl2 starts with the
		// following header: 0A 00 9F 01 FD 04 DF 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/warrior/wmd/wmdaw.cl2": {
		// The contents of the archive header at offset 0 in wmdaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in wmdaw.cl2 starts with the
		// following header: 0A 00 7B 01 8F 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmd/wmdbl.cl2": {
		// The contents of the archive header at offset 0 in wmdbl.cl2:
		//    20 00 00 00  30 00 00 00  40 00 00 00  50 00 00 00
		//    60 00 00 00  70 00 00 00  80 00 00 00  90 00 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x80 in wmdbl.cl2 starts with the
		// following header: 0A 00 E3 01 DC 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmd/wmdfm.cl2": {
		// The contents of the archive header at offset 0 in wmdfm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wmdfm.cl2 starts with the
		// following header: 0A 00 CF 01 10 05 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmd/wmdht.cl2": {
		// The contents of the archive header at offset 0 in wmdht.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in wmdht.cl2 starts with the
		// following header: 0A 00 7F 01 CF 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmd/wmdlm.cl2": {
		// The contents of the archive header at offset 0 in wmdlm.cl2:
		//    20 00 00 00  7C 00 00 00  D8 00 00 00  34 01 00 00
		//    90 01 00 00  EC 01 00 00  48 02 00 00  A4 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2E0 in wmdlm.cl2 starts with the
		// following header: 0A 00 B3 01 13 05 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmd/wmdqm.cl2": {
		// The contents of the archive header at offset 0 in wmdqm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wmdqm.cl2 starts with the
		// following header: 0A 00 2F 02 4A 05 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmd/wmdst.cl2": {
		// The contents of the archive header at offset 0 in wmdst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wmdst.cl2 starts with the
		// following header: 0A 00 5B 01 C5 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmd/wmdwl.cl2": {
		// The contents of the archive header at offset 0 in wmdwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in wmdwl.cl2 starts with the
		// following header: 0A 00 1A 01 6A 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmh/wmhas.cl2": {
		// The contents of the archive header at offset 0 in wmhas.cl2:
		//    20 00 00 00  50 00 00 00  80 00 00 00  B0 00 00 00
		//    E0 00 00 00  10 01 00 00  40 01 00 00  70 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x180 in wmhas.cl2 starts with the
		// following header: 0A 00 6E 01 88 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmh/wmhat.cl2": {
		// The contents of the archive header at offset 0 in wmhat.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in wmhat.cl2 starts with the
		// following header: 0A 00 3F 01 19 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"plrgfx/warrior/wmh/wmhaw.cl2": {
		// The contents of the archive header at offset 0 in wmhaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in wmhaw.cl2 starts with the
		// following header: 0A 00 78 01 41 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmh/wmhbl.cl2": {
		// The contents of the archive header at offset 0 in wmhbl.cl2:
		//    20 00 00 00  30 00 00 00  40 00 00 00  50 00 00 00
		//    60 00 00 00  70 00 00 00  80 00 00 00  90 00 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x80 in wmhbl.cl2 starts with the
		// following header: 0A 00 72 01 63 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmh/wmhfm.cl2": {
		// The contents of the archive header at offset 0 in wmhfm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wmhfm.cl2 starts with the
		// following header: 0A 00 68 01 89 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmh/wmhht.cl2": {
		// The contents of the archive header at offset 0 in wmhht.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in wmhht.cl2 starts with the
		// following header: 0A 00 2F 01 31 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmh/wmhlm.cl2": {
		// The contents of the archive header at offset 0 in wmhlm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wmhlm.cl2 starts with the
		// following header: 0A 00 64 01 74 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmh/wmhqm.cl2": {
		// The contents of the archive header at offset 0 in wmhqm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wmhqm.cl2 starts with the
		// following header: 0A 00 DF 01 DF 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmh/wmhst.cl2": {
		// The contents of the archive header at offset 0 in wmhst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wmhst.cl2 starts with the
		// following header: 0A 00 6B 01 F7 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmh/wmhwl.cl2": {
		// The contents of the archive header at offset 0 in wmhwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in wmhwl.cl2 starts with the
		// following header: 0A 00 3F 01 73 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmm/wmmas.cl2": {
		// The contents of the archive header at offset 0 in wmmas.cl2:
		//    20 00 00 00  50 00 00 00  80 00 00 00  B0 00 00 00
		//    E0 00 00 00  10 01 00 00  40 01 00 00  70 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x180 in wmmas.cl2 starts with the
		// following header: 0A 00 33 01 90 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmm/wmmat.cl2": {
		// The contents of the archive header at offset 0 in wmmat.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in wmmat.cl2 starts with the
		// following header: 0A 00 33 01 8F 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"plrgfx/warrior/wmm/wmmaw.cl2": {
		// The contents of the archive header at offset 0 in wmmaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in wmmaw.cl2 starts with the
		// following header: 0A 00 5F 01 E2 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmm/wmmfm.cl2": {
		// The contents of the archive header at offset 0 in wmmfm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wmmfm.cl2 starts with the
		// following header: 0A 00 4B 01 C1 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmm/wmmht.cl2": {
		// The contents of the archive header at offset 0 in wmmht.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in wmmht.cl2 starts with the
		// following header: 0A 00 2A 01 98 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmm/wmmlm.cl2": {
		// The contents of the archive header at offset 0 in wmmlm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wmmlm.cl2 starts with the
		// following header: 0A 00 38 01 9A 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmm/wmmqm.cl2": {
		// The contents of the archive header at offset 0 in wmmqm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wmmqm.cl2 starts with the
		// following header: 0A 00 B2 01 E6 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmm/wmmst.cl2": {
		// The contents of the archive header at offset 0 in wmmst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wmmst.cl2 starts with the
		// following header: 0A 00 54 01 C0 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmm/wmmwl.cl2": {
		// The contents of the archive header at offset 0 in wmmwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in wmmwl.cl2 starts with the
		// following header: 0A 00 29 01 2F 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmn/wmnas.cl2": {
		// The contents of the archive header at offset 0 in wmnas.cl2:
		//    20 00 00 00  50 00 00 00  80 00 00 00  B0 00 00 00
		//    E0 00 00 00  10 01 00 00  40 01 00 00  70 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x180 in wmnas.cl2 starts with the
		// following header: 0A 00 2B 01 3B 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmn/wmnat.cl2": {
		// The contents of the archive header at offset 0 in wmnat.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in wmnat.cl2 starts with the
		// following header: 0A 00 36 01 75 03 FA 03 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/warrior/wmn/wmnaw.cl2": {
		// The contents of the archive header at offset 0 in wmnaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in wmnaw.cl2 starts with the
		// following header: 0A 00 61 01 4A 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmn/wmndt.cl2": {
		// The contents of the archive header at offset 0 in wmndt.cl2:
		//    20 00 00 00  64 00 00 00  A8 00 00 00  EC 00 00 00
		//    30 01 00 00  74 01 00 00  B8 01 00 00  FC 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x220 in wmndt.cl2 starts with the
		// following header: 0A 00 CF 01 C7 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"plrgfx/warrior/wmn/wmnfm.cl2": {
		// The contents of the archive header at offset 0 in wmnfm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wmnfm.cl2 starts with the
		// following header: 0A 00 34 01 52 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmn/wmnht.cl2": {
		// The contents of the archive header at offset 0 in wmnht.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in wmnht.cl2 starts with the
		// following header: 0A 00 31 01 6E 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmn/wmnlm.cl2": {
		// The contents of the archive header at offset 0 in wmnlm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wmnlm.cl2 starts with the
		// following header: 0A 00 2E 01 43 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmn/wmnqm.cl2": {
		// The contents of the archive header at offset 0 in wmnqm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wmnqm.cl2 starts with the
		// following header: 0A 00 A8 01 A8 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmn/wmnst.cl2": {
		// The contents of the archive header at offset 0 in wmnst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wmnst.cl2 starts with the
		// following header: 0A 00 4D 01 85 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmn/wmnwl.cl2": {
		// The contents of the archive header at offset 0 in wmnwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in wmnwl.cl2 starts with the
		// following header: 0A 00 02 01 ED 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wms/wmsas.cl2": {
		// The contents of the archive header at offset 0 in wmsas.cl2:
		//    20 00 00 00  50 00 00 00  80 00 00 00  B0 00 00 00
		//    E0 00 00 00  10 01 00 00  40 01 00 00  70 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x180 in wmsas.cl2 starts with the
		// following header: 0A 00 90 01 09 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wms/wmsat.cl2": {
		// The contents of the archive header at offset 0 in wmsat.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in wmsat.cl2 starts with the
		// following header: 0A 00 8F 01 C8 03 A8 04 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/warrior/wms/wmsaw.cl2": {
		// The contents of the archive header at offset 0 in wmsaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in wmsaw.cl2 starts with the
		// following header: 0A 00 76 01 55 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wms/wmsfm.cl2": {
		// The contents of the archive header at offset 0 in wmsfm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wmsfm.cl2 starts with the
		// following header: 0A 00 95 01 3E 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wms/wmsht.cl2": {
		// The contents of the archive header at offset 0 in wmsht.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in wmsht.cl2 starts with the
		// following header: 0A 00 85 01 27 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wms/wmslm.cl2": {
		// The contents of the archive header at offset 0 in wmslm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wmslm.cl2 starts with the
		// following header: 0A 00 8F 01 1F 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wms/wmsqm.cl2": {
		// The contents of the archive header at offset 0 in wmsqm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wmsqm.cl2 starts with the
		// following header: 0A 00 F3 01 5F 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wms/wmsst.cl2": {
		// The contents of the archive header at offset 0 in wmsst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wmsst.cl2 starts with the
		// following header: 0A 00 44 01 7C 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wms/wmswl.cl2": {
		// The contents of the archive header at offset 0 in wmswl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in wmswl.cl2 starts with the
		// following header: 0A 00 13 01 2F 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmt/wmtas.cl2": {
		// The contents of the archive header at offset 0 in wmtas.cl2:
		//    20 00 00 00  50 00 00 00  80 00 00 00  B0 00 00 00
		//    E0 00 00 00  10 01 00 00  40 01 00 00  70 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x180 in wmtas.cl2 starts with the
		// following header: 0A 00 60 01 3B 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmt/wmtat.cl2": {
		// The contents of the archive header at offset 0 in wmtat.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in wmtat.cl2 starts with the
		// following header: 0A 00 60 01 38 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"plrgfx/warrior/wmt/wmtaw.cl2": {
		// The contents of the archive header at offset 0 in wmtaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in wmtaw.cl2 starts with the
		// following header: 0A 00 66 01 7C 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmt/wmtfm.cl2": {
		// The contents of the archive header at offset 0 in wmtfm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wmtfm.cl2 starts with the
		// following header: 0A 00 62 01 4A 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmt/wmtht.cl2": {
		// The contents of the archive header at offset 0 in wmtht.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in wmtht.cl2 starts with the
		// following header: 0A 00 58 01 02 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmt/wmtlm.cl2": {
		// The contents of the archive header at offset 0 in wmtlm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wmtlm.cl2 starts with the
		// following header: 0A 00 6B 01 44 04 11 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 12288/unknown = baz
	},
	"plrgfx/warrior/wmt/wmtqm.cl2": {
		// The contents of the archive header at offset 0 in wmtqm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wmtqm.cl2 starts with the
		// following header: 0A 00 F0 01 A8 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmt/wmtst.cl2": {
		// The contents of the archive header at offset 0 in wmtst.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wmtst.cl2 starts with the
		// following header: 0A 00 A7 01 DF 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmt/wmtwl.cl2": {
		// The contents of the archive header at offset 0 in wmtwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in wmtwl.cl2 starts with the
		// following header: 0A 00 65 01 9B 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmu/wmuas.cl2": {
		// The contents of the archive header at offset 0 in wmuas.cl2:
		//    20 00 00 00  50 00 00 00  80 00 00 00  B0 00 00 00
		//    E0 00 00 00  10 01 00 00  40 01 00 00  70 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x180 in wmuas.cl2 starts with the
		// following header: 0A 00 74 01 8E 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmu/wmuat.cl2": {
		// The contents of the archive header at offset 0 in wmuat.cl2:
		//    20 00 00 00  68 00 00 00  B0 00 00 00  F8 00 00 00
		//    40 01 00 00  88 01 00 00  D0 01 00 00  18 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x240 in wmuat.cl2 starts with the
		// following header: 0A 00 6B 01 81 04 06 05 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 16384/unknown = baz
	},
	"plrgfx/warrior/wmu/wmuaw.cl2": {
		// The contents of the archive header at offset 0 in wmuaw.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in wmuaw.cl2 starts with the
		// following header: 0A 00 57 01 49 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmu/wmubl.cl2": {
		// The contents of the archive header at offset 0 in wmubl.cl2:
		//    20 00 00 00  30 00 00 00  40 00 00 00  50 00 00 00
		//    60 00 00 00  70 00 00 00  80 00 00 00  90 00 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x80 in wmubl.cl2 starts with the
		// following header: 0A 00 88 01 7A 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmu/wmufm.cl2": {
		// The contents of the archive header at offset 0 in wmufm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wmufm.cl2 starts with the
		// following header: 0A 00 53 01 B6 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmu/wmuht.cl2": {
		// The contents of the archive header at offset 0 in wmuht.cl2:
		//    20 00 00 00  40 00 00 00  60 00 00 00  80 00 00 00
		//    A0 00 00 00  C0 00 00 00  E0 00 00 00  00 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x100 in wmuht.cl2 starts with the
		// following header: 0A 00 31 01 0C 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmu/wmulm.cl2": {
		// The contents of the archive header at offset 0 in wmulm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wmulm.cl2 starts with the
		// following header: 0A 00 62 01 C0 03 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmu/wmuqm.cl2": {
		// The contents of the archive header at offset 0 in wmuqm.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wmuqm.cl2 starts with the
		// following header: 0A 00 C4 01 A8 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmu/wmust.cl2": {
		// The contents of the archive header at offset 0 in wmust.cl2:
		//    20 00 00 00  78 00 00 00  D0 00 00 00  28 01 00 00
		//    80 01 00 00  D8 01 00 00  30 02 00 00  88 02 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x2C0 in wmust.cl2 starts with the
		// following header: 0A 00 5B 01 AE 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
	"plrgfx/warrior/wmu/wmuwl.cl2": {
		// The contents of the archive header at offset 0 in wmuwl.cl2:
		//    20 00 00 00  48 00 00 00  70 00 00 00  98 00 00 00
		//    C0 00 00 00  E8 00 00 00  10 01 00 00  38 01 00 00
		Nimgs: 8,
		// The contents of frame 0 at offset 0x140 in wmuwl.cl2 starts with the
		// following header: 0A 00 1A 01 38 04 00 00 00 00.
		Header: 10,
		W:      unknown,
		H:      unknown, // h = npixels/w = 9216/unknown = baz
	},
}

// RelPaths maps from CEL file names to "diabdat.mpq" relative paths.
var RelPaths = map[string]string{
	// CEL files.
	"golddrop.cel": "ctrlpan/golddrop.cel",
	"p8bulbs.cel":  "ctrlpan/p8bulbs.cel",
	"p8but2.cel":   "ctrlpan/p8but2.cel",
	"panel8.cel":   "ctrlpan/panel8.cel",
	"panel8bu.cel": "ctrlpan/panel8bu.cel",
	"smaltext.cel": "ctrlpan/smaltext.cel",
	"spelicon.cel": "ctrlpan/spelicon.cel",
	"talkbutt.cel": "ctrlpan/talkbutt.cel",
	"talkpanl.cel": "ctrlpan/talkpanl.cel",
	"bigtgold.cel": "data/bigtgold.cel",
	"char.cel":     "data/char.cel",
	"charbut.cel":  "data/charbut.cel",
	"diabsmal.cel": "data/diabsmal.cel",
	"inv.cel":      "data/inv/inv.cel",
	"inv_rog.cel":  "data/inv/inv_rog.cel",
	"inv_sor.cel":  "data/inv/inv_sor.cel",
	"objcurs.cel":  "data/inv/objcurs.cel",
	"medtexts.cel": "data/medtexts.cel",
	"optbar.cel":   "data/optbar.cel",
	"option.cel":   "data/option.cel",
	"pentspin.cel": "data/pentspin.cel",
	"pentspn2.cel": "data/pentspn2.cel",
	"quest.cel":    "data/quest.cel",
	"spellbk.cel":  "data/spellbk.cel",
	"spellbkb.cel": "data/spellbkb.cel",
	"spelli2.cel":  "data/spelli2.cel",
	"square.cel":   "data/square.cel",
	"textbox.cel":  "data/textbox.cel",
	"textbox2.cel": "data/textbox2.cel",
	"textslid.cel": "data/textslid.cel",
	"cut2.cel":     "gendata/cut2.cel",
	"cut3.cel":     "gendata/cut3.cel",
	"cut4.cel":     "gendata/cut4.cel",
	"cutgate.cel":  "gendata/cutgate.cel",
	"cutl1d.cel":   "gendata/cutl1d.cel",
	"cutportl.cel": "gendata/cutportl.cel",
	"cutportr.cel": "gendata/cutportr.cel",
	"cutstart.cel": "gendata/cutstart.cel",
	"cuttt.cel":    "gendata/cuttt.cel",
	"quotes.cel":   "gendata/quotes.cel",
	"armor2.cel":   "items/armor2.cel",
	"axe.cel":      "items/axe.cel",
	"axeflip.cel":  "items/axeflip.cel",
	"bldstn.cel":   "items/bldstn.cel",
	"bottle.cel":   "items/bottle.cel",
	"bow.cel":      "items/bow.cel",
	"cleaver.cel":  "items/cleaver.cel",
	"crownf.cel":   "items/crownf.cel",
	"duricons.cel": "items/duricons.cel",
	"fanvil.cel":   "items/fanvil.cel",
	"fbook.cel":    "items/fbook.cel",
	"fbow.cel":     "items/fbow.cel",
	"fbrain.cel":   "items/fbrain.cel",
	"fbttle.cel":   "items/fbttle.cel",
	"fbttlebb.cel": "items/fbttlebb.cel",
	"fbttlebl.cel": "items/fbttlebl.cel",
	"fbttlebr.cel": "items/fbttlebr.cel",
	"fbttleby.cel": "items/fbttleby.cel",
	"fbttledb.cel": "items/fbttledb.cel",
	"fbttledy.cel": "items/fbttledy.cel",
	"fbttleor.cel": "items/fbttleor.cel",
	"fbttlewh.cel": "items/fbttlewh.cel",
	"fear.cel":     "items/fear.cel",
	"feye.cel":     "items/feye.cel",
	"fheart.cel":   "items/fheart.cel",
	"flazstaf.cel": "items/flazstaf.cel",
	"fmush.cel":    "items/fmush.cel",
	"food.cel":     "items/food.cel",
	"fplatear.cel": "items/fplatear.cel",
	"goldflip.cel": "items/goldflip.cel",
	"helmut.cel":   "items/helmut.cel",
	"innsign.cel":  "items/innsign.cel",
	"larmor.cel":   "items/larmor.cel",
	"mace.cel":     "items/mace.cel",
	"manaflip.cel": "items/manaflip.cel",
	"mapz0000.cel": "items/map/mapz0000.cel",
	"mapz0001.cel": "items/map/mapz0001.cel",
	"mapz0002.cel": "items/map/mapz0002.cel",
	"mapz0003.cel": "items/map/mapz0003.cel",
	"mapz0004.cel": "items/map/mapz0004.cel",
	"mapz0005.cel": "items/map/mapz0005.cel",
	"mapz0006.cel": "items/map/mapz0006.cel",
	"mapz0007.cel": "items/map/mapz0007.cel",
	"mapz0008.cel": "items/map/mapz0008.cel",
	"mapz0009.cel": "items/map/mapz0009.cel",
	"mapz0010.cel": "items/map/mapz0010.cel",
	"mapz0011.cel": "items/map/mapz0011.cel",
	"mapz0012.cel": "items/map/mapz0012.cel",
	"mapz0013.cel": "items/map/mapz0013.cel",
	"mapz0014.cel": "items/map/mapz0014.cel",
	"mapz0015.cel": "items/map/mapz0015.cel",
	"mapz0016.cel": "items/map/mapz0016.cel",
	"mapz0017.cel": "items/map/mapz0017.cel",
	"mapz0018.cel": "items/map/mapz0018.cel",
	"mapz0019.cel": "items/map/mapz0019.cel",
	"mapz0020.cel": "items/map/mapz0020.cel",
	"mapz0021.cel": "items/map/mapz0021.cel",
	"mapz0022.cel": "items/map/mapz0022.cel",
	"mapz0023.cel": "items/map/mapz0023.cel",
	"mapz0024.cel": "items/map/mapz0024.cel",
	"mapz0025.cel": "items/map/mapz0025.cel",
	"mapz0026.cel": "items/map/mapz0026.cel",
	"mapz0027.cel": "items/map/mapz0027.cel",
	"mapz0028.cel": "items/map/mapz0028.cel",
	"mapz0029.cel": "items/map/mapz0029.cel",
	"mapz0030.cel": "items/map/mapz0030.cel",
	"mapzdoom.cel": "items/map/mapzdoom.cel",
	"ring.cel":     "items/ring.cel",
	"rock.cel":     "items/rock.cel",
	"scroll.cel":   "items/scroll.cel",
	"shield.cel":   "items/shield.cel",
	"staff.cel":    "items/staff.cel",
	"swrdflip.cel": "items/swrdflip.cel",
	"wand.cel":     "items/wand.cel",
	"wshield.cel":  "items/wshield.cel",
	"l1.cel":       "levels/l1data/l1.cel",
	"l1s.cel":      "levels/l1data/l1s.cel",
	"l2.cel":       "levels/l2data/l2.cel",
	"l2s.cel":      "levels/l2data/l2s.cel",
	"l3.cel":       "levels/l3data/l3.cel",
	"l4.cel":       "levels/l4data/l4.cel",
	"town.cel":     "levels/towndata/town.cel",
	"towns.cel":    "levels/towndata/towns.cel",
	"flamel1.cel":  "missiles/flamel1.cel",
	"flamel2.cel":  "missiles/flamel2.cel",
	"flamel3.cel":  "missiles/flamel3.cel",
	"flamel4.cel":  "missiles/flamel4.cel",
	"flamel5.cel":  "missiles/flamel5.cel",
	"flamel6.cel":  "missiles/flamel6.cel",
	"flamel7.cel":  "missiles/flamel7.cel",
	"flamel8.cel":  "missiles/flamel8.cel",
	"flamel9.cel":  "missiles/flamel9.cel",
	"flamel10.cel": "missiles/flamel10.cel",
	"flamel11.cel": "missiles/flamel11.cel",
	"flamel12.cel": "missiles/flamel12.cel",
	"flamel13.cel": "missiles/flamel13.cel",
	"flamel14.cel": "missiles/flamel14.cel",
	"flamel15.cel": "missiles/flamel15.cel",
	"flamel16.cel": "missiles/flamel16.cel",
	"flames1.cel":  "missiles/flames1.cel",
	"flames2.cel":  "missiles/flames2.cel",
	"flames3.cel":  "missiles/flames3.cel",
	"flames4.cel":  "missiles/flames4.cel",
	"flames5.cel":  "missiles/flames5.cel",
	"flames6.cel":  "missiles/flames6.cel",
	"flames7.cel":  "missiles/flames7.cel",
	"flames8.cel":  "missiles/flames8.cel",
	"flames9.cel":  "missiles/flames9.cel",
	"flames10.cel": "missiles/flames10.cel",
	"flames11.cel": "missiles/flames11.cel",
	"flames12.cel": "missiles/flames12.cel",
	"flames13.cel": "missiles/flames13.cel",
	"flames14.cel": "missiles/flames14.cel",
	"flames15.cel": "missiles/flames15.cel",
	"flames16.cel": "missiles/flames16.cel",
	"flaml1.cel":   "missiles/flaml1.cel",
	"flaml2.cel":   "missiles/flaml2.cel",
	"flaml3.cel":   "missiles/flaml3.cel",
	"flaml4.cel":   "missiles/flaml4.cel",
	"flaml5.cel":   "missiles/flaml5.cel",
	"flaml6.cel":   "missiles/flaml6.cel",
	"flaml7.cel":   "missiles/flaml7.cel",
	"flaml8.cel":   "missiles/flaml8.cel",
	"flams1.cel":   "missiles/flams1.cel",
	"flams2.cel":   "missiles/flams2.cel",
	"flams3.cel":   "missiles/flams3.cel",
	"flams4.cel":   "missiles/flams4.cel",
	"flams5.cel":   "missiles/flams5.cel",
	"flams6.cel":   "missiles/flams6.cel",
	"flams7.cel":   "missiles/flams7.cel",
	"flams8.cel":   "missiles/flams8.cel",
	"mindmace.cel": "missiles/mindmace.cel",
	"sentfr.cel":   "missiles/sentfr.cel",
	"sentout.cel":  "missiles/sentout.cel",
	"sentup.cel":   "missiles/sentup.cel",
	"acidpud.cel":  "monsters/acid/acidpud.cel",
	"magball1.cel": "monsters/magma/magball1.cel",
	"magball2.cel": "monsters/magma/magball2.cel",
	"magball3.cel": "monsters/magma/magball3.cel",
	"magball4.cel": "monsters/magma/magball4.cel",
	"magball5.cel": "monsters/magma/magball5.cel",
	"magball6.cel": "monsters/magma/magball6.cel",
	"magball7.cel": "monsters/magma/magball7.cel",
	"magball8.cel": "monsters/magma/magball8.cel",
	"magblos.cel":  "monsters/magma/magblos.cel",
	"rhinos1.cel":  "monsters/rhino/rhinos1.cel",
	"rhinos2.cel":  "monsters/rhino/rhinos2.cel",
	"rhinos3.cel":  "monsters/rhino/rhinos3.cel",
	"rhinos4.cel":  "monsters/rhino/rhinos4.cel",
	"rhinos5.cel":  "monsters/rhino/rhinos5.cel",
	"rhinos6.cel":  "monsters/rhino/rhinos6.cel",
	"rhinos7.cel":  "monsters/rhino/rhinos7.cel",
	"rhinos8.cel":  "monsters/rhino/rhinos8.cel",
	"flare.cel":    "monsters/succ/flare.cel",
	"flarexp.cel":  "monsters/succ/flarexp.cel",
	"lghning.cel":  "monsters/thin/lghning.cel",
	"unravw.cel":   "monsters/unrav/unravw.cel",
	"altboy.cel":   "objects/altboy.cel",
	"angel.cel":    "objects/angel.cel",
	"armstand.cel": "objects/armstand.cel",
	"banner.cel":   "objects/banner.cel",
	"barrel.cel":   "objects/barrel.cel",
	"barrelex.cel": "objects/barrelex.cel",
	"bcase.cel":    "objects/bcase.cel",
	"bkslbrnt.cel": "objects/bkslbrnt.cel",
	"bkurns.cel":   "objects/bkurns.cel",
	"bloodfnt.cel": "objects/bloodfnt.cel",
	"book1.cel":    "objects/book1.cel",
	"book2.cel":    "objects/book2.cel",
	"bshelf.cel":   "objects/bshelf.cel",
	"burncros.cel": "objects/burncros.cel",
	"candlabr.cel": "objects/candlabr.cel",
	"candle.cel":   "objects/candle.cel",
	"candle2.cel":  "objects/candle2.cel",
	"cauldren.cel": "objects/cauldren.cel",
	"chest1.cel":   "objects/chest1.cel",
	"chest2.cel":   "objects/chest2.cel",
	"chest3.cel":   "objects/chest3.cel",
	"cruxsk1.cel":  "objects/cruxsk1.cel",
	"cruxsk2.cel":  "objects/cruxsk2.cel",
	"cruxsk3.cel":  "objects/cruxsk3.cel",
	"decap.cel":    "objects/decap.cel",
	"dirtfall.cel": "objects/dirtfall.cel",
	"explod1.cel":  "objects/explod1.cel",
	"explod2.cel":  "objects/explod2.cel",
	"firewal1.cel": "objects/firewal1.cel",
	"flame1.cel":   "objects/flame1.cel",
	"flame3.cel":   "objects/flame3.cel",
	"ghost.cel":    "objects/ghost.cel",
	"goatshrn.cel": "objects/goatshrn.cel",
	"l1braz.cel":   "objects/l1braz.cel",
	"l1doors.cel":  "objects/l1doors.cel",
	"l2doors.cel":  "objects/l2doors.cel",
	"l3doors.cel":  "objects/l3doors.cel",
	"lever.cel":    "objects/lever.cel",
	"lshrineg.cel": "objects/lshrineg.cel",
	"lzstand.cel":  "objects/lzstand.cel",
	"mcirl.cel":    "objects/mcirl.cel",
	"mfountn.cel":  "objects/mfountn.cel",
	"miniwatr.cel": "objects/miniwatr.cel",
	"mushptch.cel": "objects/mushptch.cel",
	"nude2.cel":    "objects/nude2.cel",
	"pedistl.cel":  "objects/pedistl.cel",
	"pfountn.cel":  "objects/pfountn.cel",
	"prsrplt1.cel": "objects/prsrplt1.cel",
	"rockstan.cel": "objects/rockstan.cel",
	"rshrineg.cel": "objects/rshrineg.cel",
	"sarc.cel":     "objects/sarc.cel",
	"skulfire.cel": "objects/skulfire.cel",
	"skulpile.cel": "objects/skulpile.cel",
	"skulstik.cel": "objects/skulstik.cel",
	"switch2.cel":  "objects/switch2.cel",
	"switch3.cel":  "objects/switch3.cel",
	"switch4.cel":  "objects/switch4.cel",
	"tfountn.cel":  "objects/tfountn.cel",
	"tnudem.cel":   "objects/tnudem.cel",
	"tnudew.cel":   "objects/tnudew.cel",
	"traphole.cel": "objects/traphole.cel",
	"tsoul.cel":    "objects/tsoul.cel",
	"vapor1.cel":   "objects/vapor1.cel",
	"water.cel":    "objects/water.cel",
	"waterjug.cel": "objects/waterjug.cel",
	"weapstnd.cel": "objects/weapstnd.cel",
	"wtorch1.cel":  "objects/wtorch1.cel",
	"wtorch2.cel":  "objects/wtorch2.cel",
	"wtorch3.cel":  "objects/wtorch3.cel",
	"wtorch4.cel":  "objects/wtorch4.cel",
	"cow.cel":      "towners/animals/cow.cel",
	"deadguy.cel":  "towners/butch/deadguy.cel",
	"twndrunk.cel": "towners/drunk/twndrunk.cel",
	"healer.cel":   "towners/healer/healer.cel",
	"priest8.cel":  "towners/priest/priest8.cel",
	"smithn.cel":   "towners/smith/smithn.cel",
	"smithw.cel":   "towners/smith/smithw.cel",
	"strytell.cel": "towners/strytell/strytell.cel",
	"pegkid1.cel":  "towners/townboy/pegkid1.cel",
	"witch.cel":    "towners/townwmn1/witch.cel",
	"wmnn.cel":     "towners/townwmn1/wmnn.cel",
	"wmnw.cel":     "towners/townwmn1/wmnw.cel",
	"twnfn.cel":    "towners/twnf/twnfn.cel",
	"twnfw.cel":    "towners/twnf/twnfw.cel",

	// CL2 files.
	"acidbf1.cl2":  "missiles/acidbf1.cl2",
	"acidbf2.cl2":  "missiles/acidbf2.cl2",
	"acidbf3.cl2":  "missiles/acidbf3.cl2",
	"acidbf4.cl2":  "missiles/acidbf4.cl2",
	"acidbf5.cl2":  "missiles/acidbf5.cl2",
	"acidbf6.cl2":  "missiles/acidbf6.cl2",
	"acidbf7.cl2":  "missiles/acidbf7.cl2",
	"acidbf8.cl2":  "missiles/acidbf8.cl2",
	"acidbf9.cl2":  "missiles/acidbf9.cl2",
	"acidbf10.cl2": "missiles/acidbf10.cl2",
	"acidbf11.cl2": "missiles/acidbf11.cl2",
	"acidbf12.cl2": "missiles/acidbf12.cl2",
	"acidbf13.cl2": "missiles/acidbf13.cl2",
	"acidbf14.cl2": "missiles/acidbf14.cl2",
	"acidbf15.cl2": "missiles/acidbf15.cl2",
	"acidbf16.cl2": "missiles/acidbf16.cl2",
	"acidpud1.cl2": "missiles/acidpud1.cl2",
	"acidpud2.cl2": "missiles/acidpud2.cl2",
	"acidspla.cl2": "missiles/acidspla.cl2",
	"arrows.cl2":   "missiles/arrows.cl2",
	"bigexp.cl2":   "missiles/bigexp.cl2",
	"blodbur0.cl2": "missiles/blodbur0.cl2",
	"blodbur1.cl2": "missiles/blodbur1.cl2",
	"blodbur2.cl2": "missiles/blodbur2.cl2",
	"blodburs.cl2": "missiles/blodburs.cl2",
	"blood1.cl2":   "missiles/blood1.cl2",
	"blood2.cl2":   "missiles/blood2.cl2",
	"blood3.cl2":   "missiles/blood3.cl2",
	"blood4.cl2":   "missiles/blood4.cl2",
	"bluexbk.cl2":  "missiles/bluexbk.cl2",
	"bluexfr.cl2":  "missiles/bluexfr.cl2",
	"bone1.cl2":    "missiles/bone1.cl2",
	"bone2.cl2":    "missiles/bone2.cl2",
	"bone3.cl2":    "missiles/bone3.cl2",
	"doom1.cl2":    "missiles/doom1.cl2",
	"doom2.cl2":    "missiles/doom2.cl2",
	"doom3.cl2":    "missiles/doom3.cl2",
	"doom4.cl2":    "missiles/doom4.cl2",
	"doom5.cl2":    "missiles/doom5.cl2",
	"doom6.cl2":    "missiles/doom6.cl2",
	"doom7.cl2":    "missiles/doom7.cl2",
	"doom8.cl2":    "missiles/doom8.cl2",
	"doom9.cl2":    "missiles/doom9.cl2",
	"doomexp.cl2":  "missiles/doomexp.cl2",
	"ethrshld.cl2": "missiles/ethrshld.cl2",
	"farrow1.cl2":  "missiles/farrow1.cl2",
	"farrow2.cl2":  "missiles/farrow2.cl2",
	"farrow3.cl2":  "missiles/farrow3.cl2",
	"farrow4.cl2":  "missiles/farrow4.cl2",
	"farrow5.cl2":  "missiles/farrow5.cl2",
	"farrow6.cl2":  "missiles/farrow6.cl2",
	"farrow7.cl2":  "missiles/farrow7.cl2",
	"farrow8.cl2":  "missiles/farrow8.cl2",
	"farrow9.cl2":  "missiles/farrow9.cl2",
	"farrow10.cl2": "missiles/farrow10.cl2",
	"farrow11.cl2": "missiles/farrow11.cl2",
	"farrow12.cl2": "missiles/farrow12.cl2",
	"farrow13.cl2": "missiles/farrow13.cl2",
	"farrow14.cl2": "missiles/farrow14.cl2",
	"farrow15.cl2": "missiles/farrow15.cl2",
	"farrow16.cl2": "missiles/farrow16.cl2",
	"firarwex.cl2": "missiles/firarwex.cl2",
	"fireba1.cl2":  "missiles/fireba1.cl2",
	"fireba2.cl2":  "missiles/fireba2.cl2",
	"fireba3.cl2":  "missiles/fireba3.cl2",
	"fireba4.cl2":  "missiles/fireba4.cl2",
	"fireba5.cl2":  "missiles/fireba5.cl2",
	"fireba6.cl2":  "missiles/fireba6.cl2",
	"fireba7.cl2":  "missiles/fireba7.cl2",
	"fireba8.cl2":  "missiles/fireba8.cl2",
	"fireba9.cl2":  "missiles/fireba9.cl2",
	"fireba10.cl2": "missiles/fireba10.cl2",
	"fireba11.cl2": "missiles/fireba11.cl2",
	"fireba12.cl2": "missiles/fireba12.cl2",
	"fireba13.cl2": "missiles/fireba13.cl2",
	"fireba14.cl2": "missiles/fireba14.cl2",
	"fireba15.cl2": "missiles/fireba15.cl2",
	"fireba16.cl2": "missiles/fireba16.cl2",
	"fireplar.cl2": "missiles/fireplar.cl2",
	"firerun1.cl2": "missiles/firerun1.cl2",
	"firerun2.cl2": "missiles/firerun2.cl2",
	"firerun3.cl2": "missiles/firerun3.cl2",
	"firerun4.cl2": "missiles/firerun4.cl2",
	"firerun5.cl2": "missiles/firerun5.cl2",
	"firerun6.cl2": "missiles/firerun6.cl2",
	"firerun7.cl2": "missiles/firerun7.cl2",
	"firerun8.cl2": "missiles/firerun8.cl2",
	"firewal1.cl2": "missiles/firewal1.cl2",
	"firewal2.cl2": "missiles/firewal2.cl2",
	"flare.cl2":    "missiles/flare.cl2",
	"flareexp.cl2": "missiles/flareexp.cl2",
	"guard1.cl2":   "missiles/guard1.cl2",
	"guard2.cl2":   "missiles/guard2.cl2",
	"guard3.cl2":   "missiles/guard3.cl2",
	"holy1.cl2":    "missiles/holy1.cl2",
	"holy2.cl2":    "missiles/holy2.cl2",
	"holy3.cl2":    "missiles/holy3.cl2",
	"holy4.cl2":    "missiles/holy4.cl2",
	"holy5.cl2":    "missiles/holy5.cl2",
	"holy6.cl2":    "missiles/holy6.cl2",
	"holy7.cl2":    "missiles/holy7.cl2",
	"holy8.cl2":    "missiles/holy8.cl2",
	"holy9.cl2":    "missiles/holy9.cl2",
	"holy10.cl2":   "missiles/holy10.cl2",
	"holy11.cl2":   "missiles/holy11.cl2",
	"holy12.cl2":   "missiles/holy12.cl2",
	"holy13.cl2":   "missiles/holy13.cl2",
	"holy14.cl2":   "missiles/holy14.cl2",
	"holy15.cl2":   "missiles/holy15.cl2",
	"holy16.cl2":   "missiles/holy16.cl2",
	"holyexpl.cl2": "missiles/holyexpl.cl2",
	"inferno.cl2":  "missiles/inferno.cl2",
	"krull.cl2":    "missiles/krull.cl2",
	"larrow1.cl2":  "missiles/larrow1.cl2",
	"larrow2.cl2":  "missiles/larrow2.cl2",
	"larrow3.cl2":  "missiles/larrow3.cl2",
	"larrow4.cl2":  "missiles/larrow4.cl2",
	"larrow5.cl2":  "missiles/larrow5.cl2",
	"larrow6.cl2":  "missiles/larrow6.cl2",
	"larrow7.cl2":  "missiles/larrow7.cl2",
	"larrow8.cl2":  "missiles/larrow8.cl2",
	"larrow9.cl2":  "missiles/larrow9.cl2",
	"larrow10.cl2": "missiles/larrow10.cl2",
	"larrow11.cl2": "missiles/larrow11.cl2",
	"larrow12.cl2": "missiles/larrow12.cl2",
	"larrow13.cl2": "missiles/larrow13.cl2",
	"larrow14.cl2": "missiles/larrow14.cl2",
	"larrow15.cl2": "missiles/larrow15.cl2",
	"larrow16.cl2": "missiles/larrow16.cl2",
	"lghning.cl2":  "missiles/lghning.cl2",
	"magball1.cl2": "missiles/magball1.cl2",
	"magball2.cl2": "missiles/magball2.cl2",
	"magball3.cl2": "missiles/magball3.cl2",
	"magball4.cl2": "missiles/magball4.cl2",
	"magball5.cl2": "missiles/magball5.cl2",
	"magball6.cl2": "missiles/magball6.cl2",
	"magball7.cl2": "missiles/magball7.cl2",
	"magball8.cl2": "missiles/magball8.cl2",
	"magblos.cl2":  "missiles/magblos.cl2",
	"manashld.cl2": "missiles/manashld.cl2",
	"metlhit1.cl2": "missiles/metlhit1.cl2",
	"metlhit2.cl2": "missiles/metlhit2.cl2",
	"metlhit3.cl2": "missiles/metlhit3.cl2",
	"miniltng.cl2": "missiles/miniltng.cl2",
	"newexp.cl2":   "missiles/newexp.cl2",
	"portal.cl2":   "missiles/portal.cl2",
	"portal1.cl2":  "missiles/portal1.cl2",
	"portal2.cl2":  "missiles/portal2.cl2",
	"portalu.cl2":  "missiles/portalu.cl2",
	"ressur1.cl2":  "missiles/ressur1.cl2",
	"rportal1.cl2": "missiles/rportal1.cl2",
	"rportal2.cl2": "missiles/rportal2.cl2",
	"scbsexpb.cl2": "missiles/scbsexpb.cl2",
	"scbsexpc.cl2": "missiles/scbsexpc.cl2",
	"scbsexpd.cl2": "missiles/scbsexpd.cl2",
	"scubmisb.cl2": "missiles/scubmisb.cl2",
	"scubmisc.cl2": "missiles/scubmisc.cl2",
	"scubmisd.cl2": "missiles/scubmisd.cl2",
	"shatter1.cl2": "missiles/shatter1.cl2",
	"sklball1.cl2": "missiles/sklball1.cl2",
	"sklball2.cl2": "missiles/sklball2.cl2",
	"sklball3.cl2": "missiles/sklball3.cl2",
	"sklball4.cl2": "missiles/sklball4.cl2",
	"sklball5.cl2": "missiles/sklball5.cl2",
	"sklball6.cl2": "missiles/sklball6.cl2",
	"sklball7.cl2": "missiles/sklball7.cl2",
	"sklball8.cl2": "missiles/sklball8.cl2",
	"sklball9.cl2": "missiles/sklball9.cl2",
	"thinlght.cl2": "missiles/thinlght.cl2",
	"acida.cl2":    "monsters/acid/acida.cl2",
	"acidd.cl2":    "monsters/acid/acidd.cl2",
	"acidh.cl2":    "monsters/acid/acidh.cl2",
	"acidn.cl2":    "monsters/acid/acidn.cl2",
	"acids.cl2":    "monsters/acid/acids.cl2",
	"acidw.cl2":    "monsters/acid/acidw.cl2",
	"bata.cl2":     "monsters/bat/bata.cl2",
	"batd.cl2":     "monsters/bat/batd.cl2",
	"bath.cl2":     "monsters/bat/bath.cl2",
	"batn.cl2":     "monsters/bat/batn.cl2",
	"batw.cl2":     "monsters/bat/batw.cl2",
	"fallga.cl2":   "monsters/bigfall/fallga.cl2",
	"fallgd.cl2":   "monsters/bigfall/fallgd.cl2",
	"fallgh.cl2":   "monsters/bigfall/fallgh.cl2",
	"fallgn.cl2":   "monsters/bigfall/fallgn.cl2",
	"fallgw.cl2":   "monsters/bigfall/fallgw.cl2",
	"blacka.cl2":   "monsters/black/blacka.cl2",
	"blackd.cl2":   "monsters/black/blackd.cl2",
	"blackh.cl2":   "monsters/black/blackh.cl2",
	"blackn.cl2":   "monsters/black/blackn.cl2",
	"blackw.cl2":   "monsters/black/blackw.cl2",
	"dmagea.cl2":   "monsters/darkmage/dmagea.cl2",
	"dmaged.cl2":   "monsters/darkmage/dmaged.cl2",
	"dmageh.cl2":   "monsters/darkmage/dmageh.cl2",
	"dmagen.cl2":   "monsters/darkmage/dmagen.cl2",
	"dmages.cl2":   "monsters/darkmage/dmages.cl2",
	"dmagew.cl2":   "monsters/darkmage/dmagew.cl2",
	"demskla.cl2":  "monsters/demskel/demskla.cl2",
	"demskld.cl2":  "monsters/demskel/demskld.cl2",
	"demsklh.cl2":  "monsters/demskel/demsklh.cl2",
	"demskln.cl2":  "monsters/demskel/demskln.cl2",
	"demskls.cl2":  "monsters/demskel/demskls.cl2",
	"demsklw.cl2":  "monsters/demskel/demsklw.cl2",
	"diabloa.cl2":  "monsters/diablo/diabloa.cl2",
	"diablod.cl2":  "monsters/diablo/diablod.cl2",
	"diabloh.cl2":  "monsters/diablo/diabloh.cl2",
	"diablon.cl2":  "monsters/diablo/diablon.cl2",
	"diablos.cl2":  "monsters/diablo/diablos.cl2",
	"diablow.cl2":  "monsters/diablo/diablow.cl2",
	"phalla.cl2":   "monsters/falspear/phalla.cl2",
	"phalld.cl2":   "monsters/falspear/phalld.cl2",
	"phallh.cl2":   "monsters/falspear/phallh.cl2",
	"phalln.cl2":   "monsters/falspear/phalln.cl2",
	"phalls.cl2":   "monsters/falspear/phalls.cl2",
	"phallw.cl2":   "monsters/falspear/phallw.cl2",
	"falla.cl2":    "monsters/falsword/falla.cl2",
	"falld.cl2":    "monsters/falsword/falld.cl2",
	"fallh.cl2":    "monsters/falsword/fallh.cl2",
	"falln.cl2":    "monsters/falsword/falln.cl2",
	"falls.cl2":    "monsters/falsword/falls.cl2",
	"fallw.cl2":    "monsters/falsword/fallw.cl2",
	"fatca.cl2":    "monsters/fatc/fatca.cl2",
	"fatcd.cl2":    "monsters/fatc/fatcd.cl2",
	"fatch.cl2":    "monsters/fatc/fatch.cl2",
	"fatcn.cl2":    "monsters/fatc/fatcn.cl2",
	"fatcw.cl2":    "monsters/fatc/fatcw.cl2",
	"fata.cl2":     "monsters/fat/fata.cl2",
	"fatd.cl2":     "monsters/fat/fatd.cl2",
	"fath.cl2":     "monsters/fat/fath.cl2",
	"fatn.cl2":     "monsters/fat/fatn.cl2",
	"fats.cl2":     "monsters/fat/fats.cl2",
	"fatw.cl2":     "monsters/fat/fatw.cl2",
	"firema.cl2":   "monsters/fireman/firema.cl2",
	"firemd.cl2":   "monsters/fireman/firemd.cl2",
	"firemh.cl2":   "monsters/fireman/firemh.cl2",
	"firemn.cl2":   "monsters/fireman/firemn.cl2",
	"firems.cl2":   "monsters/fireman/firems.cl2",
	"firemw.cl2":   "monsters/fireman/firemw.cl2",
	"gargoa.cl2":   "monsters/gargoyle/gargoa.cl2",
	"gargod.cl2":   "monsters/gargoyle/gargod.cl2",
	"gargoh.cl2":   "monsters/gargoyle/gargoh.cl2",
	"gargon.cl2":   "monsters/gargoyle/gargon.cl2",
	"gargos.cl2":   "monsters/gargoyle/gargos.cl2",
	"gargow.cl2":   "monsters/gargoyle/gargow.cl2",
	"goatba.cl2":   "monsters/goatbow/goatba.cl2",
	"goatbd.cl2":   "monsters/goatbow/goatbd.cl2",
	"goatbh.cl2":   "monsters/goatbow/goatbh.cl2",
	"goatbn.cl2":   "monsters/goatbow/goatbn.cl2",
	"goatbw.cl2":   "monsters/goatbow/goatbw.cl2",
	"goatla.cl2":   "monsters/goatlord/goatla.cl2",
	"goatld.cl2":   "monsters/goatlord/goatld.cl2",
	"goatlh.cl2":   "monsters/goatlord/goatlh.cl2",
	"goatln.cl2":   "monsters/goatlord/goatln.cl2",
	"goatlw.cl2":   "monsters/goatlord/goatlw.cl2",
	"goata.cl2":    "monsters/goatmace/goata.cl2",
	"goatd.cl2":    "monsters/goatmace/goatd.cl2",
	"goath.cl2":    "monsters/goatmace/goath.cl2",
	"goatn.cl2":    "monsters/goatmace/goatn.cl2",
	"goats.cl2":    "monsters/goatmace/goats.cl2",
	"goatw.cl2":    "monsters/goatmace/goatw.cl2",
	"golema.cl2":   "monsters/golem/golema.cl2",
	"golemd.cl2":   "monsters/golem/golemd.cl2",
	"golems.cl2":   "monsters/golem/golems.cl2",
	"golemw.cl2":   "monsters/golem/golemw.cl2",
	"magea.cl2":    "monsters/mage/magea.cl2",
	"maged.cl2":    "monsters/mage/maged.cl2",
	"mageh.cl2":    "monsters/mage/mageh.cl2",
	"magen.cl2":    "monsters/mage/magen.cl2",
	"mages.cl2":    "monsters/mage/mages.cl2",
	"magew.cl2":    "monsters/mage/magew.cl2",
	"magmaa.cl2":   "monsters/magma/magmaa.cl2",
	"magmad.cl2":   "monsters/magma/magmad.cl2",
	"magmah.cl2":   "monsters/magma/magmah.cl2",
	"magman.cl2":   "monsters/magma/magman.cl2",
	"magmas.cl2":   "monsters/magma/magmas.cl2",
	"magmaw.cl2":   "monsters/magma/magmaw.cl2",
	"megaa.cl2":    "monsters/mega/megaa.cl2",
	"megad.cl2":    "monsters/mega/megad.cl2",
	"megah.cl2":    "monsters/mega/megah.cl2",
	"megan.cl2":    "monsters/mega/megan.cl2",
	"megas.cl2":    "monsters/mega/megas.cl2",
	"megaw.cl2":    "monsters/mega/megaw.cl2",
	"rhinoa.cl2":   "monsters/rhino/rhinoa.cl2",
	"rhinod.cl2":   "monsters/rhino/rhinod.cl2",
	"rhinoh.cl2":   "monsters/rhino/rhinoh.cl2",
	"rhinon.cl2":   "monsters/rhino/rhinon.cl2",
	"rhinos.cl2":   "monsters/rhino/rhinos.cl2",
	"rhinow.cl2":   "monsters/rhino/rhinow.cl2",
	"scava.cl2":    "monsters/scav/scava.cl2",
	"scavd.cl2":    "monsters/scav/scavd.cl2",
	"scavh.cl2":    "monsters/scav/scavh.cl2",
	"scavn.cl2":    "monsters/scav/scavn.cl2",
	"scavs.cl2":    "monsters/scav/scavs.cl2",
	"scavw.cl2":    "monsters/scav/scavw.cl2",
	"sklaxa.cl2":   "monsters/skelaxe/sklaxa.cl2",
	"sklaxd.cl2":   "monsters/skelaxe/sklaxd.cl2",
	"sklaxh.cl2":   "monsters/skelaxe/sklaxh.cl2",
	"sklaxn.cl2":   "monsters/skelaxe/sklaxn.cl2",
	"sklaxs.cl2":   "monsters/skelaxe/sklaxs.cl2",
	"sklaxw.cl2":   "monsters/skelaxe/sklaxw.cl2",
	"sklbwa.cl2":   "monsters/skelbow/sklbwa.cl2",
	"sklbwd.cl2":   "monsters/skelbow/sklbwd.cl2",
	"sklbwh.cl2":   "monsters/skelbow/sklbwh.cl2",
	"sklbwn.cl2":   "monsters/skelbow/sklbwn.cl2",
	"sklbws.cl2":   "monsters/skelbow/sklbws.cl2",
	"sklbww.cl2":   "monsters/skelbow/sklbww.cl2",
	"sklsra.cl2":   "monsters/skelsd/sklsra.cl2",
	"sklsrd.cl2":   "monsters/skelsd/sklsrd.cl2",
	"sklsrh.cl2":   "monsters/skelsd/sklsrh.cl2",
	"sklsrn.cl2":   "monsters/skelsd/sklsrn.cl2",
	"sklsrs.cl2":   "monsters/skelsd/sklsrs.cl2",
	"sklsrw.cl2":   "monsters/skelsd/sklsrw.cl2",
	"skinga.cl2":   "monsters/sking/skinga.cl2",
	"skingd.cl2":   "monsters/sking/skingd.cl2",
	"skingh.cl2":   "monsters/sking/skingh.cl2",
	"skingn.cl2":   "monsters/sking/skingn.cl2",
	"skings.cl2":   "monsters/sking/skings.cl2",
	"skingw.cl2":   "monsters/sking/skingw.cl2",
	"snakea.cl2":   "monsters/snake/snakea.cl2",
	"snaked.cl2":   "monsters/snake/snaked.cl2",
	"snakeh.cl2":   "monsters/snake/snakeh.cl2",
	"snaken.cl2":   "monsters/snake/snaken.cl2",
	"snakes.cl2":   "monsters/snake/snakes.cl2",
	"snakew.cl2":   "monsters/snake/snakew.cl2",
	"sneaka.cl2":   "monsters/sneak/sneaka.cl2",
	"sneakd.cl2":   "monsters/sneak/sneakd.cl2",
	"sneakh.cl2":   "monsters/sneak/sneakh.cl2",
	"sneakn.cl2":   "monsters/sneak/sneakn.cl2",
	"sneaks.cl2":   "monsters/sneak/sneaks.cl2",
	"sneakw.cl2":   "monsters/sneak/sneakw.cl2",
	"scbsa.cl2":    "monsters/succ/scbsa.cl2",
	"scbsd.cl2":    "monsters/succ/scbsd.cl2",
	"scbsh.cl2":    "monsters/succ/scbsh.cl2",
	"scbsn.cl2":    "monsters/succ/scbsn.cl2",
	"scbsw.cl2":    "monsters/succ/scbsw.cl2",
	"thina.cl2":    "monsters/thin/thina.cl2",
	"thind.cl2":    "monsters/thin/thind.cl2",
	"thinh.cl2":    "monsters/thin/thinh.cl2",
	"thinn.cl2":    "monsters/thin/thinn.cl2",
	"thins.cl2":    "monsters/thin/thins.cl2",
	"thinw.cl2":    "monsters/thin/thinw.cl2",
	"tsneaka.cl2":  "monsters/tsneak/tsneaka.cl2",
	"tsneakd.cl2":  "monsters/tsneak/tsneakd.cl2",
	"tsneakh.cl2":  "monsters/tsneak/tsneakh.cl2",
	"tsneakn.cl2":  "monsters/tsneak/tsneakn.cl2",
	"tsneakw.cl2":  "monsters/tsneak/tsneakw.cl2",
	"unrava.cl2":   "monsters/unrav/unrava.cl2",
	"unravd.cl2":   "monsters/unrav/unravd.cl2",
	"unravh.cl2":   "monsters/unrav/unravh.cl2",
	"unravn.cl2":   "monsters/unrav/unravn.cl2",
	"unravs.cl2":   "monsters/unrav/unravs.cl2",
	"zombiea.cl2":  "monsters/zombie/zombiea.cl2",
	"zombied.cl2":  "monsters/zombie/zombied.cl2",
	"zombieh.cl2":  "monsters/zombie/zombieh.cl2",
	"zombien.cl2":  "monsters/zombie/zombien.cl2",
	"zombies.cl2":  "monsters/zombie/zombies.cl2",
	"zombiew.cl2":  "monsters/zombie/zombiew.cl2",
	"rhaas.cl2":    "plrgfx/rogue/rha/rhaas.cl2",
	"rhaat.cl2":    "plrgfx/rogue/rha/rhaat.cl2",
	"rhaaw.cl2":    "plrgfx/rogue/rha/rhaaw.cl2",
	"rhafm.cl2":    "plrgfx/rogue/rha/rhafm.cl2",
	"rhaht.cl2":    "plrgfx/rogue/rha/rhaht.cl2",
	"rhalm.cl2":    "plrgfx/rogue/rha/rhalm.cl2",
	"rhaqm.cl2":    "plrgfx/rogue/rha/rhaqm.cl2",
	"rhast.cl2":    "plrgfx/rogue/rha/rhast.cl2",
	"rhawl.cl2":    "plrgfx/rogue/rha/rhawl.cl2",
	"rhbas.cl2":    "plrgfx/rogue/rhb/rhbas.cl2",
	"rhbat.cl2":    "plrgfx/rogue/rhb/rhbat.cl2",
	"rhbaw.cl2":    "plrgfx/rogue/rhb/rhbaw.cl2",
	"rhbfm.cl2":    "plrgfx/rogue/rhb/rhbfm.cl2",
	"rhbht.cl2":    "plrgfx/rogue/rhb/rhbht.cl2",
	"rhblm.cl2":    "plrgfx/rogue/rhb/rhblm.cl2",
	"rhbqm.cl2":    "plrgfx/rogue/rhb/rhbqm.cl2",
	"rhbst.cl2":    "plrgfx/rogue/rhb/rhbst.cl2",
	"rhbwl.cl2":    "plrgfx/rogue/rhb/rhbwl.cl2",
	"rhdas.cl2":    "plrgfx/rogue/rhd/rhdas.cl2",
	"rhdat.cl2":    "plrgfx/rogue/rhd/rhdat.cl2",
	"rhdaw.cl2":    "plrgfx/rogue/rhd/rhdaw.cl2",
	"rhdbl.cl2":    "plrgfx/rogue/rhd/rhdbl.cl2",
	"rhdfm.cl2":    "plrgfx/rogue/rhd/rhdfm.cl2",
	"rhdht.cl2":    "plrgfx/rogue/rhd/rhdht.cl2",
	"rhdlm.cl2":    "plrgfx/rogue/rhd/rhdlm.cl2",
	"rhdqm.cl2":    "plrgfx/rogue/rhd/rhdqm.cl2",
	"rhdst.cl2":    "plrgfx/rogue/rhd/rhdst.cl2",
	"rhdwl.cl2":    "plrgfx/rogue/rhd/rhdwl.cl2",
	"rhhas.cl2":    "plrgfx/rogue/rhh/rhhas.cl2",
	"rhhat.cl2":    "plrgfx/rogue/rhh/rhhat.cl2",
	"rhhaw.cl2":    "plrgfx/rogue/rhh/rhhaw.cl2",
	"rhhbl.cl2":    "plrgfx/rogue/rhh/rhhbl.cl2",
	"rhhfm.cl2":    "plrgfx/rogue/rhh/rhhfm.cl2",
	"rhhht.cl2":    "plrgfx/rogue/rhh/rhhht.cl2",
	"rhhlm.cl2":    "plrgfx/rogue/rhh/rhhlm.cl2",
	"rhhqm.cl2":    "plrgfx/rogue/rhh/rhhqm.cl2",
	"rhhst.cl2":    "plrgfx/rogue/rhh/rhhst.cl2",
	"rhhwl.cl2":    "plrgfx/rogue/rhh/rhhwl.cl2",
	"rhmas.cl2":    "plrgfx/rogue/rhm/rhmas.cl2",
	"rhmat.cl2":    "plrgfx/rogue/rhm/rhmat.cl2",
	"rhmaw.cl2":    "plrgfx/rogue/rhm/rhmaw.cl2",
	"rhmfm.cl2":    "plrgfx/rogue/rhm/rhmfm.cl2",
	"rhmht.cl2":    "plrgfx/rogue/rhm/rhmht.cl2",
	"rhmlm.cl2":    "plrgfx/rogue/rhm/rhmlm.cl2",
	"rhmqm.cl2":    "plrgfx/rogue/rhm/rhmqm.cl2",
	"rhmst.cl2":    "plrgfx/rogue/rhm/rhmst.cl2",
	"rhmwl.cl2":    "plrgfx/rogue/rhm/rhmwl.cl2",
	"rhnas.cl2":    "plrgfx/rogue/rhn/rhnas.cl2",
	"rhnat.cl2":    "plrgfx/rogue/rhn/rhnat.cl2",
	"rhnaw.cl2":    "plrgfx/rogue/rhn/rhnaw.cl2",
	"rhndt.cl2":    "plrgfx/rogue/rhn/rhndt.cl2",
	"rhnfm.cl2":    "plrgfx/rogue/rhn/rhnfm.cl2",
	"rhnht.cl2":    "plrgfx/rogue/rhn/rhnht.cl2",
	"rhnlm.cl2":    "plrgfx/rogue/rhn/rhnlm.cl2",
	"rhnqm.cl2":    "plrgfx/rogue/rhn/rhnqm.cl2",
	"rhnst.cl2":    "plrgfx/rogue/rhn/rhnst.cl2",
	"rhnwl.cl2":    "plrgfx/rogue/rhn/rhnwl.cl2",
	"rhsas.cl2":    "plrgfx/rogue/rhs/rhsas.cl2",
	"rhsat.cl2":    "plrgfx/rogue/rhs/rhsat.cl2",
	"rhsaw.cl2":    "plrgfx/rogue/rhs/rhsaw.cl2",
	"rhsfm.cl2":    "plrgfx/rogue/rhs/rhsfm.cl2",
	"rhsht.cl2":    "plrgfx/rogue/rhs/rhsht.cl2",
	"rhslm.cl2":    "plrgfx/rogue/rhs/rhslm.cl2",
	"rhsqm.cl2":    "plrgfx/rogue/rhs/rhsqm.cl2",
	"rhsst.cl2":    "plrgfx/rogue/rhs/rhsst.cl2",
	"rhswl.cl2":    "plrgfx/rogue/rhs/rhswl.cl2",
	"rhtas.cl2":    "plrgfx/rogue/rht/rhtas.cl2",
	"rhtat.cl2":    "plrgfx/rogue/rht/rhtat.cl2",
	"rhtaw.cl2":    "plrgfx/rogue/rht/rhtaw.cl2",
	"rhtfm.cl2":    "plrgfx/rogue/rht/rhtfm.cl2",
	"rhtht.cl2":    "plrgfx/rogue/rht/rhtht.cl2",
	"rhtlm.cl2":    "plrgfx/rogue/rht/rhtlm.cl2",
	"rhtqm.cl2":    "plrgfx/rogue/rht/rhtqm.cl2",
	"rhtst.cl2":    "plrgfx/rogue/rht/rhtst.cl2",
	"rhtwl.cl2":    "plrgfx/rogue/rht/rhtwl.cl2",
	"rhuas.cl2":    "plrgfx/rogue/rhu/rhuas.cl2",
	"rhuat.cl2":    "plrgfx/rogue/rhu/rhuat.cl2",
	"rhuaw.cl2":    "plrgfx/rogue/rhu/rhuaw.cl2",
	"rhubl.cl2":    "plrgfx/rogue/rhu/rhubl.cl2",
	"rhufm.cl2":    "plrgfx/rogue/rhu/rhufm.cl2",
	"rhuht.cl2":    "plrgfx/rogue/rhu/rhuht.cl2",
	"rhulm.cl2":    "plrgfx/rogue/rhu/rhulm.cl2",
	"rhuqm.cl2":    "plrgfx/rogue/rhu/rhuqm.cl2",
	"rhust.cl2":    "plrgfx/rogue/rhu/rhust.cl2",
	"rhuwl.cl2":    "plrgfx/rogue/rhu/rhuwl.cl2",
	"rlaas.cl2":    "plrgfx/rogue/rla/rlaas.cl2",
	"rlaat.cl2":    "plrgfx/rogue/rla/rlaat.cl2",
	"rlaaw.cl2":    "plrgfx/rogue/rla/rlaaw.cl2",
	"rlafm.cl2":    "plrgfx/rogue/rla/rlafm.cl2",
	"rlaht.cl2":    "plrgfx/rogue/rla/rlaht.cl2",
	"rlalm.cl2":    "plrgfx/rogue/rla/rlalm.cl2",
	"rlaqm.cl2":    "plrgfx/rogue/rla/rlaqm.cl2",
	"rlast.cl2":    "plrgfx/rogue/rla/rlast.cl2",
	"rlawl.cl2":    "plrgfx/rogue/rla/rlawl.cl2",
	"rlbas.cl2":    "plrgfx/rogue/rlb/rlbas.cl2",
	"rlbat.cl2":    "plrgfx/rogue/rlb/rlbat.cl2",
	"rlbaw.cl2":    "plrgfx/rogue/rlb/rlbaw.cl2",
	"rlbfm.cl2":    "plrgfx/rogue/rlb/rlbfm.cl2",
	"rlbht.cl2":    "plrgfx/rogue/rlb/rlbht.cl2",
	"rlblm.cl2":    "plrgfx/rogue/rlb/rlblm.cl2",
	"rlbqm.cl2":    "plrgfx/rogue/rlb/rlbqm.cl2",
	"rlbst.cl2":    "plrgfx/rogue/rlb/rlbst.cl2",
	"rlbwl.cl2":    "plrgfx/rogue/rlb/rlbwl.cl2",
	"rldas.cl2":    "plrgfx/rogue/rld/rldas.cl2",
	"rldat.cl2":    "plrgfx/rogue/rld/rldat.cl2",
	"rldaw.cl2":    "plrgfx/rogue/rld/rldaw.cl2",
	"rldbl.cl2":    "plrgfx/rogue/rld/rldbl.cl2",
	"rldfm.cl2":    "plrgfx/rogue/rld/rldfm.cl2",
	"rldht.cl2":    "plrgfx/rogue/rld/rldht.cl2",
	"rldlm.cl2":    "plrgfx/rogue/rld/rldlm.cl2",
	"rldqm.cl2":    "plrgfx/rogue/rld/rldqm.cl2",
	"rldst.cl2":    "plrgfx/rogue/rld/rldst.cl2",
	"rldwl.cl2":    "plrgfx/rogue/rld/rldwl.cl2",
	"rlhas.cl2":    "plrgfx/rogue/rlh/rlhas.cl2",
	"rlhat.cl2":    "plrgfx/rogue/rlh/rlhat.cl2",
	"rlhaw.cl2":    "plrgfx/rogue/rlh/rlhaw.cl2",
	"rlhbl.cl2":    "plrgfx/rogue/rlh/rlhbl.cl2",
	"rlhfm.cl2":    "plrgfx/rogue/rlh/rlhfm.cl2",
	"rlhht.cl2":    "plrgfx/rogue/rlh/rlhht.cl2",
	"rlhlm.cl2":    "plrgfx/rogue/rlh/rlhlm.cl2",
	"rlhqm.cl2":    "plrgfx/rogue/rlh/rlhqm.cl2",
	"rlhst.cl2":    "plrgfx/rogue/rlh/rlhst.cl2",
	"rlhwl.cl2":    "plrgfx/rogue/rlh/rlhwl.cl2",
	"rlmas.cl2":    "plrgfx/rogue/rlm/rlmas.cl2",
	"rlmat.cl2":    "plrgfx/rogue/rlm/rlmat.cl2",
	"rlmaw.cl2":    "plrgfx/rogue/rlm/rlmaw.cl2",
	"rlmfm.cl2":    "plrgfx/rogue/rlm/rlmfm.cl2",
	"rlmht.cl2":    "plrgfx/rogue/rlm/rlmht.cl2",
	"rlmlm.cl2":    "plrgfx/rogue/rlm/rlmlm.cl2",
	"rlmqm.cl2":    "plrgfx/rogue/rlm/rlmqm.cl2",
	"rlmst.cl2":    "plrgfx/rogue/rlm/rlmst.cl2",
	"rlmwl.cl2":    "plrgfx/rogue/rlm/rlmwl.cl2",
	"rlnas.cl2":    "plrgfx/rogue/rln/rlnas.cl2",
	"rlnat.cl2":    "plrgfx/rogue/rln/rlnat.cl2",
	"rlnaw.cl2":    "plrgfx/rogue/rln/rlnaw.cl2",
	"rlndt.cl2":    "plrgfx/rogue/rln/rlndt.cl2",
	"rlnfm.cl2":    "plrgfx/rogue/rln/rlnfm.cl2",
	"rlnht.cl2":    "plrgfx/rogue/rln/rlnht.cl2",
	"rlnlm.cl2":    "plrgfx/rogue/rln/rlnlm.cl2",
	"rlnqm.cl2":    "plrgfx/rogue/rln/rlnqm.cl2",
	"rlnst.cl2":    "plrgfx/rogue/rln/rlnst.cl2",
	"rlnwl.cl2":    "plrgfx/rogue/rln/rlnwl.cl2",
	"rlsas.cl2":    "plrgfx/rogue/rls/rlsas.cl2",
	"rlsat.cl2":    "plrgfx/rogue/rls/rlsat.cl2",
	"rlsaw.cl2":    "plrgfx/rogue/rls/rlsaw.cl2",
	"rlsfm.cl2":    "plrgfx/rogue/rls/rlsfm.cl2",
	"rlsht.cl2":    "plrgfx/rogue/rls/rlsht.cl2",
	"rlslm.cl2":    "plrgfx/rogue/rls/rlslm.cl2",
	"rlsqm.cl2":    "plrgfx/rogue/rls/rlsqm.cl2",
	"rlsst.cl2":    "plrgfx/rogue/rls/rlsst.cl2",
	"rlswl.cl2":    "plrgfx/rogue/rls/rlswl.cl2",
	"rltas.cl2":    "plrgfx/rogue/rlt/rltas.cl2",
	"rltat.cl2":    "plrgfx/rogue/rlt/rltat.cl2",
	"rltaw.cl2":    "plrgfx/rogue/rlt/rltaw.cl2",
	"rltfm.cl2":    "plrgfx/rogue/rlt/rltfm.cl2",
	"rltht.cl2":    "plrgfx/rogue/rlt/rltht.cl2",
	"rltlm.cl2":    "plrgfx/rogue/rlt/rltlm.cl2",
	"rltqm.cl2":    "plrgfx/rogue/rlt/rltqm.cl2",
	"rltst.cl2":    "plrgfx/rogue/rlt/rltst.cl2",
	"rltwl.cl2":    "plrgfx/rogue/rlt/rltwl.cl2",
	"rluas.cl2":    "plrgfx/rogue/rlu/rluas.cl2",
	"rluat.cl2":    "plrgfx/rogue/rlu/rluat.cl2",
	"rluaw.cl2":    "plrgfx/rogue/rlu/rluaw.cl2",
	"rlubl.cl2":    "plrgfx/rogue/rlu/rlubl.cl2",
	"rlufm.cl2":    "plrgfx/rogue/rlu/rlufm.cl2",
	"rluht.cl2":    "plrgfx/rogue/rlu/rluht.cl2",
	"rlulm.cl2":    "plrgfx/rogue/rlu/rlulm.cl2",
	"rluqm.cl2":    "plrgfx/rogue/rlu/rluqm.cl2",
	"rlust.cl2":    "plrgfx/rogue/rlu/rlust.cl2",
	"rluwl.cl2":    "plrgfx/rogue/rlu/rluwl.cl2",
	"rmaas.cl2":    "plrgfx/rogue/rma/rmaas.cl2",
	"rmaat.cl2":    "plrgfx/rogue/rma/rmaat.cl2",
	"rmaaw.cl2":    "plrgfx/rogue/rma/rmaaw.cl2",
	"rmafm.cl2":    "plrgfx/rogue/rma/rmafm.cl2",
	"rmaht.cl2":    "plrgfx/rogue/rma/rmaht.cl2",
	"rmalm.cl2":    "plrgfx/rogue/rma/rmalm.cl2",
	"rmaqm.cl2":    "plrgfx/rogue/rma/rmaqm.cl2",
	"rmast.cl2":    "plrgfx/rogue/rma/rmast.cl2",
	"rmawl.cl2":    "plrgfx/rogue/rma/rmawl.cl2",
	"rmbas.cl2":    "plrgfx/rogue/rmb/rmbas.cl2",
	"rmbat.cl2":    "plrgfx/rogue/rmb/rmbat.cl2",
	"rmbaw.cl2":    "plrgfx/rogue/rmb/rmbaw.cl2",
	"rmbfm.cl2":    "plrgfx/rogue/rmb/rmbfm.cl2",
	"rmbht.cl2":    "plrgfx/rogue/rmb/rmbht.cl2",
	"rmblm.cl2":    "plrgfx/rogue/rmb/rmblm.cl2",
	"rmbqm.cl2":    "plrgfx/rogue/rmb/rmbqm.cl2",
	"rmbst.cl2":    "plrgfx/rogue/rmb/rmbst.cl2",
	"rmbwl.cl2":    "plrgfx/rogue/rmb/rmbwl.cl2",
	"rmdas.cl2":    "plrgfx/rogue/rmd/rmdas.cl2",
	"rmdat.cl2":    "plrgfx/rogue/rmd/rmdat.cl2",
	"rmdaw.cl2":    "plrgfx/rogue/rmd/rmdaw.cl2",
	"rmdbl.cl2":    "plrgfx/rogue/rmd/rmdbl.cl2",
	"rmdfm.cl2":    "plrgfx/rogue/rmd/rmdfm.cl2",
	"rmdht.cl2":    "plrgfx/rogue/rmd/rmdht.cl2",
	"rmdlm.cl2":    "plrgfx/rogue/rmd/rmdlm.cl2",
	"rmdqm.cl2":    "plrgfx/rogue/rmd/rmdqm.cl2",
	"rmdst.cl2":    "plrgfx/rogue/rmd/rmdst.cl2",
	"rmdwl.cl2":    "plrgfx/rogue/rmd/rmdwl.cl2",
	"rmhas.cl2":    "plrgfx/rogue/rmh/rmhas.cl2",
	"rmhat.cl2":    "plrgfx/rogue/rmh/rmhat.cl2",
	"rmhaw.cl2":    "plrgfx/rogue/rmh/rmhaw.cl2",
	"rmhbl.cl2":    "plrgfx/rogue/rmh/rmhbl.cl2",
	"rmhfm.cl2":    "plrgfx/rogue/rmh/rmhfm.cl2",
	"rmhht.cl2":    "plrgfx/rogue/rmh/rmhht.cl2",
	"rmhlm.cl2":    "plrgfx/rogue/rmh/rmhlm.cl2",
	"rmhqm.cl2":    "plrgfx/rogue/rmh/rmhqm.cl2",
	"rmhst.cl2":    "plrgfx/rogue/rmh/rmhst.cl2",
	"rmhwl.cl2":    "plrgfx/rogue/rmh/rmhwl.cl2",
	"rmmas.cl2":    "plrgfx/rogue/rmm/rmmas.cl2",
	"rmmat.cl2":    "plrgfx/rogue/rmm/rmmat.cl2",
	"rmmaw.cl2":    "plrgfx/rogue/rmm/rmmaw.cl2",
	"rmmfm.cl2":    "plrgfx/rogue/rmm/rmmfm.cl2",
	"rmmht.cl2":    "plrgfx/rogue/rmm/rmmht.cl2",
	"rmmlm.cl2":    "plrgfx/rogue/rmm/rmmlm.cl2",
	"rmmqm.cl2":    "plrgfx/rogue/rmm/rmmqm.cl2",
	"rmmst.cl2":    "plrgfx/rogue/rmm/rmmst.cl2",
	"rmmwl.cl2":    "plrgfx/rogue/rmm/rmmwl.cl2",
	"rmnas.cl2":    "plrgfx/rogue/rmn/rmnas.cl2",
	"rmnat.cl2":    "plrgfx/rogue/rmn/rmnat.cl2",
	"rmnaw.cl2":    "plrgfx/rogue/rmn/rmnaw.cl2",
	"rmndt.cl2":    "plrgfx/rogue/rmn/rmndt.cl2",
	"rmnfm.cl2":    "plrgfx/rogue/rmn/rmnfm.cl2",
	"rmnht.cl2":    "plrgfx/rogue/rmn/rmnht.cl2",
	"rmnlm.cl2":    "plrgfx/rogue/rmn/rmnlm.cl2",
	"rmnqm.cl2":    "plrgfx/rogue/rmn/rmnqm.cl2",
	"rmnst.cl2":    "plrgfx/rogue/rmn/rmnst.cl2",
	"rmnwl.cl2":    "plrgfx/rogue/rmn/rmnwl.cl2",
	"rmsas.cl2":    "plrgfx/rogue/rms/rmsas.cl2",
	"rmsat.cl2":    "plrgfx/rogue/rms/rmsat.cl2",
	"rmsaw.cl2":    "plrgfx/rogue/rms/rmsaw.cl2",
	"rmsfm.cl2":    "plrgfx/rogue/rms/rmsfm.cl2",
	"rmsht.cl2":    "plrgfx/rogue/rms/rmsht.cl2",
	"rmslm.cl2":    "plrgfx/rogue/rms/rmslm.cl2",
	"rmsqm.cl2":    "plrgfx/rogue/rms/rmsqm.cl2",
	"rmsst.cl2":    "plrgfx/rogue/rms/rmsst.cl2",
	"rmswl.cl2":    "plrgfx/rogue/rms/rmswl.cl2",
	"rmtas.cl2":    "plrgfx/rogue/rmt/rmtas.cl2",
	"rmtat.cl2":    "plrgfx/rogue/rmt/rmtat.cl2",
	"rmtaw.cl2":    "plrgfx/rogue/rmt/rmtaw.cl2",
	"rmtfm.cl2":    "plrgfx/rogue/rmt/rmtfm.cl2",
	"rmtht.cl2":    "plrgfx/rogue/rmt/rmtht.cl2",
	"rmtlm.cl2":    "plrgfx/rogue/rmt/rmtlm.cl2",
	"rmtqm.cl2":    "plrgfx/rogue/rmt/rmtqm.cl2",
	"rmtst.cl2":    "plrgfx/rogue/rmt/rmtst.cl2",
	"rmtwl.cl2":    "plrgfx/rogue/rmt/rmtwl.cl2",
	"rmuas.cl2":    "plrgfx/rogue/rmu/rmuas.cl2",
	"rmuat.cl2":    "plrgfx/rogue/rmu/rmuat.cl2",
	"rmuaw.cl2":    "plrgfx/rogue/rmu/rmuaw.cl2",
	"rmubl.cl2":    "plrgfx/rogue/rmu/rmubl.cl2",
	"rmufm.cl2":    "plrgfx/rogue/rmu/rmufm.cl2",
	"rmuht.cl2":    "plrgfx/rogue/rmu/rmuht.cl2",
	"rmulm.cl2":    "plrgfx/rogue/rmu/rmulm.cl2",
	"rmuqm.cl2":    "plrgfx/rogue/rmu/rmuqm.cl2",
	"rmust.cl2":    "plrgfx/rogue/rmu/rmust.cl2",
	"rmuwl.cl2":    "plrgfx/rogue/rmu/rmuwl.cl2",
	"shaas.cl2":    "plrgfx/sorceror/sha/shaas.cl2",
	"shaat.cl2":    "plrgfx/sorceror/sha/shaat.cl2",
	"shaaw.cl2":    "plrgfx/sorceror/sha/shaaw.cl2",
	"shafm.cl2":    "plrgfx/sorceror/sha/shafm.cl2",
	"shaht.cl2":    "plrgfx/sorceror/sha/shaht.cl2",
	"shalm.cl2":    "plrgfx/sorceror/sha/shalm.cl2",
	"shaqm.cl2":    "plrgfx/sorceror/sha/shaqm.cl2",
	"shast.cl2":    "plrgfx/sorceror/sha/shast.cl2",
	"shawl.cl2":    "plrgfx/sorceror/sha/shawl.cl2",
	"shbas.cl2":    "plrgfx/sorceror/shb/shbas.cl2",
	"shbat.cl2":    "plrgfx/sorceror/shb/shbat.cl2",
	"shbaw.cl2":    "plrgfx/sorceror/shb/shbaw.cl2",
	"shbfm.cl2":    "plrgfx/sorceror/shb/shbfm.cl2",
	"shbht.cl2":    "plrgfx/sorceror/shb/shbht.cl2",
	"shblm.cl2":    "plrgfx/sorceror/shb/shblm.cl2",
	"shbqm.cl2":    "plrgfx/sorceror/shb/shbqm.cl2",
	"shbst.cl2":    "plrgfx/sorceror/shb/shbst.cl2",
	"shbwl.cl2":    "plrgfx/sorceror/shb/shbwl.cl2",
	"shdas.cl2":    "plrgfx/sorceror/shd/shdas.cl2",
	"shdat.cl2":    "plrgfx/sorceror/shd/shdat.cl2",
	"shdaw.cl2":    "plrgfx/sorceror/shd/shdaw.cl2",
	"shdbl.cl2":    "plrgfx/sorceror/shd/shdbl.cl2",
	"shdfm.cl2":    "plrgfx/sorceror/shd/shdfm.cl2",
	"shdht.cl2":    "plrgfx/sorceror/shd/shdht.cl2",
	"shdlm.cl2":    "plrgfx/sorceror/shd/shdlm.cl2",
	"shdqm.cl2":    "plrgfx/sorceror/shd/shdqm.cl2",
	"shdst.cl2":    "plrgfx/sorceror/shd/shdst.cl2",
	"shdwl.cl2":    "plrgfx/sorceror/shd/shdwl.cl2",
	"shhas.cl2":    "plrgfx/sorceror/shh/shhas.cl2",
	"shhat.cl2":    "plrgfx/sorceror/shh/shhat.cl2",
	"shhaw.cl2":    "plrgfx/sorceror/shh/shhaw.cl2",
	"shhbl.cl2":    "plrgfx/sorceror/shh/shhbl.cl2",
	"shhfm.cl2":    "plrgfx/sorceror/shh/shhfm.cl2",
	"shhht.cl2":    "plrgfx/sorceror/shh/shhht.cl2",
	"shhlm.cl2":    "plrgfx/sorceror/shh/shhlm.cl2",
	"shhqm.cl2":    "plrgfx/sorceror/shh/shhqm.cl2",
	"shhst.cl2":    "plrgfx/sorceror/shh/shhst.cl2",
	"shhwl.cl2":    "plrgfx/sorceror/shh/shhwl.cl2",
	"shmas.cl2":    "plrgfx/sorceror/shm/shmas.cl2",
	"shmat.cl2":    "plrgfx/sorceror/shm/shmat.cl2",
	"shmaw.cl2":    "plrgfx/sorceror/shm/shmaw.cl2",
	"shmfm.cl2":    "plrgfx/sorceror/shm/shmfm.cl2",
	"shmht.cl2":    "plrgfx/sorceror/shm/shmht.cl2",
	"shmlm.cl2":    "plrgfx/sorceror/shm/shmlm.cl2",
	"shmqm.cl2":    "plrgfx/sorceror/shm/shmqm.cl2",
	"shmst.cl2":    "plrgfx/sorceror/shm/shmst.cl2",
	"shmwl.cl2":    "plrgfx/sorceror/shm/shmwl.cl2",
	"shnas.cl2":    "plrgfx/sorceror/shn/shnas.cl2",
	"shnat.cl2":    "plrgfx/sorceror/shn/shnat.cl2",
	"shnaw.cl2":    "plrgfx/sorceror/shn/shnaw.cl2",
	"shndt.cl2":    "plrgfx/sorceror/shn/shndt.cl2",
	"shnfm.cl2":    "plrgfx/sorceror/shn/shnfm.cl2",
	"shnht.cl2":    "plrgfx/sorceror/shn/shnht.cl2",
	"shnlm.cl2":    "plrgfx/sorceror/shn/shnlm.cl2",
	"shnqm.cl2":    "plrgfx/sorceror/shn/shnqm.cl2",
	"shnst.cl2":    "plrgfx/sorceror/shn/shnst.cl2",
	"shnwl.cl2":    "plrgfx/sorceror/shn/shnwl.cl2",
	"shsas.cl2":    "plrgfx/sorceror/shs/shsas.cl2",
	"shsat.cl2":    "plrgfx/sorceror/shs/shsat.cl2",
	"shsaw.cl2":    "plrgfx/sorceror/shs/shsaw.cl2",
	"shsfm.cl2":    "plrgfx/sorceror/shs/shsfm.cl2",
	"shsht.cl2":    "plrgfx/sorceror/shs/shsht.cl2",
	"shslm.cl2":    "plrgfx/sorceror/shs/shslm.cl2",
	"shsqm.cl2":    "plrgfx/sorceror/shs/shsqm.cl2",
	"shsst.cl2":    "plrgfx/sorceror/shs/shsst.cl2",
	"shswl.cl2":    "plrgfx/sorceror/shs/shswl.cl2",
	"shtas.cl2":    "plrgfx/sorceror/sht/shtas.cl2",
	"shtat.cl2":    "plrgfx/sorceror/sht/shtat.cl2",
	"shtaw.cl2":    "plrgfx/sorceror/sht/shtaw.cl2",
	"shtfm.cl2":    "plrgfx/sorceror/sht/shtfm.cl2",
	"shtht.cl2":    "plrgfx/sorceror/sht/shtht.cl2",
	"shtlm.cl2":    "plrgfx/sorceror/sht/shtlm.cl2",
	"shtqm.cl2":    "plrgfx/sorceror/sht/shtqm.cl2",
	"shtst.cl2":    "plrgfx/sorceror/sht/shtst.cl2",
	"shtwl.cl2":    "plrgfx/sorceror/sht/shtwl.cl2",
	"shuas.cl2":    "plrgfx/sorceror/shu/shuas.cl2",
	"shuat.cl2":    "plrgfx/sorceror/shu/shuat.cl2",
	"shuaw.cl2":    "plrgfx/sorceror/shu/shuaw.cl2",
	"shubl.cl2":    "plrgfx/sorceror/shu/shubl.cl2",
	"shufm.cl2":    "plrgfx/sorceror/shu/shufm.cl2",
	"shuht.cl2":    "plrgfx/sorceror/shu/shuht.cl2",
	"shulm.cl2":    "plrgfx/sorceror/shu/shulm.cl2",
	"shuqm.cl2":    "plrgfx/sorceror/shu/shuqm.cl2",
	"shust.cl2":    "plrgfx/sorceror/shu/shust.cl2",
	"shuwl.cl2":    "plrgfx/sorceror/shu/shuwl.cl2",
	"slaas.cl2":    "plrgfx/sorceror/sla/slaas.cl2",
	"slaat.cl2":    "plrgfx/sorceror/sla/slaat.cl2",
	"slaaw.cl2":    "plrgfx/sorceror/sla/slaaw.cl2",
	"slafm.cl2":    "plrgfx/sorceror/sla/slafm.cl2",
	"slaht.cl2":    "plrgfx/sorceror/sla/slaht.cl2",
	"slalm.cl2":    "plrgfx/sorceror/sla/slalm.cl2",
	"slaqm.cl2":    "plrgfx/sorceror/sla/slaqm.cl2",
	"slast.cl2":    "plrgfx/sorceror/sla/slast.cl2",
	"slawl.cl2":    "plrgfx/sorceror/sla/slawl.cl2",
	"slbas.cl2":    "plrgfx/sorceror/slb/slbas.cl2",
	"slbat.cl2":    "plrgfx/sorceror/slb/slbat.cl2",
	"slbaw.cl2":    "plrgfx/sorceror/slb/slbaw.cl2",
	"slbfm.cl2":    "plrgfx/sorceror/slb/slbfm.cl2",
	"slbht.cl2":    "plrgfx/sorceror/slb/slbht.cl2",
	"slblm.cl2":    "plrgfx/sorceror/slb/slblm.cl2",
	"slbqm.cl2":    "plrgfx/sorceror/slb/slbqm.cl2",
	"slbst.cl2":    "plrgfx/sorceror/slb/slbst.cl2",
	"slbwl.cl2":    "plrgfx/sorceror/slb/slbwl.cl2",
	"sldas.cl2":    "plrgfx/sorceror/sld/sldas.cl2",
	"sldat.cl2":    "plrgfx/sorceror/sld/sldat.cl2",
	"sldaw.cl2":    "plrgfx/sorceror/sld/sldaw.cl2",
	"sldbl.cl2":    "plrgfx/sorceror/sld/sldbl.cl2",
	"sldfm.cl2":    "plrgfx/sorceror/sld/sldfm.cl2",
	"sldht.cl2":    "plrgfx/sorceror/sld/sldht.cl2",
	"sldlm.cl2":    "plrgfx/sorceror/sld/sldlm.cl2",
	"sldqm.cl2":    "plrgfx/sorceror/sld/sldqm.cl2",
	"sldst.cl2":    "plrgfx/sorceror/sld/sldst.cl2",
	"sldwl.cl2":    "plrgfx/sorceror/sld/sldwl.cl2",
	"slhas.cl2":    "plrgfx/sorceror/slh/slhas.cl2",
	"slhat.cl2":    "plrgfx/sorceror/slh/slhat.cl2",
	"slhaw.cl2":    "plrgfx/sorceror/slh/slhaw.cl2",
	"slhbl.cl2":    "plrgfx/sorceror/slh/slhbl.cl2",
	"slhfm.cl2":    "plrgfx/sorceror/slh/slhfm.cl2",
	"slhht.cl2":    "plrgfx/sorceror/slh/slhht.cl2",
	"slhlm.cl2":    "plrgfx/sorceror/slh/slhlm.cl2",
	"slhqm.cl2":    "plrgfx/sorceror/slh/slhqm.cl2",
	"slhst.cl2":    "plrgfx/sorceror/slh/slhst.cl2",
	"slhwl.cl2":    "plrgfx/sorceror/slh/slhwl.cl2",
	"slmas.cl2":    "plrgfx/sorceror/slm/slmas.cl2",
	"slmat.cl2":    "plrgfx/sorceror/slm/slmat.cl2",
	"slmaw.cl2":    "plrgfx/sorceror/slm/slmaw.cl2",
	"slmfm.cl2":    "plrgfx/sorceror/slm/slmfm.cl2",
	"slmht.cl2":    "plrgfx/sorceror/slm/slmht.cl2",
	"slmlm.cl2":    "plrgfx/sorceror/slm/slmlm.cl2",
	"slmqm.cl2":    "plrgfx/sorceror/slm/slmqm.cl2",
	"slmst.cl2":    "plrgfx/sorceror/slm/slmst.cl2",
	"slmwl.cl2":    "plrgfx/sorceror/slm/slmwl.cl2",
	"slnas.cl2":    "plrgfx/sorceror/sln/slnas.cl2",
	"slnat.cl2":    "plrgfx/sorceror/sln/slnat.cl2",
	"slnaw.cl2":    "plrgfx/sorceror/sln/slnaw.cl2",
	"slndt.cl2":    "plrgfx/sorceror/sln/slndt.cl2",
	"slnfm.cl2":    "plrgfx/sorceror/sln/slnfm.cl2",
	"slnht.cl2":    "plrgfx/sorceror/sln/slnht.cl2",
	"slnlm.cl2":    "plrgfx/sorceror/sln/slnlm.cl2",
	"slnqm.cl2":    "plrgfx/sorceror/sln/slnqm.cl2",
	"slnst.cl2":    "plrgfx/sorceror/sln/slnst.cl2",
	"slnwl.cl2":    "plrgfx/sorceror/sln/slnwl.cl2",
	"slsas.cl2":    "plrgfx/sorceror/sls/slsas.cl2",
	"slsat.cl2":    "plrgfx/sorceror/sls/slsat.cl2",
	"slsaw.cl2":    "plrgfx/sorceror/sls/slsaw.cl2",
	"slsfm.cl2":    "plrgfx/sorceror/sls/slsfm.cl2",
	"slsht.cl2":    "plrgfx/sorceror/sls/slsht.cl2",
	"slslm.cl2":    "plrgfx/sorceror/sls/slslm.cl2",
	"slsqm.cl2":    "plrgfx/sorceror/sls/slsqm.cl2",
	"slsst.cl2":    "plrgfx/sorceror/sls/slsst.cl2",
	"slswl.cl2":    "plrgfx/sorceror/sls/slswl.cl2",
	"sltas.cl2":    "plrgfx/sorceror/slt/sltas.cl2",
	"sltat.cl2":    "plrgfx/sorceror/slt/sltat.cl2",
	"sltaw.cl2":    "plrgfx/sorceror/slt/sltaw.cl2",
	"sltfm.cl2":    "plrgfx/sorceror/slt/sltfm.cl2",
	"sltht.cl2":    "plrgfx/sorceror/slt/sltht.cl2",
	"sltlm.cl2":    "plrgfx/sorceror/slt/sltlm.cl2",
	"sltqm.cl2":    "plrgfx/sorceror/slt/sltqm.cl2",
	"sltst.cl2":    "plrgfx/sorceror/slt/sltst.cl2",
	"sltwl.cl2":    "plrgfx/sorceror/slt/sltwl.cl2",
	"sluas.cl2":    "plrgfx/sorceror/slu/sluas.cl2",
	"sluat.cl2":    "plrgfx/sorceror/slu/sluat.cl2",
	"sluaw.cl2":    "plrgfx/sorceror/slu/sluaw.cl2",
	"slubl.cl2":    "plrgfx/sorceror/slu/slubl.cl2",
	"slufm.cl2":    "plrgfx/sorceror/slu/slufm.cl2",
	"sluht.cl2":    "plrgfx/sorceror/slu/sluht.cl2",
	"slulm.cl2":    "plrgfx/sorceror/slu/slulm.cl2",
	"sluqm.cl2":    "plrgfx/sorceror/slu/sluqm.cl2",
	"slust.cl2":    "plrgfx/sorceror/slu/slust.cl2",
	"sluwl.cl2":    "plrgfx/sorceror/slu/sluwl.cl2",
	"smaas.cl2":    "plrgfx/sorceror/sma/smaas.cl2",
	"smaat.cl2":    "plrgfx/sorceror/sma/smaat.cl2",
	"smaaw.cl2":    "plrgfx/sorceror/sma/smaaw.cl2",
	"smafm.cl2":    "plrgfx/sorceror/sma/smafm.cl2",
	"smaht.cl2":    "plrgfx/sorceror/sma/smaht.cl2",
	"smalm.cl2":    "plrgfx/sorceror/sma/smalm.cl2",
	"smaqm.cl2":    "plrgfx/sorceror/sma/smaqm.cl2",
	"smast.cl2":    "plrgfx/sorceror/sma/smast.cl2",
	"smawl.cl2":    "plrgfx/sorceror/sma/smawl.cl2",
	"smbas.cl2":    "plrgfx/sorceror/smb/smbas.cl2",
	"smbat.cl2":    "plrgfx/sorceror/smb/smbat.cl2",
	"smbaw.cl2":    "plrgfx/sorceror/smb/smbaw.cl2",
	"smbfm.cl2":    "plrgfx/sorceror/smb/smbfm.cl2",
	"smbht.cl2":    "plrgfx/sorceror/smb/smbht.cl2",
	"smblm.cl2":    "plrgfx/sorceror/smb/smblm.cl2",
	"smbqm.cl2":    "plrgfx/sorceror/smb/smbqm.cl2",
	"smbst.cl2":    "plrgfx/sorceror/smb/smbst.cl2",
	"smbwl.cl2":    "plrgfx/sorceror/smb/smbwl.cl2",
	"smdas.cl2":    "plrgfx/sorceror/smd/smdas.cl2",
	"smdat.cl2":    "plrgfx/sorceror/smd/smdat.cl2",
	"smdaw.cl2":    "plrgfx/sorceror/smd/smdaw.cl2",
	"smdbl.cl2":    "plrgfx/sorceror/smd/smdbl.cl2",
	"smdfm.cl2":    "plrgfx/sorceror/smd/smdfm.cl2",
	"smdht.cl2":    "plrgfx/sorceror/smd/smdht.cl2",
	"smdlm.cl2":    "plrgfx/sorceror/smd/smdlm.cl2",
	"smdqm.cl2":    "plrgfx/sorceror/smd/smdqm.cl2",
	"smdst.cl2":    "plrgfx/sorceror/smd/smdst.cl2",
	"smdwl.cl2":    "plrgfx/sorceror/smd/smdwl.cl2",
	"smhas.cl2":    "plrgfx/sorceror/smh/smhas.cl2",
	"smhat.cl2":    "plrgfx/sorceror/smh/smhat.cl2",
	"smhaw.cl2":    "plrgfx/sorceror/smh/smhaw.cl2",
	"smhbl.cl2":    "plrgfx/sorceror/smh/smhbl.cl2",
	"smhfm.cl2":    "plrgfx/sorceror/smh/smhfm.cl2",
	"smhht.cl2":    "plrgfx/sorceror/smh/smhht.cl2",
	"smhlm.cl2":    "plrgfx/sorceror/smh/smhlm.cl2",
	"smhqm.cl2":    "plrgfx/sorceror/smh/smhqm.cl2",
	"smhst.cl2":    "plrgfx/sorceror/smh/smhst.cl2",
	"smhwl.cl2":    "plrgfx/sorceror/smh/smhwl.cl2",
	"smmas.cl2":    "plrgfx/sorceror/smm/smmas.cl2",
	"smmat.cl2":    "plrgfx/sorceror/smm/smmat.cl2",
	"smmaw.cl2":    "plrgfx/sorceror/smm/smmaw.cl2",
	"smmfm.cl2":    "plrgfx/sorceror/smm/smmfm.cl2",
	"smmht.cl2":    "plrgfx/sorceror/smm/smmht.cl2",
	"smmlm.cl2":    "plrgfx/sorceror/smm/smmlm.cl2",
	"smmqm.cl2":    "plrgfx/sorceror/smm/smmqm.cl2",
	"smmst.cl2":    "plrgfx/sorceror/smm/smmst.cl2",
	"smmwl.cl2":    "plrgfx/sorceror/smm/smmwl.cl2",
	"smnas.cl2":    "plrgfx/sorceror/smn/smnas.cl2",
	"smnat.cl2":    "plrgfx/sorceror/smn/smnat.cl2",
	"smnaw.cl2":    "plrgfx/sorceror/smn/smnaw.cl2",
	"smndt.cl2":    "plrgfx/sorceror/smn/smndt.cl2",
	"smnfm.cl2":    "plrgfx/sorceror/smn/smnfm.cl2",
	"smnht.cl2":    "plrgfx/sorceror/smn/smnht.cl2",
	"smnlm.cl2":    "plrgfx/sorceror/smn/smnlm.cl2",
	"smnqm.cl2":    "plrgfx/sorceror/smn/smnqm.cl2",
	"smnst.cl2":    "plrgfx/sorceror/smn/smnst.cl2",
	"smnwl.cl2":    "plrgfx/sorceror/smn/smnwl.cl2",
	"smsas.cl2":    "plrgfx/sorceror/sms/smsas.cl2",
	"smsat.cl2":    "plrgfx/sorceror/sms/smsat.cl2",
	"smsaw.cl2":    "plrgfx/sorceror/sms/smsaw.cl2",
	"smsfm.cl2":    "plrgfx/sorceror/sms/smsfm.cl2",
	"smsht.cl2":    "plrgfx/sorceror/sms/smsht.cl2",
	"smslm.cl2":    "plrgfx/sorceror/sms/smslm.cl2",
	"smsqm.cl2":    "plrgfx/sorceror/sms/smsqm.cl2",
	"smsst.cl2":    "plrgfx/sorceror/sms/smsst.cl2",
	"smswl.cl2":    "plrgfx/sorceror/sms/smswl.cl2",
	"smtas.cl2":    "plrgfx/sorceror/smt/smtas.cl2",
	"smtat.cl2":    "plrgfx/sorceror/smt/smtat.cl2",
	"smtaw.cl2":    "plrgfx/sorceror/smt/smtaw.cl2",
	"smtfm.cl2":    "plrgfx/sorceror/smt/smtfm.cl2",
	"smtht.cl2":    "plrgfx/sorceror/smt/smtht.cl2",
	"smtlm.cl2":    "plrgfx/sorceror/smt/smtlm.cl2",
	"smtqm.cl2":    "plrgfx/sorceror/smt/smtqm.cl2",
	"smtst.cl2":    "plrgfx/sorceror/smt/smtst.cl2",
	"smtwl.cl2":    "plrgfx/sorceror/smt/smtwl.cl2",
	"smuas.cl2":    "plrgfx/sorceror/smu/smuas.cl2",
	"smuat.cl2":    "plrgfx/sorceror/smu/smuat.cl2",
	"smuaw.cl2":    "plrgfx/sorceror/smu/smuaw.cl2",
	"smubl.cl2":    "plrgfx/sorceror/smu/smubl.cl2",
	"smufm.cl2":    "plrgfx/sorceror/smu/smufm.cl2",
	"smuht.cl2":    "plrgfx/sorceror/smu/smuht.cl2",
	"smulm.cl2":    "plrgfx/sorceror/smu/smulm.cl2",
	"smuqm.cl2":    "plrgfx/sorceror/smu/smuqm.cl2",
	"smust.cl2":    "plrgfx/sorceror/smu/smust.cl2",
	"smuwl.cl2":    "plrgfx/sorceror/smu/smuwl.cl2",
	"whaas.cl2":    "plrgfx/warrior/wha/whaas.cl2",
	"whaat.cl2":    "plrgfx/warrior/wha/whaat.cl2",
	"whaaw.cl2":    "plrgfx/warrior/wha/whaaw.cl2",
	"whafm.cl2":    "plrgfx/warrior/wha/whafm.cl2",
	"whaht.cl2":    "plrgfx/warrior/wha/whaht.cl2",
	"whalm.cl2":    "plrgfx/warrior/wha/whalm.cl2",
	"whaqm.cl2":    "plrgfx/warrior/wha/whaqm.cl2",
	"whast.cl2":    "plrgfx/warrior/wha/whast.cl2",
	"whawl.cl2":    "plrgfx/warrior/wha/whawl.cl2",
	"whbas.cl2":    "plrgfx/warrior/whb/whbas.cl2",
	"whbat.cl2":    "plrgfx/warrior/whb/whbat.cl2",
	"whbaw.cl2":    "plrgfx/warrior/whb/whbaw.cl2",
	"whbfm.cl2":    "plrgfx/warrior/whb/whbfm.cl2",
	"whbht.cl2":    "plrgfx/warrior/whb/whbht.cl2",
	"whblm.cl2":    "plrgfx/warrior/whb/whblm.cl2",
	"whbqm.cl2":    "plrgfx/warrior/whb/whbqm.cl2",
	"whbst.cl2":    "plrgfx/warrior/whb/whbst.cl2",
	"whbwl.cl2":    "plrgfx/warrior/whb/whbwl.cl2",
	"whdas.cl2":    "plrgfx/warrior/whd/whdas.cl2",
	"whdat.cl2":    "plrgfx/warrior/whd/whdat.cl2",
	"whdaw.cl2":    "plrgfx/warrior/whd/whdaw.cl2",
	"whdbl.cl2":    "plrgfx/warrior/whd/whdbl.cl2",
	"whdfm.cl2":    "plrgfx/warrior/whd/whdfm.cl2",
	"whdht.cl2":    "plrgfx/warrior/whd/whdht.cl2",
	"whdlm.cl2":    "plrgfx/warrior/whd/whdlm.cl2",
	"whdqm.cl2":    "plrgfx/warrior/whd/whdqm.cl2",
	"whdst.cl2":    "plrgfx/warrior/whd/whdst.cl2",
	"whdwl.cl2":    "plrgfx/warrior/whd/whdwl.cl2",
	"whhas.cl2":    "plrgfx/warrior/whh/whhas.cl2",
	"whhat.cl2":    "plrgfx/warrior/whh/whhat.cl2",
	"whhaw.cl2":    "plrgfx/warrior/whh/whhaw.cl2",
	"whhbl.cl2":    "plrgfx/warrior/whh/whhbl.cl2",
	"whhfm.cl2":    "plrgfx/warrior/whh/whhfm.cl2",
	"whhht.cl2":    "plrgfx/warrior/whh/whhht.cl2",
	"whhlm.cl2":    "plrgfx/warrior/whh/whhlm.cl2",
	"whhqm.cl2":    "plrgfx/warrior/whh/whhqm.cl2",
	"whhst.cl2":    "plrgfx/warrior/whh/whhst.cl2",
	"whhwl.cl2":    "plrgfx/warrior/whh/whhwl.cl2",
	"whmas.cl2":    "plrgfx/warrior/whm/whmas.cl2",
	"whmat.cl2":    "plrgfx/warrior/whm/whmat.cl2",
	"whmaw.cl2":    "plrgfx/warrior/whm/whmaw.cl2",
	"whmfm.cl2":    "plrgfx/warrior/whm/whmfm.cl2",
	"whmht.cl2":    "plrgfx/warrior/whm/whmht.cl2",
	"whmlm.cl2":    "plrgfx/warrior/whm/whmlm.cl2",
	"whmqm.cl2":    "plrgfx/warrior/whm/whmqm.cl2",
	"whmst.cl2":    "plrgfx/warrior/whm/whmst.cl2",
	"whmwl.cl2":    "plrgfx/warrior/whm/whmwl.cl2",
	"whnas.cl2":    "plrgfx/warrior/whn/whnas.cl2",
	"whnat.cl2":    "plrgfx/warrior/whn/whnat.cl2",
	"whnaw.cl2":    "plrgfx/warrior/whn/whnaw.cl2",
	"whndt.cl2":    "plrgfx/warrior/whn/whndt.cl2",
	"whnfm.cl2":    "plrgfx/warrior/whn/whnfm.cl2",
	"whnht.cl2":    "plrgfx/warrior/whn/whnht.cl2",
	"whnlm.cl2":    "plrgfx/warrior/whn/whnlm.cl2",
	"whnqm.cl2":    "plrgfx/warrior/whn/whnqm.cl2",
	"whnst.cl2":    "plrgfx/warrior/whn/whnst.cl2",
	"whnwl.cl2":    "plrgfx/warrior/whn/whnwl.cl2",
	"whsas.cl2":    "plrgfx/warrior/whs/whsas.cl2",
	"whsat.cl2":    "plrgfx/warrior/whs/whsat.cl2",
	"whsaw.cl2":    "plrgfx/warrior/whs/whsaw.cl2",
	"whsfm.cl2":    "plrgfx/warrior/whs/whsfm.cl2",
	"whsht.cl2":    "plrgfx/warrior/whs/whsht.cl2",
	"whslm.cl2":    "plrgfx/warrior/whs/whslm.cl2",
	"whsqm.cl2":    "plrgfx/warrior/whs/whsqm.cl2",
	"whsst.cl2":    "plrgfx/warrior/whs/whsst.cl2",
	"whswl.cl2":    "plrgfx/warrior/whs/whswl.cl2",
	"whtas.cl2":    "plrgfx/warrior/wht/whtas.cl2",
	"whtat.cl2":    "plrgfx/warrior/wht/whtat.cl2",
	"whtaw.cl2":    "plrgfx/warrior/wht/whtaw.cl2",
	"whtfm.cl2":    "plrgfx/warrior/wht/whtfm.cl2",
	"whtht.cl2":    "plrgfx/warrior/wht/whtht.cl2",
	"whtlm.cl2":    "plrgfx/warrior/wht/whtlm.cl2",
	"whtqm.cl2":    "plrgfx/warrior/wht/whtqm.cl2",
	"whtst.cl2":    "plrgfx/warrior/wht/whtst.cl2",
	"whtwl.cl2":    "plrgfx/warrior/wht/whtwl.cl2",
	"whuas.cl2":    "plrgfx/warrior/whu/whuas.cl2",
	"whuat.cl2":    "plrgfx/warrior/whu/whuat.cl2",
	"whuaw.cl2":    "plrgfx/warrior/whu/whuaw.cl2",
	"whubl.cl2":    "plrgfx/warrior/whu/whubl.cl2",
	"whufm.cl2":    "plrgfx/warrior/whu/whufm.cl2",
	"whuht.cl2":    "plrgfx/warrior/whu/whuht.cl2",
	"whulm.cl2":    "plrgfx/warrior/whu/whulm.cl2",
	"whuqm.cl2":    "plrgfx/warrior/whu/whuqm.cl2",
	"whust.cl2":    "plrgfx/warrior/whu/whust.cl2",
	"whuwl.cl2":    "plrgfx/warrior/whu/whuwl.cl2",
	"wlaas.cl2":    "plrgfx/warrior/wla/wlaas.cl2",
	"wlaat.cl2":    "plrgfx/warrior/wla/wlaat.cl2",
	"wlaaw.cl2":    "plrgfx/warrior/wla/wlaaw.cl2",
	"wlafm.cl2":    "plrgfx/warrior/wla/wlafm.cl2",
	"wlaht.cl2":    "plrgfx/warrior/wla/wlaht.cl2",
	"wlalm.cl2":    "plrgfx/warrior/wla/wlalm.cl2",
	"wlaqm.cl2":    "plrgfx/warrior/wla/wlaqm.cl2",
	"wlast.cl2":    "plrgfx/warrior/wla/wlast.cl2",
	"wlawl.cl2":    "plrgfx/warrior/wla/wlawl.cl2",
	"wlbas.cl2":    "plrgfx/warrior/wlb/wlbas.cl2",
	"wlbat.cl2":    "plrgfx/warrior/wlb/wlbat.cl2",
	"wlbaw.cl2":    "plrgfx/warrior/wlb/wlbaw.cl2",
	"wlbfm.cl2":    "plrgfx/warrior/wlb/wlbfm.cl2",
	"wlbht.cl2":    "plrgfx/warrior/wlb/wlbht.cl2",
	"wlblm.cl2":    "plrgfx/warrior/wlb/wlblm.cl2",
	"wlbqm.cl2":    "plrgfx/warrior/wlb/wlbqm.cl2",
	"wlbst.cl2":    "plrgfx/warrior/wlb/wlbst.cl2",
	"wlbwl.cl2":    "plrgfx/warrior/wlb/wlbwl.cl2",
	"wldas.cl2":    "plrgfx/warrior/wld/wldas.cl2",
	"wldat.cl2":    "plrgfx/warrior/wld/wldat.cl2",
	"wldaw.cl2":    "plrgfx/warrior/wld/wldaw.cl2",
	"wldbl.cl2":    "plrgfx/warrior/wld/wldbl.cl2",
	"wldfm.cl2":    "plrgfx/warrior/wld/wldfm.cl2",
	"wldht.cl2":    "plrgfx/warrior/wld/wldht.cl2",
	"wldlm.cl2":    "plrgfx/warrior/wld/wldlm.cl2",
	"wldqm.cl2":    "plrgfx/warrior/wld/wldqm.cl2",
	"wldst.cl2":    "plrgfx/warrior/wld/wldst.cl2",
	"wldwl.cl2":    "plrgfx/warrior/wld/wldwl.cl2",
	"wlhas.cl2":    "plrgfx/warrior/wlh/wlhas.cl2",
	"wlhat.cl2":    "plrgfx/warrior/wlh/wlhat.cl2",
	"wlhaw.cl2":    "plrgfx/warrior/wlh/wlhaw.cl2",
	"wlhbl.cl2":    "plrgfx/warrior/wlh/wlhbl.cl2",
	"wlhfm.cl2":    "plrgfx/warrior/wlh/wlhfm.cl2",
	"wlhht.cl2":    "plrgfx/warrior/wlh/wlhht.cl2",
	"wlhlm.cl2":    "plrgfx/warrior/wlh/wlhlm.cl2",
	"wlhqm.cl2":    "plrgfx/warrior/wlh/wlhqm.cl2",
	"wlhst.cl2":    "plrgfx/warrior/wlh/wlhst.cl2",
	"wlhwl.cl2":    "plrgfx/warrior/wlh/wlhwl.cl2",
	"wlmas.cl2":    "plrgfx/warrior/wlm/wlmas.cl2",
	"wlmat.cl2":    "plrgfx/warrior/wlm/wlmat.cl2",
	"wlmaw.cl2":    "plrgfx/warrior/wlm/wlmaw.cl2",
	"wlmfm.cl2":    "plrgfx/warrior/wlm/wlmfm.cl2",
	"wlmht.cl2":    "plrgfx/warrior/wlm/wlmht.cl2",
	"wlmlm.cl2":    "plrgfx/warrior/wlm/wlmlm.cl2",
	"wlmqm.cl2":    "plrgfx/warrior/wlm/wlmqm.cl2",
	"wlmst.cl2":    "plrgfx/warrior/wlm/wlmst.cl2",
	"wlmwl.cl2":    "plrgfx/warrior/wlm/wlmwl.cl2",
	"wlnas.cl2":    "plrgfx/warrior/wln/wlnas.cl2",
	"wlnat.cl2":    "plrgfx/warrior/wln/wlnat.cl2",
	"wlnaw.cl2":    "plrgfx/warrior/wln/wlnaw.cl2",
	"wlndt.cl2":    "plrgfx/warrior/wln/wlndt.cl2",
	"wlnfm.cl2":    "plrgfx/warrior/wln/wlnfm.cl2",
	"wlnht.cl2":    "plrgfx/warrior/wln/wlnht.cl2",
	"wlnlm.cl2":    "plrgfx/warrior/wln/wlnlm.cl2",
	"wlnqm.cl2":    "plrgfx/warrior/wln/wlnqm.cl2",
	"wlnst.cl2":    "plrgfx/warrior/wln/wlnst.cl2",
	"wlnwl.cl2":    "plrgfx/warrior/wln/wlnwl.cl2",
	"wlsas.cl2":    "plrgfx/warrior/wls/wlsas.cl2",
	"wlsat.cl2":    "plrgfx/warrior/wls/wlsat.cl2",
	"wlsaw.cl2":    "plrgfx/warrior/wls/wlsaw.cl2",
	"wlsfm.cl2":    "plrgfx/warrior/wls/wlsfm.cl2",
	"wlsht.cl2":    "plrgfx/warrior/wls/wlsht.cl2",
	"wlslm.cl2":    "plrgfx/warrior/wls/wlslm.cl2",
	"wlsqm.cl2":    "plrgfx/warrior/wls/wlsqm.cl2",
	"wlsst.cl2":    "plrgfx/warrior/wls/wlsst.cl2",
	"wlswl.cl2":    "plrgfx/warrior/wls/wlswl.cl2",
	"wltas.cl2":    "plrgfx/warrior/wlt/wltas.cl2",
	"wltat.cl2":    "plrgfx/warrior/wlt/wltat.cl2",
	"wltaw.cl2":    "plrgfx/warrior/wlt/wltaw.cl2",
	"wltfm.cl2":    "plrgfx/warrior/wlt/wltfm.cl2",
	"wltht.cl2":    "plrgfx/warrior/wlt/wltht.cl2",
	"wltlm.cl2":    "plrgfx/warrior/wlt/wltlm.cl2",
	"wltqm.cl2":    "plrgfx/warrior/wlt/wltqm.cl2",
	"wltst.cl2":    "plrgfx/warrior/wlt/wltst.cl2",
	"wltwl.cl2":    "plrgfx/warrior/wlt/wltwl.cl2",
	"wluas.cl2":    "plrgfx/warrior/wlu/wluas.cl2",
	"wluat.cl2":    "plrgfx/warrior/wlu/wluat.cl2",
	"wluaw.cl2":    "plrgfx/warrior/wlu/wluaw.cl2",
	"wlubl.cl2":    "plrgfx/warrior/wlu/wlubl.cl2",
	"wlufm.cl2":    "plrgfx/warrior/wlu/wlufm.cl2",
	"wluht.cl2":    "plrgfx/warrior/wlu/wluht.cl2",
	"wlulm.cl2":    "plrgfx/warrior/wlu/wlulm.cl2",
	"wluqm.cl2":    "plrgfx/warrior/wlu/wluqm.cl2",
	"wlust.cl2":    "plrgfx/warrior/wlu/wlust.cl2",
	"wluwl.cl2":    "plrgfx/warrior/wlu/wluwl.cl2",
	"wmaas.cl2":    "plrgfx/warrior/wma/wmaas.cl2",
	"wmaat.cl2":    "plrgfx/warrior/wma/wmaat.cl2",
	"wmaaw.cl2":    "plrgfx/warrior/wma/wmaaw.cl2",
	"wmafm.cl2":    "plrgfx/warrior/wma/wmafm.cl2",
	"wmaht.cl2":    "plrgfx/warrior/wma/wmaht.cl2",
	"wmalm.cl2":    "plrgfx/warrior/wma/wmalm.cl2",
	"wmaqm.cl2":    "plrgfx/warrior/wma/wmaqm.cl2",
	"wmast.cl2":    "plrgfx/warrior/wma/wmast.cl2",
	"wmawl.cl2":    "plrgfx/warrior/wma/wmawl.cl2",
	"wmbas.cl2":    "plrgfx/warrior/wmb/wmbas.cl2",
	"wmbat.cl2":    "plrgfx/warrior/wmb/wmbat.cl2",
	"wmbaw.cl2":    "plrgfx/warrior/wmb/wmbaw.cl2",
	"wmbfm.cl2":    "plrgfx/warrior/wmb/wmbfm.cl2",
	"wmbht.cl2":    "plrgfx/warrior/wmb/wmbht.cl2",
	"wmblm.cl2":    "plrgfx/warrior/wmb/wmblm.cl2",
	"wmbqm.cl2":    "plrgfx/warrior/wmb/wmbqm.cl2",
	"wmbst.cl2":    "plrgfx/warrior/wmb/wmbst.cl2",
	"wmbwl.cl2":    "plrgfx/warrior/wmb/wmbwl.cl2",
	"wmdas.cl2":    "plrgfx/warrior/wmd/wmdas.cl2",
	"wmdat.cl2":    "plrgfx/warrior/wmd/wmdat.cl2",
	"wmdaw.cl2":    "plrgfx/warrior/wmd/wmdaw.cl2",
	"wmdbl.cl2":    "plrgfx/warrior/wmd/wmdbl.cl2",
	"wmdfm.cl2":    "plrgfx/warrior/wmd/wmdfm.cl2",
	"wmdht.cl2":    "plrgfx/warrior/wmd/wmdht.cl2",
	"wmdlm.cl2":    "plrgfx/warrior/wmd/wmdlm.cl2",
	"wmdqm.cl2":    "plrgfx/warrior/wmd/wmdqm.cl2",
	"wmdst.cl2":    "plrgfx/warrior/wmd/wmdst.cl2",
	"wmdwl.cl2":    "plrgfx/warrior/wmd/wmdwl.cl2",
	"wmhas.cl2":    "plrgfx/warrior/wmh/wmhas.cl2",
	"wmhat.cl2":    "plrgfx/warrior/wmh/wmhat.cl2",
	"wmhaw.cl2":    "plrgfx/warrior/wmh/wmhaw.cl2",
	"wmhbl.cl2":    "plrgfx/warrior/wmh/wmhbl.cl2",
	"wmhfm.cl2":    "plrgfx/warrior/wmh/wmhfm.cl2",
	"wmhht.cl2":    "plrgfx/warrior/wmh/wmhht.cl2",
	"wmhlm.cl2":    "plrgfx/warrior/wmh/wmhlm.cl2",
	"wmhqm.cl2":    "plrgfx/warrior/wmh/wmhqm.cl2",
	"wmhst.cl2":    "plrgfx/warrior/wmh/wmhst.cl2",
	"wmhwl.cl2":    "plrgfx/warrior/wmh/wmhwl.cl2",
	"wmmas.cl2":    "plrgfx/warrior/wmm/wmmas.cl2",
	"wmmat.cl2":    "plrgfx/warrior/wmm/wmmat.cl2",
	"wmmaw.cl2":    "plrgfx/warrior/wmm/wmmaw.cl2",
	"wmmfm.cl2":    "plrgfx/warrior/wmm/wmmfm.cl2",
	"wmmht.cl2":    "plrgfx/warrior/wmm/wmmht.cl2",
	"wmmlm.cl2":    "plrgfx/warrior/wmm/wmmlm.cl2",
	"wmmqm.cl2":    "plrgfx/warrior/wmm/wmmqm.cl2",
	"wmmst.cl2":    "plrgfx/warrior/wmm/wmmst.cl2",
	"wmmwl.cl2":    "plrgfx/warrior/wmm/wmmwl.cl2",
	"wmnas.cl2":    "plrgfx/warrior/wmn/wmnas.cl2",
	"wmnat.cl2":    "plrgfx/warrior/wmn/wmnat.cl2",
	"wmnaw.cl2":    "plrgfx/warrior/wmn/wmnaw.cl2",
	"wmndt.cl2":    "plrgfx/warrior/wmn/wmndt.cl2",
	"wmnfm.cl2":    "plrgfx/warrior/wmn/wmnfm.cl2",
	"wmnht.cl2":    "plrgfx/warrior/wmn/wmnht.cl2",
	"wmnlm.cl2":    "plrgfx/warrior/wmn/wmnlm.cl2",
	"wmnqm.cl2":    "plrgfx/warrior/wmn/wmnqm.cl2",
	"wmnst.cl2":    "plrgfx/warrior/wmn/wmnst.cl2",
	"wmnwl.cl2":    "plrgfx/warrior/wmn/wmnwl.cl2",
	"wmsas.cl2":    "plrgfx/warrior/wms/wmsas.cl2",
	"wmsat.cl2":    "plrgfx/warrior/wms/wmsat.cl2",
	"wmsaw.cl2":    "plrgfx/warrior/wms/wmsaw.cl2",
	"wmsfm.cl2":    "plrgfx/warrior/wms/wmsfm.cl2",
	"wmsht.cl2":    "plrgfx/warrior/wms/wmsht.cl2",
	"wmslm.cl2":    "plrgfx/warrior/wms/wmslm.cl2",
	"wmsqm.cl2":    "plrgfx/warrior/wms/wmsqm.cl2",
	"wmsst.cl2":    "plrgfx/warrior/wms/wmsst.cl2",
	"wmswl.cl2":    "plrgfx/warrior/wms/wmswl.cl2",
	"wmtas.cl2":    "plrgfx/warrior/wmt/wmtas.cl2",
	"wmtat.cl2":    "plrgfx/warrior/wmt/wmtat.cl2",
	"wmtaw.cl2":    "plrgfx/warrior/wmt/wmtaw.cl2",
	"wmtfm.cl2":    "plrgfx/warrior/wmt/wmtfm.cl2",
	"wmtht.cl2":    "plrgfx/warrior/wmt/wmtht.cl2",
	"wmtlm.cl2":    "plrgfx/warrior/wmt/wmtlm.cl2",
	"wmtqm.cl2":    "plrgfx/warrior/wmt/wmtqm.cl2",
	"wmtst.cl2":    "plrgfx/warrior/wmt/wmtst.cl2",
	"wmtwl.cl2":    "plrgfx/warrior/wmt/wmtwl.cl2",
	"wmuas.cl2":    "plrgfx/warrior/wmu/wmuas.cl2",
	"wmuat.cl2":    "plrgfx/warrior/wmu/wmuat.cl2",
	"wmuaw.cl2":    "plrgfx/warrior/wmu/wmuaw.cl2",
	"wmubl.cl2":    "plrgfx/warrior/wmu/wmubl.cl2",
	"wmufm.cl2":    "plrgfx/warrior/wmu/wmufm.cl2",
	"wmuht.cl2":    "plrgfx/warrior/wmu/wmuht.cl2",
	"wmulm.cl2":    "plrgfx/warrior/wmu/wmulm.cl2",
	"wmuqm.cl2":    "plrgfx/warrior/wmu/wmuqm.cl2",
	"wmust.cl2":    "plrgfx/warrior/wmu/wmust.cl2",
	"wmuwl.cl2":    "plrgfx/warrior/wmu/wmuwl.cl2",
}
