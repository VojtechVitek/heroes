package main

import (
	"bytes"
	"fmt"
	"image/png"
	"log"
	"math/rand"
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

	r.Get("/H3sprite.lod/{defFile}.def", srv.HandleDef)

	if err := http.ListenAndServe("0.0.0.0:3003", r); err != nil {
		log.Fatal(err)
	}
}

const usage = `lod H3sprite.lod
`

type Server struct {
	lod *lod.LOD
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

	frame := def.Frames[rand.Intn(len(def.Frames))]

	if frame.Format != 0 {
		w.WriteHeader(500)
		fmt.Fprintf(w, "unsupported format %v", frame.Format)
		return
	}

	img := frame.Image()

	w.Header().Set("Content-Type", "image/png")
	if err := png.Encode(w, img); err != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w, err)
		return
	}
}

// // Sprite
// sprite := &Sprite{
// 	Format:     get.Int(4),
// 	FullWidth:  get.Int(4),
// 	FullHeight: get.Int(4),
// 	Width:      get.Int(4),
// 	Height:     get.Int(4),
// 	LeftMargin: get.Int(4),
// 	TopMargin:  get.Int(4),
// 	//DataImage [][]byte
// }

// switch sprite.Format {
// default:
// 	log.Fatalf("sprite.Format: %v", sprite.Format)
// }

// fmt.Printf("%v %#v\n", err, def)
