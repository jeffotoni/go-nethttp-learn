// -> curl -i localhost:8080/public/ping
// -> curl -i localhost:8080/admin/report
// -> curl -i localhost:8080/admin/report -H "X-API-Key: secret"
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Middleware func(http.Handler) http.Handler

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

func chain(middlewares ...Middleware) Middleware {
	return func(final http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			final = middlewares[i](final)
		}
		return final
	}
}

func requestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-Id")
		if requestID == "" {
			requestID = fmt.Sprintf("req-%d", time.Now().UnixNano())
		}
		w.Header().Set("X-Request-Id", requestID)
		next.ServeHTTP(w, r)
	})
}

func accessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
		started := time.Now()
		next.ServeHTTP(rec, r)
		log.Printf("method=%s path=%s status=%d took=%s", r.Method, r.URL.Path, rec.status, time.Since(started))
	})
}

func apiKeyMiddleware(expected string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("X-API-Key") != expected {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				_, _ = w.Write([]byte(`{"error":"unauthorized"}`))
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func publicPing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"scope":"public","message":"pong"}`))
}

func adminReport(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"scope":"admin","report":"ok"}`))
}

func main() {
	publicMux := http.NewServeMux()
	publicMux.HandleFunc("GET /ping", publicPing)

	adminMux := http.NewServeMux()
	adminMux.HandleFunc("GET /report", adminReport)

	root := http.NewServeMux()
	root.Handle("/public/", http.StripPrefix("/public", chain(requestIDMiddleware, accessLogMiddleware)(publicMux)))
	root.Handle("/admin/", http.StripPrefix("/admin", chain(requestIDMiddleware, accessLogMiddleware, apiKeyMiddleware("secret"))(adminMux)))

	log.Fatal(http.ListenAndServe(":8080", root))
}
