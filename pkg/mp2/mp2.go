package mp2

import (
	"encoding/binary"
	"io"

	"github.com/pkg/errors"
)

type Map struct {
	MagicByte           [4]byte
	Level               Level
	Width               uint8
	Height              uint8
	KingdomColors       AllowColors
	AllowHumanColors    AllowColors
	AllowComputerColors AllowColors
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
