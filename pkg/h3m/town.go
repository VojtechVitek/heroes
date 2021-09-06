package h3m

import "fmt"

type Town int

const (
	Castle     Town = 0
	Rampart    Town = 1
	Tower      Town = 2
	Inferno    Town = 3
	_          Town = 4
	Dungeon    Town = 5
	Stronghold Town = 6
	Random     Town = 0xFF
)

func (t Town) String() string {
	switch t {
	case Castle:
		return "Castle"
	case Rampart:
		return "Rampart"
	case Tower:
		return "Tower"
	case Inferno:
		return "Inferno"
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
