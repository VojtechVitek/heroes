package rpc

import "context"

func (rpc *RPC) GetMap(ctx context.Context) (m *Map, err error) {
	ma, _ := rpc.Maps["./maps/Loss of Innocence.h3m"]
	return &Map{ma}, nil
}
