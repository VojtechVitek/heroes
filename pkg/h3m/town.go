package h3m

import "fmt"

type Town int

const (
	Zero    Town = 0
	One     Town = 1
	Two     Town = 2
	Three   Town = 3
	Four    Town = 4
	Dungeon Town = 5
)

func (t Town) String() string {
	switch t {
	case Zero:
		return "0"
	case One:
		return "1"
	case Two:
		return "2"
	case Four:
		return "3"
	case Dungeon:
		return "Dungeon"
	default:
		return fmt.Sprintf("unknown (%x)", int(t))
	}
}
