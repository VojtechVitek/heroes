package h3m

type MapInfo struct {
	HasHero      bool
	MapSize      int
	HasTwoLevels bool
	Name         string
	Desc         string
	Difficulty   int
	MasteryCap   int // Only set on ArmageddonsBlade and ShadowOfDeath maps.
}
