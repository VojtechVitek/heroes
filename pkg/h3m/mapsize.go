package h3m

import "fmt"

type MapSize int

const (
	S  MapSize = 0x24
	M  MapSize = 0x48
	L  MapSize = 0x6c
	XL MapSize = 0x04
)

func (t MapSize) String() string {
	switch t {
	case S:
		return "S"
	case M:
		return "M"
	case L:
		return "L"
	case XL:
		return "XL"
	default:
		return fmt.Sprintf("unknown (%x)", int(t))
	}
}

func (t MapSize) Size() (x int, y int) {
	switch t {
	case S:
		return 36, 36
	case M:
		return 72, 72
	case L:
		return 108, 108
	case XL:
		return 144, 144
	}

	panic(fmt.Sprintf("unknown map size (%x)", int(t)))
}
