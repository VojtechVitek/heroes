package rpc

import (
	"context"

	"github.com/VojtechVitek/heroes/pkg/h3m"
)

type API interface {
	GetMapTiles(ctx context.Context) (tiles []*MapTile, err error)
}

type MapTile struct {
	*h3m.Tile
}

// Define empty response type, since TypeScript WebRPC client doesn't handle empty response body.
type Empty struct{}

// RPC implements API.
type RPC struct {
	Maps map[string]*h3m.H3M
}
