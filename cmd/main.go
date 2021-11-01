package main

import (
	"log"

	"github.com/broswen/tqs/internal/server"
)

func main() {

	server, err := server.New()
	if err != nil {
		log.Fatalf("init server: %v\n", err)
	}

	if err := server.Start(); err != nil {
		log.Fatalf("start server: %v\n", err)
	}
}
