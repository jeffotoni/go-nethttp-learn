// Run with: PORT=9090 DATABASE_URL=postgres://localhost/mydb go run main.go
// -> curl -i localhost:9090/ping
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Port with fallback
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Mandatory variable: stops if not defined
	dbURL, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		log.Fatal("DATABASE_URL not defined")
	}

	// Optional variable with fallback
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	fmt.Printf("Starting at :%s (env=%s db=%s)\n", port, env, dbURL)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("pong"))
	})

	log.Fatal(http.ListenAndServe(":"+port, mux))
}
