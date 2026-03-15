package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandlers(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", healthz)
	mux.HandleFunc("GET /readyz", readyz)
	mux.HandleFunc("GET /livez", livez)

	tests := []struct {
		url    string
		status int
	}{
		{"/healthz", http.StatusOK},
		{"/readyz", http.StatusOK},
		{"/livez", http.StatusOK},
	}

	for _, tt := range tests {
		req := httptest.NewRequest("GET", tt.url, nil)
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if rr.Code != tt.status {
			t.Errorf("%s: got status %v want %v", tt.url, rr.Code, tt.status)
		}
	}
}
