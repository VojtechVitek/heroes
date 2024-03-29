package def

// https://github.com/vcmi/vcmi/blob/bc1d99431d4b6f075fce3b551a6891fdb4ad5dd1/client/gui/CAnimation.cpp

import (
	"encoding/binary"
	"io"

	"github.com/VojtechVitek/heroes/pkg/bytestream"
	"github.com/VojtechVitek/heroes/pkg/palette"
	"github.com/pkg/errors"
)

type Def struct {
	Type        int
	TotalBlocks int
	Width       int
	Height      int

	Palette palette.Palette

	Frames []*Frame

	Data []byte
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
}

func Parse(r io.ReadSeeker) (*Def, error) {
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

	def.Palette = get.Bytes(256 * 3)

	for i := 0; i < def.TotalBlocks; i++ {
		blockId := get.Int(4)
		totalFrames := get.Int(4)
		if totalFrames > 1000 {
			return nil, errors.Errorf("too many block entries: %v", totalFrames)
		}

		_ = get.Bytes(8) // unknown

		frames := make([]*Frame, totalFrames)

		// Get name of each frame (13 bytes).
		for i := 0; i < totalFrames; i++ {
			frames[i] = &Frame{
				BlockId: blockId,
				Name:    get.String(13),
				Palette: &def.Palette,
				Width:   def.Width,
				Height:  def.Height,
			}
		}
		// Get offset of each frame (1 byte).
		for i := 0; i < totalFrames; i++ {
			frames[i].Offset = get.Int(4)
		}

		// TODO: Split into [blocks][frames] ?

		def.Frames = append(def.Frames, frames...)
	}

	for _, frame := range def.Frames {
		r.Seek(int64(frame.Offset), 0)
		get := bytestream.New(r, binary.LittleEndian)

		frame.size = get.Int(4)
		frame.Format = get.Int(4)
		frame.FullWidth = get.Int(4)
		frame.FullHeight = get.Int(4)
		frame.Width2 = get.Int(4)
		frame.Height2 = get.Int(4)
		frame.LeftMargin = get.Int(4)
		frame.RightMargin = get.Int(4)

		frame.data = get.Bytes(frame.size)
	}

	return def, get.Error()
}
