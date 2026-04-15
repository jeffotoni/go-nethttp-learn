package main

import (
	"bufio"
	"fmt"
	"os"

	"gonethttplocal/examples/10-dual-listenandserve-goroutines/internal/dualserver"
)

func main() {
	s := dualserver.New()
	errCh := make(chan error, 2)
	s.StartBoth(errCh)

	fmt.Println("servers started on :8080 and :3000")
	fmt.Println("press ENTER to stop")
	_, _ = bufio.NewReader(os.Stdin).ReadString('\n')
}
