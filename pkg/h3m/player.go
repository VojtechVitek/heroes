package h3m

type Player struct {
	CanBeHuman             bool
	CanBeComputer          bool
	Behavior               int
	AllowedAlignments      int // ShadowOfDeath only.
	TownTypes              int
	Unknown1_HasRandomTown bool
	HasMainTown            bool

	StartingTownCreateHero bool
	StartingTownType       int
	StartingTownPos        Position

	StartingHeroIsRandom bool
	StartingHeroType     int
	StartingHeroFace     int
	StartingHeroName     string
}

func (h3m *H3M) NumberOfPlayers() int {
	var num int
	for _, player := range h3m.Players {
		if player.CanBeHuman || player.CanBeComputer {
			num++
		}
	}
	return num
}
