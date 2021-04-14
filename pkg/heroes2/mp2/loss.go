package mp2

import (
	"fmt"
)

type LossConditions uint8

const (
	LooseAll  LossConditions = iota // 0
	LooseTown                       // 1
	LooseHero                       // 2
	TimeLimit                       // 3
)

var lossConditionsString = []string{
	"Loose All",
	"Loose Town",
	"Loose Hero",
	"Time Limit",
}

func (l LossConditions) String() string {
	if int(l) > len(lossConditionsString)-1 {
		return fmt.Sprintf("%v", int(l))
	}
	return lossConditionsString[l]
}
