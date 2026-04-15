package main

import (
	"log"
	"sync"

	"gonethttplocal/examples/10-dual-listenandserve-goroutines/internal/dualserver"
)

func main() {
	s := dualserver.New()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		if err := s.RunMain(); err != nil {
			log.Printf("server :8080 error: %v", err)
		}
	}()

	go func() {
		defer wg.Done()
		if err := s.RunMock(); err != nil {
			log.Printf("server :3000 error: %v", err)
		}
	}()

	log.Println("servers started on :8080 and :3000")
	wg.Wait()
}
