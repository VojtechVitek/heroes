package h3m

// The binary format sourced from
// https://github.com/potmdehex/homm3tools/blob/master/h3m/h3mlib/h3m_structures/h3m.h
// The MIT License (MIT)
// Copyright (c) 2016 John Åkerblom

// Maps commonly end with 124 bytes of null padding. Extra content at end
// is ok.
type H3M struct {
	Format FileFormat
	MapInfo
	Players []Player
}

type MapInfo struct {
	// Basic info.
	HasHero      bool
	MapSize      MapSize
	HasTwoLevels bool
	Name         string
	Desc         string
	Difficulty   int
	MasteryCap   int // Only set on ArmageddonsBlade and ShadowOfDeath maps.

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
	Teams      [8]int // RED, BLUE, TAN, GREEN, ORANGE, PURPLE, TEAL, PINK

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
