package main

import (
	"bytes"
	"fmt"
	"image/png"
	"log"
	"net/http"
	"os"

	"github.com/VojtechVitek/heroes/pkg/def"
	"github.com/VojtechVitek/heroes/pkg/lod"
	"github.com/go-chi/chi/v5"
)

const VERSION = "v0.0.1"

func main() {
	//lodFilename := "./lod/H3bitmap.lod"
	lodFilename := "./lod/H3sprite.lod"

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

	//fmt.Println(lod.Files())

	srv := &Server{
		lod: lod,
	}

	r := chi.NewRouter()

	//r.Get("/favicon.ico", )

	r.Get("/{lodFile}.lod", srv.HandleLod)
	r.Get("/{lodFile}.lod/{defFile}.def", srv.HandleDef)
	r.Get("/{lodFile}.lod/{defFile}.def/palette", srv.HandleDefPalette)

	if err := http.ListenAndServe("0.0.0.0:3003", r); err != nil {
		log.Fatal(err)
	}
}

const usage = `lod H3sprite.lod
`

type Server struct {
	lod *lod.LOD
}

func (s *Server) HandleLod(w http.ResponseWriter, r *http.Request) {
	for _, file := range s.lod.Files() {
		fmt.Fprintf(w, `<a href="/H3sprite.lod/%v">%v</a><br />`, file, file)
	}
}

func (s *Server) HandleDef(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	defFilename := fmt.Sprintf("%v.def", chi.URLParamFromCtx(ctx, "defFile"))
	defData, err := s.lod.ReadFile(defFilename)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w, err)
		return
	}

	def, err := def.Parse(bytes.NewReader(defData))
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "failed to load %q: %v", defFilename, err)
		return
	}

	frame := def.Frames[0] //rand.Intn(len(def.Frames))]

	fmt.Println(frame)

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

func (s *Server) HandleDefPalette(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	defFilename := fmt.Sprintf("%v.def", chi.URLParamFromCtx(ctx, "defFile"))
	defData, err := s.lod.ReadFile(defFilename)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w, err)
		return
	}

	def, err := def.Parse(bytes.NewReader(defData))
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "failed to load %q: %v", defFilename, err)
		return
	}

	frame := def.Frames[0] //rand.Intn(len(def.Frames))]

	fmt.Println(frame)

	img, err := frame.PaletteImage()
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
