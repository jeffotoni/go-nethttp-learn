// -> curl -i localhost:8080/api/v1/user
// -> curl -i -u admin:s3cr3t localhost:8080/api/v1/user
package main

import (
	"crypto/subtle"
	"log"
	"net/http"
	"os"
)

func unauthorized(w http.ResponseWriter) {
	w.Header().Set("WWW-Authenticate", `Basic realm="api", charset="UTF-8"`)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	_, _ = w.Write([]byte(`{"error":"unauthorized"}`))
}

func basicAuthMiddleware(next http.Handler) http.Handler {
	// Reads credentials from environment variables
	expectedUser := os.Getenv("API_USER")
	expectedPass := os.Getenv("API_PASS")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok {
			unauthorized(w)
			return
		}
		userOK := subtle.ConstantTimeCompare([]byte(user), []byte(expectedUser)) == 1
		passOK := subtle.ConstantTimeCompare([]byte(pass), []byte(expectedPass)) == 1
		if !userOK || !passOK {
			unauthorized(w)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"name":"Jeff","email":"jeff@email.com"}`))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/user", getUser)
	finalHandler := basicAuthMiddleware(mux)
	log.Fatal(http.ListenAndServe(":8080", finalHandler))
}
