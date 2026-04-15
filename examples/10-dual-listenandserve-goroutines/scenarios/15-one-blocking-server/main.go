package main

import (
	"log"

	"gonethttplocal/examples/10-dual-listenandserve-goroutines/internal/dualserver"
)

func main() {
	s := dualserver.New()
	go func() {
		if err := s.RunMock(); err != nil {
			log.Printf("server :3000 error: %v", err)
		}
	}()

	if err := s.RunMain(); err != nil {
		log.Fatalf("server :8080 error: %v", err)
	}
}
