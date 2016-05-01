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
	// CEL images.
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
	"missiles/flames1.cel": {
		// The contents of frame 0 at offset 0x18 in flames1.cel starts with the
		// following header: 0A 00 2A 00 7D 00 00 00 00 00.
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
}

// RelPaths maps from CEL file names to "diabdat.mpq" relative paths.
var RelPaths = map[string]string{
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
	"flamel10.cel": "missiles/flamel10.cel",
	"flamel11.cel": "missiles/flamel11.cel",
	"flamel12.cel": "missiles/flamel12.cel",
	"flamel13.cel": "missiles/flamel13.cel",
	"flamel14.cel": "missiles/flamel14.cel",
	"flamel15.cel": "missiles/flamel15.cel",
	"flamel16.cel": "missiles/flamel16.cel",
	"flamel2.cel":  "missiles/flamel2.cel",
	"flamel3.cel":  "missiles/flamel3.cel",
	"flamel4.cel":  "missiles/flamel4.cel",
	"flamel5.cel":  "missiles/flamel5.cel",
	"flamel6.cel":  "missiles/flamel6.cel",
	"flamel7.cel":  "missiles/flamel7.cel",
	"flamel8.cel":  "missiles/flamel8.cel",
	"flamel9.cel":  "missiles/flamel9.cel",
	"flames1.cel":  "missiles/flames1.cel",
	"flames10.cel": "missiles/flames10.cel",
	"flames11.cel": "missiles/flames11.cel",
	"flames12.cel": "missiles/flames12.cel",
	"flames13.cel": "missiles/flames13.cel",
	"flames14.cel": "missiles/flames14.cel",
	"flames15.cel": "missiles/flames15.cel",
	"flames16.cel": "missiles/flames16.cel",
	"flames2.cel":  "missiles/flames2.cel",
	"flames3.cel":  "missiles/flames3.cel",
	"flames4.cel":  "missiles/flames4.cel",
	"flames5.cel":  "missiles/flames5.cel",
	"flames6.cel":  "missiles/flames6.cel",
	"flames7.cel":  "missiles/flames7.cel",
	"flames8.cel":  "missiles/flames8.cel",
	"flames9.cel":  "missiles/flames9.cel",
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
}
