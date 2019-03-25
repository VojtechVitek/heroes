package mp2

type Level uint16

const (
	Easy   Level = iota // 0
	Normal              // 1
	Hard                // 2
	Expert              // 3
)

var levelString = []string{
	"easy",
	"normal",
	"hard",
	"expert",
}

func (l Level) String() string { return levelString[l] }
