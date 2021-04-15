package h3m

import "fmt"

type Condition int

const (
	ACQUIRE_ARTIFACT     Condition = 0x00
	ACCUMULATE_CREATURES Condition = 0x01
	ACCUMULATE_RESOURCES Condition = 0x02
	UPGRADE_TOWN         Condition = 0x03
	BUILD_GRAIL          Condition = 0x04
	DEFEAT_HERO          Condition = 0x05
	CAPTURE_TOWN         Condition = 0x06
	DEFEAT_MONSTER       Condition = 0x07
	FLAG_DWELLINGS       Condition = 0x08
	FLAG_MINES           Condition = 0x09
	TRANSPORT_ARTIFACT   Condition = 0x0A
)

func (c Condition) String() string {
	switch c {
	case ACQUIRE_ARTIFACT:
		return "ACQUIRE_ARTIFACT"
	case ACCUMULATE_CREATURES:
		return "ACCUMULATE_CREATURES"
	case ACCUMULATE_RESOURCES:
		return "ACCUMULATE_RESOURCES"
	case UPGRADE_TOWN:
		return "UPGRADE_TOWN"
	case BUILD_GRAIL:
		return "BUILD_GRAIL"
	case DEFEAT_HERO:
		return "DEFEAT_HERO"
	case CAPTURE_TOWN:
		return "CAPTURE_TOWN"
	case DEFEAT_MONSTER:
		return "DEFEAT_MONSTER"
	case FLAG_DWELLINGS:
		return "FLAG_DWELLINGS"
	case FLAG_MINES:
		return "FLAG_MINES"
	case TRANSPORT_ARTIFACT:
		return "TRANSPORT_ARTIFACT"
	default:
		return fmt.Sprintf("unknown condition (%x)", c)
	}
}
