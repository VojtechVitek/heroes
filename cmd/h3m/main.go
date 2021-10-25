package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/VojtechVitek/heroes/pkg/h3m"
	"github.com/VojtechVitek/heroes/router"
	"github.com/VojtechVitek/heroes/rpc"
)

const VERSION = "v0.0.1"

func main() {
	mapFileName := "./maps/Loss of Innocence.h3m"

	if len(os.Args) >= 2 {
		mapFileName = os.Args[1]
	}

	f, err := os.Open(mapFileName)
	if err != nil {
		log.Fatal(err)
	}

	m, err := h3m.Parse(f)
	if err != nil {
		log.Fatal(err)
	}

	heroes := &rpc.RPC{
		Maps: map[string]*h3m.H3M{
			mapFileName: m,
		},
	}

	// Create HTTP server.
	bind := fmt.Sprintf("0.0.0.0:%d", 7777)
	srv := &http.Server{
		Addr:              bind,
		Handler:           router.Router(heroes),
		IdleTimeout:       60 * time.Second, // idle connections
		ReadHeaderTimeout: 10 * time.Second, // request header
		ReadTimeout:       5 * time.Minute,  // request body
		WriteTimeout:      5 * time.Minute,  // response body
		MaxHeaderBytes:    1 << 20,          // 1 MB
	}

	log.Printf("Archive API (%v) serving at %v", VERSION, bind)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

const usage = `h3m map.h3m
`
