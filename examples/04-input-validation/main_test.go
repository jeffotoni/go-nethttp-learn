package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateUser(t *testing.T) {
	tests := []struct {
		name        string
		body        string
		contentType string
		wantStatus  int
	}{
		{"Valid JSON", `{"name":"Jeff","email":"jeff@email.com"}`, "application/json", http.StatusCreated},
		{"Invalid JSON", `{"name":"Jeff","email":"missing-quote}`, "application/json", http.StatusBadRequest},
		{"Missing Content-Type", `{"name":"Jeff"}`, "", http.StatusUnsupportedMediaType},
		{"Unknown Fields", `{"name":"Jeff","email":"a@b.com","extra":"x"}`, "application/json", http.StatusBadRequest},
		{"Empty Fields", `{"name":"","email":""}`, "application/json", http.StatusUnprocessableEntity},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/api/v1/users", strings.NewReader(tt.body))
			if tt.contentType != "" {
				req.Header.Set("Content-Type", tt.contentType)
			}
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(createUser)

			handler.ServeHTTP(rr, req)

			if rr.Code != tt.wantStatus {
				t.Errorf("%s: got status %v want %v", tt.name, rr.Code, tt.wantStatus)
			}
		})
	}
}
