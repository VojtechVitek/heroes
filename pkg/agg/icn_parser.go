package agg

import "fmt"

type spriteParser struct {
	data      []byte
	lineWidth int
	pallete   pallete
	pixels    []uint8
	pos       int
}

func (p *spriteParser) getPixels() []uint8 {
	for {
		cmd := p.nextByte()

		switch {
		case cmd == 0x00:
			// EOL. Fill the rest of the line with transparent color.
			for i := p.pos % p.lineWidth; i < p.lineWidth; i++ {
				p.pixels = append(p.pixels, 0, 0, 0, 0)
			}

		case cmd >= 0x01 && cmd <= 0x7F:
			// Number of pixels to fill with a specific color (next byte) from the pallete.
			for i := 0; i < int(cmd); i++ {
				r, g, b := p.pallete.RGB(p.nextByte())
				p.pixels = append(p.pixels, r, g, b, opaqueAlpha)
			}

		case cmd == 0x80:
			// EOF.

			// Fill in the missing pixels.
			for i := 0; i < cap(p.pixels)-len(p.pixels); i++ {
				p.pixels = append(p.pixels, 0, 0, 0, 0)
			}

			return p.pixels

		case cmd >= 0x81 && cmd <= 0xBF:
			// Number of pixels to skip. Fill with transparent color.
			for i := 0; i < int(cmd-0x80); i++ {
				p.pixels = append(p.pixels, 0, 0, 0, 0)
			}

		case cmd == 0xC0:
			// Number (next byte or two) of shadow pixels.
			n := int(p.nextByte()) % 4
			if n == 0 {
				n = int(p.nextByte())
			}

			for i := 0; i < n; i++ {
				p.pixels = append(p.pixels, 0, 0, 0, 64)
			}

		case cmd == 0xC1:
			// Number (next byte) of pixels to fill with a specific color (second next byte) from the pallete.
			n := int(p.nextByte())
			r, g, b := p.pallete.RGB(p.nextByte())

			for i := 0; i < n; i++ {
				p.pixels = append(p.pixels, r, g, b, opaqueAlpha)
			}

		case cmd >= 0xC2 && cmd <= 0xFF:
			// Number of pixels of same color (next byte) shifted by 0xC0.
			n := int(cmd) - 0xC0

			r, g, b := p.pallete.RGB(p.nextByte())

			for i := 0; i < n; i++ {
				p.pixels = append(p.pixels, r, g, b, opaqueAlpha)
			}

		default:
			panic(fmt.Sprintf("unknown cmd 0x%X", cmd))
		}
	}
}

func (p *spriteParser) nextByte() uint8 {
	b := p.data[p.pos]
	p.pos++
	return b
}
