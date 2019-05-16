package agg

import (
	"image/color"

	"github.com/pkg/errors"
)

// 256 * 3 bytes
type palette []byte

func NewPalette(data []byte) (palette, error) {
	if len(data) != 256*3 {
		return nil, errors.Errorf("failed to create color palette: expected %v bytes, got %v bytes", 256*3, len(data))
	}
	return palette(data), nil
}

func (p palette) RGB(index int) (r, g, b uint8) {
	return p.color(index * 3), p.color(index*3 + 1), p.color(index*3 + 2)
}

func (p palette) color(c int) uint8 {
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

	return p[c] << 2 // The palette is dark, need to multiply by 4 to get the real color.
}

func (p palette) GifPalette() color.Palette {
	palette := color.Palette{}
	uniq := map[color.RGBA]struct{}{}
	for i := 0; i <= 255; i++ {
		r, g, b := p.RGB(i)
		c := color.RGBA{r, g, b, 255}
		if _, ok := uniq[c]; !ok {
			palette = append(palette, c)
			uniq[c] = struct{}{}
		}
	}

	return palette
}
