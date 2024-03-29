package h3m

type Player struct {
	MasteryCap int // Only set on ArmageddonsBlade and ShadowOfDeath maps.

	CanBeHuman        bool
	CanBeComputer     bool
	Behavior          int
	AllowedAlignments int // ShadowOfDeath only.
	TownTypes         Town
	HasRandomTown     bool
	HasMainTown       bool

	StartingTownCreateHero bool
	StartingTown           Town
	StartingTownPos        Position

	StartingHeroIsRandom bool
	StartingHeroType     int
	StartingHeroFace     int
	StartingHeroName     string
}

type Heroes struct {
	Type int
	Name string
}

type CustomHeroes struct {
	Type           int
	Face           int
	Name           string
	AllowedPlayers int
}

// type (p Player) String() string {
// 	if !p.CanBeHuman && !p.CanBeComputer {
// 		return ""
// 	}
// 	return ""
// }

func (h3m *H3M) NumberOfPlayers() int {
	var num int
	for _, player := range h3m.Players {
		if player.CanBeHuman || player.CanBeComputer {
			num++
		}
	}
	return num
}
