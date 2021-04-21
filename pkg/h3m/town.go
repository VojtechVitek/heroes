package h3m

import "fmt"

type Town int

const (
	None    Town = 0
	Rampart Town = 1
	Two     Town = 2
	Three   Town = 3

	Dungeon    Town = 5
	Stronghold Town = 6
	Random     Town = 0xFF
)

func (t Town) String() string {
	switch t {
	case None:
		return "<None>"
	case Rampart:
		return "Rampart"
	case Two:
		return "2"
	case Dungeon:
		return "Dungeon"
	case Stronghold:
		return "Stronghold"
	case Random:
		return "<Random>"
	default:
		return fmt.Sprintf("unknown (%x)", int(t))
	}
}
