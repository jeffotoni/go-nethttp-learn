package main

import "gonethttplocal/examples/10-dual-listenandserve-goroutines/internal/dualserver"

func main() {
	s := dualserver.New()
	errCh := make(chan error, 2)
	s.StartBoth(errCh)

	ch := make(chan struct{})
	for range ch {
	}
}
