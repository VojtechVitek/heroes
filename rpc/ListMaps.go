package rpc

import (
	"context"
	"io/ioutil"
)

func (rpc *RPC) ListMaps(ctx context.Context) (maps []string, err error) {
	files, err := ioutil.ReadDir("./maps/")
	if err != nil {
		return nil, WrapError(ErrInternal, err, "failed to list maps")
	}

	for _, fileInfo := range files {
		maps = append(maps, fileInfo.Name())
	}

	return maps, nil
}
