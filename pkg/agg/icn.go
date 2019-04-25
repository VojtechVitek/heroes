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
// 0x02  Data size                 4 bytes
// 0x06  Sprites headers + data    (each header's dataOffset points here)
func NewICN(data []byte, pallete pallete) (*ICN, error) {
	var u16 uint16
	if err := binary.Read(bytes.NewReader(data[0:2]), binary.LittleEndian, &u16); err != nil {
		return nil, errors.Wrap(err, "failed to read number of sprites")
	}
	numSprites := int(u16)

	var u32 uint32
	if err := binary.Read(bytes.NewReader(data[2:6]), binary.LittleEndian, &u32); err != nil {
		return nil, errors.Wrap(err, "failed to read data size")
	}
	dataSize := int(u32)

	data = data[6:] // Sprites headers + data only. That's what sprite headers offsets reference.
	if len(data) != dataSize {
		return nil, errors.Errorf("expected data size %v bytes, got %v bytes", dataSize, len(data))
	}

	sprites := make([]*Sprite, 0, numSprites)
	for i := 0; i < numSprites; i++ {
		sprite, err := NewSprite(data, i)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse sprite")
		}
		sprites = append(sprites, sprite)
	}

	return &ICN{sprites: sprites}, nil
}

type ICN struct {
	sprites []*Sprite
}

type Sprite struct {
	x      int
	y      int
	width  int
	height int
	typ    uint8
	data   []byte
}

// 0x00  x                      2 bytes
// 0x02  y                      2 bytes
// 0x04  width                  2 bytes
// 0x06  height                 2 bytes
// 0x08  type                   1 byte
// 0x09  data offset            4 bytes
func NewSprite(data []byte, index int) (*Sprite, error) {
	var (
		s   Sprite
		s16 int16
		u16 uint16
		u32 uint32
	)

	header := data[index*13 : (index+1)*13]

	if err := binary.Read(bytes.NewReader(header[0:2]), binary.LittleEndian, &s16); err != nil {
		return nil, errors.Wrap(err, "failed to read x")
	}
	s.x = int(s16)

	if err := binary.Read(bytes.NewReader(header[2:4]), binary.LittleEndian, &s16); err != nil {
		return nil, errors.Wrap(err, "failed to read y")
	}
	s.y = int(s16)

	if err := binary.Read(bytes.NewReader(header[4:6]), binary.LittleEndian, &u16); err != nil {
		return nil, errors.Wrap(err, "failed to read width")
	}
	s.width = int(u16)

	if err := binary.Read(bytes.NewReader(header[6:8]), binary.LittleEndian, &u16); err != nil {
		return nil, errors.Wrap(err, "failed to read height")
	}
	s.height = int(u16)

	s.typ = uint8(header[9])

	if err := binary.Read(bytes.NewReader(header[9:13]), binary.LittleEndian, &u32); err != nil {
		return nil, errors.Wrap(err, "failed to read data offset")
	}
	dataOffset := int(u32)

	s.data = data[dataOffset:]

	return &s, nil
}

func (icn *ICN) Sprite(index int) *Sprite {
	return icn.sprites[index]
}

func (icn *ICN) Sprites() []*Sprite {
	return icn.sprites
}

func (icn *ICN) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Number of sprites: %v\n", len(icn.sprites))

	for _, header := range icn.sprites {
		fmt.Fprintf(&b, "%+v\n", header)
	}

	return b.String()
}

func (s *Sprite) Image(pallete pallete) (*image.RGBA, error) {
	rect := image.Rect(0, 0, s.width, s.height)

	parser := &spriteParser{
		data:      s.data,
		lineWidth: s.width,
		pallete:   pallete,
		pixels:    make([]uint8, 0, 4*s.width*s.height),
	}
	img := &image.RGBA{parser.getPixels(), 4 * s.width, rect}

	return img, nil
}

func (s *Sprite) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "x, y = %v, %v\n", s.x, s.y)
	fmt.Fprintf(&b, "width, height = %v, %v\n", s.width, s.height)
	fmt.Fprintf(&b, "typ = %v\n", s.typ)

	return b.String()
}
