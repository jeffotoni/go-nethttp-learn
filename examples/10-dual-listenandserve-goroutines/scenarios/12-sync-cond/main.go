package main

import (
	"sync"

	"gonethttplocal/examples/10-dual-listenandserve-goroutines/internal/dualserver"
)

func main() {
	s := dualserver.New()
	errCh := make(chan error, 2)
	s.StartBoth(errCh)

	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	mu.Lock()
	cond.Wait()
}
