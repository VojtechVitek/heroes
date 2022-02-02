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
	defFilename := "AVXsirn0.def"

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

	fmt.Println(lod)

	defData, err := lod.ReadFile(defFilename)
	if err != nil {
		log.Fatal(err)
	}

	def, err := def.Parse(bytes.NewReader(defData))
	if err != nil {
		log.Fatalf("failed to load %q: %v", defFilename, err)
	}

	fmt.Printf("%v %#v\n", err, def)
}

const usage = `lod H3sprite.lod
`
