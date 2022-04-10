package rpc

import (
	"context"
	"os"
	"path/filepath"

	"github.com/VojtechVitek/heroes/pkg/h3m"
	"github.com/pkg/errors"
)

func (rpc *RPC) GetMap(ctx context.Context, filename string) (m *Map, err error) {
	h3map, ok := rpc.Maps[filename]
	if !ok {
		f, err := os.Open(filepath.Join("./maps/", filename))
		if err != nil {
			return nil, errors.Wrapf(err, "failed to open %v", filename)
		}

		h3map, err = h3m.Parse(f)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to parse %v", filename)
		}

		rpc.Maps[filename] = h3map
	}

	return &Map{h3map}, nil
}
