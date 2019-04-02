package agg

import (
	"bytes"
	"encoding/binary"
)

// 0x00  Number of tiles  2 bytes
// 0x02  Tile width       2 bytes
// 0x04  Tile height      2 bytes
// 0x06  Data
type Tiles []byte

func (t Tiles) NumTiles() int { return t.uint16ToInt(0, 2) }
func (t Tiles) Width() int    { return t.uint16ToInt(2, 4) }
func (t Tiles) Height() int   { return t.uint16ToInt(4, 6) }

func (t Tiles) uint16ToInt(from, to int) int {
	var v uint16
	_ = binary.Read(bytes.NewReader(t[from:to]), binary.LittleEndian, &v)
	return int(v)
}
