package palette

type Palette []byte // 256 * {R, G, B}

func (p Palette) RGBA(index int) (r, g, b, a uint8) {
	return uint8(p[index*3]), uint8(p[index*3+1]), uint8(p[index*3+2]), OpaqueAlpha
}

const OpaqueAlpha = 255
