package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestBasicAuthMiddleware(t *testing.T) {
	// Setup env vars for the test
	os.Setenv("API_USER", "admin")
	os.Setenv("API_PASS", "s3cr3t")
	defer os.Unsetenv("API_USER")
	defer os.Unsetenv("API_PASS")

	mux := http.NewServeMux()
	mux.HandleFunc("GET /user", getUser)
	handler := basicAuthMiddleware(mux)

	tests := []struct {
		name       string
		user       string
		pass       string
		withAuth   bool
		wantStatus int
	}{
		{"No Auth", "", "", false, http.StatusUnauthorized},
		{"Wrong Credentials", "wrong", "wrong", true, http.StatusUnauthorized},
		{"Correct Credentials", "admin", "s3cr3t", true, http.StatusOK},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/user", nil)
			if tt.withAuth {
				req.SetBasicAuth(tt.user, tt.pass)
			}
			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			if rr.Code != tt.wantStatus {
				t.Errorf("%s: got status %v want %v", tt.name, rr.Code, tt.wantStatus)
			}
		})
	}
}
