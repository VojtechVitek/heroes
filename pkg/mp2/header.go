package mp2

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"strings"

	"github.com/pkg/errors"
)

// Header
//
// We have to unpack these bytes manually. We can't
// use binary.Read() and unmarshal the data
// automatically into struct fields, since Go adds
// struct padding whenever it feels like when you
// combine uint16 and uint8. ¯\_(ツ)_/¯
//
// 0x0		0		MagicByte			4 bytes
// 0x4		4		Difficulty			2 bytes
// 0x6		6		Width				1 byte
// 0x7		7		Height				1 byte
// 0x8		8		KingdomColors		6 bytes
// 0xE		14		AllowHumanColors	6 bytes
// 0x14		20		AllowAIColors		6 bytes
// 0x1A		26		_ (Kingdom Count?)	3 bytes
// 0x1D		29		VictoryConditions	1 byte
// 0x1E		30		AIAlsoWins			1 byte
// 0x1F		31		AllowNormalVictory	1 byte
// 0x20		32		VictoryData1		2 bytes
// 0x22		34		LossConditions		1 byte
// 0x23		35		LossData1			2 bytes
// 0x25		37		StartWithHeroes		1 byte
// 0x26		38		Races				6 bytes
// 0x2C		44		VictoryData2		2 bytes
// 0x2e		46		LossData2			2 bytes
// 0x30		48		_					10 bytes
// 0x3A		58		Name				16 bytes
// 0x4A		74		_					44 bytes
// 0x76		118		Description			143 bytes
// 0x105	261		_					159 bytes
// 0x1A4	420		Width (duplicate)	4 bytes
// 0x1A8	424		Height (duplicate)	4 bytes
type Header [428]byte

func LoadHeader(r io.Reader) (*Header, error) {
	header := &Header{}
	if err := binary.Read(r, binary.LittleEndian, header); err != nil {
		return nil, errors.Wrap(err, "failed to deserialize map info data")
	}

	if err := header.Validate(); err != nil {
		return nil, errors.Wrap(err, "failed to validate map")
	}

	return header, nil
}

func (h Header) Validate() error {
	magicByte := binary.BigEndian.Uint32(h[0:4])
	if magicByte != uint32(0x5C000000) {
		return errors.Errorf("expected magic byte %v, got %v", uint32(0x5C000000), magicByte)
	}

	// Duplicated Width and Height fields must match.
	widthDuplicate := int(binary.LittleEndian.Uint32(h[420:424]))
	if h.Width() != widthDuplicate {
		return errors.Errorf("map width mismatch: %v != %v", h.Width(), widthDuplicate)
	}

	heightDuplicate := int(binary.LittleEndian.Uint32(h[424:428]))
	if h.Height() != heightDuplicate {
		return errors.Errorf("map height mismatch: %v != %v", h.Height(), heightDuplicate)
	}

	// Map Width must be same as Height. It's always a square.
	if widthDuplicate != heightDuplicate {
		return errors.Errorf("map must be a square: got width=%v, height=%v", widthDuplicate, heightDuplicate)
	}

	return nil
}

func (h Header) Difficulty() Difficulty { return Difficulty(binary.LittleEndian.Uint16(h[4:6])) }
func (h Header) Width() int             { return int(h[6]) }
func (h Header) Height() int            { return int(h[7]) }
func (h Header) KingdomColors() (colors AllowColors) {
	_ = binary.Read(bytes.NewReader(h[8:14]), binary.LittleEndian, &colors)
	return
}
func (h Header) AllowHumanColors() (colors AllowColors) {
	_ = binary.Read(bytes.NewReader(h[14:20]), binary.LittleEndian, &colors)
	return
}
func (h Header) AllowAIColors() (colors AllowColors) {
	_ = binary.Read(bytes.NewReader(h[20:26]), binary.LittleEndian, &colors)
	return
}
func (h Header) VictoryConditions() VictoryConditions { return VictoryConditions(h[29]) }
func (h Header) AIAlsoWins() bool                     { return uint8(h[30]) > 0 }
func (h Header) AllowNormalVictory() bool             { return uint8(h[31]) > 0 }
func (h Header) VictoryData1() uint16                 { return binary.LittleEndian.Uint16(h[32:34]) }
func (h Header) LossConditions() LossConditions       { return LossConditions(h[34]) }
func (h Header) LossData1() uint16                    { return binary.LittleEndian.Uint16(h[35:37]) }
func (h Header) StartWithHeroes() bool                { return uint8(h[37]) > 0 }
func (h Header) Races() (races Races) {
	_ = binary.Read(bytes.NewReader(h[38:44]), binary.LittleEndian, &races)
	return
}
func (h Header) VictoryData2() uint16 { return binary.LittleEndian.Uint16(h[44:46]) }
func (h Header) LossData2() uint16    { return binary.LittleEndian.Uint16(h[46:48]) }
func (h Header) Name() string         { return nullTerminatedString(h[58:74]) }
func (h Header) Description() string  { return nullTerminatedString(h[118:261]) }

func (h Header) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Name: %s\n", h.Name())
	fmt.Fprintf(&b, "Difficulty: %v\nWidth: %v, Height: %v\n", h.Difficulty(), h.Width(), h.Height())
	fmt.Fprintf(&b, "Description: %s\n", h.Description())

	fmt.Fprintf(&b, "Players %v, Human: %v, AI: %v\n", h.KingdomColors(), h.AllowHumanColors(), h.AllowAIColors())
	fmt.Fprintf(&b, "Races: %v\n", h.Races())

	fmt.Fprintf(&b, "Victory conditions: %v\n", h.VictoryConditions())
	//fmt.Fprintf(&b, "AIAlsoWins: %v, AllowNormalVictory: %v\n", h.AIAlsoWins(), h.AllowNormalVictory())
	//fmt.Fprintf(&b, "Victory data: %v, %v\n", h.VictoryData1(), h.VictoryData2())
	fmt.Fprintf(&b, "Loss conditions: %v\n", h.LossConditions())
	//fmt.Fprintf(&b, "Loss data: %v, %v\n", h.LossData1(), h.LossData2())
	fmt.Fprintf(&b, "StartWithHeroes: %v\n", h.StartWithHeroes())

	return b.String()
}
