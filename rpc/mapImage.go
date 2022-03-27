package rpc

import (
	"fmt"
	"image/png"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (rpc RPC) HandleMap(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	mapName := chi.URLParamFromCtx(ctx, "mapName")
	m, err := rpc.GetMap(ctx, mapName)
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprintf(w, "failed to find map: %s: %v", mapName, err)
		return
	}

	img, err := m.Image()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "failed to get pcx image: %v", err)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	if err := png.Encode(w, img); err != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w, err)
		return
	}
}
