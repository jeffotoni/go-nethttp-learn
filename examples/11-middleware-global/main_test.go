package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMiddlewareChainAddsRequestID(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/ping", pingHandler)

	h := chain(recoverMiddleware, requestIDMiddleware, accessLogMiddleware)(mux)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/ping", nil)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
	if rr.Header().Get("X-Request-Id") == "" {
		t.Fatal("expected X-Request-Id header")
	}
}

func TestRecoverMiddlewareHandlesPanic(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /panic", panicHandler)

	h := chain(recoverMiddleware, requestIDMiddleware, accessLogMiddleware)(mux)

	req := httptest.NewRequest(http.MethodGet, "/panic", nil)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 500, got %d", rr.Code)
	}
	if rr.Body.String() != `{"error":"internal_server_error"}` {
		t.Fatalf("unexpected body: %s", rr.Body.String())
	}
}
