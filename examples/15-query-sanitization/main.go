// -> curl -i "localhost:8080/api/v1/logs?offset=0&limit=10&type=event&from=2026-03-01&to=2026-03-31"
// -> curl -i "localhost:8080/api/v1/logs?limit=9999"
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type LogFilters struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Type   string `json:"type,omitempty"`
	From   string `json:"from,omitempty"`
	To     string `json:"to,omitempty"`
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	b, err := json.Marshal(payload)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error":{"code":"internal_error","message":"json_encode_failed"}}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(b)
}

func writeError(w http.ResponseWriter, status int, code, message string) {
	writeJSON(w, status, map[string]any{
		"error": APIError{Code: code, Message: message},
	})
}

func parseIntParam(raw string, def, min, max int) (int, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return def, nil
	}
	n, err := strconv.Atoi(raw)
	if err != nil {
		return 0, fmt.Errorf("must be an integer")
	}
	if n < min || n > max {
		return 0, fmt.Errorf("must be between %d and %d", min, max)
	}
	return n, nil
}

func parseDateParam(raw string) (string, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return "", nil
	}
	if _, err := time.Parse("2006-01-02", raw); err != nil {
		return "", fmt.Errorf("must use YYYY-MM-DD")
	}
	return raw, nil
}

func parseEnumParam(raw string, allowed ...string) (string, error) {
	raw = strings.ToLower(strings.TrimSpace(raw))
	if raw == "" {
		return "", nil
	}
	for _, v := range allowed {
		if raw == v {
			return raw, nil
		}
	}
	return "", fmt.Errorf("must be one of: %s", strings.Join(allowed, ", "))
}

func logsHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	offset, err := parseIntParam(q.Get("offset"), 0, 0, 100000)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid_offset", err.Error())
		return
	}

	limit, err := parseIntParam(q.Get("limit"), 50, 1, 500)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid_limit", err.Error())
		return
	}

	typeValue, err := parseEnumParam(q.Get("type"), "event", "status")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid_type", err.Error())
		return
	}

	from, err := parseDateParam(q.Get("from"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid_from", err.Error())
		return
	}

	to, err := parseDateParam(q.Get("to"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid_to", err.Error())
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"filters": LogFilters{
			Offset: offset,
			Limit:  limit,
			Type:   typeValue,
			From:   from,
			To:     to,
		},
		"data": []map[string]any{},
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/logs", logsHandler)

	http.ListenAndServe(":8080", mux)
}
