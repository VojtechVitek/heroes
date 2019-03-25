package mp2

import (
	"encoding/binary"
	"fmt"
	"io"
	"strings"

	"github.com/pkg/errors"
)

type Map struct {
	MagicByte          [4]byte
	Level              Level
	Width              uint8
	Height             uint8
	KingdomColors      AllowColors
	AllowHumanColors   AllowColors
	AllowAIColors      AllowColors
	_                  [0x1D]byte // TODO: Kingdom count on 0x1A?
	ConditionsWins     uint8
	AIAlsoWins         Bool
	AllowNormalVictory Bool
	WinsData1          uint16
	_                  [0x2c]byte
	WinsData2          uint16
	_                  [0x22]byte
	ConditionsLoss     uint8
	LossData1          uint16
	_                  [0x2e]byte
	LossData2          uint16
	_                  [0x25]byte
	DontStartWithHero  Bool
	Race               Race
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
	fmt.Fprintf(&b, "DontStartWithHero: %v\n", m.DontStartWithHero)

	fmt.Fprintf(&b, "Race: %v\n", m.Race)
	return b.String()
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
