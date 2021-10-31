package rpc

import (
	"context"
	"fmt"
)

func (rpc *RPC) GetMap(ctx context.Context) (m *Map, err error) {
	ma, _ := rpc.Maps["./maps/Loss of Innocence.h3m"]

	fmt.Println(ma)
	return &Map{ma}, nil
}
