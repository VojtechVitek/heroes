package h3m

type Terrain int

const (
	Dirt         Terrain = 0
	Sand         Terrain = 1
	Grass        Terrain = 2
	Snow         Terrain = 3
	Swamp        Terrain = 4
	Rough        Terrain = 5
	Subterranean Terrain = 6
	Lava         Terrain = 7
	Water        Terrain = 8
	Rock         Terrain = 9
)

func (t Terrain) RGB() (r uint8, g uint8, b uint8) {
	switch t {
	case 0:
		return 0x0F, 0x3F, 0x50
	case 1:
		return 0x8F, 0xCF, 0xDF
	case 2:
		return 0x00, 0x40, 0x00
	case 3:
		return 0xC0, 0xC0, 0xB0
	case 4:
		return 0x6F, 0x80, 0x4F
	case 5:
		return 0x30, 0x70, 0x80
	case 6:
		return 0x30, 0x80, 0x00
	case 7:
		return 0x4F, 0x4F, 0x4F
	case 8:
		return 0x90, 0x50, 0x0F
	case 9:
		return 0x00, 0x00, 0x00
	}

	return 0xFF, 0xFF, 0xFF
}
