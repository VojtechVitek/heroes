package mp2

import "fmt"

type Race uint16

const (
	Knight      Race = iota // 0
	Barbarian               // 1
	Sorcerer                // 2
	Warlock                 // 3
	Wizzard                 // 4
	Necromancer             // 5
	Multi                   // 6
	Random                  // 7
)

var raceString = []string{
	"knight",
	"barbarian",
	"sorcerer",
	"warlock",
	"wizzard",
	"necromancer",
	"multi",
	"random",
}

func (l Race) String() string {
	if int(l) > len(raceString)-1 {
		return fmt.Sprintf("%v", int(l))
	}
	return raceString[l]
}
