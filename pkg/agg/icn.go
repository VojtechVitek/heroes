package agg

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"image"
	"strings"

	"github.com/pkg/errors"
)

// 0x00  Number of sprites         2 bytes
// 0x02  Data size (headers+data)  2 bytes
// 0x04  Sprite headers            Number of sprites * 13 bytes
// 0x06  Sprite data
func NewICN(data []byte, pallete pallete) (*ICN, error) {
	icn := ICN{
		pallete: pallete,
	}

	var v uint16
	if err := binary.Read(bytes.NewReader(data[0:2]), binary.LittleEndian, &v); err != nil {
		return nil, errors.Wrap(err, "failed to read number of sprites")
	}
	icn.NumSprites = int(v)

	if err := binary.Read(bytes.NewReader(data[2:4]), binary.LittleEndian, &v); err != nil {
		return nil, errors.Wrap(err, "failed to read data size")
	}
	icn.dataSize = int(v)

	spriteHeaders := make([]*SpriteHeader, icn.NumSprites)
	for i := 0; i < icn.NumSprites; i++ {
		var header SpriteHeader
		if err := binary.Read(bytes.NewReader(data[4:4+icn.NumSprites*13]), binary.LittleEndian, &header); err != nil {
			return nil, errors.Wrap(err, "failed to parse sprite header")
		}
		spriteHeaders = append(spriteHeaders, &header)
	}

	icn.spritesData = data[4+icn.NumSprites*13:]

	return &icn, nil
}

type ICN struct {
	NumSprites    int
	dataSize      int
	spriteHeaders []*SpriteHeader
	spritesData   []byte
	pallete       pallete
}

type SpriteHeader struct {
	_ [13]byte
}

func NewSpriteHeader(data []byte) (*SpriteHeader, error) {
	var header SpriteHeader
	// s16 offsetX; // positionning offset of the sprite on X axis
	// s16 offsetY; // positionning offset of the sprite on Y axis
	// u16 width; // sprite's width
	// u16 height; // sprite's height
	// u8 type; // type of sprite : 0 = Normal, 32 = Monochromatic shape
	// u32 offsetData; // beginning of the data
	return &header, nil
}

func (icn *ICN) Images() ([]*image.RGBA, error) {
	images := make([]*image.RGBA, 0, icn.NumSprites)

	// data := t.data[6:] // Pixels only, strip off the header.

	// numTiles := t.NumTiles()
	// width := t.TileWidth()
	// height := t.TileHeight()
	// rect := image.Rect(0, 0, width, height*numTiles)

	// pixels := make([]uint8, 0, numTiles*width*height*4)
	// for i := 0; i < numTiles*width*height; i++ {
	// 	r, g, b := t.pallete.RGB(data[i])
	// 	pixels = append(pixels, r, g, b, opaqueAlpha)
	// }

	return images, nil
}

func (icn *ICN) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Number of sprites: %v\n", icn.NumSprites)

	for _, header := range icn.spriteHeaders {
		fmt.Fprintln(&b, header)
	}

	return b.String()
}
