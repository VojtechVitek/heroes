package h3m

import "fmt"

type WinCondition int

const (
	WIN_ACQUIRE_ARTIFACT     WinCondition = 0x00
	WIN_ACCUMULATE_CREATURES WinCondition = 0x01
	WIN_ACCUMULATE_RESOURCES WinCondition = 0x02
	WIN_UPGRADE_TOWN         WinCondition = 0x03
	WIN_BUILD_GRAIL          WinCondition = 0x04 // position
	WIN_DEFEAT_HERO          WinCondition = 0x05 // position
	WIN_CAPTURE_TOWN         WinCondition = 0x06
	WIN_DEFEAT_MONSTER       WinCondition = 0x07
	WIN_FLAG_DWELLINGS       WinCondition = 0x08
	WIN_FLAG_MINES           WinCondition = 0x09
	WIN_TRANSPORT_ARTIFACT   WinCondition = 0x0A
	WIN_DEFEAT_ALL_ENEMIES   WinCondition = 0xFF
)

func (c WinCondition) String() string {
	switch c {
	case WIN_ACQUIRE_ARTIFACT:
		return "ACQUIRE_ARTIFACT"
	case WIN_ACCUMULATE_CREATURES:
		return "ACCUMULATE_CREATURES"
	case WIN_ACCUMULATE_RESOURCES:
		return "ACCUMULATE_RESOURCES"
	case WIN_UPGRADE_TOWN:
		return "UPGRADE_TOWN"
	case WIN_BUILD_GRAIL:
		return "BUILD_GRAIL"
	case WIN_DEFEAT_HERO:
		return "DEFEAT_HERO"
	case WIN_CAPTURE_TOWN:
		return "CAPTURE_TOWN"
	case WIN_DEFEAT_MONSTER:
		return "DEFEAT_MONSTER"
	case WIN_FLAG_DWELLINGS:
		return "FLAG_DWELLINGS"
	case WIN_FLAG_MINES:
		return "FLAG_MINES"
	case WIN_TRANSPORT_ARTIFACT:
		return "TRANSPORT_ARTIFACT"
	case WIN_DEFEAT_ALL_ENEMIES:
		return "WIN_DEFEAT_ALL_ENEMIES"
	default:
		return fmt.Sprintf("unknown condition (%X)", int(c))
	}
}

func (c WinCondition) Is(conditions ...WinCondition) bool {
	for _, condition := range conditions {
		if c == condition {
			return true
		}
	}
	return false
}

type LoseCondition int

const (
	LOSE_TOWN                LoseCondition = 0x00
	LOSE_HERO                LoseCondition = 0x01
	LOSE_TIME                LoseCondition = 0x02
	LOSE_ALL_TOWNS_OR_HEROES               = 0xEF
)

func (c LoseCondition) String() string {
	switch c {
	case LOSE_TOWN:
		return "LOSE_TOWN"
	case LOSE_HERO:
		return "LOSE_HERO"
	case LOSE_TIME:
		return "TIME"
	case LOSE_ALL_TOWNS_OR_HEROES:
		return "LOSE_ALL_TOWNS_OR_HEROES"
	default:
		return fmt.Sprintf("unknown condition (%X)", int(c))
	}
}
