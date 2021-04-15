package main

import (
	"fmt"
	"log"
	"os"

	"github.com/VojtechVitek/heroes/pkg/h3m"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal(usage)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	m, err := h3m.Parse(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", m)
}

const usage = `h3m map.h3m
`
