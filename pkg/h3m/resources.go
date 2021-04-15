package h3m

type Resource int

const (
	Wood    Resource = 0
	Mercury Resource = 1
	Ore     Resource = 2
	Sulfur  Resource = 3
	Crystal Resource = 4
	Gems    Resource = 5
	Gold    Resource = 6
)

func (r Resource) String() string {
	switch r {
	case Wood:
		return "Wood"
	case Mercury:
		return "Mercury"
	case Ore:
		return "Ore"
	case Sulfur:
		return "Sulfur"
	case Gems:
		return "Gems"
	case Gold:
		return "Gold"
	default:
		panic("unknown resource %q")
	}
}
