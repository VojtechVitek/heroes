package mp2

import (
	"encoding/binary"
	"fmt"
	"io"
	"strings"

	"github.com/pkg/errors"
)

type Map struct {
	*Header           // 0x0
	Tiles   []MapTile // 0x1AC (428)

	// *Addons

	// Uniq uint32 // EOF - 0x4
	// EOF
}

func LoadMap(r io.Reader) (*Map, error) {
	header, err := LoadHeader(r)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse info")
	}

	tiles := make([]MapTile, header.Width()*header.Height())
	if err := binary.Read(r, binary.LittleEndian, tiles); err != nil {
		return nil, errors.Wrap(err, "failed to parse tiles")
	}

	m := &Map{
		Header: header,
		Tiles:  tiles,
	}

	return m, nil
}

func (m *Map) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Level: %v\nWidth: %v, Height: %v\n", m.Level(), m.Width(), m.Height())
	fmt.Fprintf(&b, "Kingdom colors: %v\nHuman colors: %v\nAI colors: %v\n", m.KingdomColors(), m.AllowHumanColors(), m.AllowAIColors())

	fmt.Fprintf(&b, "Conditions Wins: %v\n", m.ConditionsWins())
	fmt.Fprintf(&b, "AIAlsoWins: %v, AllowNormalVictory: %v\n", m.AIAlsoWins(), m.AllowNormalVictory())
	fmt.Fprintf(&b, "Wins data: %v, %v\n", m.WinsData1(), m.WinsData2())
	fmt.Fprintf(&b, "Conditions Loss: %v\n", m.ConditionsLoss())
	fmt.Fprintf(&b, "Loss data: %v, %v\n", m.LossData1(), m.LossData2())
	fmt.Fprintf(&b, "StartWithHeroes: %v\n", m.StartWithHeroes())

	fmt.Fprintf(&b, "Races: %v\n", m.Races())

	fmt.Fprintf(&b, "Name: %s\n", m.Name())
	fmt.Fprintf(&b, "Description: %s\n", m.Description())

	fmt.Fprintf(&b, "Tiles: %v\n", m.Tiles)

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
