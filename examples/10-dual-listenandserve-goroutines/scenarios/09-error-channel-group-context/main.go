package main

import (
	"context"
	"log"
	"time"

	"gonethttplocal/examples/10-dual-listenandserve-goroutines/internal/dualserver"
)

func main() {
	s := dualserver.New()
	errCh := make(chan error, 2)
	s.StartBoth(errCh)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Println("servers started on :8080 and :3000")

	select {
	case err := <-errCh:
		if err != nil {
			log.Printf("error received: %v", err)
			cancel()
		}
	case <-ctx.Done():
	}

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()
	if err := s.Shutdown(shutdownCtx); err != nil {
		log.Printf("shutdown error: %v", err)
	}
}
