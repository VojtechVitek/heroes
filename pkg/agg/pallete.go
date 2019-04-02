package agg

import (
	"github.com/pkg/errors"
)

// 256 * 3 bytes
type pallete []byte

func NewPallete(data []byte) (pallete, error) {
	if len(data) != 256*3 {
		return nil, errors.Errorf("failed to create color pallete: expected %v bytes, got %v bytes", 256*3, len(data))
	}
	return pallete(data), nil
}

func (p pallete) RGB(index uint8) (r, g, b uint8) {
	return p.color(index * 3), p.color(index*3 + 1), p.color(index*3 + 2)
}

func (p pallete) color(c uint8) uint8 {
	if c >= 214 && c <= 241 {
		switch c {
		case 214, 215, 216, 217:
			// red
		case 218, 219, 220, 221:
			// yellow
		case 231, 232, 233, 234, 235, 236, 237:
			// water
		case 238, 239, 240, 241:
			// blue
		}
	}

	return p[c] << 2 // The pallete is dark, need to multiply by 4 to get the real color.
}
