package agg

import "fmt"

type spriteParser struct {
	data      []byte
	lineWidth int
	pallete   pallete
	pixels    []uint8
}

func (p *spriteParser) getPixels() []uint8 {
	pos := 0
	for {
		cmd := p.data[pos]
		pos++
		switch {
		case cmd == 0x00:
			// EOL. Fill line with transparent color.
			for i := pos % p.lineWidth; i < p.lineWidth; i++ {
				p.pixels = append(p.pixels, 0, 0, 0, 0)
			}

		case cmd >= 0x01 && cmd <= 0x7F:
			// Number of pixels to fill with a pallete color.
			for i := 0; i < int(cmd); i++ {
				r, g, b := p.pallete.RGB(p.data[pos])
				pos++
				p.pixels = append(p.pixels, r, g, b, opaqueAlpha)
			}

		case cmd >= 0x81 && cmd <= 0xBF:
			// Number of pixels to skip. Fill with tranparent color.
			for i := 0; i < int(cmd-0x80); i++ {
				p.pixels = append(p.pixels, 0, 0, 0, 0)
			}

		case cmd == 0xC0:
			// Number (next byte or two) of shadow pixels.
			var n int

			nextByte := p.data[pos]
			pos++

			n = int(nextByte) % 4
			if n == 0 {
				secondNextByte := p.data[pos]
				pos++
				n = int(secondNextByte)
			}

			for i := 0; i < n; i++ {
				p.pixels = append(p.pixels, 0, 0, 0, 0x40)
			}

		case cmd == 0x80:
			// EOF.

			// Fill in the missing pixels.
			for i := 0; i < cap(p.pixels)-len(p.pixels); i++ {
				p.pixels = append(p.pixels, 0, 0, 0, 0x40)
			}

			return p.pixels

		case cmd >= 0xC2 && cmd <= 0xFF:
			// Number of pixels of same color (next byte) shifted by 0xC0.
			n := int(cmd) - 0xC0

			r, g, b := p.pallete.RGB(p.data[pos])
			pos++

			for i := 0; i < n; i++ {
				p.pixels = append(p.pixels, r, g, b, opaqueAlpha)
			}

		default:
			panic(fmt.Sprintf("unknown cmd 0x%X", cmd))
		}
	}
}
