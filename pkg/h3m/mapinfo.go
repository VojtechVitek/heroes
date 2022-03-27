package h3m

type MapInfo struct {
	// Basic info.
	HasHero      bool
	MapSize      MapSize
	HasTwoLevels bool
	Name         string
	Desc         string
	Difficulty   int

	// Additional info.
	WinCondition                   WinCondition
	WinConditionAllowNormalWin     bool
	WinConditionAppliesToComputer  bool
	WinConditionType               int // Might be type of Resource, ie. "Gems".
	WinConditionAmount             int
	WinConditionUpgradeHallLevel   int
	WinConditionUpgradeCastleLevel int
	WinConditionPos                Position

	LoseCondition     LoseCondition
	LoseConditionPos  Position
	LoseConditionDays int

	TeamsCount int
	Teams      []int // len=8 - RED, BLUE, TAN, GREEN, ORANGE, PURPLE, TEAL, PINK

	AvailableHeroes   []byte // 20
	CustomHeroesCount int
	CustomHeroes      []CustomHeroes
}

type Position struct {
	X int
	Y int
	Z int
}

type Rumor struct {
	Name string
	Desc string
}
