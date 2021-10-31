package main

import (
	"fmt"
	"log"
	"os"

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

	fmt.Println(lod)

	def, err := lod.ReadFile("AB04_.def")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(def)
}

const usage = `lod H3sprite.lod
`
