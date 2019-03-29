package mp2

import (
	"fmt"
	"math"
	"strings"
)

type Tile struct {
	// Note: The struct padding is correct on 64bit processors..
	// but we should probably rewrite this into []byte and parse
	// fields individually by hand like for the Header.
	TileIndex     uint16 // 0x00
	ObjectName1   uint8  // 0x02 .. why double?
	IndexName1    uint8  // 0x03
	Quantity1     uint8  // 0x04
	Quantity2     uint8  // 0x05
	ObjectName2   uint8  // 0x06
	IndexName2    uint8  // 0x07
	Shape         uint8  // 0x08
	GeneralObject Object // 0x09
	IndexAddon    uint16 // 0x0A
	UniqueNumber1 uint32 // 0x0C
	UniqueNumber2 uint32 // 0x10
} // sizeof: 0x14 (20)

func (t Tile) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "%v\n", t.GeneralObject)

	return b.String()
}

type Tiles []Tile

func (t Tiles) String() string {
	var b strings.Builder

	// side of a square
	width := int(math.Sqrt(float64(len(t))))

	for x := 0; x < width; x++ {
		fmt.Fprintf(&b, "%v:\n", x)

		for y := 0; y < width; y++ {
			fmt.Fprintf(&b, "%8v ", t[x*width+y].GeneralObject.String())
		}
		// fmt.Fprintf(&b, "\n")
		// for y := 0; y < width; y++ {
		// 	fmt.Fprintf(&b, "%v ", t[x*width+y].ObjectName1)
		// }
		// fmt.Fprintf(&b, "\n")
		// for y := 0; y < width; y++ {
		// 	fmt.Fprintf(&b, "%v ", t[x*width+y].ObjectName2)
		// }
		// for y := 0; y < width; y++ {
		// 	fmt.Fprintf(&b, "%v ", t[x*width+y].IndexName1)
		// }
		// fmt.Fprintf(&b, "\n")
		// for y := 0; y < width; y++ {
		// 	fmt.Fprintf(&b, "%v ", t[x*width+y].IndexName2)
		// }
		// fmt.Fprintf(&b, "\n\n")
	}

	return b.String()
}
