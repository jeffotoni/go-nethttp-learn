package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func newRootHandler() http.Handler {
	publicMux := http.NewServeMux()
	publicMux.HandleFunc("GET /ping", publicPing)

	adminMux := http.NewServeMux()
	adminMux.HandleFunc("GET /report", adminReport)

	root := http.NewServeMux()
	root.Handle("/public/", http.StripPrefix("/public", chain(requestIDMiddleware, accessLogMiddleware)(publicMux)))
	root.Handle("/admin/", http.StripPrefix("/admin", chain(requestIDMiddleware, accessLogMiddleware, apiKeyMiddleware("secret"))(adminMux)))
	return root
}

func TestPublicGroup(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/public/ping", nil)
	rr := httptest.NewRecorder()

	newRootHandler().ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
	if rr.Header().Get("X-Request-Id") == "" {
		t.Fatal("expected X-Request-Id header")
	}
}

func TestAdminGroupRequiresAPIKey(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/admin/report", nil)
	rr := httptest.NewRecorder()

	newRootHandler().ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Fatalf("expected status 401, got %d", rr.Code)
	}
}

func TestAdminGroupWithAPIKey(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/admin/report", nil)
	req.Header.Set("X-API-Key", "secret")
	rr := httptest.NewRecorder()

	newRootHandler().ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
}
