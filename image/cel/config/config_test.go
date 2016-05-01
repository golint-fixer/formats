package config

import (
	"fmt"
	"sort"
	"testing"
)

// TODO: Extend the test case to verify that npixelsMapping contains the same
// paths as confs.

func TestConfs(t *testing.T) {
	var relCelPaths []string
	for relCelPath := range confs {
		relCelPaths = append(relCelPaths, relCelPath)
	}
	sort.Strings(relCelPaths)

	for _, relCelPath := range relCelPaths {
		npixels, ok := npixelsMapping[relCelPath]
		if !ok {
			t.Errorf("unable to locate pixel count for %q", relCelPath)
			continue
		}

		// Get config and frame numbers with specific image dimensions.
		conf := confs[relCelPath]

		// TODO: Check if this test is redundant, and may therefore be removed.

		// Verify pixel count for each frame with specific image dimensions.
		m := make(map[int]bool)
		for frameNum := range conf.FrameWidth {
			m[frameNum] = true
		}
		for frameNum := range conf.FrameHeight {
			m[frameNum] = true
		}
		var frameNums []int
		for frameNum := range m {
			frameNums = append(frameNums, frameNum)
		}
		sort.Ints(frameNums)
		for _, frameNum := range frameNums {
			width, ok := conf.FrameWidth[frameNum]
			if !ok {
				width = conf.W
			}
			height, ok := conf.FrameHeight[frameNum]
			if !ok {
				height = conf.H
			}
			got := width * height
			want := npixels[frameNum]
			if got != want {
				t.Errorf("%q: pixel count mismatch for frame number %d; expected %d, got %d", relCelPath, frameNum, want, got)
				continue
			}
		}

		// Verify pixel count for default dimension frames.
		for frameNum, want := range npixels {
			width, ok := conf.FrameWidth[frameNum]
			if !ok {
				width = conf.W
			}
			height, ok := conf.FrameHeight[frameNum]
			if !ok {
				height = conf.H
			}
			got := width * height
			if got != want {
				t.Errorf("%q: pixel count mismatch for frame number %d; expected %d, got %d", relCelPath, frameNum, want, got)
				continue
			} else {
				// TODO: Remove once the image configs have matured.
				fmt.Println("PASS:", relCelPath)
			}
		}
	}
}

