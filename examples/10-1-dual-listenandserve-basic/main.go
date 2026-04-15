package main

import (
	"log"
	"net/http"
)

// newMainMux registers routes for port 8080.
func newMainMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/user", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"message":"user read from :8080"}`))
	})
	mux.HandleFunc("POST /api/v1/user", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"message":"user created on :8080"}`))
	})
	return mux
}

// newMockMux registers routes for port 3000.
func newMockMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/mock/user", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"message":"mock user read from :3000"}`))
	})
	mux.HandleFunc("POST /api/v1/mock/user", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"message":"mock user created on :3000"}`))
	})
	return mux
}

// startServer starts one HTTP server in a goroutine.
func startServer(srv *http.Server) {
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server %s error: %v", srv.Addr, err)
		}
	}()
}

// main starts both servers and blocks the process.
func main() {
	mainServer := &http.Server{Addr: ":8080", Handler: newMainMux()}
	mockServer := &http.Server{Addr: ":3000", Handler: newMockMux()}

	startServer(mainServer)
	startServer(mockServer)

	select {} // 01) Simple, direct, less controllable.
	// done := make(chan struct{}); <-done // 02) Explicit blocking channel.
	// var wg sync.WaitGroup; wg.Add(2); ...; wg.Wait() // 03) WaitGroup.
	// stop := make(chan os.Signal, 1); signal.Notify(stop, os.Interrupt, syscall.SIGTERM); <-stop // 04) Signal channel.
	// time.Sleep(10 * time.Minute) // 05) Fixed sleep window.
	// ctx, cancel := context.WithCancel(context.Background()); <-ctx.Done(); cancel() // 06) Context cancel.
	// for { time.Sleep(time.Hour) } // 07) Infinite loop.
	// errCh := make(chan error, 2); if err := <-errCh; err != nil { ... } // 08) Error channel group.
	// ctx, cancel := context.WithCancel(context.Background()); select { case err := <-errCh: cancel() } // 09) Error channel + context.
	// var mu sync.Mutex; mu.Lock(); mu.Lock() // 10) Mutex deadlock (didactic).
	// runtime.Goexit() // 11) Ends main goroutine only.
	// var mu sync.Mutex; cond := sync.NewCond(&mu); mu.Lock(); cond.Wait() // 12) sync.Cond wait.
	// ch := make(chan struct{}); for range ch {} // 13) Channel range block.
	// fmt.Scanln() // 14) Block on stdin.
	// go runMock(); runMain() // 15) One blocking server + one goroutine.
	// ticker := time.NewTicker(time.Hour); defer ticker.Stop(); for range ticker.C {} // 16) Infinite ticker.
	// stop := make(chan struct{}); select { case <-stop: case err := <-errCh: _ = err } // 17) Select with multiple channels.
}
