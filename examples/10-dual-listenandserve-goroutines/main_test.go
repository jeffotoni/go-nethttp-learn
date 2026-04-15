package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMainServerRoutes(t *testing.T) {
	mux := newMainMux()

	tests := []struct {
		method string
		path   string
		status int
	}{
		{method: http.MethodGet, path: "/api/v1/user", status: http.StatusOK},
		{method: http.MethodPost, path: "/api/v1/user", status: http.StatusCreated},
	}

	for _, tt := range tests {
		req := httptest.NewRequest(tt.method, tt.path, nil)
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if rr.Code != tt.status {
			t.Errorf("%s %s: got status %d want %d", tt.method, tt.path, rr.Code, tt.status)
		}
	}
}

func TestMockServerRoutes(t *testing.T) {
	mux := newMockMux()

	tests := []struct {
		method string
		path   string
		status int
	}{
		{method: http.MethodGet, path: "/api/v1/mock/user", status: http.StatusOK},
		{method: http.MethodPost, path: "/api/v1/mock/user", status: http.StatusCreated},
	}

	for _, tt := range tests {
		req := httptest.NewRequest(tt.method, tt.path, nil)
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if rr.Code != tt.status {
			t.Errorf("%s %s: got status %d want %d", tt.method, tt.path, rr.Code, tt.status)
		}
	}
}
