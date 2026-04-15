package main

import "net/http"
import "log"

func main() {
    addr := ":3000"
    log.Printf("Serving on http://localhost%s\n", addr)
    log.Fatal(http.ListenAndServe(addr, http.FileServer(http.Dir("."))))
}
