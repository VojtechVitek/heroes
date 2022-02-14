package def

// https://github.com/vcmi/vcmi/blob/bc1d99431d4b6f075fce3b551a6891fdb4ad5dd1/client/gui/CAnimation.cpp

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/VojtechVitek/heroes/pkg/bytestream"
	"github.com/pkg/errors"
)

type Def struct {
	Type        int
	TotalBlocks int
	Width       int
	Height      int

	Frames []*Frame
}

type Frame struct {
	BlockId int
	Name    string
	Offset  int
	Data    []byte
}

func (frame *Frame) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "frame %v (block %v)", frame.Name, frame.BlockId)

	return b.String()
}

// https://github.com/vcmi/vcmi/blob/bc1d99431d4b6f075fce3b551a6891fdb4ad5dd1/client/gui/CAnimation.cpp
type Sprite struct {
	Size       int
	Format     int
	FullWidth  int
	FullHeight int
	Width      int
	Height     int
	LeftMargin int
	TopMargin  int

	Data []byte
	//Palette
}

type RGBA struct {
	r int
	g int
	b int
	a int
}

func Parse(r io.Reader) (*Def, error) {
	get := bytestream.New(r, binary.LittleEndian)
	def := &Def{
		Type:        get.Int(4), // https://github.com/vcmi/vcmi/blob/bc1d99431d4b6f075fce3b551a6891fdb4ad5dd1/client/gui/CAnimation.cpp#L243
		Width:       get.Int(4),
		Height:      get.Int(4),
		TotalBlocks: get.Int(4),
	}
	if def.TotalBlocks > 1000 {
		return nil, errors.Errorf("too many blocks: %v", def.TotalBlocks)
	}

	var palette [256]RGBA
	for i := 0; i < 256; i++ {
		palette[i].r = get.Int(1)
		palette[i].g = get.Int(1)
		palette[i].b = get.Int(1)
		palette[i].a = 255 // Alpha Opaque
	}

	log.Printf("blocks: %v", def.TotalBlocks)

	for i := 0; i < def.TotalBlocks; i++ {
		blockId := get.Int(4)
		totalFrames := get.Int(4)
		if totalFrames > 1000 {
			return nil, errors.Errorf("too many block entries: %v", totalFrames)
		}

		_ = get.Bytes(8) // unknown

		frames := make([]*Frame, totalFrames)

		for i := 0; i < totalFrames; i++ {
			frames[i] = &Frame{
				BlockId: blockId,
				Name:    get.String(13),
			}
		}
		for i := 0; i < totalFrames; i++ {
			frames[i].Offset = get.Int(4)
			frames[i].Data = get.Bytes(def.TotalBlocks * def.Width * def.Height)
		}

		// TODO: Split into [blocks][frames] ?

		def.Frames = append(def.Frames, frames...)

		for _, frame := range frames {
			fmt.Printf("%v\n\n", frame)
		}
	}

	return def, nil
}
