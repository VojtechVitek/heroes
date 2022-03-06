package main

import (
	"bytes"
	"fmt"
	"image/png"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/VojtechVitek/heroes/pkg/def"
	"github.com/VojtechVitek/heroes/pkg/lod"
	"github.com/VojtechVitek/heroes/pkg/pcx"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
)

const VERSION = "v0.0.1"

func main() {
	//lodFilename := "./lod/H3sprite.lod"
	//lodFilename := "./lod/H3ab_spr.lod"
	lodFilename := "./lod/H3bitmap.lod"

	if len(os.Args) >= 2 {
		lodFilename = os.Args[1]
	}

	f, err := os.Open(lodFilename)
	if err != nil {
		log.Fatal(err)
	}

	lod, err := lod.Parse(f)
	if err != nil {
		log.Fatal(err)
	}

	srv := &Server{
		lod: lod,
	}

	r := chi.NewRouter()

	logger := httplog.NewLogger("heroes", httplog.Options{JSON: false})
	r.Use(httplog.RequestLogger(logger))

	//r.Get("/favicon.ico", )

	r.Get("/", srv.HandleFS)
	r.Get("/{lodFile}", srv.HandleLod)
	r.Get("/{lodFile}/{filename}", srv.HandleLodFile)

	if err := http.ListenAndServe("0.0.0.0:3003", r); err != nil {
		log.Fatal(err)
	}
}

const usage = `lod H3sprite.lod
`

type Server struct {
	lod *lod.LOD
}

func (s *Server) HandleFS(w http.ResponseWriter, r *http.Request) {
	for _, file := range []string{"H3ab_bmp.lod", "H3ab_spr.lod", "H3bitmap.lod", "H3sprite.lod"} {
		fmt.Fprintf(w, `<a href="/%v">%v</a><br`, file, file)
	}
}

func (s *Server) HandleLod(w http.ResponseWriter, r *http.Request) {
	for _, file := range s.lod.Files() {
		if strings.ToLower(filepath.Ext(file)) == ".def" {
			for i := 0; i < 10; i++ {
				fmt.Fprintf(w, `<a href="/H3sprite.lod/%v?frame=%v"><img src="/H3sprite.lod/%v?frame=%v" /></a> `, file, i, file, i)
			}
			fmt.Fprintf(w, `<br /><br />`)
		} else if strings.ToLower(filepath.Ext(file)) == ".pcx" {
			fmt.Fprintf(w, `<a href="/H3bitmap.lod/%v"><img src="/H3bitmap.lod/%v" /></a><br />`, file, file)
		}
	}
}

func (s *Server) HandleLodFile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	filename := chi.URLParamFromCtx(ctx, "filename")

	switch strings.ToLower(filepath.Ext(filename)) {
	case ".def":
		s.handleDef(w, r)

	case ".pcx":
		s.handlePcx(w, r)

	default:
		w.WriteHeader(500)
		fmt.Fprintf(w, "cannot handle %q", filename)
	}
}

func (s *Server) handleDef(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	filename := chi.URLParamFromCtx(ctx, "filename")

	lodData, err := s.lod.ReadFile(filename)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w, err)
		return
	}

	def, err := def.Parse(bytes.NewReader(lodData))
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "failed to load %q: %v", filename, err)
		return
	}

	frameNo := 0
	if parseFrameNo, err := strconv.Atoi(r.URL.Query().Get("frame")); err == nil {
		frameNo = parseFrameNo
	}

	if len(def.Frames) <= frameNo {
		w.WriteHeader(404)
		fmt.Fprintf(w, "failed to get frame no: %v", frameNo)
		return
	}

	frame := def.Frames[frameNo]

	img, err := frame.Image()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "failed to get frame image: %v", err)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	if err := png.Encode(w, img); err != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w, err)
		return
	}
}

func (s *Server) handlePcx(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	filename := chi.URLParamFromCtx(ctx, "filename")

	lodData, err := s.lod.ReadFile(filename)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w, err)
		return
	}

	pcx, err := pcx.Parse(bytes.NewReader(lodData))
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "failed to load %q: %v", filename, err)
		return
	}

	img, err := pcx.Image()
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