// npixelsMapping maps from CEL file name to pixel count. The pixel count slice
// maps from frame number to pixel count of the given frame. If each frame of
// the CEL has the same pixel count, then the pixel count slice contains a
// single element.
var npixelsMapping = map[string][]int{
	"ctrlpan/golddrop.cel":          {35496},
	"ctrlpan/p8bulbs.cel":           {7744},
	"ctrlpan/p8but2.cel":            {1056},
	"ctrlpan/panel8.cel":            {92160},
	"ctrlpan/panel8bu.cel":          {1349},
	"ctrlpan/smaltext.cel":          {143},
	"ctrlpan/spelicon.cel":          {3136},
	"ctrlpan/talkbutt.cel":          {976},
	"ctrlpan/talkpanl.cel":          {92160},
	"data/bigtgold.cel":             {2070},
	"data/char.cel":                 {112640},
	"data/charbut.cel":              {2090, 902, 902, 902, 902, 902, 902, 902, 902},
	"data/diabsmal.cel":             {29600},
	"data/inv/inv.cel":              {112640},
	"data/inv/inv_rog.cel":          {112640},
	"data/inv/inv_sor.cel":          {112640},
	"data/inv/objcurs.cel":          {957, 1024, 1024, 1024, 1024, 1024, 1024, 1024, 1024, 1024, 805, 784, 784, 784, 784, 784, 784, 784, 784, 784, 784, 784, 784, 784, 784, 784, 784, 784, 784, 784, 784, 784, 784, 784, 784, 784, 784, 784, 784, 784, 784, 784, 784, 784, 784, 784, 784, 784, 784, 784, 784, 784, 784, 784, 784, 784, 784, 784, 784, 784, 784, 1568, 1568, 1568, 1568, 1568, 1568, 2352, 2352, 2352, 2352, 2352, 2352, 2352, 2352, 2352, 2352, 2352, 2352, 2352, 2352, 2352, 2352, 2352, 2352, 2352, 3136, 3136, 3136, 3136, 3136, 3136, 3136, 3136, 3136, 3136, 3136, 3136, 3136, 3136, 3136, 3136, 3136, 3136, 3136, 3136, 3136, 3136, 3136, 3136, 3136, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704, 4704},
	"data/medtexts.cel":             {484},
	"data/optbar.cel":               {9184},
	"data/option.cel":               {756},
	"data/pentspin.cel":             {2304},
	"data/pentspn2.cel":             {144},
	"data/quest.cel":                {112640},
	"data/spellbk.cel":              {112640},
	"data/spellbkb.cel":             {2204},
	"data/spelli2.cel":              {1406},
	"data/square.cel":               {8192},
	"data/textbox.cel":              {179073},
	"data/textbox2.cel":             {82113},
	"data/textslid.cel":             {144},
	"gendata/cut2.cel":              {307200},
	"gendata/cut3.cel":              {307200},
	"gendata/cut4.cel":              {307200},
	"gendata/cutgate.cel":           {307200},
	"gendata/cutl1d.cel":            {307200},
	"gendata/cutportl.cel":          {307200},
	"gendata/cutportr.cel":          {307200},
	"gendata/cutstart.cel":          {307200},
	"gendata/cuttt.cel":             {307200},
	"gendata/quotes.cel":            {307200},
	"items/armor2.cel":              {15360},
	"items/axe.cel":                 {15360},
	"items/axeflip.cel":             {15360},
	"items/bldstn.cel":              {15360},
	"items/bottle.cel":              {15360},
	"items/bow.cel":                 {15360},
	"items/cleaver.cel":             {15360},
	"items/crownf.cel":              {12288},
	"items/duricons.cel":            {1024},
	"items/fanvil.cel":              {15360},
	"items/fbook.cel":               {15360},
	"items/fbow.cel":                {15360},
	"items/fbrain.cel":              {15360},
	"items/fbttle.cel":              {15360},
	"items/fbttlebb.cel":            {15360},
	"items/fbttlebl.cel":            {15360},
	"items/fbttlebr.cel":            {15360},
	"items/fbttleby.cel":            {15360},
	"items/fbttledb.cel":            {15360},
	"items/fbttledy.cel":            {15360},
	"items/fbttleor.cel":            {15360},
	"items/fbttlewh.cel":            {15360},
	"items/fear.cel":                {12288},
	"items/feye.cel":                {15360},
	"items/fheart.cel":              {15360},
	"items/flazstaf.cel":            {15360},
	"items/fmush.cel":               {15360},
	"items/food.cel":                {12288},
	"items/fplatear.cel":            {15360},
	"items/goldflip.cel":            {15360},
	"items/helmut.cel":              {15360},
	"items/innsign.cel":             {15360},
	"items/larmor.cel":              {12288},
	"items/mace.cel":                {15360},
	"items/manaflip.cel":            {15360},
	"items/map/mapz0000.cel":        {225280},
	"items/map/mapz0001.cel":        {225280},
	"items/map/mapz0002.cel":        {225280},
	"items/map/mapz0003.cel":        {225280},
	"items/map/mapz0004.cel":        {225280},
	"items/map/mapz0005.cel":        {225280},
	"items/map/mapz0006.cel":        {225280},
	"items/map/mapz0007.cel":        {225280},
	"items/map/mapz0008.cel":        {225280},
	"items/map/mapz0009.cel":        {225280},
	"items/map/mapz0010.cel":        {225280},
	"items/map/mapz0011.cel":        {225280},
	"items/map/mapz0012.cel":        {225280},
	"items/map/mapz0013.cel":        {225280},
	"items/map/mapz0014.cel":        {225280},
	"items/map/mapz0015.cel":        {225280},
	"items/map/mapz0016.cel":        {225280},
	"items/map/mapz0017.cel":        {225280},
	"items/map/mapz0018.cel":        {225280},
	"items/map/mapz0019.cel":        {225280},
	"items/map/mapz0020.cel":        {225280},
	"items/map/mapz0021.cel":        {225280},
	"items/map/mapz0022.cel":        {225280},
	"items/map/mapz0023.cel":        {225280},
	"items/map/mapz0024.cel":        {225280},
	"items/map/mapz0025.cel":        {225280},
	"items/map/mapz0026.cel":        {225280},
	"items/map/mapz0027.cel":        {225280},
	"items/map/mapz0028.cel":        {225280},
	"items/map/mapz0029.cel":        {225280},
	"items/map/mapz0030.cel":        {225280},
	"items/map/mapzdoom.cel":        {225280},
	"items/ring.cel":                {15360},
	"items/rock.cel":                {9216},
	"items/scroll.cel":              {15360},
	"items/shield.cel":              {15360},
	"items/staff.cel":               {15360},
	"items/swrdflip.cel":            {15360},
	"items/wand.cel":                {15360},
	"items/wshield.cel":             {12288},
	"levels/l1data/l1s.cel":         {10240},
	"levels/l2data/l2s.cel":         {10240},
	"levels/towndata/towns.cel":     {14336},
	"missiles/flamel1.cel":          {12288},
	"missiles/flamel10.cel":         {9216},
	"missiles/flamel11.cel":         {12288},
	"missiles/flamel12.cel":         {9216},
	"missiles/flamel13.cel":         {12288},
	"missiles/flamel14.cel":         {9216},
	"missiles/flamel15.cel":         {12288},
	"missiles/flamel16.cel":         {9216},
	"missiles/flamel2.cel":          {9216},
	"missiles/flamel3.cel":          {12288},
	"missiles/flamel4.cel":          {9216},
	"missiles/flamel5.cel":          {12288},
	"missiles/flamel6.cel":          {9216},
	"missiles/flamel7.cel":          {12288},
	"missiles/flamel8.cel":          {9216},
	"missiles/flamel9.cel":          {12288},
	"missiles/flames1.cel":          {12288},
	"missiles/flames10.cel":         {9216},
	"missiles/flames11.cel":         {12288},
	"missiles/flames12.cel":         {9216},
	"missiles/flames13.cel":         {12288},
	"missiles/flames14.cel":         {9216},
	"missiles/flames15.cel":         {12288},
	"missiles/flames16.cel":         {9216},
	"missiles/flames2.cel":          {9216},
	"missiles/flames3.cel":          {12288},
	"missiles/flames4.cel":          {9216},
	"missiles/flames5.cel":          {12288},
	"missiles/flames6.cel":          {9216},
	"missiles/flames7.cel":          {12288},
	"missiles/flames8.cel":          {9216},
	"missiles/flames9.cel":          {12288},
	"missiles/flaml1.cel":           {16384},
	"missiles/flaml2.cel":           {16384},
	"missiles/flaml3.cel":           {16384},
	"missiles/flaml4.cel":           {16384},
	"missiles/flaml5.cel":           {16384},
	"missiles/flaml6.cel":           {16384},
	"missiles/flaml7.cel":           {16384},
	"missiles/flaml8.cel":           {16384},
	"missiles/flams1.cel":           {16384},
	"missiles/flams2.cel":           {16384},
	"missiles/flams3.cel":           {16384},
	"missiles/flams4.cel":           {16384},
	"missiles/flams5.cel":           {16384},
	"missiles/flams6.cel":           {16384},
	"missiles/flams7.cel":           {16384},
	"missiles/flams8.cel":           {16384},
	"missiles/mindmace.cel":         {9216},
	"missiles/sentfr.cel":           {9216},
	"missiles/sentout.cel":          {9216},
	"missiles/sentup.cel":           {9216},
	"monsters/acid/acidpud.cel":     {12288},
	"monsters/magma/magball1.cel":   {16384},
	"monsters/magma/magball2.cel":   {16384},
	"monsters/magma/magball3.cel":   {16384},
	"monsters/magma/magball4.cel":   {16384},
	"monsters/magma/magball5.cel":   {16384},
	"monsters/magma/magball6.cel":   {16384},
	"monsters/magma/magball7.cel":   {16384},
	"monsters/magma/magball8.cel":   {16384},
	"monsters/magma/magblos.cel":    {16384},
	"monsters/rhino/rhinos1.cel":    {20480},
	"monsters/rhino/rhinos2.cel":    {20480},
	"monsters/rhino/rhinos3.cel":    {20480},
	"monsters/rhino/rhinos4.cel":    {20480},
	"monsters/rhino/rhinos5.cel":    {20480},
	"monsters/rhino/rhinos6.cel":    {20480},
	"monsters/rhino/rhinos7.cel":    {20480},
	"monsters/rhino/rhinos8.cel":    {20480},
	"monsters/succ/flare.cel":       {16384},
	"monsters/succ/flarexp.cel":     {16384},
	"monsters/thin/lghning.cel":     {9216},
	"monsters/unrav/unravw.cel":     {9216},
	"objects/altboy.cel":            {16384},
	"objects/angel.cel":             {12288},
	"objects/armstand.cel":          {9216},
	"objects/banner.cel":            {9216},
	"objects/barrel.cel":            {9216},
	"objects/barrelex.cel":          {9216},
	"objects/bcase.cel":             {12288},
	"objects/bkslbrnt.cel":          {9216},
	"objects/bkurns.cel":            {9216},
	"objects/bloodfnt.cel":          {9216},
	"objects/book1.cel":             {9216},
	"objects/book2.cel":             {9216},
	"objects/bshelf.cel":            {12288},
	"objects/burncros.cel":          {25600},
	"objects/candlabr.cel":          {9216},
	"objects/candle.cel":            {9216},
	"objects/candle2.cel":           {9216},
	"objects/cauldren.cel":          {9216},
	"objects/chest1.cel":            {9216},
	"objects/chest2.cel":            {9216},
	"objects/chest3.cel":            {9216},
	"objects/cruxsk1.cel":           {12288},
	"objects/cruxsk2.cel":           {12288},
	"objects/cruxsk3.cel":           {12288},
	"objects/decap.cel":             {9216},
	"objects/dirtfall.cel":          {15360},
	"objects/explod1.cel":           {16384},
	"objects/explod2.cel":           {16384},
	"objects/firewal1.cel":          {20480},
	"objects/flame1.cel":            {9216},
	"objects/flame3.cel":            {12288},
	"objects/ghost.cel":             {16384},
	"objects/goatshrn.cel":          {9216},
	"objects/l1braz.cel":            {10240},
	"objects/l1doors.cel":           {10240},
	"objects/l2doors.cel":           {8192},
	"objects/l3doors.cel":           {8192},
	"objects/lever.cel":             {9216},
	"objects/lshrineg.cel":          {16384},
	"objects/lzstand.cel":           {16384},
	"objects/mcirl.cel":             {9216},
	"objects/mfountn.cel":           {16384},
	"objects/miniwatr.cel":          {8192},
	"objects/mushptch.cel":          {9216},
	"objects/nude2.cel":             {16384},
	"objects/pedistl.cel":           {9216},
	"objects/pfountn.cel":           {16384},
	"objects/prsrplt1.cel":          {9216},
	"objects/rockstan.cel":          {9216},
	"objects/rshrineg.cel":          {16384},
	"objects/sarc.cel":              {12288},
	"objects/skulfire.cel":          {9216},
	"objects/skulpile.cel":          {9216},
	"objects/skulstik.cel":          {9216},
	"objects/switch2.cel":           {9216},
	"objects/switch3.cel":           {9216},
	"objects/switch4.cel":           {9216},
	"objects/tfountn.cel":           {12288},
	"objects/tnudem.cel":            {16384},
	"objects/tnudew.cel":            {16384},
	"objects/traphole.cel":          {9216},
	"objects/tsoul.cel":             {12288},
	"objects/vapor1.cel":            {16384},
	"objects/water.cel":             {20480},
	"objects/waterjug.cel":          {9216},
	"objects/weapstnd.cel":          {9216},
	"objects/wtorch1.cel":           {12288},
	"objects/wtorch2.cel":           {12288},
	"objects/wtorch3.cel":           {12288},
	"objects/wtorch4.cel":           {12288},
	"towners/animals/cow.cel":       {16384},
	"towners/butch/deadguy.cel":     {9216},
	"towners/drunk/twndrunk.cel":    {9216},
	"towners/healer/healer.cel":     {9216},
	"towners/priest/priest8.cel":    {9216},
	"towners/smith/smithn.cel":      {9216},
	"towners/smith/smithw.cel":      {9216},
	"towners/strytell/strytell.cel": {9216},
	"towners/townboy/pegkid1.cel":   {6144},
	"towners/townwmn1/witch.cel":    {9216},
	"towners/townwmn1/wmnn.cel":     {9216},
	"towners/townwmn1/wmnw.cel":     {9216},
	"towners/twnf/twnfn.cel":        {9216},
	"towners/twnf/twnfw.cel":        {9216},
}
