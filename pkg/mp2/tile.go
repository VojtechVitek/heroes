package mp2

import (
	"fmt"
	"strings"
)

type MapTile struct {
	// Note: The struct padding is correct on 64bit processors..
	// but we should probably rewrite this into []byte and parse
	// fields individually by hand like for the Header.
	TileIndex     uint16 // 0x00
	ObjectName1   uint8  // 0x02
	IndexName1    uint8  // 0x03
	Quantity1     uint8  // 0x04
	Quantity2     uint8  // 0x05
	ObjectName2   uint8  // 0x06
	IndexName2    uint8  // 0x07
	Shape         uint8  // 0x08
	GeneralObject uint8  // 0x09
} // sizeof: 0x14 (20)

func (t MapTile) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "index: %v\n", t.TileIndex)

	return b.String()
}
