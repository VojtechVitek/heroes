package mp2

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"strings"

	"github.com/pkg/errors"
)

type Map struct {
	*MapInfo // 0x0

	MapTiles []MapTile // 0x1AC (428)

	// *MapAddons //

	// Uniq uint32 // EOF - 0x4
	// EOF
}

// MapInfo
// Slice these bytes manually, can't unmarshall
// the data automatically into struct fields, since
// Go adds padding whenever it feels like ¯\_(ツ)_/¯
//
// 0x0		0		MagicByte			4 bytes
// 0x4		4		Level				2 bytes
// 0x6		6		Width				1 byte
// 0x7		7		Height				1 byte
// 0x8		8		KingdomColors		6 bytes
// 0xE		14		AllowHumanColors	6 bytes
// 0x14		20		AllowAIColors		6 bytes
// 0x1A		26		_ (Kingdom Count?)	3 bytes
// 0x1D		29		ConditionsWins		1 byte
// 0x1E		30		AIAlsoWins			1 byte
// 0x1F		31		AllowNormalVictory	1 byte
// 0x20		32		WinsData1			2 bytes
// 0x22		34		ConditionsLoss		1 byte
// 0x23		35		LossData1			2 bytes
// 0x25		37		StartWithHeroes		1 byte
// 0x26		38		Races				6 bytes
// 0x2C		44		WinsData2			2 bytes
// 0x2e		46		LossData2			2 bytes
// 0x30		48		_					10 bytes
// 0x3A		58		Name				16 bytes
// 0x4A		74		_					44 bytes
// 0x76		118		Description			143 bytes
// 0x105	261		_					159 bytes
// 0x1A4	420		Width duplicate		4 bytes
// 0x1A8	424		Height duplicate	4 bytes
type MapInfo [428]byte

func (i MapInfo) Validate() error {
	magicByte := binary.BigEndian.Uint32(i[0:4])
	if magicByte != uint32(0x5C000000) {
		return errors.Errorf("expected magic byte %v, got %v", uint32(0x5C000000), magicByte)
	}

	// Duplicated Width and Height fields must match.
	widthDuplicate := int(binary.LittleEndian.Uint32(i[420:424]))
	if i.Width() != widthDuplicate {
		return errors.Errorf("map width mismatch: %v != %v", i.Width(), widthDuplicate)
	}

	heightDuplicate := int(binary.LittleEndian.Uint32(i[424:428]))
	if i.Height() != heightDuplicate {
		return errors.Errorf("map height mismatch: %v != %v", i.Height(), heightDuplicate)
	}

	// Map Width must be same as Height. It's always a square.
	if widthDuplicate != heightDuplicate {
		return errors.Errorf("map must be a square: got width=%v, height:%v", widthDuplicate, heightDuplicate)
	}

	return nil
}
func (i MapInfo) Level() Level { return Level(binary.LittleEndian.Uint16(i[4:6])) }
func (i MapInfo) Width() int   { return int(i[6]) }
func (i MapInfo) Height() int  { return int(i[7]) }
func (i MapInfo) KingdomColors() (colors AllowColors) {
	if err := binary.Read(bytes.NewReader(i[8:14]), binary.LittleEndian, &colors); err != nil {
		panic(err)
	}
	return
}
func (i MapInfo) AllowHumanColors() (colors AllowColors) {
	if err := binary.Read(bytes.NewReader(i[14:20]), binary.LittleEndian, &colors); err != nil {
		panic(err)
	}
	return
}
func (i MapInfo) AllowAIColors() (colors AllowColors) {
	if err := binary.Read(bytes.NewReader(i[20:26]), binary.LittleEndian, &colors); err != nil {
		panic(err)
	}
	return
}
func (i MapInfo) ConditionsWins() uint8    { return uint8(i[29]) }
func (i MapInfo) AIAlsoWins() bool         { return uint8(i[30]) > 0 }
func (i MapInfo) AllowNormalVictory() bool { return uint8(i[31]) > 0 }
func (i MapInfo) WinsData1() uint32        { return binary.LittleEndian.Uint32(i[32:34]) }
func (i MapInfo) ConditionsLoss() uint8    { return uint8(i[34]) }
func (i MapInfo) LossData1() uint32        { return binary.LittleEndian.Uint32(i[35:37]) }
func (i MapInfo) StartWithHeroes() bool    { return uint8(i[37]) > 0 }
func (i MapInfo) Races() (races [6]Race) {
	if err := binary.Read(bytes.NewReader(i[38:44]), binary.LittleEndian, &races); err != nil {
		panic(err)
	}
	return
}
func (i MapInfo) WinsData2() uint32   { return binary.LittleEndian.Uint32(i[44:46]) }
func (i MapInfo) LossData2() uint32   { return binary.LittleEndian.Uint32(i[46:48]) }
func (i MapInfo) Name() string        { return nullTerminatedString(i[58:74]) }
func (i MapInfo) Description() string { return nullTerminatedString(i[118:261]) }

type MapTile struct {
	TileIndex     uint16 // 0x00
	ObjectName1   uint8  // 0x02
	IndexName1    uint8  // 0x03
	Quantity1     uint8  // 0x04
	Quantity2     uint8  // 0x05
	ObjectName2   uint8  // 0x06
	IndexName2    uint8  // 0x07
	Shape         uint8  // 0x08
	GeneralObject uint8  // 0x09
} // sizeof: 0x14 (20)

func LoadMap(r io.Reader) (*Map, error) {
	mapInfo, err := LoadMapInfo(r)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load map info")
	}

	mapTiles := make([]MapTile, mapInfo.Width()*mapInfo.Height())
	if err := binary.Read(r, binary.LittleEndian, mapTiles); err != nil {
		return nil, errors.Wrap(err, "failed to deserialize map tiles")
	}

	m := &Map{
		MapInfo:  mapInfo,
		MapTiles: mapTiles,
	}

	return m, nil
}

func LoadMapInfo(r io.Reader) (*MapInfo, error) {
	mapInfo := &MapInfo{}
	if err := binary.Read(r, binary.LittleEndian, mapInfo); err != nil {
		return nil, errors.Wrap(err, "failed to deserialize map info data")
	}

	if err := mapInfo.Validate(); err != nil {
		return nil, errors.Wrap(err, "failed to validate map")
	}

	return mapInfo, nil
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

	fmt.Fprintf(&b, "MapTiles: %v\n", m.MapTiles)

	return b.String()
}

func (t MapTile) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "index: %v\n", t.TileIndex)

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
