package mp2

type Difficulty uint16

const (
	Easy   Difficulty = iota // 0
	Normal                   // 1
	Hard                     // 2
	Expert                   // 3
)

var difficultyString = []string{
	"easy",
	"normal",
	"hard",
	"expert",
}

func (l Difficulty) String() string { return difficultyString[l] }
