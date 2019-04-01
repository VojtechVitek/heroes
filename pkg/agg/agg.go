package agg

import (
	"encoding/binary"
	"io"

	"github.com/pkg/errors"
)

type AGG struct {
}

func Load(r io.Reader) (*AGG, error) {
	agg := &AGG{}

	if err := binary.Read(r, binary.LittleEndian, agg); err != nil {
		return nil, errors.Wrap(err, "failed to parse AGG header")
	}

	// tiles := make(Tiles, header.Width()*header.Height())
	// if err := binary.Read(r, binary.LittleEndian, tiles); err != nil {
	// 	return nil, errors.Wrap(err, "failed to parse tiles")
	// }

	return agg, nil
}
