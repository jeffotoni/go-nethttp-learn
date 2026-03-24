package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogsHandlerValidQuery(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/logs?offset=10&limit=25&type=event&from=2026-03-01&to=2026-03-31", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(logsHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
}

func TestLogsHandlerInvalidLimit(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/logs?limit=9999", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(logsHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rr.Code)
	}
}

func TestLogsHandlerInvalidDate(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/logs?from=03-01-2026", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(logsHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rr.Code)
	}
}

func TestLogsHandlerInvalidEnum(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/logs?type=unknown", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(logsHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rr.Code)
	}
}
