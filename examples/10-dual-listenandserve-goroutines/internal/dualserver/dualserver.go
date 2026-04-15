package dualserver

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type payload map[string]string

type Servers struct {
	Main *http.Server
	Mock *http.Server
}

func New() *Servers {
	return &Servers{
		Main: &http.Server{
			Addr:              ":8080",
			Handler:           newMainMux(),
			ReadHeaderTimeout: 5 * time.Second,
		},
		Mock: &http.Server{
			Addr:              ":3000",
			Handler:           newMockMux(),
			ReadHeaderTimeout: 5 * time.Second,
		},
	}
}

func (s *Servers) RunMain() error {
	return listen(s.Main)
}

func (s *Servers) RunMock() error {
	return listen(s.Mock)
}

func (s *Servers) StartBoth(errCh chan<- error) {
	go func() {
		if err := s.RunMain(); err != nil {
			errCh <- err
		}
	}()

	go func() {
		if err := s.RunMock(); err != nil {
			errCh <- err
		}
	}()
}

func (s *Servers) Shutdown(ctx context.Context) error {
	return errors.Join(
		s.Main.Shutdown(ctx),
		s.Mock.Shutdown(ctx),
	)
}

func listen(srv *http.Server) error {
	err := srv.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}
	return err
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

func writeJSON(w http.ResponseWriter, status int, body payload) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}
