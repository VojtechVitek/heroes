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
// 0x02  Data size (headers+data)  4 bytes
// 0x06  Sprites headers           Number of sprites * 13 bytes
// 0x??  Sprites data
func NewICN(data []byte, pallete pallete) (*ICN, error) {
	icn := ICN{
		pallete: pallete,
	}

	var u16 uint16
	if err := binary.Read(bytes.NewReader(data[0:2]), binary.LittleEndian, &u16); err != nil {
		return nil, errors.Wrap(err, "failed to read number of sprites")
	}
	icn.NumSprites = int(u16)

	var u32 uint32
	if err := binary.Read(bytes.NewReader(data[2:6]), binary.LittleEndian, &u32); err != nil {
		return nil, errors.Wrap(err, "failed to read data size")
	}
	dataSize := int(u32)

	headerData := data[6 : 6+icn.NumSprites*13]
	spriteHeaders := make([]*SpriteHeader, 0, icn.NumSprites)
	for i := 0; i < icn.NumSprites; i++ {
		header, err := NewSpriteHeader(headerData[i*13 : (i+1)*13])
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse sprite header")
		}
		spriteHeaders = append(spriteHeaders, header)
	}
	icn.spriteHeaders = spriteHeaders

	icn.spritesData = data[6+icn.NumSprites*13:]

	if len(icn.spritesData)+len(headerData) != dataSize {
		return nil, errors.Errorf("expected data size %v bytes, got %v bytes", dataSize, len(icn.spritesData)+len(headerData))
	}

	return &icn, nil
}

type ICN struct {
	NumSprites    int
	spriteHeaders []*SpriteHeader
	spritesData   []byte
	pallete       pallete
}

type SpriteHeader struct {
	x      int
	y      int
	width  int
	height int
	typ    uint8
	dataAt int
}

// 0x00  x                      2 bytes
// 0x02  y                      2 bytes
// 0x04  width                  2 bytes
// 0x06  height                 2 bytes
// 0x08  type                   1 byte
// 0x09  data offset            4 bytes
func NewSpriteHeader(data []byte) (*SpriteHeader, error) {
	if len(data) != 13 {
		return nil, errors.Errorf("expected 13 bytes of sprite header, got %v bytes", len(data))
	}

	var (
		h   SpriteHeader
		s16 int16
		u16 uint16
		u32 uint32
	)

	if err := binary.Read(bytes.NewReader(data[0:2]), binary.LittleEndian, &s16); err != nil {
		return nil, errors.Wrap(err, "failed to read x")
	}
	h.x = int(s16)

	if err := binary.Read(bytes.NewReader(data[2:4]), binary.LittleEndian, &s16); err != nil {
		return nil, errors.Wrap(err, "failed to read y")
	}
	h.y = int(s16)

	if err := binary.Read(bytes.NewReader(data[4:6]), binary.LittleEndian, &u16); err != nil {
		return nil, errors.Wrap(err, "failed to read width")
	}
	h.width = int(u16)

	if err := binary.Read(bytes.NewReader(data[6:8]), binary.LittleEndian, &u16); err != nil {
		return nil, errors.Wrap(err, "failed to read height")
	}
	h.height = int(u16)

	h.typ = uint8(data[9])

	if err := binary.Read(bytes.NewReader(data[9:13]), binary.LittleEndian, &u32); err != nil {
		return nil, errors.Wrap(err, "failed to read height")
	}
	h.dataAt = int(u32)

	return &h, nil
}

func (icn *ICN) Images() ([]*image.RGBA, error) {
	images := make([]*image.RGBA, 0, icn.NumSprites)

	for i, header := range icn.spriteHeaders {
		if i > 0 {
			return images, nil
		}

		_ = header
		//rect := image.Rect(0, 0, width, height*numTiles)

		// pixels := make([]uint8, 0, numTiles*width*height*4)
		// for i := 0; i < numTiles*width*height; i++ {
		// 	r, g, b := t.pallete.RGB(data[i])
		// 	pixels = append(pixels, r, g, b, opaqueAlpha)
		// }

	}

	return images, nil
}

func (icn *ICN) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Number of sprites: %v\n", icn.NumSprites)

	for _, header := range icn.spriteHeaders {
		fmt.Fprintf(&b, "%+v\n", header)
	}

	return b.String()
}
