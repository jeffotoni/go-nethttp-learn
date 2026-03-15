package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestOkHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/ok", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(okHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"data":{"course":"net/http"},"message":"ok"}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestBadHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/bad", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(badHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}
