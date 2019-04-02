package agg

type Pallete []byte

func (p Pallete) RGB(index uint8) (r, g, b uint8) {
	rgb := p[index*3 : index*3+3]
	return rgb[0] * 4, rgb[1] * 4, rgb[2] * 4
}
