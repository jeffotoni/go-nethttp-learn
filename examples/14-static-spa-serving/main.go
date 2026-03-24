// -> curl -i localhost:8080/app/
// -> curl -i localhost:8080/app/assets/app.js
// -> curl -i localhost:8080/app/dashboard
package main

import (
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func serveSPA(prefix, distDir string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rel := strings.TrimPrefix(r.URL.Path, prefix)
		rel = strings.TrimPrefix(rel, "/")
		rel = path.Clean("/" + rel)
		rel = strings.TrimPrefix(rel, "/")

		if strings.Contains(rel, "..") {
			http.NotFound(w, r)
			return
		}

		if rel == "" {
			rel = "index.html"
		}

		fullPath := filepath.Join(distDir, filepath.FromSlash(rel))
		if info, err := os.Stat(fullPath); err == nil && !info.IsDir() {
			http.ServeFile(w, r, fullPath)
			return
		}

		if filepath.Ext(rel) != "" {
			http.NotFound(w, r)
			return
		}

		http.ServeFile(w, r, filepath.Join(distDir, "index.html"))
	})
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/app/", http.StripPrefix("/app", serveSPA("", "examples/14-static-spa-serving/web")))

	http.ListenAndServe(":8080", mux)
}
