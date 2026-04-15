package main

import (
	"sync"

	"gonethttplocal/examples/10-dual-listenandserve-goroutines/internal/dualserver"
)

func main() {
	s := dualserver.New()
	errCh := make(chan error, 2)
	_ = errCh

	var mu sync.Mutex
	mu.Lock()

	go func() {
		errCh <- s.RunMain()
	}()

	go func() {
		errCh <- s.RunMock()
	}()

	mu.Lock()
}
