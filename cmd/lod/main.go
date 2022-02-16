package main

import (
	"bytes"
	"image/png"
	"log"
	"net/http"
	"os"

	"github.com/VojtechVitek/heroes/pkg/def"
	"github.com/VojtechVitek/heroes/pkg/lod"
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

	http.ListenAndServe("0.0.0.0:1001", srv)
}

const usage = `lod H3sprite.lod
`

type Server struct {
	lod *lod.LOD
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		return
	}

	defFilename := "AVXsirn0.def"
	defData, err := s.lod.ReadFile(defFilename)
	if err != nil {
		log.Print(err)
		w.WriteHeader(500)
		return
	}

	def, err := def.Parse(bytes.NewReader(defData))
	if err != nil {
		log.Printf("failed to load %q: %v", defFilename, err)
		w.WriteHeader(500)
		return
	}

	frame := def.Frames[0]

	img := frame.Image()

	w.Header().Set("Content-Type", "image/png")
	if err := png.Encode(w, img); err != nil {
		log.Print(err)
		w.WriteHeader(500)
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
