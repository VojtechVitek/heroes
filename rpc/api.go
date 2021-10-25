package rpc

import (
	"context"

	"github.com/VojtechVitek/heroes/pkg/h3m"
)

type API interface {
	GetMap(ctx context.Context) (Map *Map, err error)
}

type Map struct {
	*h3m.H3M
}

// Define empty response type, since TypeScript WebRPC client doesn't handle empty response body.
type Empty struct{}

// RPC implements API.
type RPC struct {
	Maps map[string]*h3m.H3M
}
