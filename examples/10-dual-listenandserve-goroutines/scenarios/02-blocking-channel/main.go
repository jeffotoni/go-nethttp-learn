package main

import (
	"log"

	"gonethttplocal/examples/10-dual-listenandserve-goroutines/internal/dualserver"
)

func main() {
	s := dualserver.New()
	errCh := make(chan error, 2)
	s.StartBoth(errCh)

	log.Println("servers started on :8080 and :3000")
	done := make(chan struct{})
	<-done
}
