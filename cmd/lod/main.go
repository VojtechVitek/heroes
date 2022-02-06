package main

import (
	"bytes"
	"fmt"
	"log"
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

	defFilename := "AVXsirn0.def"

	defData, err := lod.ReadFile(defFilename)
	if err != nil {
		log.Fatal(err)
	}

	def, err := def.Parse(bytes.NewReader(defData))
	if err != nil {
		log.Fatalf("failed to load %q: %v", defFilename, err)
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

	fmt.Printf("%v %#v\n", err, def)
}

const usage = `lod H3sprite.lod
`
