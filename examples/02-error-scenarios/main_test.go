package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestErrorByScenario(t *testing.T) {
	tests := []struct {
		scenario string
		wantCode int
	}{
		{"bad-json", http.StatusBadRequest},
		{"not-found", http.StatusNotFound},
		{"unknown", http.StatusInternalServerError},
	}

	for _, tt := range tests {
		t.Run(tt.scenario, func(t *testing.T) {
			mux := http.NewServeMux()
			mux.HandleFunc("GET /error/{scenario}", errorByScenario)

			req := httptest.NewRequest("GET", "/error/"+tt.scenario, nil)
			rr := httptest.NewRecorder()

			mux.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.wantCode {
				t.Errorf("handler for %s returned wrong status code: got %v want %v", tt.scenario, status, tt.wantCode)
			}
		})
	}
}
