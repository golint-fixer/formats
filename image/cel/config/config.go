// Package config specifies the data required for decoding CEL images.
package config

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
}

// TODO: Remove unknown once no longer needed.

// unknown tracks unknown data.
const unknown = 0

// Confs specifies the data required for decoding
var Confs = map[string]*Config{
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
	"gendata/cut2.cel":              {W: unknown, H: unknown},
	"gendata/cut3.cel":              {W: unknown, H: unknown},
	"gendata/cut4.cel":              {W: unknown, H: unknown},
	"gendata/cutgate.cel":           {W: unknown, H: unknown},
	"gendata/cutl1d.cel":            {W: unknown, H: unknown},
	"gendata/cutportl.cel":          {W: unknown, H: unknown},
	"gendata/cutportr.cel":          {W: unknown, H: unknown},
	"gendata/cutstart.cel":          {W: unknown, H: unknown},
	"gendata/cuttt.cel":             {W: unknown, H: unknown},
	"gendata/quotes.cel":            {W: unknown, H: unknown},
	"items/armor2.cel":              {W: unknown, H: unknown},
	"items/axe.cel":                 {W: unknown, H: unknown},
	"items/axeflip.cel":             {W: unknown, H: unknown},
	"items/bldstn.cel":              {W: unknown, H: unknown},
	"items/bottle.cel":              {W: unknown, H: unknown},
	"items/bow.cel":                 {W: unknown, H: unknown},
	"items/cleaver.cel":             {W: unknown, H: unknown},
	"items/crownf.cel":              {W: unknown, H: unknown},
	"items/duricons.cel":            {W: unknown, H: unknown},
	"items/fanvil.cel":              {W: unknown, H: unknown},
	"items/fbook.cel":               {W: unknown, H: unknown},
	"items/fbow.cel":                {W: unknown, H: unknown},
	"items/fbrain.cel":              {W: unknown, H: unknown},
	"items/fbttle.cel":              {W: unknown, H: unknown},
	"items/fbttlebb.cel":            {W: unknown, H: unknown},
	"items/fbttlebl.cel":            {W: unknown, H: unknown},
	"items/fbttlebr.cel":            {W: unknown, H: unknown},
	"items/fbttleby.cel":            {W: unknown, H: unknown},
	"items/fbttledb.cel":            {W: unknown, H: unknown},
	"items/fbttledy.cel":            {W: unknown, H: unknown},
	"items/fbttleor.cel":            {W: unknown, H: unknown},
	"items/fbttlewh.cel":            {W: unknown, H: unknown},
	"items/fear.cel":                {W: unknown, H: unknown},
	"items/feye.cel":                {W: unknown, H: unknown},
	"items/fheart.cel":              {W: unknown, H: unknown},
	"items/flazstaf.cel":            {W: unknown, H: unknown},
	"items/fmush.cel":               {W: unknown, H: unknown},
	"items/food.cel":                {W: unknown, H: unknown},
	"items/fplatear.cel":            {W: unknown, H: unknown},
	"items/goldflip.cel":            {W: unknown, H: unknown},
	"items/helmut.cel":              {W: unknown, H: unknown},
	"items/innsign.cel":             {W: unknown, H: unknown},
	"items/larmor.cel":              {W: unknown, H: unknown},
	"items/mace.cel":                {W: unknown, H: unknown},
	"items/manaflip.cel":            {W: unknown, H: unknown},
	"items/map/mapz0000.cel":        {W: unknown, H: unknown},
	"items/map/mapz0001.cel":        {W: unknown, H: unknown},
	"items/map/mapz0002.cel":        {W: unknown, H: unknown},
	"items/map/mapz0003.cel":        {W: unknown, H: unknown},
	"items/map/mapz0004.cel":        {W: unknown, H: unknown},
	"items/map/mapz0005.cel":        {W: unknown, H: unknown},
	"items/map/mapz0006.cel":        {W: unknown, H: unknown},
	"items/map/mapz0007.cel":        {W: unknown, H: unknown},
	"items/map/mapz0008.cel":        {W: unknown, H: unknown},
	"items/map/mapz0009.cel":        {W: unknown, H: unknown},
	"items/map/mapz0010.cel":        {W: unknown, H: unknown},
	"items/map/mapz0011.cel":        {W: unknown, H: unknown},
	"items/map/mapz0012.cel":        {W: unknown, H: unknown},
	"items/map/mapz0013.cel":        {W: unknown, H: unknown},
	"items/map/mapz0014.cel":        {W: unknown, H: unknown},
	"items/map/mapz0015.cel":        {W: unknown, H: unknown},
	"items/map/mapz0016.cel":        {W: unknown, H: unknown},
	"items/map/mapz0017.cel":        {W: unknown, H: unknown},
	"items/map/mapz0018.cel":        {W: unknown, H: unknown},
	"items/map/mapz0019.cel":        {W: unknown, H: unknown},
	"items/map/mapz0020.cel":        {W: unknown, H: unknown},
	"items/map/mapz0021.cel":        {W: unknown, H: unknown},
	"items/map/mapz0022.cel":        {W: unknown, H: unknown},
	"items/map/mapz0023.cel":        {W: unknown, H: unknown},
	"items/map/mapz0024.cel":        {W: unknown, H: unknown},
	"items/map/mapz0025.cel":        {W: unknown, H: unknown},
	"items/map/mapz0026.cel":        {W: unknown, H: unknown},
	"items/map/mapz0027.cel":        {W: unknown, H: unknown},
	"items/map/mapz0028.cel":        {W: unknown, H: unknown},
	"items/map/mapz0029.cel":        {W: unknown, H: unknown},
	"items/map/mapz0030.cel":        {W: unknown, H: unknown},
	"items/map/mapzdoom.cel":        {W: unknown, H: unknown},
	"items/ring.cel":                {W: unknown, H: unknown},
	"items/rock.cel":                {W: unknown, H: unknown},
	"items/scroll.cel":              {W: unknown, H: unknown},
	"items/shield.cel":              {W: unknown, H: unknown},
	"items/staff.cel":               {W: unknown, H: unknown},
	"items/swrdflip.cel":            {W: unknown, H: unknown},
	"items/wand.cel":                {W: unknown, H: unknown},
	"items/wshield.cel":             {W: unknown, H: unknown},
	"levels/l1data/l1.cel":          {W: unknown, H: unknown},
	"levels/l1data/l1s.cel":         {W: unknown, H: unknown},
	"levels/l2data/l2.cel":          {W: unknown, H: unknown},
	"levels/l2data/l2s.cel":         {W: unknown, H: unknown},
	"levels/l3data/l3.cel":          {W: unknown, H: unknown},
	"levels/l4data/l4.cel":          {W: unknown, H: unknown},
	"levels/towndata/town.cel":      {W: unknown, H: unknown},
	"levels/towndata/towns.cel":     {W: unknown, H: unknown},
	"missiles/flamel1.cel":          {W: unknown, H: unknown},
	"missiles/flamel10.cel":         {W: unknown, H: unknown},
	"missiles/flamel11.cel":         {W: unknown, H: unknown},
	"missiles/flamel12.cel":         {W: unknown, H: unknown},
	"missiles/flamel13.cel":         {W: unknown, H: unknown},
	"missiles/flamel14.cel":         {W: unknown, H: unknown},
	"missiles/flamel15.cel":         {W: unknown, H: unknown},
	"missiles/flamel16.cel":         {W: unknown, H: unknown},
	"missiles/flamel2.cel":          {W: unknown, H: unknown},
	"missiles/flamel3.cel":          {W: unknown, H: unknown},
	"missiles/flamel4.cel":          {W: unknown, H: unknown},
	"missiles/flamel5.cel":          {W: unknown, H: unknown},
	"missiles/flamel6.cel":          {W: unknown, H: unknown},
	"missiles/flamel7.cel":          {W: unknown, H: unknown},
	"missiles/flamel8.cel":          {W: unknown, H: unknown},
	"missiles/flamel9.cel":          {W: unknown, H: unknown},
	"missiles/flames1.cel":          {W: unknown, H: unknown},
	"missiles/flames10.cel":         {W: unknown, H: unknown},
	"missiles/flames11.cel":         {W: unknown, H: unknown},
	"missiles/flames12.cel":         {W: unknown, H: unknown},
	"missiles/flames13.cel":         {W: unknown, H: unknown},
	"missiles/flames14.cel":         {W: unknown, H: unknown},
	"missiles/flames15.cel":         {W: unknown, H: unknown},
	"missiles/flames16.cel":         {W: unknown, H: unknown},
	"missiles/flames2.cel":          {W: unknown, H: unknown},
	"missiles/flames3.cel":          {W: unknown, H: unknown},
	"missiles/flames4.cel":          {W: unknown, H: unknown},
	"missiles/flames5.cel":          {W: unknown, H: unknown},
	"missiles/flames6.cel":          {W: unknown, H: unknown},
	"missiles/flames7.cel":          {W: unknown, H: unknown},
	"missiles/flames8.cel":          {W: unknown, H: unknown},
	"missiles/flames9.cel":          {W: unknown, H: unknown},
	"missiles/flaml1.cel":           {W: unknown, H: unknown},
	"missiles/flaml2.cel":           {W: unknown, H: unknown},
	"missiles/flaml3.cel":           {W: unknown, H: unknown},
	"missiles/flaml4.cel":           {W: unknown, H: unknown},
	"missiles/flaml5.cel":           {W: unknown, H: unknown},
	"missiles/flaml6.cel":           {W: unknown, H: unknown},
	"missiles/flaml7.cel":           {W: unknown, H: unknown},
	"missiles/flaml8.cel":           {W: unknown, H: unknown},
	"missiles/flams1.cel":           {W: unknown, H: unknown},
	"missiles/flams2.cel":           {W: unknown, H: unknown},
	"missiles/flams3.cel":           {W: unknown, H: unknown},
	"missiles/flams4.cel":           {W: unknown, H: unknown},
	"missiles/flams5.cel":           {W: unknown, H: unknown},
	"missiles/flams6.cel":           {W: unknown, H: unknown},
	"missiles/flams7.cel":           {W: unknown, H: unknown},
	"missiles/flams8.cel":           {W: unknown, H: unknown},
	"missiles/mindmace.cel":         {W: unknown, H: unknown},
	"missiles/sentfr.cel":           {W: unknown, H: unknown},
	"missiles/sentout.cel":          {W: unknown, H: unknown},
	"missiles/sentup.cel":           {W: unknown, H: unknown},
	"monsters/acid/acidpud.cel":     {W: unknown, H: unknown},
	"monsters/magma/magball1.cel":   {W: unknown, H: unknown},
	"monsters/magma/magball2.cel":   {W: unknown, H: unknown},
	"monsters/magma/magball3.cel":   {W: unknown, H: unknown},
	"monsters/magma/magball4.cel":   {W: unknown, H: unknown},
	"monsters/magma/magball5.cel":   {W: unknown, H: unknown},
	"monsters/magma/magball6.cel":   {W: unknown, H: unknown},
	"monsters/magma/magball7.cel":   {W: unknown, H: unknown},
	"monsters/magma/magball8.cel":   {W: unknown, H: unknown},
	"monsters/magma/magblos.cel":    {W: unknown, H: unknown},
	"monsters/rhino/rhinos1.cel":    {W: unknown, H: unknown},
	"monsters/rhino/rhinos2.cel":    {W: unknown, H: unknown},
	"monsters/rhino/rhinos3.cel":    {W: unknown, H: unknown},
	"monsters/rhino/rhinos4.cel":    {W: unknown, H: unknown},
	"monsters/rhino/rhinos5.cel":    {W: unknown, H: unknown},
	"monsters/rhino/rhinos6.cel":    {W: unknown, H: unknown},
	"monsters/rhino/rhinos7.cel":    {W: unknown, H: unknown},
	"monsters/rhino/rhinos8.cel":    {W: unknown, H: unknown},
	"monsters/succ/flare.cel":       {W: unknown, H: unknown},
	"monsters/succ/flarexp.cel":     {W: unknown, H: unknown},
	"monsters/thin/lghning.cel":     {W: unknown, H: unknown},
	"monsters/unrav/unravw.cel":     {W: unknown, H: unknown},
	"objects/altboy.cel":            {W: unknown, H: unknown},
	"objects/angel.cel":             {W: unknown, H: unknown},
	"objects/armstand.cel":          {W: unknown, H: unknown},
	"objects/banner.cel":            {W: unknown, H: unknown},
	"objects/barrel.cel":            {W: unknown, H: unknown},
	"objects/barrelex.cel":          {W: unknown, H: unknown},
	"objects/bcase.cel":             {W: unknown, H: unknown},
	"objects/bkslbrnt.cel":          {W: unknown, H: unknown},
	"objects/bkurns.cel":            {W: unknown, H: unknown},
	"objects/bloodfnt.cel":          {W: unknown, H: unknown},
	"objects/book1.cel":             {W: unknown, H: unknown},
	"objects/book2.cel":             {W: unknown, H: unknown},
	"objects/bshelf.cel":            {W: unknown, H: unknown},
	"objects/burncros.cel":          {W: unknown, H: unknown},
	"objects/candlabr.cel":          {W: unknown, H: unknown},
	"objects/candle.cel":            {W: unknown, H: unknown},
	"objects/candle2.cel":           {W: unknown, H: unknown},
	"objects/cauldren.cel":          {W: unknown, H: unknown},
	"objects/chest1.cel":            {W: unknown, H: unknown},
	"objects/chest2.cel":            {W: unknown, H: unknown},
	"objects/chest3.cel":            {W: unknown, H: unknown},
	"objects/cruxsk1.cel":           {W: unknown, H: unknown},
	"objects/cruxsk2.cel":           {W: unknown, H: unknown},
	"objects/cruxsk3.cel":           {W: unknown, H: unknown},
	"objects/decap.cel":             {W: unknown, H: unknown},
	"objects/dirtfall.cel":          {W: unknown, H: unknown},
	"objects/explod1.cel":           {W: unknown, H: unknown},
	"objects/explod2.cel":           {W: unknown, H: unknown},
	"objects/firewal1.cel":          {W: unknown, H: unknown},
	"objects/flame1.cel":            {W: unknown, H: unknown},
	"objects/flame3.cel":            {W: unknown, H: unknown},
	"objects/ghost.cel":             {W: unknown, H: unknown},
	"objects/goatshrn.cel":          {W: unknown, H: unknown},
	"objects/l1braz.cel":            {W: unknown, H: unknown},
	"objects/l1doors.cel":           {W: unknown, H: unknown},
	"objects/l2doors.cel":           {W: unknown, H: unknown},
	"objects/l3doors.cel":           {W: unknown, H: unknown},
	"objects/lever.cel":             {W: unknown, H: unknown},
	"objects/lshrineg.cel":          {W: unknown, H: unknown},
	"objects/lzstand.cel":           {W: unknown, H: unknown},
	"objects/mcirl.cel":             {W: unknown, H: unknown},
	"objects/mfountn.cel":           {W: unknown, H: unknown},
	"objects/miniwatr.cel":          {W: unknown, H: unknown},
	"objects/mushptch.cel":          {W: unknown, H: unknown},
	"objects/nude2.cel":             {W: unknown, H: unknown},
	"objects/pedistl.cel":           {W: unknown, H: unknown},
	"objects/pfountn.cel":           {W: unknown, H: unknown},
	"objects/prsrplt1.cel":          {W: unknown, H: unknown},
	"objects/rockstan.cel":          {W: unknown, H: unknown},
	"objects/rshrineg.cel":          {W: unknown, H: unknown},
	"objects/sarc.cel":              {W: unknown, H: unknown},
	"objects/skulfire.cel":          {W: unknown, H: unknown},
	"objects/skulpile.cel":          {W: unknown, H: unknown},
	"objects/skulstik.cel":          {W: unknown, H: unknown},
	"objects/switch2.cel":           {W: unknown, H: unknown},
	"objects/switch3.cel":           {W: unknown, H: unknown},
	"objects/switch4.cel":           {W: unknown, H: unknown},
	"objects/tfountn.cel":           {W: unknown, H: unknown},
	"objects/tnudem.cel":            {W: unknown, H: unknown},
	"objects/tnudew.cel":            {W: unknown, H: unknown},
	"objects/traphole.cel":          {W: unknown, H: unknown},
	"objects/tsoul.cel":             {W: unknown, H: unknown},
	"objects/vapor1.cel":            {W: unknown, H: unknown},
	"objects/water.cel":             {W: unknown, H: unknown},
	"objects/waterjug.cel":          {W: unknown, H: unknown},
	"objects/weapstnd.cel":          {W: unknown, H: unknown},
	"objects/wtorch1.cel":           {W: unknown, H: unknown},
	"objects/wtorch2.cel":           {W: unknown, H: unknown},
	"objects/wtorch3.cel":           {W: unknown, H: unknown},
	"objects/wtorch4.cel":           {W: unknown, H: unknown},
	"towners/animals/cow.cel":       {W: unknown, H: unknown},
	"towners/butch/deadguy.cel":     {W: unknown, H: unknown},
	"towners/drunk/twndrunk.cel":    {W: unknown, H: unknown},
	"towners/healer/healer.cel":     {W: unknown, H: unknown},
	"towners/priest/priest8.cel":    {W: unknown, H: unknown},
	"towners/smith/smithn.cel":      {W: unknown, H: unknown},
	"towners/smith/smithw.cel":      {W: unknown, H: unknown},
	"towners/strytell/strytell.cel": {W: unknown, H: unknown},
	"towners/townboy/pegkid1.cel":   {W: unknown, H: unknown},
	"towners/townwmn1/witch.cel":    {W: unknown, H: unknown},
	"towners/townwmn1/wmnn.cel":     {W: unknown, H: unknown},
	"towners/townwmn1/wmnw.cel":     {W: unknown, H: unknown},
	"towners/twnf/twnfn.cel":        {W: unknown, H: unknown},
	"towners/twnf/twnfw.cel":        {W: unknown, H: unknown},
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
