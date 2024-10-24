package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	serveStatic(mux)

	port := "8080"
	log.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))

}

func serveStatic(r *http.ServeMux) {
	staticDir := "./static"
	fs := http.FileServer(http.Dir(staticDir))
	r.Handle("/web/", http.StripPrefix("/web/", fs))
}
