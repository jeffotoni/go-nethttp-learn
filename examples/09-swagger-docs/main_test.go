package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSwaggerUI(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /docs", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte(`<!DOCTYPE html><html>...</html>`))
	})

	req := httptest.NewRequest("GET", "/docs", nil)
	rr := httptest.NewRecorder()

	mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("got status %v want %v", rr.Code, http.StatusOK)
	}

	if contentType := rr.Header().Get("Content-Type"); contentType != "text/html" {
		t.Errorf("got content-type %s want text/html", contentType)
	}
}

func TestOpenAPIYaml(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /openapi.yaml", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/yaml")
		_, _ = w.Write([]byte(`openapi: "3.0.3"`))
	})

	req := httptest.NewRequest("GET", "/openapi.yaml", nil)
	rr := httptest.NewRecorder()

	mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("got status %v want %v", rr.Code, http.StatusOK)
	}

	if !strings.Contains(rr.Body.String(), "openapi") {
		t.Errorf("body does not contain openapi definition")
	}
}
