package mp2

import (
	"fmt"
)

type VictoryConditions uint8

const (
	DefeatAll      VictoryConditions = iota // 0
	CaptureTown                             // 1
	DefeatHero                              // 2
	FindArtifact                            // 3
	DefeatTeam                              // 4
	AccumulateGold                          // 5
)

var victoryConditionsString = []string{
	"Defeat All",
	"Capture Town",
	"Defeat Hero",
	"Find Artifact",
	"Defeat Team",
	"Accumulate Gold",
}

func (l VictoryConditions) String() string {
	if int(l) > len(victoryConditionsString)-1 {
		return fmt.Sprintf("%v", int(l))
	}
	return victoryConditionsString[l]
}
