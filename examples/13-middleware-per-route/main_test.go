package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func newMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("GET /public/ping", Use(
		requestIDMiddleware,
		accessLogMiddleware,
	)(http.HandlerFunc(publicPing)))

	mux.Handle("GET /admin/report", Use(
		recoverMiddleware,
		requestIDMiddleware,
		accessLogMiddleware,
		apiKeyMiddleware("secret"),
	)(http.HandlerFunc(adminReport)))

	mux.Handle("GET /panic", Use(
		recoverMiddleware,
		requestIDMiddleware,
		accessLogMiddleware,
	)(http.HandlerFunc(panicHandler)))

	return mux
}

func TestPerRoutePublic(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/public/ping", nil)
	rr := httptest.NewRecorder()

	newMux().ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
}

func TestPerRouteAdminRequiresAPIKey(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/admin/report", nil)
	rr := httptest.NewRecorder()

	newMux().ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Fatalf("expected status 401, got %d", rr.Code)
	}
}

func TestPerRouteAdminWithAPIKey(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/admin/report", nil)
	req.Header.Set("X-API-Key", "secret")
	rr := httptest.NewRecorder()

	newMux().ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
}

func TestPerRouteRecover(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/panic", nil)
	rr := httptest.NewRecorder()

	newMux().ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 500, got %d", rr.Code)
	}
}
