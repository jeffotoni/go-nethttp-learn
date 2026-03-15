// -> curl -i localhost:8080/docs
// -> curl -i localhost:8080/openapi.yaml
package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// API Routes
	mux.HandleFunc("GET /api/v1/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`[{"name":"Jeff","email":"jeff@email.com"}]`))
	})

	// Serves openapi.yaml file directly
	mux.HandleFunc("GET /openapi.yaml", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/yaml")
		// Correct path relative to where you run the server
		http.ServeFile(w, r, "examples/09-swagger-docs/docs/openapi.yaml")
	})

	// Serves Swagger UI via CDN (pure HTML, no Go dependency)
	mux.HandleFunc("GET /docs", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte(`<!DOCTYPE html>
<html>
<head>
  <title>API Docs</title>
  <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist/swagger-ui.css"/>
</head>
<body>
  <div id="swagger-ui"></div>
  <script src="https://unpkg.com/swagger-ui-dist/swagger-ui-bundle.js"></script>
  <script>
    SwaggerUIBundle({ url: "/openapi.yaml", dom_id: "#swagger-ui" })
  </script>
</body>
</html>`))
	})

	log.Println("Docs at http://localhost:8080/docs")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
