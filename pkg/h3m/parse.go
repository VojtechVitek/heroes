package h3m

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"io"

	"github.com/VojtechVitek/heroes/pkg/bytestream"
)

func Parse(r io.Reader) (*H3M, error) {
	r, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	_, err = b.ReadFrom(r)
	if err != nil {
		return nil, err
	}

	get := bytestream.New(b.Bytes(), binary.LittleEndian)

	h3m := &H3M{
		Format: FileFormat(get.Int(4)),
	}

	h3m.MapInfo.HasHero = get.Bool(1)
	h3m.MapInfo.MapSize = get.Int(4)
	h3m.MapInfo.HasTwoLevels = get.Bool(1)
	nameSize := get.Int(4)
	h3m.MapInfo.Name = get.ReadString(nameSize)
	descSize := get.Int(4)
	h3m.MapInfo.Desc = get.ReadString(descSize)
	h3m.MapInfo.Difficulty = get.Int(1)

	switch h3m.Format {
	case ArmageddonsBlade, ShadowOfDeath:
		h3m.MapInfo.MasteryCap = get.Int(1)
	}

	for i := 0; i < len(h3m.Players); i++ {
		player := &h3m.Players[i]
		player.CanBeHuman = get.Bool(1)
		player.CanBeHuman = get.Bool(1)
		player.Behavior = get.Int(1)

		if h3m.Format == ShadowOfDeath {
			player.AllowedAlignments = get.Int(1)
		}

		player.TownTypes = get.Int(1)

		switch h3m.Format {
		case ArmageddonsBlade, ShadowOfDeath:
			townConflux := get.Int(1)
			_ = townConflux // not used for now..
		}

		player.HasRandomTown = get.Bool(1)
		player.HasMainTown = get.Bool(1)
	}

	return h3m, get.Error()
}
