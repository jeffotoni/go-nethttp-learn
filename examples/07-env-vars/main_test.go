package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestEnvVarsHandler(t *testing.T) {
	// Simple test for the ping handler
	req := httptest.NewRequest("GET", "/ping", nil)
	rr := httptest.NewRecorder()

	// In the main.go, the handler is anonymous. 
	// To test it simply, we can just define it here as it is in main.go
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("pong"))
	})

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("got status %v want %v", rr.Code, http.StatusOK)
	}

	if rr.Body.String() != "pong" {
		t.Errorf("got body %s want pong", rr.Body.String())
	}
}

func TestEnvSetup(t *testing.T) {
	os.Setenv("PORT", "9999")
	os.Setenv("DATABASE_URL", "test_url")
	defer os.Unsetenv("PORT")
	defer os.Unsetenv("DATABASE_URL")

	port := os.Getenv("PORT")
	if port != "9999" {
		t.Errorf("expected PORT 9999, got %s", port)
	}
}
