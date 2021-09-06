package rpc

type API interface {
	//GetMapTiles(ctx context.Context) (tiles []*h3m.MapTile, err error)
}

// Define empty response type, since TypeScript WebRPC client doesn't handle empty response body.
type Empty struct{}

// RPC implements API.
type RPC struct{}
