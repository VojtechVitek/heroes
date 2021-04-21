package h3m

import "fmt"

type MapSize int

const (
	S  MapSize = 70
	M  MapSize = 0x48
	L  MapSize = 0x6c
	XL MapSize = 4
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
