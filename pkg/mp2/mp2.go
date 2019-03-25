package mp2

import (
	"encoding/binary"
	"fmt"
	"io"
	"strings"

	"github.com/pkg/errors"
)

type Map struct {
	MagicByte          [4]byte     // 0x0
	Level              Level       // 0x4
	Width              uint8       // 0x6
	Height             uint8       // 0x7
	KingdomColors      AllowColors // 0x8
	AllowHumanColors   AllowColors // 0xE  (14)
	AllowAIColors      AllowColors // 0x14 (20)
	_                  [3]byte     // 0x1A (26) TODO: Kingdom count?
	ConditionsWins     uint8       // 0x1D (29)
	AIAlsoWins         Bool        // 0x1E (30)
	AllowNormalVictory Bool        // 0x1F (31)
	WinsData1          uint16      // 0x20 (32)
	ConditionsLoss     uint8       // 0x22 (34)
	LossData1          uint16      // 0x23 (35)
	StartWithHeroes    Bool        // 0x25 (37)
	Race               RaceColor   // 0x26 (38)
	WinsData2          uint16      // 0x2C (44)
	LossData2          uint16      // 0x2e (46)
	_                  [10]byte    // 0x30 (48)
	Name               [16]byte    // 0x3A (58)
	_                  [44]byte    // 0x4A (74)
	Description        [143]byte   // 0x76 (118)
}

func (m *Map) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Level: %v\nWidth: %v, Height: %v\n", m.Level, m.Width, m.Height)
	fmt.Fprintf(&b, "Kingdom colors: %v\nHuman colors: %v\nAI colors: %v\n", m.KingdomColors, m.AllowHumanColors, m.AllowAIColors)

	fmt.Fprintf(&b, "Conditions Wins: %v\n", m.ConditionsWins)
	fmt.Fprintf(&b, "AIAlsoWins: %v, AllowNormalVictory: %v\n", m.AIAlsoWins, m.AllowNormalVictory)
	fmt.Fprintf(&b, "Wins data: %v, %v\n", m.WinsData1, m.WinsData2)
	fmt.Fprintf(&b, "Conditions Loss: %v\n", m.ConditionsLoss)
	fmt.Fprintf(&b, "Loss data: %v, %v\n", m.LossData1, m.LossData2)
	fmt.Fprintf(&b, "StartWithHeroes: %v\n", m.StartWithHeroes)

	fmt.Fprintf(&b, "Race: %v\n", m.Race)

	fmt.Fprintf(&b, "Name: %s\n", nullTerminatedString(m.Name[:]))
	fmt.Fprintf(&b, "Description: %s\n", nullTerminatedString(m.Description[:]))

	return b.String()
}

func nullTerminatedString(nullTerminatedString []byte) string {
	return string(nullTerminatedString[:nullTerminatedStringLen(nullTerminatedString)])
}

func nullTerminatedStringLen(nullTerminatedString []byte) int {
	for i := 0; i < len(nullTerminatedString); i++ {
		if nullTerminatedString[i] == 0 {
			return i
		}
	}
	return len(nullTerminatedString)
}

func LoadMap(r io.Reader) (*Map, error) {
	m := &Map{}
	if err := binary.Read(r, binary.LittleEndian, m); err != nil {
		return nil, errors.Wrap(err, "failed to deserialize map data")
	}

	if magicByte := [4]byte{0x5C, 0x00, 0x00, 0x00}; m.MagicByte != magicByte {
		return nil, errors.Errorf("expected magic byte %v, got %v", magicByte, m.MagicByte)
	}

	return m, nil
}
