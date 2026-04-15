// -> curl -i localhost:8080/api/v1/user
// -> curl -i -X POST localhost:8080/api/v1/user
// -> curl -i localhost:3000/api/v1/mock/user
// -> curl -i -X POST localhost:3000/api/v1/mock/user
package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type payload map[string]string

func writeJSON(w http.ResponseWriter, status int, body payload) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, payload{
		"scope":   "main",
		"method":  r.Method,
		"message": "user read from :8080",
	})
}

func postUser(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusCreated, payload{
		"scope":   "main",
		"method":  r.Method,
		"message": "user created on :8080",
	})
}

func getMockUser(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, payload{
		"scope":   "mock",
		"method":  r.Method,
		"message": "mock user read from :3000",
	})
}

func postMockUser(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusCreated, payload{
		"scope":   "mock",
		"method":  r.Method,
		"message": "mock user created on :3000",
	})
}

func newMainMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/user", getUser)
	mux.HandleFunc("POST /api/v1/user", postUser)
	return mux
}

func newMockMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/mock/user", getMockUser)
	mux.HandleFunc("POST /api/v1/mock/user", postMockUser)
	return mux
}

func startServer(srv *http.Server, errCh chan<- error) {
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		errCh <- err
	}
}

func main() {
	mainServer := &http.Server{
		Addr:              ":8080",
		Handler:           newMainMux(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	mockServer := &http.Server{
		Addr:              ":3000",
		Handler:           newMockMux(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	errCh := make(chan error, 2)
	go startServer(mainServer, errCh)
	go startServer(mockServer, errCh)

	log.Println("server :8080 started")
	log.Println("server :3000 started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-quit:
		log.Println("shutdown signal received")
	case err := <-errCh:
		log.Fatalf("server failed: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := mainServer.Shutdown(ctx); err != nil {
		log.Printf("error shutting down :8080: %v", err)
	}
	if err := mockServer.Shutdown(ctx); err != nil {
		log.Printf("error shutting down :3000: %v", err)
	}
}
