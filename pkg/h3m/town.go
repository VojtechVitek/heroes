package h3m

import "fmt"

type Town int

const (
	Castle     Town = 0
	Rampart    Town = 1
	Tower      Town = 2
	Inferno    Town = 3
	Necropolis Town = 4
	Dungeon    Town = 5
	Stronghold Town = 6
	Fortress   Town = 7
	Conflux    Town = 8
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
	case Necropolis:
		return "Necropolis"
	case Dungeon:
		return "Dungeon"
	case Stronghold:
		return "Stronghold"
	case Fortress:
		return "Fortress"
	case Conflux:
		return "Conflux"
	case Random:
		return "<Random>"
	default:
		return fmt.Sprintf("unknown town %x", int(t))
	}
}
