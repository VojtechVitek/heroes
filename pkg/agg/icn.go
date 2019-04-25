package agg

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type ICN struct {
	sprites []*Sprite
}

// 0x00  Number of sprites         2 bytes
// 0x02  Data size                 4 bytes
// 0x06  Sprites headers + data    (each header's dataOffset points here)
func NewICN(data []byte, pallete pallete) (*ICN, error) {
	var u16 uint16
	if err := binary.Read(bytes.NewReader(data[0:2]), binary.LittleEndian, &u16); err != nil {
		return nil, errors.Wrap(err, "failed to read number of sprites")
	}
	numSprites := int(u16)

	var u32 uint32
	if err := binary.Read(bytes.NewReader(data[2:6]), binary.LittleEndian, &u32); err != nil {
		return nil, errors.Wrap(err, "failed to read data size")
	}
	dataSize := int(u32)

	data = data[6:] // Sprites headers + data only. That's what sprite headers offsets reference.
	if len(data) != dataSize {
		return nil, errors.Errorf("expected data size %v bytes, got %v bytes", dataSize, len(data))
	}

	sprites := make([]*Sprite, 0, numSprites)
	for i := 0; i < numSprites; i++ {
		sprite, err := NewSprite(data, i)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse sprite")
		}
		sprites = append(sprites, sprite)
	}

	return &ICN{sprites: sprites}, nil
}

func (icn *ICN) Sprite(index int) *Sprite {
	return icn.sprites[index]
}

func (icn *ICN) Sprites() []*Sprite {
	return icn.sprites
}

func (icn *ICN) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Number of sprites: %v\n", len(icn.sprites))

	for _, header := range icn.sprites {
		fmt.Fprintf(&b, "%+v\n", header)
	}

	return b.String()
}
