package def

type Palette [256]RGBA

type RGBA struct {
	r int
	g int
	b int
	a int
}

func (p Palette) RGB(index int) (r, g, b uint8) {
	return uint8(index * 3), uint8(index*3 + 1), uint8(index*3 + 2)
}
