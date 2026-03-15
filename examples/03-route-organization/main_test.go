package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoutes(t *testing.T) {
	mux := http.NewServeMux()
	registerRoutes(mux)

	tests := []struct {
		method string
		url    string
		status int
	}{
		{"GET", "/healthz", http.StatusOK},
		{"GET", "/api/v1/users/123", http.StatusOK},
		{"POST", "/api/v1/users", http.StatusCreated},
	}

	for _, tt := range tests {
		req := httptest.NewRequest(tt.method, tt.url, nil)
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if rr.Code != tt.status {
			t.Errorf("%s %s: got status %v want %v", tt.method, tt.url, rr.Code, tt.status)
		}
	}
}
