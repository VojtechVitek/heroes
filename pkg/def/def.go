package def

// https://github.com/vcmi/vcmi/blob/bc1d99431d4b6f075fce3b551a6891fdb4ad5dd1/client/gui/CAnimation.cpp

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/VojtechVitek/heroes/pkg/bytestream"
	"github.com/pkg/errors"
)

type Def struct {
	Type        int
	TotalBlocks int
	Width       int
	Height      int
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
	now := time.Now()
	def := &Def{
		Type:        get.Int(4), // https://github.com/vcmi/vcmi/blob/bc1d99431d4b6f075fce3b551a6891fdb4ad5dd1/client/gui/CAnimation.cpp#L243
		Width:       get.Int(4), // not used?
		Height:      get.Int(4), // not used?
		TotalBlocks: get.Int(4),
	}
	if def.TotalBlocks > 1000 {
		return nil, errors.Errorf("too many blocks: %v", def.TotalBlocks)
	}

	log.Println("pre-pallete:", time.Since(now))
	var palette [256]RGBA
	for i := 0; i < 256; i++ {
		palette[i].r = get.Int(1)
		palette[i].g = get.Int(1)
		palette[i].b = get.Int(1)
		palette[i].a = 255 // Alpha Opaque
	}
	log.Println("post-pallete:", time.Since(now))

	fmt.Printf("total blocks: %v\n", def.TotalBlocks)

	for i := 0; i < def.TotalBlocks; i++ {
		blockId := get.Int(4)
		totalEntries := get.Int(4)
		_ = get.Bytes(8)

		fmt.Println("total entries:", totalEntries)

		if totalEntries > 1000 {
			return nil, errors.Errorf("too many block entries: %v", totalEntries)
		}

		for i := 0; i < totalEntries; i++ {
			frameName := get.String(13)
			if err := get.Error(); err != nil {
				return nil, errors.Wrap(err, "failed to iterate")
			}
			fmt.Println("frame name:", frameName)
		}

		for i := 0; i < totalEntries; i++ {
			offset := get.Int(4)
			_ = offset
		}

		log.Printf("blockId=%v, totalEntries=%v", blockId, totalEntries)
	}

	//def.Size = get.Int(4)

	// Sprite
	sprite := &Sprite{
		Format:     get.Int(4),
		FullWidth:  get.Int(4),
		FullHeight: get.Int(4),
		Width:      get.Int(4),
		Height:     get.Int(4),
		LeftMargin: get.Int(4),
		TopMargin:  get.Int(4),
		//DataImage [][]byte
	}

	switch sprite.Format {
	default:
		log.Fatalf("sprite.Format: %v", sprite.Format)
	}

	return def, nil
}
