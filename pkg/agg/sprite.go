package agg

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"image"
	"strings"

	"github.com/pkg/errors"
)

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

func (s *Sprite) RenderImage(pallete pallete) (*image.RGBA, error) {
	r := bytes.NewReader(s.data)
	pixels := make([]uint8, 0, 4*s.width*s.height)

	pos := 0
	nextByte := func() byte {
		b, err := r.ReadByte()
		if err != nil {
			panic(err)
		}
		return b
	}

	for {
		cmd := nextByte()

		switch {
		case cmd == 0x00:
			// EOL. Fill the rest of the line with transparent color.
			for i := pos % s.width; i <= s.width; i++ {
				pixels = append(pixels, 0, 0, 0, 0)
				pos++
			}

		case cmd >= 0x01 && cmd <= 0x7F:
			// Number of pixels to fill with a specific color (next byte) from the pallete.
			for i := 0; i < int(cmd); i++ {
				r, g, b := pallete.RGB(nextByte())
				pixels = append(pixels, r, g, b, opaqueAlpha)
				pos++
			}

		case cmd == 0x80:
			// EOF.

			// Fill in the missing pixels.
			for i := 0; i < cap(pixels)-len(pixels); i++ {
				pixels = append(pixels, 0, 0, 0, 0)
				pos++
			}

			img := &image.RGBA{pixels, 4 * s.width, image.Rect(0, 0, s.width, s.height)}
			return img, nil

		case cmd >= 0x81 && cmd <= 0xBF:
			// Number of pixels to skip. Fill with transparent color.
			for i := 0; i < int(cmd-0x80); i++ {
				pixels = append(pixels, 0, 0, 0, 0)
				pos++
			}

		case cmd == 0xC0:
			// Number (next byte or two) of shadow pixels.
			n := int(nextByte()) % 4
			if n == 0 {
				n = int(nextByte())
			}

			for i := 0; i < n; i++ {
				pixels = append(pixels, 0, 0, 0, 64)
				pos++
			}

		case cmd == 0xC1:
			// Number (next byte) of pixels to fill with a specific color (second next byte) from the pallete.
			n := int(nextByte())
			r, g, b := pallete.RGB(nextByte())

			for i := 0; i < n; i++ {
				pixels = append(pixels, r, g, b, opaqueAlpha)
				pos++
			}

		case cmd >= 0xC2 && cmd <= 0xFF:
			// Number of pixels of same color (next byte) shifted by 0xC0.
			n := int(cmd) - 0xC0

			r, g, b := pallete.RGB(nextByte())

			for i := 0; i < n; i++ {
				pixels = append(pixels, r, g, b, opaqueAlpha)
				pos++
			}

		default:
			return nil, fmt.Errorf("unknown cmd 0x%X", cmd)
		}
	}
}

func (s *Sprite) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "x, y = %v, %v\n", s.x, s.y)
	fmt.Fprintf(&b, "width, height = %v, %v\n", s.width, s.height)
	fmt.Fprintf(&b, "typ = %v\n", s.typ)

	return b.String()
}
