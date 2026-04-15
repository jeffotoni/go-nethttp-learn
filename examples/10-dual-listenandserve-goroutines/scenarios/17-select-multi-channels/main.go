package main

import (
	"log"

	"gonethttplocal/examples/10-dual-listenandserve-goroutines/internal/dualserver"
)

func main() {
	s := dualserver.New()
	errCh := make(chan error, 2)
	s.StartBoth(errCh)

	stop := make(chan struct{})
	select {
	case <-stop:
		log.Println("stop channel closed")
	case err := <-errCh:
		log.Printf("error channel received: %v", err)
	}
}
